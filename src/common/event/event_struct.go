package event

import "github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"

// mq事件结构
type SingleMessage struct {
	UserInfo *im_home_proto.ClientOnlineInfo
	Message  *im_home_proto.ChatMessage
}
