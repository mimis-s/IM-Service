package service

import (
	"encoding/json"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_main_proto"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/golang_tools/net"
	"github.com/mimis-s/golang_tools/net/clientConn"
)

// 服务器统一处理客户端消息的函数(暂时放在这里，这个后面会挪到lobby服务里面去)
func HandlerRespone(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
	if reqClient.Tag == -1 {
		// 心跳包
		fmt.Printf("client send heartCheack\n")
		return &clientConn.ClientMsg{
			Tag: -1,
		}, nil
	}
	fmt.Printf("client send tag:%v message:%s\n", reqClient.Tag, reqClient.Msg)
	req := &im_main_proto.ChatSingleReq{}
	err := json.Unmarshal(reqClient.Msg, req)
	if err != nil {
		errStr := fmt.Sprintf("json Marshal[%v] is err:%v", req, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}

	res := &im_main_proto.ChatSingleRes{
		TestStr: req.TestStr,
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

type Service struct {
	// 大厅服务
	Dao *dao.Dao
}

func Init(addr, webAddr string) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	httpServer := net.InitServer(webAddr, "http", HandlerRespone)

	go func() {
		err := httpServer.Listen()
		if err != nil {
			panic(err)
		}
	}()

	// 客户端连接的TCP连接
	// s := net.InitServer(addr, "tcp", HandlerRespone)
	// s.Listen()

	S = &Service{
		Dao: d,
	}
	return S
}
