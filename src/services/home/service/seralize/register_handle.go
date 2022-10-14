package seralize

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mimis-s/IM-Service/src/common/commonproto/github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
)

/*
	大厅主要的作用就是将客户端发送的消息分发到不同的服务里面去
	所以每一次请求都要生成一个新的调用函数,调用到不同的服务里面
*/

// 消息原始结构
type msgMetaData struct {
	OriginalStruct interface{}
	TypeOf         reflect.Type
	Crc32ID        uint32
	Name           string
}

func newMsgMetaData(s interface{}) *msgMetaData {
	if s == nil {
		return nil
	}
	typeOf := reflect.TypeOf(s)
	fmt.Printf("name:%v crc32:%v", typeOf.Name(), GetMsgIDByName(typeOf.Name()))
	return &msgMetaData{
		OriginalStruct: s,
		TypeOf:         typeOf,
		Crc32ID:        GetMsgIDByName(typeOf.Name()),
		Name:           typeOf.Name(),
	}
}

type handlerFunc func(ctx context.Context, req, res Message) im_error_proto.ErrCode

type MsgHandler struct {
	ReqClient   *msgMetaData
	ResClient   *msgMetaData
	FuncHandler handlerFunc
}

var RegisterHandlerMap map[uint32]*MsgHandler

// 注册handler
func RegisterHandler(req interface{}, res interface{}, f handlerFunc) {
	handler := new(MsgHandler)
	handler.ReqClient = newMsgMetaData(req)
	handler.ResClient = newMsgMetaData(res)
	handler.FuncHandler = f

	_, find := RegisterHandlerMap[handler.ReqClient.Crc32ID]
	if find {
		errStr := fmt.Sprintf("重复注册相同的协议:%v", handler.ReqClient.Name)
		panic(fmt.Errorf(errStr))
	}
	RegisterHandlerMap[handler.ReqClient.Crc32ID] = handler
}
