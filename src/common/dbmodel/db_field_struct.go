package dbmodel

// 数据库嵌套结构

type TBJsonField1_AccountRoleInfo struct {
	RoleID int64 `json:"role_id,omitempty"`
	Zone   int   `json:"zone,omitempty"`
}

type TBJsonField_AccountRolesInfo struct {
	Roles []*TBJsonField1_AccountRoleInfo `json:"roles"`
}
