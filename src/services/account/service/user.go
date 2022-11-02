package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
)

func (s *Service) Login(ctx context.Context, req *api_account.LoginReq, res *api_account.LoginRes) error {
	// res.Data = req.Data
	return nil
}

// 用户注册账户
func (s *Service) Register(ctx context.Context, req *api_account.RegisterReq, res *api_account.RegisterRes) error {
	// res.Data = req.Data
	// 判断是否有重名
	_, find, err := s.Dao.GetUserInfoFromName(req.Data.UserName)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("role name[%v] register db is err:%v", req.Data.UserName, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	if find {
		// 重名
		res.ErrCode = im_error_proto.ErrCode_account_user_name_repeat
		errStr := fmt.Sprintf("role name[%v] register db, but name is repeat", req.Data.UserName)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}
	userInfo := &dbmodel.AccountUser{
		UserID:   time.Now().Unix(), // 先用时间戳代替id, 之后再用算法修正
		UserName: req.Data.UserName,
		Password: req.Data.Password,
		UserExtraInfo: dbmodel.TBJsonField_UserExtraInfo{
			NikeName: req.Data.UserName,
			Nation:   1, // 国家暂时固定ID为1
		},
		Level: 1,
		Exp:   0,
	}

	err = s.Dao.InsertUserInfo(userInfo)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		errStr := fmt.Sprintf("role name[%v] register db write is err:%v", req.Data.UserName, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	res.Data.UserID = userInfo.UserID
	res.Data.UserName = userInfo.UserName

	return nil
}
