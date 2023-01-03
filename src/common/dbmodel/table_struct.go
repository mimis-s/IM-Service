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

// 分库名字
func ShardDatabaseName() map[string]bool {
	return map[string]bool{
		DbSubTreasury_Account: true,
		DbSubTreasury_Chat:    true,
		DbSubTreasury_Relay:   true,
		DbSubTreasury_Friends: true,
		DbSubTreasury_Message: true,
	}
}

// 分表
type DbTableInterface interface {
	SubName() string  // 表名
	SubTableNum() int // 分表数量
	BindSubTreasury() DbSubTreasury
}
