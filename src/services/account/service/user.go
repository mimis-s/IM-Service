package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
)

func (s *Service) Login(ctx context.Context, req *api_account.LoginReq, res *api_account.LoginRes) error {
	// res.Data = req.Data
	userInfo, find, err := s.Dao.GetUserInfoFromID(req.Data.UserID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("role id[%v] register db is err:%v", req.Data.UserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	if !find {
		// 该账号不存在
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("role id[%v] login but account is not find", req.Data.UserID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	// 判断密码是否正确
	if req.Data.Password != userInfo.Password {
		res.ErrCode = im_error_proto.ErrCode_account_password_incorrect
		errStr := fmt.Sprintf("role id[%v] login but password is incorrect", req.Data.UserID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	// 成功登录
	res.Data = &im_home_proto.LoginRes{
		UserID:   userInfo.UserId,
		UserName: userInfo.UserName,
	}

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
		UserId:   time.Now().Unix(), // 先用时间戳代替id, 之后再用算法修正
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

	res.Data = &im_home_proto.RegisterRes{
		UserID:   userInfo.UserId,
		UserName: userInfo.UserName,
	}

	return nil
}
