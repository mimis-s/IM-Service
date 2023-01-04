package dao

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
)

/*
	存储数据的层级关系:前缀key + 用户ID -> 好友ID -> 消息(转化为base64编码存储)；消息
*/

type cacheMessageInfo struct {
	Messages []*im_home_proto.ChatMessage
}

// 编码
func encodeCacheMessage(data *cacheMessageInfo) (string, error) {
	byteMessage, err := json.Marshal(data)
	if err != nil {
		errStr := fmt.Sprintf("json marshal chat message data[%v] is err:%v", data, err)
		im_log.Warn(errStr)
		return "", fmt.Errorf(errStr)
	}
	return base64.StdEncoding.EncodeToString(byteMessage), nil
}

// 解码
func decodeCacheMessage(data string) (*cacheMessageInfo, error) {

	byteResult, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		errStr := fmt.Sprintf("base64 decode chat message data[%v] is err:%v", data, err)
		im_log.Warn(errStr)
		return nil, fmt.Errorf(errStr)
	}

	messageInfo := &cacheMessageInfo{}
	err = json.Unmarshal(byteResult, messageInfo)
	if err != nil {
		errStr := fmt.Sprintf("json unmarshal chat message data[%v] is err:%v", data, err)
		im_log.Warn(errStr)
		return nil, fmt.Errorf(errStr)
	}
	return messageInfo, nil
}

const (
	userOfflineMessagePrefix = "user.message.off.line."
)

// 获取用户离线消息
func (d *Dao) GetUserAllOfflineMessage(userID int64) (map[int64][]*im_home_proto.ChatMessage, error) {
	userOfflineMessage := userOfflineMessagePrefix + strconv.FormatInt(userID, 10)
	// redis事务
	data, err := d.cache.Transaction(userOfflineMessage, func(pipe redis.Pipeliner) (interface{}, error) {

		allMessageData, err := pipe.HGetAll(context.Background(), userOfflineMessage).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get all off line message is err:%v", userID, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		mapChatMessage := make(map[int64][]*im_home_proto.ChatMessage)
		for strSender, messageData := range allMessageData {
			messageInfo, err := decodeCacheMessage(messageData)
			if err != nil {
				errStr := fmt.Sprintf("user[%v] redis decode sender[%v] off line message is err:%v", userID, strSender, err)
				im_log.Warn(errStr)
				return nil, fmt.Errorf(errStr)
			}

			if messageInfo == nil || messageInfo.Messages == nil || len(messageInfo.Messages) == 0 {
				continue
			}

			senderID, err := strconv.ParseInt(strSender, 10, 64)
			if err != nil {
				errStr := fmt.Sprintf("user[%v] redis decode sender[%v] parse int64 is err:%v", userID, strSender, err)
				im_log.Warn(errStr)
				return nil, fmt.Errorf(errStr)
			}

			mapChatMessage[senderID] = messageInfo.Messages
		}

		return mapChatMessage, nil
	})
	return data.(map[int64][]*im_home_proto.ChatMessage), err
}

// 添加用户离线消息
func (d *Dao) AddUserOneOfflineMessage(sender, receiver int64, chatMessage *im_home_proto.ChatMessage) error {

	userOfflineMessage := userOfflineMessagePrefix + strconv.FormatInt(receiver, 10)
	// redis事务
	_, err := d.cache.Transaction(userOfflineMessage, func(pipe redis.Pipeliner) (interface{}, error) {

		messageData, err := pipe.HGet(context.Background(), userOfflineMessage, strconv.FormatInt(sender, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get off line message is err:%v", receiver, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}
		messageInfo, err := decodeCacheMessage(messageData)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode off line message is err:%v", receiver, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		if messageInfo == nil || messageInfo.Messages == nil {
			messageInfo = &cacheMessageInfo{
				Messages: make([]*im_home_proto.ChatMessage, 0),
			}
		}

		messageInfo.Messages = append(messageInfo.Messages, chatMessage)
		info, err := encodeCacheMessage(messageInfo)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get off line message encode is err:%v", receiver, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}
		_, err = pipe.HSet(context.Background(), userOfflineMessage, strconv.FormatInt(sender, 10), info).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis set user[%v] off line message encode is err:%v", receiver, sender, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		return nil, nil
	})

	return err
}

// 把离线消息删除
func (d *Dao) DelUserOneOfflineMessage(senderID, receiverID int64) error {
	userOfflineMessage := userOfflineMessagePrefix + strconv.FormatInt(receiverID, 10)
	// redis事务
	_, err := d.cache.Transaction(userOfflineMessage, func(pipe redis.Pipeliner) (interface{}, error) {

		messageData, err := pipe.HGet(context.Background(), userOfflineMessage, strconv.FormatInt(senderID, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get off line message is err:%v", receiverID, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}
		messageInfo, err := decodeCacheMessage(messageData)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode off line message is err:%v", receiverID, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		if messageInfo == nil || messageInfo.Messages == nil {
			messageInfo = &cacheMessageInfo{
				Messages: make([]*im_home_proto.ChatMessage, 0),
			}
		}

		for index, info := range messageInfo.Messages {
			if info.SenderID == senderID && info.ReceiverID == receiverID {
				messageInfo.Messages = append(messageInfo.Messages[:index], messageInfo.Messages[index+1:]...)
			}
		}

		info, err := encodeCacheMessage(messageInfo)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get off line message encode is err:%v", receiverID, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}
		_, err = pipe.HSet(context.Background(), userOfflineMessage, strconv.FormatInt(senderID, 10), info).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis set user[%v] off line message encode is err:%v", receiverID, senderID, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		return nil, nil
	})

	return err
}
