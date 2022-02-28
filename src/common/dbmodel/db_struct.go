package dbmodel

// 数据库字段

import "time"

type AccountAccount struct {
	Id        int                          `xorm:"not null pk autoincr INT(10)"`
	Account   string                       `xorm:"not null unique VARCHAR(128)"`
	Channel   string                       `xorm:"comment('渠道') VARCHAR(128)"`
	Platform  string                       `xorm:"comment('平台') VARCHAR(128)"`
	SdkId     int64                        `xorm:"comment('sdk唯一id') BIGINT(20)"`
	Country   string                       `xorm:"comment('国家') VARCHAR(128)"`
	Roles     TBJsonField_AccountRolesInfo `xorm:"comment('#TBJsonField_AccountRolesInfo#拥有角色') JSON"`
	BanReason int                          `xorm:"comment('封禁理由') INT(10)"`
	BanTime   int64                        `xorm:"comment('封禁到时') BIGINT(20)"`
	CreatedAt time.Time                    `xorm:"created"`
	UpdatedAt time.Time                    `xorm:"updated"`
	DeletedAt time.Time                    `xorm:"deleted"`
}
