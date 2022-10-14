// Code generated by protoc-gen-go. DO NOT EDIT.
// source: home.proto

package api_lobby

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	math "math"
)

import (
	context "context"
	client "github.com/mimis-s/golang_tools/rpcx/client"
	service "github.com/mimis-s/golang_tools/rpcx/service"
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

var serverName string = "main"

type MainClientInterface interface {
	ClientRequestHandleProto(context.Context, *ClientRequestHandleReq) (*ClientRequestHandleRes, error)
	ClientRequestHandleJson(context.Context, *ClientRequestHandleReq) (*ClientRequestHandleRes, error)
}

func NewMainClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string) MainClientInterface {
	c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)

	return &MainClient{
		c: c,
	}
}

type MainClient struct {
	c *client.ClientManager
}

func (c *MainClient) ClientRequestHandleProto(ctx context.Context,
	in *ClientRequestHandleReq) (*ClientRequestHandleRes, error) {
	out := new(ClientRequestHandleRes)
	err := c.c.Call(ctx, "ClientRequestHandleProto", in, out)
	return out, err
}

func (c *MainClient) ClientRequestHandleJson(ctx context.Context,
	in *ClientRequestHandleReq) (*ClientRequestHandleRes, error) {
	out := new(ClientRequestHandleRes)
	err := c.c.Call(ctx, "ClientRequestHandleJson", in, out)
	return out, err
}

type MainServiceInterface interface {
	ClientRequestHandleProto(context.Context, *ClientRequestHandleReq, *ClientRequestHandleRes) error
	ClientRequestHandleJson(context.Context, *ClientRequestHandleReq, *ClientRequestHandleRes) error
}

func RegisterMainService(s *service.ServerManage, hdlr MainServiceInterface) error {
	return s.RegisterOneService(serverName, hdlr)
}

func NewMainServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler MainServiceInterface, etcdBasePath string) (*service.ServerManage, error) {
	s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath)
	if err != nil {
		return nil, err
	}

	err = RegisterMainService(s, handler)
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