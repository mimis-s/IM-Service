package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
)

func (d *Dao) GetFriends(userID int64) (*dbmodel.Friends, bool, error) {
	info := &dbmodel.Friends{}
	find, err := d.Session.Friends.Where("user_id = ?", userID).Get(info)
	if err != nil {
		errStr := fmt.Sprintf("user ID[%v] get friends info is err:%v", userID, err)
		fmt.Println(errStr)
		return nil, find, fmt.Errorf(errStr)
	}
	return info, find, err
}

func (d *Dao) InsertFriends(user *im_home_proto.ClientOnlineInfo, info *dbmodel.Friends) error {
	_, err := d.Session.Friends.InsertOne(info)
	if err != nil {
		errStr := fmt.Sprintf("user id[%v] insert friends is err:%v", user.UserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}

func (d *Dao) UpdateFriends(user *im_home_proto.ClientOnlineInfo, info *dbmodel.Friends) error {
	_, err := d.Session.Friends.Where("user_id=?", user.UserID).Cols(dbmodel.TFriends.Friends).Update(info)
	if err != nil {
		errStr := fmt.Sprintf("user id[%v] update friends is err:%v", user.UserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
