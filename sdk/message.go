package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendTextMessage 发送文本消息
// https://developers.weixin.qq.com/doc/service/api/customer/message/api_sendcustommessage.html
// toUser (string): 接收消息的用户的 OpenID。
// content (string): 发送的文本内容。
func (w *WechatSDKImpl) SendTextMessage(toUser, content string) (int, error) {
	data := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	}

	// 创建请求
	url := fmt.Sprintf("%s/cgi-bin/message/custom/send?access_token=%s", Domain, w.AccessToken)
	client := newHttpClient()
	body, err := client.Request(http.MethodPost, url, data)
	if err != nil {
		return 0, err
	}

	var rspJson SendMessageRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return 0, err
	}

	if rspJson.Errcode == 0 {
		return rspJson.MsgID, nil
	}

	return 0, formatError(rspJson.Errcode, rspJson.Errmsg)
}

// SendMiniprogramMessage 发送小程序卡片
// https://developers.weixin.qq.com/doc/service/api/customer/message/api_sendcustommessage.html
// toUser (string): 接收消息的用户的 OpenID。
// title (string): 小程序卡片的标题。
// appid (string): 小程序的 AppID。
// pagePath (string): 小程序页面路径。
// mediaId (string): 小程序卡片的封面图片的 media_id。
func (w *WechatSDKImpl) SendMiniprogramMessage(toUser, title, appid, pagePath, mediaId string) (int, error) {
	data := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "miniprogrampage",
		"miniprogrampage": map[string]interface{}{
			"title":          title,
			"appid":          appid,
			"pagepath":       pagePath,
			"thumb_media_id": mediaId,
		},
	}

	// 创建请求
	url := fmt.Sprintf("%s/cgi-bin/message/custom/send?access_token=%s", Domain, w.AccessToken)
	client := newHttpClient()
	body, err := client.Request(http.MethodPost, url, data)
	if err != nil {
		return 0, err
	}

	var rspJson SendMessageRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return 0, err
	}

	if rspJson.Errcode == 0 {
		return rspJson.MsgID, nil
	}

	return 0, formatError(rspJson.Errcode, rspJson.Errmsg)
}

// SendTemplateMessage 发送模版消息
// https://developers.weixin.qq.com/doc/service/api/notify/template/api_sendtemplatemessage.html
func (w *WechatSDKImpl) SendTemplateMessage(message *TmplMessage) (int, error) {
	url := fmt.Sprintf("%s/cgi-bin/message/template/send?access_token=%s", Domain, w.AccessToken)
	client := newHttpClient()
	body, err := client.Request(http.MethodPost, url, message)
	if err != nil {
		return 0, err
	}

	var rspJson SendMessageRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return 0, err
	}

	if rspJson.Errcode == 0 {
		return rspJson.MsgID, nil
	}

	return 0, formatError(rspJson.Errcode, rspJson.Errmsg)
}
