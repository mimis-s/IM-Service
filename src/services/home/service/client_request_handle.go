package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
)

var commonErrMsgID = seralize.GetMsgIdByStruct(im_error_proto.CommonError{})

func dispatchMsg(ctx context.Context, req *api_home.ClientRequestHandleReq, res *api_home.ClientRequestHandleRes,
	decodeFun func([]byte, interface{}) error, encodeFun func(interface{}) ([]byte, error)) error {
	// 通过msg_id拿到注册的handler结构
	handler, find := seralize.RegisterHandlerMap[req.MsgID]
	if !find {
		errStr := fmt.Sprintf("user[%v] req msg_id[%v] not define", req.Client.UserID, req.MsgID)
		fmt.Println(errStr)
		res.ErrCode = int32(im_error_proto.ErrCode_common_unexpected_err)
		return fmt.Errorf(errStr)
	}

	// 通过msg_id拿到原始的req请求
	reqStruct := reflect.New(handler.ReqClient.TypeOf).Interface()
	err := decodeFun(req.Payload, reqStruct)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] req msg_id[%v] decode is err:%v", req.Client.UserID, req.MsgID, err)
		fmt.Println(errStr)
		res.ErrCode = int32(im_error_proto.ErrCode_common_unexpected_err)
		return fmt.Errorf(errStr)
	}

	// 新建一个res请求
	resStruct := reflect.New(handler.ResClient.TypeOf).Interface()

	reqMessage := reqStruct.(seralize.Message)
	resMessage := resStruct.(seralize.Message)

	// 调用单独消息处理函数
	errCode := handler.FuncHandler(ctx, req, reqMessage, resMessage)
	if errCode == im_error_proto.ErrCode_success {
		// 调用没有报错,有返回值
		if resMessage != nil {
			resPayload, err := encodeFun(resMessage)
			if err != nil {
				// 解析失败的话一般是服务器自己的协议写错了
				errStr := fmt.Sprintf("user[%v] ip[%v] req[%v] but service marshal res[%v] is err:%v",
					req.Client.UserID, req.Client.IPAddr, req.MsgID, resMessage, err)
				fmt.Println(errStr)
				errCode = im_error_proto.ErrCode_common_unexpected_err
			} else {
				res.MsgID = handler.ResClient.Crc32ID
				res.Payload = resPayload
				return nil
			}
		}
	}

	// 如果errcode不是0, 则说明要返回错误码
	commonErr := im_error_proto.CommonError{
		Code:       errCode,
		ReqMsgID:   req.MsgID,
		ReqPayload: string(req.Payload),
		ResMsgID:   commonErrMsgID,
	}

	resPayload, _ := encodeFun(commonErr)
	res.MsgID = seralize.GetMsgIdByStruct(im_error_proto.CommonError{})
	res.Payload = resPayload

	return nil
}

// 处理分发客户端发送的消息
func (s *Service) ClientRequestHandleProto(ctx context.Context, req *api_home.ClientRequestHandleReq, res *api_home.ClientRequestHandleRes) error {
	decodeFun := func(buf []byte, m interface{}) error {
		return seralize.Unmarshal(buf, m.(seralize.Message))
	}
	encodeFun := func(m interface{}) ([]byte, error) {
		return seralize.Marshal(m.(seralize.Message))
	}

	return dispatchMsg(ctx, req, res, decodeFun, encodeFun)
}

func (s *Service) ClientRequestHandleJson(ctx context.Context, req *api_home.ClientRequestHandleReq, res *api_home.ClientRequestHandleRes) error {
	return dispatchMsg(ctx, req, res, json.Unmarshal, json.Marshal)
}
