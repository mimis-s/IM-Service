package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/golang_tools/net/clientConn"
	"google.golang.org/protobuf/proto"
)

/*
	现在有两种类型 1:服务器主动推送 2:服务器转发客户端数据
*/

func (s *Service) SendToClient(senderID, receiverID int64, msgID uint32, msgStruct interface{}) error {
	getClientConnTypeReq := &api_gateway.GetClientConnTypeReq{}
	getClientConnTypeRes, err := api_gateway.GetClientConnType(context.Background(), getClientConnTypeReq)
	if err != nil {
		im_log.Warn("senderID[%v] get receiverID[%v] conn type is err:%v", senderID, receiverID, err)
		return err
	}
	var payLoad []byte

	if clientConn.ClientConn_Enum(getClientConnTypeRes.ConnType) == clientConn.ClientConn_TCP_Enum {
		payLoad, err = proto.Marshal(msgStruct.(proto.Message))
		if err != nil {
			im_log.Warn("proto marshal[%v] is err:%v", msgStruct, err)
			return err
		}
	} else if clientConn.ClientConn_Enum(getClientConnTypeRes.ConnType) == clientConn.ClientConn_HTTP_Enum {
		payLoad, err = json.Marshal(msgStruct)
		if err != nil {
			im_log.Warn("json marshal[%v] is err:%v", msgStruct, err)
			return err
		}
	} else {
		errStr := fmt.Sprintf("senderID[%v] getreceiverID[%v] conn type[%v] is not define", senderID, receiverID, getClientConnTypeRes.ConnType)
		im_log.Warn(errStr)
		return fmt.Errorf(errStr)
	}

	sendToClientReq := &api_gateway.SendToClientReq{
		SenderID:   senderID,
		ReceiverID: receiverID,
		MsgTag:     msgID,
		Payload:    payLoad,
	}
	api_gateway.SendToClient(context.Background(), sendToClientReq)
	return nil
}
