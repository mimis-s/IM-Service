package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/golang_tools/lib"
)

/*
	历史记录表userfrist是选择发送者和接受者id最小的来查询
*/

// 存储历史记录
func (d *Dao) AddHistoryMessage(senderID, receiverID int64, chatMessage *im_home_proto.ChatMessage) error {

	userFrist := lib.MinInt64(senderID, receiverID)
	userSecond := lib.MaxInt64(senderID, receiverID)

	historyMessage := &dbmodel.HistoryMessage{
		UserIdFrist:  userFrist,
		UserIdSecond: userSecond,
		MessageId:    int64(chatMessage.MessageID),
		MessageData: dbmodel.TBJsonField_HistoryMessage{
			HistoryData: chatMessage,
		},
	}

	_, err := d.Session.HistoryMessage.InsertOne(historyMessage)
	if err != nil {
		errStr := fmt.Sprintf("insert history message sender[%v] receiver[%v] messageID[%v] is err:%v",
			senderID, receiverID, chatMessage.MessageID, err)
		im_log.Warn(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}

// 更新历史记录(更新状态)
func (d *Dao) UpdateHistoryMessage(senderID, receiverID, messageID int64, historyMessage *dbmodel.HistoryMessage) error {

	userFrist := lib.MinInt64(senderID, receiverID)
	userSecond := lib.MaxInt64(senderID, receiverID)

	_, err := d.Session.HistoryMessage.Where("user_id_frist=? and user_id_second=? and message_id=?",
		userFrist, userSecond, messageID).Cols(dbmodel.THistoryMessage.MessageData).Update(historyMessage)
	if err != nil {
		errStr := fmt.Sprintf("insert history message sender[%v] receiver[%v] messageID[%v] is err:%v",
			senderID, receiverID, messageID, err)
		im_log.Warn(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
