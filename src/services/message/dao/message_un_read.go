package dao

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mimis-s/IM-Service/src/common/im_log"
)

/*
	存储数据的层级关系:前缀key + 用户ID -> 好友ID -> 消息(转化为base64编码存储)；消息
*/

const (
	userUnReadMessagePrefix = "user.message.unRead.line.sum."
)

// 获取用户所有离线消息个数(返回好友对应消息数量)
func (d *Dao) GetUserAllUnReadMessage(userID int64) (map[int64]int, error) {
	userUnReadMessage := userUnReadMessagePrefix + strconv.FormatInt(userID, 10)

	allMessageData, err := d.cache.Client.HGetAll(context.Background(), userUnReadMessage).Result()
	if err != nil {
		errStr := fmt.Sprintf("user[%v] redis get all off line message is err:%v", userID, err)
		im_log.Warn(errStr)
		return nil, fmt.Errorf(errStr)
	}

	mapFriendMessageSum := make(map[int64]int)
	for strFriend, messageSum := range allMessageData {

		friendsID, err := strconv.ParseInt(strFriend, 10, 64)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode friend[%v] off line message is err:%v", userID, strFriend, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		data, err := strconv.Atoi(messageSum)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode friend[%v] off line message is err:%v", userID, strFriend, err)
			im_log.Warn(errStr)
			return nil, fmt.Errorf(errStr)
		}

		if friendsID == 0 || data <= 0 {
			continue
		}

		mapFriendMessageSum[friendsID] = data
	}

	return mapFriendMessageSum, nil
}

// 获取用户未读消息
func (d *Dao) GetUserUnReadMessage(userID int64, friendID int64) (int, error) {
	userUnReadMessage := userUnReadMessagePrefix + strconv.FormatInt(userID, 10)

	exists, err := d.cache.Client.HExists(context.Background(), userUnReadMessage, strconv.FormatInt(friendID, 10)).Result()
	if err != nil {
		im_log.Warn("user[%v] friend[%v] redis get off line message Exists is err:%v", userID, friendID, err)
		return 0, err
	}

	if exists {

		messageSum, err := d.cache.Client.HGet(context.Background(), userUnReadMessage, strconv.FormatInt(friendID, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get friend[%v] off line message is err:%v", userID, friendID, err)
			im_log.Warn(errStr)
			return 0, fmt.Errorf(errStr)
		}

		data, err := strconv.Atoi(messageSum)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode friend[%v] off line message is err:%v", userID, friendID, err)
			im_log.Warn(errStr)
			return 0, fmt.Errorf(errStr)
		}

		return data, nil
	}
	return 0, nil
}

// 添加用户离线消息
func (d *Dao) AddUserOneUnReadMessage(sender, receiver int64) error {

	userUnReadMessage := userUnReadMessagePrefix + strconv.FormatInt(receiver, 10)

	exists, err := d.cache.Client.HExists(context.Background(), userUnReadMessage, strconv.FormatInt(sender, 10)).Result()
	if err != nil {
		im_log.Warn("user[%v] sender[%v] redis get off line message Exists is err:%v", receiver, sender, err)
		return err
	}

	UnReadSum := 0

	if exists {
		messageSum, err := d.cache.Client.HGet(context.Background(), userUnReadMessage, strconv.FormatInt(sender, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis get off line message is err:%v", receiver, err)
			im_log.Warn(errStr)
			return fmt.Errorf(errStr)
		}

		data, err := strconv.Atoi(messageSum)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis decode friend[%v] off line message is err:%v", sender, receiver, err)
			im_log.Warn(errStr)
			return fmt.Errorf(errStr)
		}
		UnReadSum = data
	}

	UnReadSum++

	_, err = d.cache.Client.HSet(context.Background(), userUnReadMessage, strconv.FormatInt(sender, 10), strconv.Itoa(UnReadSum)).Result()
	if err != nil {
		errStr := fmt.Sprintf("user[%v] redis set user[%v] off line message encode is err:%v", receiver, sender, err)
		im_log.Warn(errStr)
		return fmt.Errorf(errStr)
	}

	return nil
}

// 把离线消息删除
func (d *Dao) DelUserOneUnReadMessage(senderID, receiverID int64) error {
	userUnReadMessage := userUnReadMessagePrefix + strconv.FormatInt(senderID, 10)

	exists, err := d.cache.Client.HExists(context.Background(), userUnReadMessage, strconv.FormatInt(receiverID, 10)).Result()
	if err != nil {
		im_log.Warn("user[%v] sender[%v] redis get off line message Exists is err:%v", receiverID, senderID, err)
		return err
	}

	if !exists {
		im_log.Info("sender[%v] receiver[%v] not off line message need del", senderID, receiverID)
		return nil
	} else {
		_, err = d.cache.Client.HDel(context.Background(), userUnReadMessage, strconv.FormatInt(receiverID, 10)).Result()
		if err != nil {
			errStr := fmt.Sprintf("user[%v] redis set user[%v] off line message encode is err:%v", receiverID, senderID, err)
			im_log.Warn(errStr)
			return fmt.Errorf(errStr)
		}
	}

	return nil
}
