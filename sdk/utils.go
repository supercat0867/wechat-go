package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

// BuildTemplateMessage 构造模版消息
func BuildTemplateMessage(touser, templateID, url, clientMsgID string,
	msgData map[string]string, miniprogram *TmplMessageMiniProgram) *TmplMessage {
	var data = make(map[string]TmplMessageData)
	for key, value := range msgData {
		data[key] = TmplMessageData{value}
	}
	tmplMessage := TmplMessage{
		ToUser:      touser,
		TemplateID:  templateID,
		URL:         url,
		ClientMsgID: clientMsgID,
		Data:        data,
	}
	if miniprogram != nil {
		tmplMessage.MiniProgram = *miniprogram
	}

	return &tmplMessage
}

// BuildTextResponse 构造被动回复文本消息xml
func BuildTextResponse(toUser, fromUser, content string) string {
	return fmt.Sprintf(`<xml>
<ToUserName><![CDATA[%s]]></ToUserName>
<FromUserName><![CDATA[%s]]></FromUserName>
<CreateTime>%d</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[%s]]></Content>
</xml>`, toUser, fromUser, time.Now().Unix(), content)
}

// formatError 格式化错误信息
func formatError(code int, msg string) error {
	return errors.Wrapf(errors.New(msg), "errcode: %d", code)
}

// httpClient 结构体，包含通用 HTTP 客户端
type httpClient struct {
	client *http.Client
}

// newHttpClient 创建一个新的 HttpClient
func newHttpClient() *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Request 封装了通用的 GET 和 POST 请求方法
func (hc *httpClient) Request(method, url string, params interface{}) ([]byte, error) {
	var req *http.Request
	var err error

	// 根据请求方法选择请求类型
	if method == http.MethodGet {
		// 构造 GET 请求
		req, err = http.NewRequest(http.MethodGet, url, nil)
	} else if method == http.MethodPost {
		// 对 params 进行 JSON 编码
		data, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		// 构造 POST 请求
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
	} else {
		return nil, fmt.Errorf("unsupported method: %s", method)
	}

	if err != nil {
		return nil, err
	}

	// 发送请求
	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 如果响应是错误码，直接返回错误信息
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s, response: %s", resp.Status, string(body))
	}

	return body, nil
}
