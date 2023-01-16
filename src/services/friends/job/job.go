package job

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/friends/service"
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
	friendsStatusList, errCode, err := j.s.GetFriendsStatusList(context.Background(), userLogin.UserInfo)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] login ok, but get friends staus list is errCode[%v] err:%v",
			userLogin.UserInfo.UserID, errCode, err)
		im_log.Warn(errStr)
		return fmt.Errorf(errStr)
	}

	relay_api.NotifyUser(userLogin.UserInfo.UserID, friendsStatusList)

	return nil
}
