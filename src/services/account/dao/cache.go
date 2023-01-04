package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
)

/*
	登录之后的用户信息缓存在redis里面
*/

const (
	userStatusRedis = "user.status"
)

func (d *Dao) CacheUserLogin(userID int64) error {
	_, err := d.cache.Client.HSet(context.Background(), userStatusRedis, strconv.FormatInt(userID, 10), "1").Result()
	return err
}

func (d *Dao) CacheUserLogOut(userID int64) error {
	_, err := d.cache.Client.HDel(context.Background(), userStatusRedis, strconv.FormatInt(userID, 10)).Result()
	return err
}

func (d *Dao) CacheGetUserStatus(userIDs []int64) ([]im_home_proto.Enum_UserStatus, error) {
	// cmd, err := d.Cache.HExists(userStatusRedis, strconv.FormatInt(userID, 10)).Result()
	// if err != nil {
	// 	return 0, err
	// }
	// if cmd {
	// 	return im_home_proto.Enum_UserStatus_Enum_UserStatus_Online, nil
	// }
	// return im_home_proto.Enum_UserStatus_Enum_UserStatus_Outline, nil
	var luaScript = redis.NewScript(
		`
		local function split(str,reps)
    		local resultStrList = {}
    		string.gsub(str,'[^'..reps..']+',function ( w )
        	table.insert(resultStrList,w)
   			end)
    		return resultStrList
		end
		local array_args = split(ARGV[1], ",")
		local array = {}
		for i, v in ipairs(array_args)
		do
			if redis.call("hexists", KEYS[1], v) == 0 then
				array[i] = "0"
			else
				array[i] = "1"
			end
		end
		return array
	`)

	str := ""
	for i, v := range userIDs {
		if i == 0 {
			str += strconv.FormatInt(v, 10)
		} else {
			str += "," + strconv.FormatInt(v, 10)
		}
	}

	var res interface{}
	res, err := luaScript.Run(context.Background(), d.cache.Client, []string{userStatusRedis}, str).Result()
	if err != nil {
		return nil, err
	}

	res1, ok := res.([]interface{})
	if !ok {
		return nil, fmt.Errorf("interface to []interface is fail")
	}
	status := make([]im_home_proto.Enum_UserStatus, 0, len(res1))
	for _, v := range res1 {
		if v.(string) == "1" {
			status = append(status, im_home_proto.Enum_UserStatus_Enum_UserStatus_Online)
		} else {
			status = append(status, im_home_proto.Enum_UserStatus_Enum_UserStatus_Outline)
		}
	}
	return status, nil
}
