// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package api_gateway

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

var serverName string = "gateway"

type GatewayClientInterface interface {
	NotifyClient(context.Context, *NotifyClientReq) (*NotifyClientRes, error)
}

func NewGatewayClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string) GatewayClientInterface {
	c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)

	return &GatewayClient{
		c: c,
	}
}

type GatewayClient struct {
	c *client.ClientManager
}

func (c *GatewayClient) NotifyClient(ctx context.Context,
	in *NotifyClientReq) (*NotifyClientRes, error) {
	out := new(NotifyClientRes)
	err := c.c.Call(ctx, "NotifyClient", in, out)
	return out, err
}

type GatewayServiceInterface interface {
	NotifyClient(context.Context, *NotifyClientReq, *NotifyClientRes) error
}

func RegisterGatewayService(s *service.ServerManage, hdlr GatewayServiceInterface) error {
	return s.RegisterOneService(serverName, hdlr)
}

func NewGatewayServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler GatewayServiceInterface, etcdBasePath string) (*service.ServerManage, error) {
	s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath)
	if err != nil {
		return nil, err
	}

	err = RegisterGatewayService(s, handler)
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
