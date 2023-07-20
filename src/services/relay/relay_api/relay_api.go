package relay_api

import (
	context "context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/golang_tools/net/clientConn"
	"github.com/mimis-s/golang_tools/zlog"
)

func NotifyUser(userID int64, msgStruct seralize.Message) error {
	getClientConnTypeReq := &api_gateway.GetClientConnTypeReq{
		UserID: userID,
	}
	getClientConnTypeRes, err := api_gateway.GetClientConnType(context.Background(), getClientConnTypeReq)
	if err != nil {
		zlog.Warn("notify user[%v] get conn type is err:%v", userID, err)
		return err
	}
	var payLoad []byte

	if clientConn.ClientConn_Enum(getClientConnTypeRes.ConnType) == clientConn.ClientConn_TCP_Enum {
		payLoad, err = seralize.Marshal(msgStruct)
		if err != nil {
			zlog.Warn("proto marshal[%v] is err:%v", msgStruct, err)
			return err
		}
	} else if clientConn.ClientConn_Enum(getClientConnTypeRes.ConnType) == clientConn.ClientConn_HTTP_Enum {
		payLoad, err = json.Marshal(msgStruct)
		if err != nil {
			zlog.Warn("json marshal[%v] is err:%v", msgStruct, err)
			return err
		}
	} else {
		errStr := fmt.Sprintf("notify user[%v] conn type[%v] is not define", userID, getClientConnTypeRes.ConnType)
		zlog.Warn(errStr)
		return fmt.Errorf(errStr)
	}

	msgID := seralize.GetMsgIDByName(reflect.TypeOf(msgStruct).Elem().Name())

	notifyClientReq := &api_gateway.NotifyClientReq{
		UserID:  userID,
		MsgTag:  msgID,
		Payload: payLoad,
	}
	api_gateway.NotifyClient(context.Background(), notifyClientReq)
	return nil
}
