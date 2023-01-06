package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/golang_tools/lib"
)

// 获取message_id并且自增1

const (
	userMessageIDPrefix = "user.message.id"
)

func (d *Dao) GetMessageID(senderID, receiverID int64) (int64, error) {
	userFrist := lib.MinInt64(senderID, receiverID)
	userSecond := lib.MaxInt64(senderID, receiverID)

	keySecondary := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10)

	var luaScript = redis.NewScript(`
		if redis.call("hexists", KEYS[1], ARGV[1]) == 0 then
			redis.call("hset", KEYS[1], ARGV[1], 1)
			return "0"
		else
			local value = redis.call("hget", KEYS[1], ARGV[1])
			redis.call("hset", KEYS[1], ARGV[1], value + 1)
			return value
		end
	`)

	var res interface{}
	res, err := luaScript.Run(context.Background(), d.cache.Client, []string{userMessageIDPrefix}, keySecondary).Result()
	if err != nil {
		return 0, err
	}

	res1, ok := res.(string)
	if !ok {
		return 0, fmt.Errorf("interface to string is fail")
	}

	messageID, err := strconv.ParseInt(res1, 10, 64)
	if err != nil {
		return 0, err
	}

	return messageID, nil
}
