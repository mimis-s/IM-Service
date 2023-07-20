package job

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/message/service"
	"github.com/mimis-s/IM-Service/src/services/relay/relay_api"
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
	"github.com/mimis-s/golang_tools/zlog"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service) *Job {
	url := boot_config.BootConfigData.MQ.Url
	durable := boot_config.BootConfigData.MQ.Durable

	j := &Job{s}

	event.InitConsumers(url, durable,
		[]*rabbitmq.ConsumersQueue{
			{event.Event_UserLogin, j.userLoginOk},
		},
	)

	return j
}

func (j *Job) userLoginOk(payload interface{}) error {
	userLogin := payload.(*event.UserLogin)
	// 登录成功下发未读消息
	userChatMessage, err := j.s.Dao.GetUserAllUnReadMessage(userLogin.UserInfo.UserID)
	if err != nil {
		zlog.Warn("user[%v] login ok, but get off line message is err:%v", userLogin.UserInfo.UserID, err)
		return err
	}

	if len(userChatMessage) == 0 {
		return nil
	}

	notifyUserMessage := &im_home_proto.NotifyUserMessage{
		UnReadSingleChat: make([]*im_home_proto.NotifyUnReadMessage, 0),
	}
	usersInfoReq := &api_account.GetUsersInfoServiceReq{
		ClientInfo: &im_home_proto.ClientOnlineInfo{
			UserID: userLogin.UserInfo.UserID,
		},
		UserIDs: make([]int64, 0, len(userChatMessage)),
	}
	for senderID, _ := range userChatMessage {
		usersInfoReq.UserIDs = append(usersInfoReq.UserIDs, senderID)
	}

	usersInfoRes, err := api_account.GetUsersInfoService(context.Background(), usersInfoReq)
	if err != nil {
		zlog.Warn("user[%v] login ok, but friends[%v] info is err:%v", userLogin.UserInfo.UserID, usersInfoReq.UserIDs, err)
		return err
	}

	for _, friend := range usersInfoRes.Datas {
		notifyUnReadMessage := &im_home_proto.NotifyUnReadMessage{
			User:             friend,
			UnReadMessageSum: int32(userChatMessage[friend.UserID]),
		}

		notifyUserMessage.UnReadSingleChat = append(notifyUserMessage.UnReadSingleChat, notifyUnReadMessage)
	}

	relay_api.NotifyUser(userLogin.UserInfo.UserID, notifyUserMessage)

	return nil
}
