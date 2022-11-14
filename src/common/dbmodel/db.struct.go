// 数据库表结构定义
package dbmodel

import (
	"time"
)

type AccountUser struct {
	UserId           int64                     `xorm:"not null pk comment('PARTITION:HASH|10') BIGINT"`
	UserName         string                    `xorm:"not null index VARCHAR(255)"`
	Password         string                    `xorm:"not null index VARCHAR(255)"`
	UserExtraInfo    TBJsonField_UserExtraInfo `xorm:"not null comment('#TBJsonField_UserExtraInfo#额外数据') JSON"`
	LatestLoginTime  int64                     `xorm:"comment('最近登录时间戳') BIGINT"`
	LatestLogoutTime int64                     `xorm:"comment('最近登出时间戳') BIGINT"`
	UserPayInfo      TBJsonField_UserPayInfo   `xorm:"not null comment('#TBJsonField_UserPayInfo#用户付费数据') JSON"`
	Level            int                       `xorm:"not null default 0 comment('等级') INT"`
	Exp              int                       `xorm:"not null default 0 comment('经验值') INT"`
	CreatedAt        time.Time                 `xorm:"created"`
	UpdatedAt        time.Time                 `xorm:"updated"`
	DeletedAt        time.Time                 `xorm:"deleted"`
}

type Friends struct {
	Id        int64               `xorm:"pk autoincr comment('主键') BIGINT"`
	UserId    int64               `xorm:"not null comment('玩家ID') unique BIGINT"`
	Friends   TBJsonField_Friends `xorm:"not null comment('#TBJsonField_Friends# 好友') JSON"`
	CreatedAt time.Time           `xorm:"created"`
	UpdatedAt time.Time           `xorm:"updated"`
	DeletedAt time.Time           `xorm:"deleted"`
}
