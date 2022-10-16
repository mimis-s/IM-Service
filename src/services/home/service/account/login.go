package account

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_home_proto.LoginReq{}, im_home_proto.LoginRes{}, Login)
}

func Login(ctx context.Context, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.LoginReq)
	resMsg := res.(*im_home_proto.LoginRes)
	resMsg.UserID = reqMsg.UserID
	resMsg.UserName = reqMsg.Password
	return 0
}
