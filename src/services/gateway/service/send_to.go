package service

import (
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/golang_tools/net/clientConn"
)

/*
	主要用于给对应userID的客户端发送消息, 现在不考虑历史消息延迟推送, 只考虑实时在线的消息推送
	如果接收端不在线，则直接丢弃包, 这个过程要在返回发送端消息之前，确保接收端在线
*/

func SendToUser(sendUserID, receiverUserID int64, msg interface{}) error {
	msg_id := seralize.GetMsgIdByStruct(msg)
	c, ok := cacheClient.Load(receiverUserID)
	if !ok {
		// 接收方未登录, 消息无法到达
		errStr := fmt.Sprintf("user[%v] send msg[%v] to user[%v], but receiver not online",
			sendUserID, msg_id, receiverUserID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	session := c.(*Session)
	var err error
	data := make([]byte, 0)
	if session.GetClientConn().GetConnType() == clientConn.ClientConn_HTTP_Enum {
		data, err = json.Marshal(msg)
	} else if session.GetClientConn().GetConnType() == clientConn.ClientConn_TCP_Enum {
		data, err = proto.Marshal(msg.(proto.Message))
	}
	if err != nil {
		errStr := fmt.Sprintf("user[%v] send msg[%v] to user[%v], but is err:%v",
			sendUserID, msg_id, receiverUserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	resClientMsg := &clientConn.ClientMsg{
		Tag: int(msg_id),
		Msg: data,
	}
	err = session.GetClientConn().SendMsg(resClientMsg)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] send msg[%v] to user[%v], but is err:%v",
			sendUserID, msg_id, receiverUserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
