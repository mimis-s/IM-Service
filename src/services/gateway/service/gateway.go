package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
)

// 获取连接类型
func (s *Service) GetClientConnType(ctx context.Context, req *api_gateway.GetClientConnTypeReq, res *api_gateway.GetClientConnTypeRes) error {
	c, ok := cacheClient.Load(req.UserID)
	if !ok {
		// 接收方未登录, 消息无法到达
		res.ErrCode = im_error_proto.ErrCode_common_unexpected_err
		errStr := fmt.Sprintf("user[%v] not online", req.UserID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	session := c.(*Session)

	res.ConnType = int32(session.clientConn.GetConnType())
	return nil
}
