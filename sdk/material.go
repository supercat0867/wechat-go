package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// AddMaterial 上传永久素材
// https://developers.weixin.qq.com/doc/service/api/material/permanent/api_addmaterial.html
// mediaType: 媒体类型，图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (w *WechatSDKImpl) AddMaterial(mediaType, fileName string, media []byte) (string, error) {
	url := fmt.Sprintf("%s/cgi-bin/material/add_material?access_token=%s&type=%s", Domain,
		w.AccessToken, mediaType)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	// 将文件字节流写入表单
	_, err = io.Copy(part, bytes.NewReader(media))
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}

	// 关闭multipart写入器，最终生成整个表单数据
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %v", err)
	}

	// 创建HTTP POST请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	// 设置multipart表单的Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}
	var response AddMediaRsp
	if err = json.Unmarshal(respBody, &response); err != nil {
		return "", err
	}

	if response.Errcode == 0 {
		return response.MediaId, nil
	}

	return "", formatError(response.Errcode, response.Errmsg)
}
