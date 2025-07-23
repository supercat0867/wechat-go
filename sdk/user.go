package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetFans 获取关注用户列表
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_getfans.html
func (w *WechatSDKImpl) GetFans(nextOpenID string) (*GetFansRsp, error) {
	url := fmt.Sprintf("%s/cgi-bin/user/get?access_token=%s&next_openid=%s", Domain,
		w.AccessToken, nextOpenID)

	client := newHttpClient()
	body, err := client.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var rspJson GetFansRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return nil, err
	}

	if rspJson.Errcode == 0 {
		return &rspJson, nil
	}

	return nil, formatError(rspJson.Errcode, rspJson.Errmsg)
}

// GetUserInfo 获取用户基本信息
// https://developers.weixin.qq.com/doc/service/api/usermanage/userinfo/api_userinfo.html
func (w *WechatSDKImpl) GetUserInfo(openID string) (*GetUserInfoRsp, error) {
	// 接口地址
	url := fmt.Sprintf("%s/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN",
		Domain, w.AccessToken, openID)

	client := newHttpClient()
	body, err := client.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var rspJson GetUserInfoRsp
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return nil, err
	}

	if rspJson.Errcode == 0 {
		return &rspJson, nil
	}

	return nil, formatError(rspJson.Errcode, rspJson.Errmsg)
}
