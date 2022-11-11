package account

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

func init() {
	seralize.RegisterHandler(im_home_proto.LoginReq{}, im_home_proto.LoginRes{}, Login)
	seralize.RegisterHandler(im_home_proto.RegisterReq{}, im_home_proto.RegisterRes{}, Register)

}

func Login(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.LoginReq)
	resMsg := res.(*im_home_proto.LoginRes)

	resRpc, err := api_account.Login(ctx, &api_account.LoginReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})
	if err != nil {
		im_log.Error("login user[%v] is err:%v", reqMsg.UserID, err)
		return resRpc.ErrCode
	}

	resMsg.UserID = resRpc.Data.UserID
	resMsg.UserName = resRpc.Data.UserName
	return 0
}

func Register(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.RegisterReq)
	resMsg := res.(*im_home_proto.RegisterRes)

	resRpc, err := api_account.Register(ctx, &api_account.RegisterReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})

	if err != nil {
		im_log.Error("register user[%v] is err:%v", reqMsg.UserName, err)
		return resRpc.ErrCode
	}

	resMsg.UserID = resRpc.Data.UserID
	resMsg.UserName = resRpc.Data.UserName
	return 0
}
