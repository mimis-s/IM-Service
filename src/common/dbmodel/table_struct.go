package dbmodel

// 分库
type DbSubTreasury string

const (
	DbSubTreasury_Account = "account"
	DbSubTreasury_Chat    = "chat"
	DbSubTreasury_Relay   = "relay"
	DbSubTreasury_Friends = "friends"
	DbSubTreasury_Message = "message"
)

// 分表
type DbTableInterface interface {
	SubName() string  // 表名
	SubTableNum() int // 分表数量
	BindSubTreasury() DbSubTreasury
}
