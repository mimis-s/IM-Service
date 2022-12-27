package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/services/overrall/api_overrall"
	"github.com/mimis-s/IM-Service/src/services/overrall/view"
)

var GenerateUserIDIdgenObject *view.IdgenObject

// 雪花算法生成唯一用户id
func (s *Service) GenerateUniqueUserID(ctx context.Context, req *api_overrall.GenerateUniqueUserIDReq, res *api_overrall.GenerateUniqueUserIDRes) error {
	var newId = GenerateUserIDIdgenObject.NextId()
	res.UserID = newId
	return nil
}
