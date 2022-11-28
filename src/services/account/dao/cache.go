package dao

import (
	"strconv"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
)

/*
	登录之后的用户信息缓存在redis里面
*/

const (
	userStatusRedis = "user.status"
)

func (d *Dao) CacheUserLogin(userID int64) error {
	_, err := d.Cache.HSet(userStatusRedis, strconv.FormatInt(userID, 10), "1").Result()
	return err
}

func (d *Dao) CacheUserLogOut(userID int64) error {
	_, err := d.Cache.HDel(userStatusRedis, strconv.FormatInt(userID, 10)).Result()
	return err
}

func (d *Dao) CacheGetUserStatus(userID int64) (im_home_proto.Enum_UserStatus, error) {
	cmd, err := d.Cache.HExists(userStatusRedis, strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		return 0, err
	}
	if cmd {
		return im_home_proto.Enum_UserStatus_Enum_UserStatus_Online, nil
	}
	return im_home_proto.Enum_UserStatus_Enum_UserStatus_Outline, nil
}
