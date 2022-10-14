package account

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_main_proto"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_main_proto.LoginReq{}, im_main_proto.LoginRes{}, Login)
}

func Login(ctx context.Context, req, res seralize.Message) im_error_proto.ErrCode {
	return 0
}
