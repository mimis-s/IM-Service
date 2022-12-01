package dbmodel

type TBJsonField_UserExtraInfo struct {
	NikeName          string `json:"nike_name"`          // 昵称
	Nation            int    `json:"nation"`             // 国家
	PersonalSignature string `json:"personal_signature"` // 个性签名
	HeadUrl           string `json:"head_url"`           // 头像地址(这个地址好像没有用,先不管)
	PhoneNumber       string `json:"phone_number"`       // 电话号码
}

// 用户付费数据
type TBJsonField_UserPayInfo struct {
	VipYellow int `json:"vip_yellow"` // 黄钻
	VipGreen  int `json:"vip_green"`  // 绿钻
}

type TBJsonField_Friends struct {
	IDs             []int64
	ApplyFriendIDs  []int64 // 发送的好友申请
	ReceiveApplyIDs []int64 // 收到的好友申请
}
