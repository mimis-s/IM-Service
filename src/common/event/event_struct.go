package event

import "github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"

// mq事件结构

type UserLogin struct {
	UserInfo *im_home_proto.UserInfo
}

type SingleMessage struct {
	UserInfo *im_home_proto.ClientOnlineInfo
	Message  *im_home_proto.ChatMessage
}

type ApplyFriend struct {
	UserInfo *im_home_proto.ClientOnlineInfo
	Message  *im_home_proto.ApplyFriendsToReceiver
}

type AgreeApplyFriend struct {
	UserInfo *im_home_proto.ClientOnlineInfo
	Message  *im_home_proto.AgreeApplyFriendsToReceiver
}
