package account

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/golang_tools/zlog"
)

func init() {
	seralize.RegisterHandler(im_home_proto.LoginReq{}, im_home_proto.LoginRes{}, Login)
	seralize.RegisterHandler(im_home_proto.LogoutReq{}, im_home_proto.LogoutRes{}, Logout)
	seralize.RegisterHandler(im_home_proto.RegisterReq{}, im_home_proto.RegisterRes{}, Register)
	seralize.RegisterHandler(im_home_proto.GetUserInfoReq{}, im_home_proto.GetUserInfoRes{}, GetUserInfo)
	seralize.RegisterHandler(im_home_proto.ModifyUserInfoReq{}, im_home_proto.ModifyUserInfoRes{}, ModifyUserInfo)

}

func Login(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.LoginReq)
	resMsg := res.(*im_home_proto.LoginRes)

	resRpc, err := api_account.Login(ctx, &api_account.LoginReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})
	if err != nil {
		zlog.Error("login user[%v] is err:%v", reqMsg.UserID, err)
		return resRpc.ErrCode
	}

	resMsg.Info = resRpc.Data.Info
	return 0
}

func Logout(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.LogoutReq)

	resRpc, err := api_account.Logout(ctx, &api_account.LogoutReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})
	if err != nil {
		zlog.Error("logout user[%v] is err:%v", clientInfo.Client.UserID, err)
		return resRpc.ErrCode
	}

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
		zlog.Error("register user[%v] is err:%v", reqMsg.UserName, err)
		return resRpc.ErrCode
	}

	resMsg.UserID = resRpc.Data.UserID
	resMsg.UserName = resRpc.Data.UserName
	return 0
}

func GetUserInfo(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.GetUserInfoReq)
	resMsg := res.(*im_home_proto.GetUserInfoRes)

	resRpc, err := api_account.GetUserInfo(ctx, &api_account.GetUserInfoReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})

	if err != nil {
		zlog.Error("user[%v] get user[%v] info is err:%v", clientInfo.Client.UserID, reqMsg.UserID, err)
		return resRpc.ErrCode
	}

	resMsg.Data = resRpc.Data.Data
	resMsg.Relation = resRpc.Data.Relation
	return 0
}

func ModifyUserInfo(ctx context.Context, clientInfo *api_home.ClientRequestHandleReq, req, res seralize.Message) im_error_proto.ErrCode {
	reqMsg := req.(*im_home_proto.ModifyUserInfoReq)
	resMsg := res.(*im_home_proto.ModifyUserInfoRes)

	resRpc, err := api_account.ModifyUserInfo(ctx, &api_account.ModifyUserInfoReq{
		ClientInfo: clientInfo.Client,
		Data:       reqMsg,
	})

	if err != nil {
		zlog.Error("user[%v] modify user[%v] info is err:%v", clientInfo.Client.UserID, reqMsg.Data.UserID, err)
		return resRpc.ErrCode
	}

	resMsg.Data = resRpc.Data.Data
	return 0
}
