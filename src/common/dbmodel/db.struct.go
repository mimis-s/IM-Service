// 数据库表结构定义
package dbmodel

import (
	"time"
)

type AccountUser struct {
	UserID           int64                     `xorm:"not null pk comment('PARTITION:HASH|10') BIGINT"`
	UserExtraInfo    TBJsonField_UserExtraInfo `xorm:"not null comment('#TBJsonField_UserExtraInfo#额外数据') TEXT"`
	LatestLoginTime  int64                     `xorm:"comment('最近登录时间戳') BIGINT"`
	LatestLogoutTime int64                     `xorm:"comment('最近登出时间戳') BIGINT"`
	UserPayInfo      TBJsonField_UserPayInfo   `xorm:"not null comment('#TBJsonField_UserPayInfo#用户付费数据') TEXT"`
	Level            int                       `xorm:"not null default 0 comment('等级') INT"`
	Exp              int                       `xorm:"not null default 0 comment('经验值') INT"`
	CreatedAt        time.Time                 `xorm:"created"`
	UpdatedAt        time.Time                 `xorm:"updated"`
	DeletedAt        time.Time                 `xorm:"deleted"`
}
