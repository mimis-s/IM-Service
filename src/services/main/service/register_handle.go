package service

import (
	"context"
	"reflect"
)

/*
	大厅主要的作用就是将客户端发送的消息分发到不同的服务里面去
	所以每一次请求都要生成一个新的调用函数,调用到不同的服务里面
*/

// 消息原始结构
type msgMetaData struct {
	originalStruct interface{}
	typeOf         reflect.Type
	crc32ID        uint32
	name           string
}

type MsgHandler struct {
	reqClient   *msgMetaData
	resClient   *msgMetaData
	FuncHandler func(ctx context.Context, req Message, res Message)
}
