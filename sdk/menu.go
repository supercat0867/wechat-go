package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateCustomMenu 创建自定义菜单
// https://developers.weixin.qq.com/doc/service/api/custommenu/api_createcustommenu.html
func (w *WechatSDKImpl) CreateCustomMenu(menu *Menu) error {
	url := fmt.Sprintf("%s/cgi-bin/menu/create?access_token=%s", Domain, w.AccessToken)
	client := newHttpClient()
	body, err := client.Request(http.MethodPost, url, menu)
	if err != nil {
		return err
	}

	var rspJson Error
	if err = json.Unmarshal(body, &rspJson); err != nil {
		return err
	}

	if rspJson.Errcode == 0 {
		return nil
	}

	return formatError(rspJson.Errcode, rspJson.Errmsg)
}
