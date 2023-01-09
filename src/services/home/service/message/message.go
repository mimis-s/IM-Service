package message

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
)

func init() {
	seralize.RegisterHandler(im_home_proto.GetSingleChatHistoryReq{}, im_home_proto.GetSingleChatHistoryRes{}, GetSingleChatHistory)
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
		im_log.Error("user[%v] get friend[%v] Single Chat History is err:%v", clientInfo.Client.UserID,
			reqMsg.FriendID, err)
		return resRpc.ErrCode
	}
	resMsg.Data = resRpc.Data.Data
	return 0
}
