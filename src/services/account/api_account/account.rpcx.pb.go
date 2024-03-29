// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package api_account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	_ "github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	math "math"
)

import (
	context "context"
	client "github.com/mimis-s/golang_tools/rpcx/client"
	service "github.com/mimis-s/golang_tools/rpcx/service"
	"sync"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var serverName string = "account"

var callSingleMethodFunc func()

var AccountClientInstance AccountClientInterface
var AccountClientOnce = new(sync.Once)

func newAccountClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) AccountClientInterface {
	if !isLocal {
		c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)
		return &AccountRpcxClient{
			c: c,
		}
	} else {
		return &AccountLocalClient{}
	}
}

func SingleNewAccountClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) {
	callSingleMethodFunc = func() {
		c := newAccountClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		AccountClientInstance = c
	}
}

// 外部调用函数

func Login(ctx context.Context,
	in *LoginReq) (*LoginRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(LoginRes)
	out, err := AccountClientInstance.Login(ctx, in)
	return out, err
}

func Logout(ctx context.Context,
	in *LogoutReq) (*LogoutRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(LogoutRes)
	out, err := AccountClientInstance.Logout(ctx, in)
	return out, err
}

func Register(ctx context.Context,
	in *RegisterReq) (*RegisterRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(RegisterRes)
	out, err := AccountClientInstance.Register(ctx, in)
	return out, err
}

func GetUserInfo(ctx context.Context,
	in *GetUserInfoReq) (*GetUserInfoRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(GetUserInfoRes)
	out, err := AccountClientInstance.GetUserInfo(ctx, in)
	return out, err
}

func GetUserInfoService(ctx context.Context,
	in *GetUserInfoServiceReq) (*GetUserInfoServiceRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(GetUserInfoServiceRes)
	out, err := AccountClientInstance.GetUserInfoService(ctx, in)
	return out, err
}

func GetUsersInfoService(ctx context.Context,
	in *GetUsersInfoServiceReq) (*GetUsersInfoServiceRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(GetUsersInfoServiceRes)
	out, err := AccountClientInstance.GetUsersInfoService(ctx, in)
	return out, err
}

func ModifyUserInfo(ctx context.Context,
	in *ModifyUserInfoReq) (*ModifyUserInfoRes, error) {

	if callSingleMethodFunc != nil {
		AccountClientOnce.Do(callSingleMethodFunc)
	}

	out := new(ModifyUserInfoRes)
	out, err := AccountClientInstance.ModifyUserInfo(ctx, in)
	return out, err
}

type AccountClientInterface interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Logout(context.Context, *LogoutReq) (*LogoutRes, error)
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRes, error)
	GetUserInfoService(context.Context, *GetUserInfoServiceReq) (*GetUserInfoServiceRes, error)
	GetUsersInfoService(context.Context, *GetUsersInfoServiceReq) (*GetUsersInfoServiceRes, error)
	ModifyUserInfo(context.Context, *ModifyUserInfoReq) (*ModifyUserInfoRes, error)
}

// rpcx客户端
type AccountRpcxClient struct {
	c *client.ClientManager
}

func (c *AccountRpcxClient) Login(ctx context.Context,
	in *LoginReq) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.c.Call(ctx, "Login", in, out)
	return out, err
}

func (c *AccountRpcxClient) Logout(ctx context.Context,
	in *LogoutReq) (*LogoutRes, error) {
	out := new(LogoutRes)
	err := c.c.Call(ctx, "Logout", in, out)
	return out, err
}

func (c *AccountRpcxClient) Register(ctx context.Context,
	in *RegisterReq) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.c.Call(ctx, "Register", in, out)
	return out, err
}

func (c *AccountRpcxClient) GetUserInfo(ctx context.Context,
	in *GetUserInfoReq) (*GetUserInfoRes, error) {
	out := new(GetUserInfoRes)
	err := c.c.Call(ctx, "GetUserInfo", in, out)
	return out, err
}

