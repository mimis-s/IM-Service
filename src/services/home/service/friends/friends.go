package friends

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_home_proto.GetFriendsListReq{}, im_home_proto.GetFriendsListRes{}, GetFriendsList)
	seralize.RegisterHandler(im_home_proto.ApplyFriendsReq{}, im_home_proto.ApplyFriendsRes{}, ApplyFriends)
	seralize.RegisterHandler(im_home_proto.AgreeFriendApplyReq{}, im_home_proto.AgreeFriendApplyRes{}, AgreeFriendApply)
	seralize.RegisterHandler(im_home_proto.DelFriendsReq{}, im_home_proto.DelFriendsRes{}, DelFriends)
}

func GetFriendsList(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.GetFriendsListReq)
	resMsg := res.(*im_home_proto.GetFriendsListRes)
	reqRpc := &api_friends.GetFriendsListReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_friends.GetFriendsList(context.Background(), reqRpc)
	if err != nil {
		im_log.Error("get friends list rpc is err:%v", err)
		return resRpc.ErrCode
	}
	resMsg.List = resRpc.Data.List
	return 0
}

func ApplyFriends(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.ApplyFriendsReq)
	resMsg := res.(*im_home_proto.ApplyFriendsRes)
	reqRpc := &api_friends.ApplyFriendsReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_friends.ApplyFriends(context.Background(), reqRpc)
	if err != nil {
		im_log.Error("apply friends[%v] rpc is err:%v", reqMsg.ApplyFriendsID, err)
		return resRpc.ErrCode
	}
	resMsg.FriendInfo = resRpc.Data.FriendInfo
	return 0
}

func AgreeFriendApply(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.AgreeFriendApplyReq)
	resMsg := res.(*im_home_proto.AgreeFriendApplyRes)
	reqRpc := &api_friends.AgreeFriendApplyReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_friends.AgreeFriendApply(context.Background(), reqRpc)
	if err != nil {
		im_log.Error("agree friends[%v] apply rpc is err:%v", reqMsg.FriendsID, err)
		return resRpc.ErrCode
	}
	resMsg.FriendsID = resRpc.Data.FriendsID
	return 0
}

func DelFriends(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.DelFriendsReq)
	resMsg := res.(*im_home_proto.DelFriendsRes)
	reqRpc := &api_friends.DelFriendsReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_friends.DelFriends(context.Background(), reqRpc)
	if err != nil {
		im_log.Error("del friends[%v] rpc is err:%v", reqMsg.FriendsID, err)
		return resRpc.ErrCode
	}
	resMsg.FriendsID = resRpc.Data.FriendsID
	return 0
}
