package dbmodel

/*
	分表, 绑定分库
*/

// 账户表
func (tb *AccountUser) SubTableNum() int {
	return 1
}

func (tb *AccountUser) SubTable(baseValue int64) string {
	return subName(tb.SubName(), baseValue, tb.SubTableNum())
}

func (tb *AccountUser) BindSubTreasury() DbSubTreasury {
	return DbSubTreasury_Account
}

// 好友表
func (tb *Friends) SubTableNum() int {
	return 1
}

func (tb *Friends) SubTable(baseValue int64) string {
	return subName(tb.SubName(), baseValue, tb.SubTableNum())
}

func (tb *Friends) BindSubTreasury() DbSubTreasury {
	return DbSubTreasury_Friends
}

// 聊天记录表
func (tb *HistoryMessage) SubTableNum() int {
	return 10
}

func (tb *HistoryMessage) SubTable(baseValue int64) string {
	return subName(tb.SubName(), baseValue, tb.SubTableNum())
}

func (tb *HistoryMessage) BindSubTreasury() DbSubTreasury {
	return DbSubTreasury_Message
}
