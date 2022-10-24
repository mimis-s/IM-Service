package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/golang_tools/net"
	"github.com/mimis-s/golang_tools/net/clientConn"
)

// 服务器处理客户端消息的回调函数(现在的简单框架足够了, 但是后面会集成到session里面)
func (s *Service) HandlerHttpRespone(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
	if reqClient.Tag == -1 {
		// 心跳包
		fmt.Printf("client send heartCheack\n")
		return &clientConn.ClientMsg{
			Tag: -1,
		}, nil
	}

	// 先不记录client连接
	req := &api_home.ClientRequestHandleReq{
		MsgID:   uint32(reqClient.Tag),
		Payload: reqClient.Msg,
		Client: &im_home_proto.ClientOnlineInfo{
			UserID:   1,
			UserName: "test",
		},
	}
	res := &api_home.ClientRequestHandleRes{}

	res, err := api_home.ClientRequestHandleJson(context.Background(), req)
	if err != nil {
		errStr := fmt.Sprintf("Client Request Handle Json[%v] is err:%v", req, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}

	return &clientConn.ClientMsg{
		Tag: int(res.MsgID),
		Msg: res.Payload,
	}, nil
}

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao *dao.Dao
}

func Init(addr, webAddr string) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao: d,
	}

	httpServer := net.InitServer(webAddr, "http", S.HandlerHttpRespone)

	go func() {
		err := httpServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	// 客户端连接的TCP连接
	// s := net.InitServer(addr, "tcp", HandlerRespone)
	// s.Listen()

	return S
}