func (c *AccountRpcxClient) GetUserInfoService(ctx context.Context,
	in *GetUserInfoServiceReq) (*GetUserInfoServiceRes, error) {
	out := new(GetUserInfoServiceRes)
	err := c.c.Call(ctx, "GetUserInfoService", in, out)
	return out, err
}

func (c *AccountRpcxClient) GetUsersInfoService(ctx context.Context,
	in *GetUsersInfoServiceReq) (*GetUsersInfoServiceRes, error) {
	out := new(GetUsersInfoServiceRes)
	err := c.c.Call(ctx, "GetUsersInfoService", in, out)
	return out, err
}

func (c *AccountRpcxClient) ModifyUserInfo(ctx context.Context,
	in *ModifyUserInfoReq) (*ModifyUserInfoRes, error) {
	out := new(ModifyUserInfoRes)
	err := c.c.Call(ctx, "ModifyUserInfo", in, out)
	return out, err
}

// 本地调用客户端
type AccountLocalClient struct {
}

func (c *AccountLocalClient) Login(ctx context.Context,
	in *LoginReq) (*LoginRes, error) {
	out := new(LoginRes)
	err := AccountServiceLocal.Login(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) Logout(ctx context.Context,
	in *LogoutReq) (*LogoutRes, error) {
	out := new(LogoutRes)
	err := AccountServiceLocal.Logout(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) Register(ctx context.Context,
	in *RegisterReq) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := AccountServiceLocal.Register(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) GetUserInfo(ctx context.Context,
	in *GetUserInfoReq) (*GetUserInfoRes, error) {
	out := new(GetUserInfoRes)
	err := AccountServiceLocal.GetUserInfo(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) GetUserInfoService(ctx context.Context,
	in *GetUserInfoServiceReq) (*GetUserInfoServiceRes, error) {
	out := new(GetUserInfoServiceRes)
	err := AccountServiceLocal.GetUserInfoService(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) GetUsersInfoService(ctx context.Context,
	in *GetUsersInfoServiceReq) (*GetUsersInfoServiceRes, error) {
	out := new(GetUsersInfoServiceRes)
	err := AccountServiceLocal.GetUsersInfoService(ctx, in, out)
	return out, err
}

func (c *AccountLocalClient) ModifyUserInfo(ctx context.Context,
	in *ModifyUserInfoReq) (*ModifyUserInfoRes, error) {
	out := new(ModifyUserInfoRes)
	err := AccountServiceLocal.ModifyUserInfo(ctx, in, out)
	return out, err
}

type AccountServiceInterface interface {
	Login(context.Context, *LoginReq, *LoginRes) error
	Logout(context.Context, *LogoutReq, *LogoutRes) error
	Register(context.Context, *RegisterReq, *RegisterRes) error
	GetUserInfo(context.Context, *GetUserInfoReq, *GetUserInfoRes) error
	GetUserInfoService(context.Context, *GetUserInfoServiceReq, *GetUserInfoServiceRes) error
	GetUsersInfoService(context.Context, *GetUsersInfoServiceReq, *GetUsersInfoServiceRes) error
	ModifyUserInfo(context.Context, *ModifyUserInfoReq, *ModifyUserInfoRes) error
}

var AccountServiceLocal AccountServiceInterface

func RegisterAccountService(s *service.ServerManage, hdlr AccountServiceInterface) error {
	// 本地调用的时候使用(rpc本地客户端对应调用本地服务器)
	AccountServiceLocal = hdlr
	return s.RegisterOneService(serverName, hdlr)
}

func NewAccountServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler AccountServiceInterface, etcdBasePath string, isLocal bool) (*service.ServerManage, error) {
	if !isLocal {
		s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath, listenAddr)
		if err != nil {
			return nil, err
		}

		err = RegisterAccountService(s, handler)
		if err != nil {
			return nil, err
		}

		return s, nil
	}

	// 本地调用的时候使用
	AccountServiceLocal = handler
	return nil, nil
}
