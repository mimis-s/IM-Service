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
	userMessageIDPrefix        = "user.message.id"
	userHistoryMessageIDPrefix = "user.message.history.id"
)

// 生成+获取messageid
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

// 保存更新历史数据最新的messageid
func (d *Dao) UpdateHistoryMessageID(senderID, receiverID int64, messageID int64) error {
	userFrist := lib.MinInt64(senderID, receiverID)
	userSecond := lib.MaxInt64(senderID, receiverID)

	keySecondary := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10)

	var luaScript = redis.NewScript(`
		if redis.call("hexists", KEYS[1], ARGV[1]) == 0 then
			redis.call("hset", KEYS[1], ARGV[1], ARGV[2])
		else
			local value = redis.call("hget", KEYS[1], ARGV[1])
			if tonumber(value) < tonumber(ARGV[2]) then
				redis.call("hset", KEYS[1], ARGV[1], ARGV[2])
			end
		end
		return redis.call("hget", KEYS[1], ARGV[1])
	`)

	_, err := luaScript.Run(context.Background(), d.cache.Client, []string{userHistoryMessageIDPrefix}, keySecondary, strconv.FormatInt(messageID, 10)).Result()
	if err != nil {
		return err
	}
	return nil
}

// 获取更新历史数据最新的messageid
func (d *Dao) GetHistoryMessageID(senderID, receiverID int64) (int64, error) {
	userFrist := lib.MinInt64(senderID, receiverID)
	userSecond := lib.MaxInt64(senderID, receiverID)

	keySecondary := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10)

	var luaScript = redis.NewScript(`
		if redis.call("hexists", KEYS[1], ARGV[1]) == 0 then
			return "-1"
		else
			local value = redis.call("hget", KEYS[1], ARGV[1])
			return value
		end
	`)

	var res interface{}
	res, err := luaScript.Run(context.Background(), d.cache.Client, []string{userHistoryMessageIDPrefix}, keySecondary).Result()
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
