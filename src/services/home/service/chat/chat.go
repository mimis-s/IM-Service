package chat

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_home_proto.ChatSingleReq{}, im_home_proto.ChatSingleRes{}, ChatSingle)
}

func ChatSingle(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.ChatSingleReq)
	resMsg := res.(*im_home_proto.ChatSingleRes)
	reqRpc := &api_chat.ChatSingleReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg.Data,
	}
	resRpc, err := api_chat.ChatSingle(context.Background(), reqRpc)
	if err != nil {
		fmt.Printf("chat single rpc is err:%v", err)
		return im_error_proto.ErrCode_common_unexpected_err
	}
	resMsg.Data = resRpc.Data
	return 0
}
