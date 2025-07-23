package sdk

const (
	// Domain 微信公众号域名
	Domain = "https://api.weixin.qq.com"
)

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// GetAccessTokenRsp 获取access_token响应
type GetAccessTokenRsp struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
	Error
}

// SendMessageRsp 发送模版消息响应
type SendMessageRsp struct {
	MsgID int `json:"msgid"` // 消息ID
	Error
}

// Menu 菜单
type Menu struct {
	Button []MenuButton `json:"button"`
}

// MenuButton 菜单按钮
type MenuButton struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name,omitempty"`
	Key       string       `json:"key,omitempty"`
	Url       string       `json:"url,omitempty"`
	AppID     string       `json:"appid,omitempty"`
	PagePath  string       `json:"pagepath,omitempty"`
	SubButton []MenuButton `json:"sub_button,omitempty"`
}

// AddMediaRsp 上传永久素材响应
type AddMediaRsp struct {
	MediaId string `json:"media_id"` // 媒体文件上传后，获取标识
	Error
}

// GetFansRsp 获取关注用户列表响应
type GetFansRsp struct {
	Total int `json:"total"` // 关注该公众账号的总用户数
	Count int `json:"count"` // 拉取的OPENID个数，最大值为10000
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"` // 列表数据，OPENID的列表
	NextOpenID string `json:"next_openid"` // 拉取列表的最后一个用户的OPENID
	Error
}

// GetUserInfoRsp 获取用户基本信息响应
type GetUserInfoRsp struct {
	Subscribe      int    `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenID         string `json:"openid"`          // 用户的标识，对当前公众号唯一
	Language       string `json:"language"`        // 用户的语言，简体中文为zh_CN
	SubscribeTime  int    `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionID        string `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台账号后，才会出现该字段。
	Remark         string `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID        int    `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagIDList      []int  `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubScribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载 ,ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号 , ADD_SCENE_OTHERS 其他
	QRScene        int    `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QRSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
	Error
}

// TmplMessage 模版消息通用格式
type TmplMessage struct {
	ToUser      string                     `json:"touser"`        // 接收者openid
	TemplateID  string                     `json:"template_id"`   // 模板ID
	URL         string                     `json:"url"`           // 模板跳转链接（海外账号没有跳转能力）
	MiniProgram TmplMessageMiniProgram     `json:"miniprogram"`   // 跳小程序所需数据，不需跳小程序可不用传该数据
	ClientMsgID string                     `json:"client_msg_id"` // 防重入id。对于同一个openid + client_msg_id, 只发送一条消息,10分钟有效,超过10分钟不保证效果。若无防重入需求，可不填
	Data        map[string]TmplMessageData `json:"data"`          // 模板数据
}
type TmplMessageData struct {
	Value string `json:"value"`
}
type TmplMessageMiniProgram struct {
	AppID    string `json:"app_id"`   // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagePath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}
