# wechat-go

`wechat-go`是一个基于微信公众号官方文档构建的Golang版本SDK。

## 接口

已对接的接口有：

| 模块    | 功能        | 方法                                                                                  |
|-------|-----------|-------------------------------------------------------------------------------------|
| 基础接口  | 获取接口调用凭据  | GetAccessToken() (*GetAccessTokenRsp, error)                                        |
| 自定义菜单 | 创建自定义菜单   | CreateCustomMenu(menu *Menu) error                                                  |
| 用户管理  | 获取关注用户列表  | GetFans(nextOpenID string) (*GetFansRsp, error)                                     |
|       | 获取用户基本信息  | GetUserInfo(openID string) (*GetUserInfoRsp, error)                                 |
| 模版消息  | 发送模版消息    | SendTemplateMessage(message *TmplMessage) (int, error)                              |
| 客服消息  | 发送文本消息    | SendTextMessage(toUser, content string) (int, error)                                |
|       | 发送小程序卡片消息 | SendMiniprogramMessage(toUser, title, appid, pagePath, mediaId string) (int, error) |
| 素材管理  | 上传永久素材    | AddMaterial(mediaType, fileName string, media []byte) (string, error)               |

## 快速开始

1. 引入本项目：
   ```bash
   go get github.com/supercat0867/wechat-go
2. 事例：
   ```go
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

3. 更多功能参考源码示例...   