package main

import (
	"github.com/supercat0867/wechat-go/sdk"
)

func main() {
	wechat := sdk.NewWechatSDK("your-appid", "your-appsecret", "")

	// 获取access_token
	rsp, err := wechat.GetAccessToken()
	if err != nil {
		panic(err)
	}
	wechat = sdk.NewWechatSDK("your-appid", "your-appsecret", rsp.AccessToken)

	// 发送文本消息
	if _, err = wechat.SendTextMessage("o_Z5Z5xj_wXZ5Z5xj_wXZ5Z5xj_wX", "hello world"); err != nil {
		panic(err)
	}

	// 发送模版消息
	data := map[string]string{
		"thing2":   "请假流程通知",
		"time15":   "2012-01-02",
		"phrase10": "小明",
		"thing16":  "扶老奶奶过马路",
	}
	tempMessage := sdk.BuildTemplateMessage("obIt16lHlQiZpT5MYC_lTfFv7ZSA", "IWMM8w9XD3jqc01gXyisvG6Y6yPMfGhlGyLPWimAN2w",
		"www.baidu.com", "", data, nil)
	if _, err = wechat.SendTemplateMessage(tempMessage); err != nil {
		panic(err)
	}
}
