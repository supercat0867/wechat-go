package msg_handler

import (
	"encoding/xml"
	"log"
	"net/http"
)

type Message struct {
	Type         MessageType // 消息类型
	Content      string      // 消息内容
	ToUserName   string      // 开发者微信号
	FromUserName string      // 发送方openid
	MediaId      string      // 素材ID
	Event        string      // 事件类型
}

type MessageType string

// 消息类型
const (
	TextMessage       MessageType = "text"       // 文本消息
	VoiceMessage      MessageType = "voice"      // 语音消息
	VideoMessage      MessageType = "video"      // 视频消息
	ShortVideoMessage MessageType = "shortvideo" // 小视频消息
	LocationMessage   MessageType = "location"   // 地理位置消息
	LinkMessage       MessageType = "link"       // 链接消息
	EventMessage      MessageType = "event"      // 事件消息
)

// XMLMessage 微信xml消息格式
type XMLMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`            // 开发者微信号
	FromUserName string   `xml:"FromUserName"`          // 发送方账号（一个OpenID）
	CreateTime   int64    `xml:"CreateTime"`            // 消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`               // 消息类型，文本为text
	Content      string   `xml:"Content"`               // 文本消息内容
	MsgId        int64    `xml:"MsgId"`                 // 消息id，64位整型
	MsgDataId    string   `xml:"MsgDataId,omitempty"`   // 消息的数据ID（消息如果来自文章时才有）
	Idx          string   `xml:"Idx,omitempty"`         // 多图文时第几篇文章，从1开始（消息如果来自文章时才有）
	PicUrl       string   `xml:"PicUrl,omitempty"`      // 图片链接（由系统生成）
	MediaId      string   `xml:"MediaId,omitempty"`     // 图片消息媒体id或语音消息媒体id，可以调用获取临时素材接口拉取数据。
	Format       string   `xml:"Format,omitempty"`      // 语音格式，如amr，speex等
	Recognition  string   `xml:"Recognition,omitempty"` // 语音识别结果，UTF8编码 (已废弃)
	Event        string   `xml:"Event,omitempty"`       // 事件类型
}

type MessageHandler func(msg *Message, w http.ResponseWriter)

var messageHandler = make(map[MessageType]MessageHandler)

// RegisterHandler 注册消息处理方法
func RegisterHandler(msgType MessageType, handler MessageHandler) {
	messageHandler[msgType] = handler
}

// 解析微信xml消息到结构体
func parseWeChatMessage(data []byte) (*XMLMessage, error) {
	var msg XMLMessage
	err := xml.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// HandleWeChatMessage 处理消息
func HandleWeChatMessage(data []byte, w http.ResponseWriter) {
	msg, err := parseWeChatMessage(data)
	if err != nil {
		log.Printf("微信消息xml解析失败: %v", err)
		return
	}

	genericMsg := &Message{
		ToUserName:   msg.ToUserName,
		FromUserName: msg.FromUserName,
	}

	switch msg.MsgType {
	case "text":
		genericMsg.Type = TextMessage
		genericMsg.Content = msg.Content
	case "voice":
		genericMsg.Type = VoiceMessage
		// 语音自动转文字能力被官方移除
		genericMsg.Content = msg.Recognition
		genericMsg.MediaId = msg.MediaId
	case "event":
		genericMsg.Type = EventMessage
		genericMsg.Event = msg.Event
	// 添加其他消息类型的转换
	default:
		// 处理未知消息类型
		return
	}

	// 调用对应类型的处理器
	if handler, ok := messageHandler[genericMsg.Type]; ok {
		handler(genericMsg, w)
		return
	}
	log.Printf("未注册此消息类型对应的handler: %s", genericMsg.Type)
}
