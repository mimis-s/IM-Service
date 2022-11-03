package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/dbmodel"
)

func (d *Dao) GetUserInfoFromID(roleID int64) (*dbmodel.AccountUser, bool, error) {
	info := &dbmodel.AccountUser{}
	find, err := d.Engine.Where("user_id = ?", roleID).Get(info)
	if err != nil {
		errStr := fmt.Sprintf("role ID[%v] get info is err:%v", roleID, err)
		fmt.Println(errStr)
		return nil, find, fmt.Errorf(errStr)
	}
	return info, find, err
}

func (d *Dao) GetUserInfoFromName(userName string) (*dbmodel.AccountUser, bool, error) {
	info := &dbmodel.AccountUser{}
	find, err := d.Engine.Table("account_user").Where("user_name = ?", userName).Get(info)
	if err != nil {
		errStr := fmt.Sprintf("role name[%v] get info is err:%v", userName, err)
		fmt.Println(errStr)
		return nil, find, fmt.Errorf(errStr)
	}
	return info, find, err
}

func (d *Dao) InsertUserInfo(info *dbmodel.AccountUser) error {
	_, err := d.Engine.InsertOne(info)
	if err != nil {
		errStr := fmt.Sprintf("user id[%v] name[%v] insert userinfo is err:%v", info.UserId, info.UserName, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
