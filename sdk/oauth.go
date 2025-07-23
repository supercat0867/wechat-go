package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WechatSDK interface {
	// GetAccessToken 获取接口调用凭据
	GetAccessToken() (*GetAccessTokenRsp, error)
	// SendTemplateMessage 发送模板消息
	SendTemplateMessage(message *TmplMessage) (int, error)
	// SendTextMessage 发送文本消息
	SendTextMessage(toUser, content string) (int, error)
	// SendMiniprogramMessage 发送小程序卡片
	SendMiniprogramMessage(toUser, title, appid, pagePath, mediaId string) (int, error)
	// CreateCustomMenu 创建自定义菜单
	CreateCustomMenu(menu *Menu) error
	// AddMaterial 上传永久素材
	AddMaterial(mediaType, fileName string, media []byte) (string, error)
	// GetFans 获取关注用户列表
	GetFans(nextOpenID string) (*GetFansRsp, error)
	// GetUserInfo 获取用户基本信息
	GetUserInfo(openID string) (*GetUserInfoRsp, error)
}

type WechatSDKImpl struct {
	AppID       string
	AppSecret   string
	AccessToken string
}

func NewWechatSDK(appID, appSecret, accessToken string) WechatSDK {
	return &WechatSDKImpl{
		AppID:       appID,
		AppSecret:   appSecret,
		AccessToken: accessToken,
	}
}

// GetAccessToken 获取接口调用凭据
// https://developers.weixin.qq.com/doc/service/api/base/api_getaccesstoken.html
func (w *WechatSDKImpl) GetAccessToken() (*GetAccessTokenRsp, error) {
	// 接口地址
	url := fmt.Sprintf("%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		Domain, w.AppID, w.AppSecret)

	client := newHttpClient()
	body, err := client.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var rspJson GetAccessTokenRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return nil, err
	}

	if rspJson.Errcode == 0 {
		return &rspJson, nil
	}

	return nil, formatError(rspJson.Errcode, rspJson.Errmsg)
}
