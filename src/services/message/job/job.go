package job

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/message/service"
	"github.com/mimis-s/IM-Service/src/services/relay/relay_api"
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service, configOptions *boot_config.ConfigOptions) *Job {
	url := configOptions.BootConfigFile.MQ.Url
	durable := configOptions.BootConfigFile.MQ.Durable

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
	// 登录成功下发离线消息
	userChatMessage, err := j.s.Dao.GetUserAllOfflineMessage(userLogin.UserInfo.UserID)
	if err != nil {
		im_log.Warn("user[%v] login ok, but get off line message is err:%v", userLogin.UserInfo.UserID, err)
		return err
	}

	notifyUserMessage := im_home_proto.NotifyUserMessage{
		OfflineSingleChat: make([]*im_home_proto.NotifyOfflineMessage, 0),
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
		im_log.Warn("user[%v] login ok, but friends[%v] info is err:%v", userLogin.UserInfo.UserID, err)
		return err
	}

	for _, friend := range usersInfoRes.Datas {
		notifyOfflineMessage := &im_home_proto.NotifyOfflineMessage{
			User: friend,
			Data: make([]*im_home_proto.ChatMessage, 0),
		}
		for senderID, chatMessage := range userChatMessage {
			if friend.UserID == senderID {
				notifyOfflineMessage.Data = append(notifyOfflineMessage.Data, chatMessage...)
				break
			}
		}

		notifyUserMessage.OfflineSingleChat = append(notifyUserMessage.OfflineSingleChat, notifyOfflineMessage)
	}

	relay_api.NotifyUser(userLogin.UserInfo.UserID, notifyUserMessage)

	return nil
}
