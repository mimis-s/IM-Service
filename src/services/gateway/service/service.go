package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	home_service "github.com/mimis-s/IM-Service/src/services/home/service"
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

	req := &api_home.ClientRequestHandleReq{
		MsgID:   uint32(reqClient.Tag),
		Payload: reqClient.Msg,
	}
	res := &api_home.ClientRequestHandleRes{}

	err := s.Main.ClientRequestHandleJson(context.Background(), req, res)
	if err != nil {
		errStr := fmt.Sprintf("Client Request Handle Json[%v] is err:%v", req, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		errStr := fmt.Sprintf("json Marshal[%v] is err:%v", res, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}
	return &clientConn.ClientMsg{
		Tag: reqClient.Tag,
		Msg: msg,
	}, nil
}

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao  *dao.Dao
	Main *home_service.Service // 大厅服务
}

func Init(addr, webAddr string) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao:  d,
		Main: new(home_service.Service),
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

// func (s *Service) testHandlerHttpRespone(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
// 	if reqClient.Tag == -1 {
// 		// 心跳包
// 		fmt.Printf("client send heartCheack\n")
// 		return &clientConn.ClientMsg{
// 			Tag: -1,
// 		}, nil
// 	}
// 	fmt.Printf("client send tag:%v message:%s\n", reqClient.Tag, reqClient.Msg)
// 	req := &im_main_proto.ChatSingleReq{}
// 	err := json.Unmarshal(reqClient.Msg, req)
// 	if err != nil {
// 		errStr := fmt.Sprintf("json Marshal[%v] is err:%v", req, err)
// 		fmt.Println(errStr)
// 		return nil, fmt.Errorf(errStr)
// 	}

// 	res := &im_main_proto.ChatSingleRes{
// 		TestStr: req.TestStr,
// 	}

// 	msg, err := json.Marshal(res)
// 	if err != nil {
// 		errStr := fmt.Sprintf("json Marshal[%v] is err:%v", res, err)
// 		fmt.Println(errStr)
// 		return nil, fmt.Errorf(errStr)
// 	}
// 	return &clientConn.ClientMsg{
// 		Tag: reqClient.Tag,
// 		Msg: msg,
// 	}, nil
// }
