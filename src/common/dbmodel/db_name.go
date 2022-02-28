package dbmodel

type tAccountAccount struct {
	Id        string
	Account   string
	Channel   string
	Platform  string
	SdkId     string
	Country   string
	Roles     string
	Vip       string
	BanReason string
	BanTime   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

var TAccountAccount = &tAccountAccount{
	Id:        "id",
	Account:   "account",
	Channel:   "channel",
	Platform:  "platform",
	SdkId:     "sdk_id",
	Country:   "country",
	Roles:     "roles",
	Vip:       "vip",
	BanReason: "ban_reason",
	BanTime:   "ban_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}
