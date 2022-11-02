package dbmodel

type TBJsonField_UserExtraInfo struct {
	NikeName          string `json:"nike_name"`          // 昵称
	Nation            int    `json:"nation"`             // 国家
	PersonalSignature string `json:"personal_signature"` // 个性签名
	HeadUrl           string `json:"head_url"`           // 头像地址
	PhoneNumber       string `json:"phone_number"`       // 电话号码
}

// 用户付费数据
type TBJsonField_UserPayInfo struct {
	VipYellow int `json:"vip_yellow"` // 黄钻
	VipGreen  int `json:"vip_green"`  // 绿钻
}
