package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/im_log"
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

	messageID, err := d.cache.Transaction(userMessageIDPrefix, func(pipe redis.Pipeliner) (interface{}, error) {

		strMessageID, err := pipe.HGet(context.Background(), userMessageIDPrefix, keySecondary).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v][%v] redis get message id is err:%v", senderID, receiverID, err)
			im_log.Warn(errStr)
			return 0, fmt.Errorf(errStr)
		}
		messageID, err := strconv.ParseInt(strMessageID, 10, 64)
		if err != nil {
			errStr := fmt.Sprintf("user[%v][%v] redis get message id is err:%v", senderID, receiverID, err)
			im_log.Warn(errStr)
			return 0, fmt.Errorf(errStr)
		}

		_, err = pipe.HSet(context.Background(), userMessageIDPrefix, keySecondary, strconv.FormatInt(messageID+1, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf(" redis set user[%v][%v] message id[%v] is err:%v", receiverID, senderID, messageID+1, err)
			im_log.Warn(errStr)
			return 0, fmt.Errorf(errStr)
		}

		return messageID, nil
	})

	return messageID.(int64), err
}
