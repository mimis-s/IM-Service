package dbmodel

/*
	分表, 绑定分库
*/

// 账户表
func (tb *AccountUser) SubTableNum() int {
	return 1
}

func (tb *AccountUser) BindSubTreasury() DbSubTreasury {
	return DbSubTreasury_Account
}
