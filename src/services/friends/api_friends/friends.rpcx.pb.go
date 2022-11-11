// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friends.proto

package api_friends

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

var serverName string = "friends"

var callSingleMethodFunc func()

var FriendsClientInstance FriendsClientInterface
var FriendsClientOnce = new(sync.Once)

func newFriendsClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) FriendsClientInterface {
	if !isLocal {
		c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)
		return &FriendsRpcxClient{
			c: c,
		}
	} else {
		return &FriendsLocalClient{}
	}
}

func SingleNewFriendsClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) {
	callSingleMethodFunc = func() {
		c := newFriendsClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		FriendsClientInstance = c
	}
}

// 外部调用函数

func GetFriendsList(ctx context.Context,
	in *GetFriendsListReq) (*GetFriendsListRes, error) {

	if callSingleMethodFunc != nil {
		FriendsClientOnce.Do(callSingleMethodFunc)
	}

	out := new(GetFriendsListRes)
	out, err := FriendsClientInstance.GetFriendsList(ctx, in)
	return out, err
}

func ApplyFriends(ctx context.Context,
	in *ApplyFriendsReq) (*ApplyFriendsRes, error) {

	if callSingleMethodFunc != nil {
		FriendsClientOnce.Do(callSingleMethodFunc)
	}

	out := new(ApplyFriendsRes)
	out, err := FriendsClientInstance.ApplyFriends(ctx, in)
	return out, err
}

func DelFriends(ctx context.Context,
	in *DelFriendsReq) (*DelFriendsRes, error) {

	if callSingleMethodFunc != nil {
		FriendsClientOnce.Do(callSingleMethodFunc)
	}

	out := new(DelFriendsRes)
	out, err := FriendsClientInstance.DelFriends(ctx, in)
	return out, err
}

type FriendsClientInterface interface {
	GetFriendsList(context.Context, *GetFriendsListReq) (*GetFriendsListRes, error)
	ApplyFriends(context.Context, *ApplyFriendsReq) (*ApplyFriendsRes, error)
	DelFriends(context.Context, *DelFriendsReq) (*DelFriendsRes, error)
}

// rpcx客户端
type FriendsRpcxClient struct {
	c *client.ClientManager
}

func (c *FriendsRpcxClient) GetFriendsList(ctx context.Context,
	in *GetFriendsListReq) (*GetFriendsListRes, error) {
	out := new(GetFriendsListRes)
	err := c.c.Call(ctx, "GetFriendsList", in, out)
	return out, err
}

func (c *FriendsRpcxClient) ApplyFriends(ctx context.Context,
	in *ApplyFriendsReq) (*ApplyFriendsRes, error) {
	out := new(ApplyFriendsRes)
	err := c.c.Call(ctx, "ApplyFriends", in, out)
	return out, err
}

func (c *FriendsRpcxClient) DelFriends(ctx context.Context,
	in *DelFriendsReq) (*DelFriendsRes, error) {
	out := new(DelFriendsRes)
	err := c.c.Call(ctx, "DelFriends", in, out)
	return out, err
}

// 本地调用客户端
type FriendsLocalClient struct {
}

func (c *FriendsLocalClient) GetFriendsList(ctx context.Context,
	in *GetFriendsListReq) (*GetFriendsListRes, error) {
	out := new(GetFriendsListRes)
	err := FriendsServiceLocal.GetFriendsList(ctx, in, out)
	return out, err
}

func (c *FriendsLocalClient) ApplyFriends(ctx context.Context,
	in *ApplyFriendsReq) (*ApplyFriendsRes, error) {
	out := new(ApplyFriendsRes)
	err := FriendsServiceLocal.ApplyFriends(ctx, in, out)
	return out, err
}

func (c *FriendsLocalClient) DelFriends(ctx context.Context,
	in *DelFriendsReq) (*DelFriendsRes, error) {
	out := new(DelFriendsRes)
	err := FriendsServiceLocal.DelFriends(ctx, in, out)
	return out, err
}

type FriendsServiceInterface interface {
	GetFriendsList(context.Context, *GetFriendsListReq, *GetFriendsListRes) error
	ApplyFriends(context.Context, *ApplyFriendsReq, *ApplyFriendsRes) error
	DelFriends(context.Context, *DelFriendsReq, *DelFriendsRes) error
}

var FriendsServiceLocal FriendsServiceInterface

func RegisterFriendsService(s *service.ServerManage, hdlr FriendsServiceInterface) error {
	return s.RegisterOneService(serverName, hdlr)
}

func NewFriendsServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler FriendsServiceInterface, etcdBasePath string, isLocal bool) (*service.ServerManage, error) {
	if !isLocal {
		s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath)
		if err != nil {
			return nil, err
		}

		err = RegisterFriendsService(s, handler)
		if err != nil {
			return nil, err
		}

		go func() {
			err = s.Run(listenAddr)
			if err != nil {
				panic(fmt.Errorf("listen(%v) error(%v)", listenAddr, err))
			}
		}()
		return s, nil
	}

	// 本地调用的时候使用
	FriendsServiceLocal = handler
	return nil, nil
}
