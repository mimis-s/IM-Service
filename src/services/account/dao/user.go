package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/dbmodel"
)

func (d *Dao) GetUserInfoFromID(userID int64) (*dbmodel.AccountUser, bool, error) {
	info := &dbmodel.AccountUser{}
	find, err := d.db.Table((*dbmodel.AccountUser).SubTable(nil, 0)).Where("user_id = ?", userID).Get(info)
	if err != nil {
		errStr := fmt.Sprintf("user ID[%v] get info is err:%v", userID, err)
		fmt.Println(errStr)
		return nil, find, fmt.Errorf(errStr)
	}
	return info, find, err
}

func (d *Dao) GetUserInfoFromIDs(userID []int64) ([]*dbmodel.AccountUser, error) {
	listUser := make([]*dbmodel.AccountUser, 0)
	err := d.db.Table((*dbmodel.AccountUser).SubTable(nil, 0)).In("user_id", userID).Find(&listUser)
	if err != nil {
		errStr := fmt.Sprintf("user ID[%v] get info is err:%v", userID, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}
	return listUser, err
}

func (d *Dao) GetUserInfoFromName(userName string) (*dbmodel.AccountUser, bool, error) {
	info := &dbmodel.AccountUser{}
	find, err := d.db.Table((*dbmodel.AccountUser).SubTable(nil, 0)).Table("account_user").Where("user_name = ?", userName).Get(info)
	if err != nil {
		errStr := fmt.Sprintf("role name[%v] get info is err:%v", userName, err)
		fmt.Println(errStr)
		return nil, find, fmt.Errorf(errStr)
	}
	return info, find, err
}

func (d *Dao) InsertUserInfo(info *dbmodel.AccountUser) error {
	_, err := d.db.Table((*dbmodel.AccountUser).SubTable(nil, 0)).InsertOne(info)
	if err != nil {
		errStr := fmt.Sprintf("user id[%v] name[%v] insert userinfo is err:%v", info.UserId, info.UserName, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}

func (d *Dao) UpdateUserInfo(info *dbmodel.AccountUser) error {
	_, err := d.db.Table((*dbmodel.AccountUser).SubTable(nil, 0)).Where("user_id=?", info.UserId).Update(info)
	if err != nil {
		errStr := fmt.Sprintf("user id[%v] name[%v] update userinfo is err:%v", info.UserId, info.UserName, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
