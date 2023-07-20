package message

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/golang_tools/zlog"
)

func init() {
	seralize.RegisterHandler(im_home_proto.GetSingleChatHistoryReq{}, im_home_proto.GetSingleChatHistoryRes{}, GetSingleChatHistory)
	seralize.RegisterHandler(im_home_proto.UnReadMessageReq{}, im_home_proto.UnReadMessageRes{}, UnReadMessage)
	seralize.RegisterHandler(im_home_proto.DownLoadFileMessageReq{}, im_home_proto.DownLoadFileMessageRes{}, DownLoadFileMessage)
}

func GetSingleChatHistory(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.GetSingleChatHistoryReq)
	resMsg := res.(*im_home_proto.GetSingleChatHistoryRes)

	reqRpc := &api_message.GetSingleChatHistoryReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_message.GetSingleChatHistory(context.Background(), reqRpc)
	if err != nil {
		zlog.Error("user[%v] get friend[%v] Single Chat History is err:%v", clientInfo.Client.UserID,
			reqMsg.FriendID, err)
		return resRpc.ErrCode
	}
	resMsg.Data = resRpc.Data.Data
	return 0
}

func UnReadMessage(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.UnReadMessageReq)
	resMsg := res.(*im_home_proto.UnReadMessageRes)

	reqRpc := &api_message.UnReadMessageReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_message.UnReadMessage(context.Background(), reqRpc)
	if err != nil {
		zlog.Error("user[%v] Read friend[%v] un read Message Single Chat is err:%v", clientInfo.Client.UserID,
			reqMsg.FriendID, err)
		return resRpc.ErrCode
	}
	resMsg.FriendID = resRpc.Data.FriendID
	return 0
}

func DownLoadFileMessage(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.DownLoadFileMessageReq)
	resMsg := res.(*im_home_proto.DownLoadFileMessageRes)

	reqRpc := &api_message.DownLoadFileMessageReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	}
	resRpc, err := api_message.DownLoadFileMessage(context.Background(), reqRpc)
	if err != nil {
		zlog.Error("user[%v] down load friend[%v] file Message Single Chat is err:%v", clientInfo.Client.UserID,
			reqMsg.FriendID, err)
		return resRpc.ErrCode
	}

	resMsg.FriendID = resRpc.Data.FriendID
	resMsg.MessageID = resRpc.Data.MessageID
	resMsg.FileIndex = resRpc.Data.FileIndex
	resMsg.FileData = resRpc.Data.FileData
	resMsg.MessageFileType = resRpc.Data.MessageFileType

	return 0
}
