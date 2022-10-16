package chat

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_home_proto.ChatSingleReq{}, im_home_proto.ChatSingleRes{}, ChatSingle)
}

func ChatSingle(ctx context.Context, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.ChatSingleReq)
	resMsg := res.(*im_home_proto.ChatSingleRes)
	resMsg.TestStr = reqMsg.TestStr
	return 0
}
