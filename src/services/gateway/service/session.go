package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/golang_tools/net/clientConn"
	"github.com/mimis-s/golang_tools/zlog"
)

var cacheClient = sync.Map{}

// session记录每一个客户端的连接情况
type Session struct {
	clientInfo *im_home_proto.ClientOnlineInfo
	clientConn clientConn.ClientConn
}

func NewSession(clientConn clientConn.ClientConn) clientConn.ClientSession {
	return &Session{
		clientConn: clientConn,
	}
}

func (s *Session) GetClientConn() clientConn.ClientConn {
	return s.clientConn
}

func (s *Session) ConnectCallBack() {

}

func (s *Session) RequestCallBack(reqClient *clientConn.ClientMsg) (*clientConn.ClientMsg, error) {
	if reqClient.Tag == -1 {
		// 心跳包
		zlog.Info("client send heartCheack\n")
		return &clientConn.ClientMsg{
			Tag: -1,
		}, nil
	}
	// 如果是登录则记录用户ID和conn的映射
	loginReqID := seralize.GetMsgIdByStruct(im_home_proto.LoginReq{})
	if uint32(reqClient.Tag) == loginReqID {
		// 这里比较完善的做法应该是需要存入redis，但是现在可以先存入map里面

		// 判断当前用户是否已经在线，如果在线，则删除之前的保留现在的
		loginReq := &im_home_proto.LoginReq{}
		if s.GetClientConn().GetConnType() == clientConn.ClientConn_HTTP_Enum {
			json.Unmarshal(reqClient.Msg, loginReq)
		} else if s.GetClientConn().GetConnType() == clientConn.ClientConn_TCP_Enum {
			proto.Unmarshal(reqClient.Msg, loginReq)
		}
		s.clientInfo = &im_home_proto.ClientOnlineInfo{
			UserID: loginReq.UserID,
		}
		s.clientInfo.UserID = loginReq.UserID
		zlog.Info("用户[%v] IP[%v]尝试登录", loginReq.UserID, s.GetClientConn().GetIP())
		cacheClient.Store(loginReq.UserID, s)
	}

	// 先不记录client连接
	req := &api_home.ClientRequestHandleReq{
		MsgID:   uint32(reqClient.Tag),
		Payload: reqClient.Msg,
		Client:  s.clientInfo,
	}

	if s.GetClientConn().GetConnType() == clientConn.ClientConn_HTTP_Enum {

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

	res, err := api_home.ClientRequestHandleProto(context.Background(), req)
	if err != nil {
		errStr := fmt.Sprintf("Client Request Handle proto[%v] is err:%v", req, err)
		fmt.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}

	return &clientConn.ClientMsg{
		Tag: int(res.MsgID),
		Msg: res.Payload,
	}, nil
}

func (s *Session) DisConnectCallBack() {
	if s.clientInfo == nil {
		// 用户没有登录
		zlog.Info("IP[%v]断开连接", s.GetClientConn().GetIP())
		return
	}
	// 用户登出
	logoutReq := &api_account.LogoutReq{
		ClientInfo: s.clientInfo,
		Data:       &im_home_proto.LogoutReq{},
	}
	_, err := api_account.Logout(context.Background(), logoutReq)
	if err != nil {
		zlog.Error("用户[%v] IP[%v]登出 err:%v", s.clientInfo.UserID, s.GetClientConn().GetIP(), err)
	}
	cacheClient.Delete(s.clientInfo.UserID)
	zlog.Info("用户[%v] IP[%v]登出", s.clientInfo.UserID, s.GetClientConn().GetIP())
}
