package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
)

func (s *Service) Login(ctx context.Context, req *api_account.LoginReq, res *api_account.LoginRes) error {
	// res.Data = req.Data
	userInfo, find, err := s.Dao.GetUserInfoFromID(req.Data.UserID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("role id[%v] register db is err:%v", req.Data.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if !find {
		// 该账号不存在
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("role id[%v] login but account is not find", req.Data.UserID)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 判断密码是否正确
	if req.Data.Password != userInfo.Password {
		res.ErrCode = im_error_proto.ErrCode_account_password_incorrect
		errStr := fmt.Sprintf("role id[%v] login but password is incorrect", req.Data.UserID)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 成功登录
	res.Data = &im_home_proto.LoginRes{
		Info: &im_home_proto.UserInfo{
			UserID:      userInfo.UserId,
			UserName:    userInfo.UserName,
			Region:      int32(userInfo.UserExtraInfo.Nation),
			Autograph:   userInfo.UserExtraInfo.PersonalSignature,
			Status:      im_home_proto.Enum_UserStatus_Enum_UserStatus_Online,
			PhoneNumber: userInfo.UserExtraInfo.PhoneNumber,
		},
	}

	// 下载头像
	headImg, err := s.Dao.DownLoadUserHead(userInfo.UserId)
	if err != nil {
		// 这个地方日志打印错误, 但是这个错误不返回给客户端,不阻断主流程
		errStr := fmt.Sprintf("user[%v] login, but dfs down load head is err:%v", userInfo.UserId, err)
		im_log.Warn(errStr)
	} else {
		res.Data.Info.HeadImg = string(headImg)
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
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if find {
		// 重名
		res.ErrCode = im_error_proto.ErrCode_account_user_name_repeat
		errStr := fmt.Sprintf("role name[%v] register db, but name is repeat", req.Data.UserName)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}
	userInfo := &dbmodel.AccountUser{
		UserId:   time.Now().Unix(), // 先用时间戳代替id, 之后再用算法修正
		UserName: req.Data.UserName,
		Password: req.Data.Password,
		UserExtraInfo: dbmodel.TBJsonField_UserExtraInfo{
			NikeName:          req.Data.UserName,
			Nation:            int(req.Data.Region),
			PersonalSignature: req.Data.Autograph,
			PhoneNumber:       req.Data.PhoneNumber,
		},
		Level: 1,
		Exp:   0,
	}

	err = s.Dao.InsertUserInfo(userInfo)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		errStr := fmt.Sprintf("role name[%v] register db write is err:%v", req.Data.UserName, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 上传头像
	err = s.Dao.UpLoadUserHead(userInfo.UserId, []byte(req.Data.HeadImg))
	if err != nil {
		// 这个地方日志打印错误, 但是这个错误不返回给客户端,不阻断主流程
		errStr := fmt.Sprintf("user id[%v] register, but dfs up load head is err:%v", userInfo.UserId, err)
		im_log.Warn(errStr)
	}

	res.Data = &im_home_proto.RegisterRes{
		UserID:   userInfo.UserId,
		UserName: userInfo.UserName,
	}

	return nil
}

// 获取用户信息(客户端)
func (s *Service) GetUserInfo(ctx context.Context, req *api_account.GetUserInfoReq, res *api_account.GetUserInfoRes) error {

	userInfo, find, err := s.Dao.GetUserInfoFromID(req.Data.UserID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but db is err:%v", req.ClientInfo.UserID, req.Data.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if !find {
		// 没有找到说明没有这个人
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but db is not found", req.ClientInfo.UserID, req.Data.UserID)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 获取在线状态
	status, err := s.Dao.CacheGetUserStatus([]int64{userInfo.UserId})
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but redis db is err:%v", req.ClientInfo.UserID, req.Data.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	res.Data = &im_home_proto.GetUserInfoRes{
		Relation: im_home_proto.Enum_UserRelation_Enum_UserRelation_Stranger, // 暂时固定为非好友
		Data: &im_home_proto.UserInfo{
			UserID:      userInfo.UserId,
			UserName:    userInfo.UserName,
			Region:      int32(userInfo.UserExtraInfo.Nation),
			Autograph:   userInfo.UserExtraInfo.PersonalSignature,
			Status:      status[0],
			PhoneNumber: userInfo.UserExtraInfo.PhoneNumber,
		},
	}

	// 下载头像
	headImg, err := s.Dao.DownLoadUserHead(userInfo.UserId)
	if err != nil {
		// 这个地方日志打印错误, 但是这个错误不返回给客户端,不阻断主流程
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but dfs down load head is err:%v", req.ClientInfo.UserID, userInfo.UserId, err)
		im_log.Warn(errStr)
	}
	res.Data.Data.HeadImg = string(headImg)

	return nil
}

// 获取用户信息(服务器)
func (s *Service) GetUserInfoService(ctx context.Context, req *api_account.GetUserInfoServiceReq, res *api_account.GetUserInfoServiceRes) error {

	userInfo, find, err := s.Dao.GetUserInfoFromID(req.UserID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but db is err:%v", req.ClientInfo.UserID, req.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if !find {
		// 没有找到说明没有这个人
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but db is not found", req.ClientInfo.UserID, req.UserID)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 获取在线状态
	status, err := s.Dao.CacheGetUserStatus([]int64{userInfo.UserId})
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but redis db is err:%v", req.ClientInfo.UserID, req.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	res.Data = &im_home_proto.UserInfo{
		UserID:      userInfo.UserId,
		UserName:    userInfo.UserName,
		Region:      int32(userInfo.UserExtraInfo.Nation),
		Autograph:   userInfo.UserExtraInfo.PersonalSignature,
		Status:      status[0],
		PhoneNumber: userInfo.UserExtraInfo.PhoneNumber,
	}

	return nil
}

// 获取用户信息(服务器),多个用户
func (s *Service) GetUsersInfoService(ctx context.Context, req *api_account.GetUsersInfoServiceReq, res *api_account.GetUsersInfoServiceRes) error {

	userInfos, err := s.Dao.GetUserInfoFromIDs(req.UserIDs)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get users[%v] info, but db is err:%v", req.ClientInfo.UserID, req.UserIDs, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if len(userInfos) == 0 {
		// 没有找到说明没有这个人
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("user[%v] get users[%v] info, but db is not found", req.ClientInfo.UserID, req.UserIDs)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	userIDs := make([]int64, 0, len(userInfos))
	for _, v := range userInfos {
		userIDs = append(userIDs, v.UserId)
	}

	// 获取在线状态
	status, err := s.Dao.CacheGetUserStatus(userIDs)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but redis db is err:%v", req.ClientInfo.UserID, req.UserIDs, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	res.Datas = make([]*im_home_proto.UserInfo, 0, len(userInfos))

	for index, userInfo := range userInfos {
		res.Datas = append(res.Datas, &im_home_proto.UserInfo{
			UserID:      userInfo.UserId,
			UserName:    userInfo.UserName,
			Region:      int32(userInfo.UserExtraInfo.Nation),
			Autograph:   userInfo.UserExtraInfo.PersonalSignature,
			Status:      status[index],
			PhoneNumber: userInfo.UserExtraInfo.PhoneNumber,
		})
	}

	return nil
}

// 修改用户信息
func (s *Service) ModifyUserInfo(ctx context.Context, req *api_account.ModifyUserInfoReq, res *api_account.ModifyUserInfoRes) error {
	userInfo, find, err := s.Dao.GetUserInfoFromID(req.Data.Data.UserID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		errStr := fmt.Sprintf("user[%v] modify user[%v] info, but db is err:%v", req.ClientInfo.UserID, req.Data.Data.UserID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if !find {
		// 没有找到说明没有这个人
		res.ErrCode = im_error_proto.ErrCode_account_account_not_found
		errStr := fmt.Sprintf("user[%v] modify user[%v] info, but db is not found", req.ClientInfo.UserID, req.Data.Data.UserID)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	bUpdate := false

	// 如果修改了用户名,判断是否有重名
	if req.Data.Data.UserName != "" && userInfo.UserName != req.Data.Data.UserName {
		_, find, err = s.Dao.GetUserInfoFromName(req.Data.Data.UserName)
		if err != nil {
			res.ErrCode = im_error_proto.ErrCode_db_read_err
			errStr := fmt.Sprintf("role name[%v] get db is err:%v", req.Data.Data.UserName, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}

		if find {
			// 重名
			res.ErrCode = im_error_proto.ErrCode_account_user_name_repeat
			errStr := fmt.Sprintf("role name[%v] modify db, but name is repeat", req.Data.Data.UserName)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
		bUpdate = true
		userInfo.UserName = req.Data.Data.UserName
	}
	// 国家
	if req.Data.Data.Region != 0 && userInfo.UserExtraInfo.Nation != int(req.Data.Data.Region) {
		userInfo.UserExtraInfo.Nation = int(req.Data.Data.Region)
		bUpdate = true
	}
	// 签名
	if req.Data.Data.Autograph != "" && userInfo.UserExtraInfo.PersonalSignature != req.Data.Data.Autograph {
		userInfo.UserExtraInfo.PersonalSignature = req.Data.Data.Autograph
		bUpdate = true
	}

	// 电话
	if req.Data.Data.PhoneNumber != "" && userInfo.UserExtraInfo.PhoneNumber != req.Data.Data.PhoneNumber {
		userInfo.UserExtraInfo.PhoneNumber = req.Data.Data.PhoneNumber
		bUpdate = true
	}

	if bUpdate {
		err = s.Dao.UpdateUserInfo(userInfo)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] modify user info, but update db[%v] is err:%v", userInfo.UserId, userInfo, err)
			im_log.Warn(errStr)
			res.ErrCode = im_error_proto.ErrCode_db_write_err
			return fmt.Errorf(errStr)
		}
	}

	res.Data = &im_home_proto.ModifyUserInfoRes{
		Data: &im_home_proto.UserInfo{
			UserID:      userInfo.UserId,
			UserName:    userInfo.UserName,
			Region:      int32(userInfo.UserExtraInfo.Nation),
			Autograph:   userInfo.UserExtraInfo.PersonalSignature,
			Status:      im_home_proto.Enum_UserStatus_Enum_UserStatus_Online,
			PhoneNumber: userInfo.UserExtraInfo.PhoneNumber,
		},
	}

	// 头像如果变动,只需要上传文件服务器
	if req.Data.Data.HeadImg != "" {
		// 上传头像
		err = s.Dao.UpLoadUserHead(userInfo.UserId, []byte(req.Data.Data.HeadImg))
		if err != nil {
			// 这个地方日志打印错误, 但是这个错误不返回给客户端,不阻断主流程
			errStr := fmt.Sprintf("user[%v] modify head imag, but dfs up load head is err:%v", userInfo.UserId, err)
			im_log.Warn(errStr)
		}
		res.Data.Data.HeadImg = string(req.Data.Data.HeadImg)
	} else {
		// 下载头像
		headImg, err := s.Dao.DownLoadUserHead(userInfo.UserId)
		if err != nil {
			// 这个地方日志打印错误, 但是这个错误不返回给客户端,不阻断主流程
			errStr := fmt.Sprintf("user[%v] login, but dfs down load head is err:%v", userInfo.UserId, err)
			im_log.Warn(errStr)
		} else {
			res.Data.Data.HeadImg = string(headImg)
		}
	}

	return nil
}
