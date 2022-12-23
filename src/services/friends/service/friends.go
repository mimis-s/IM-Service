package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
)

/*
	好友列表只会给出简略信息, 详细信息存储在客户端, 客户端登录时检查比对好友信息
	网页客户端只能实时获取
*/

func (s *Service) InitFriendsList(userID int64) (*dbmodel.Friends, im_error_proto.ErrCode, error) {
	dbFriends, find, err := s.Dao.GetFriends(userID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but db is err:%v", userID, err)
		im_log.Error(errStr)
		return nil, im_error_proto.ErrCode_db_read_err, fmt.Errorf(errStr)
	}

	if !find {
		// 找不到说明是新账号, 插入一个
		dbFriends = &dbmodel.Friends{
			UserId: userID,
			Friends: dbmodel.TBJsonField_Friends{
				IDs: make([]int64, 0),
			},
		}
		err = s.Dao.InsertFriends(userID, dbFriends)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get friends list, but insert db is err:%v", userID, err)
			im_log.Error(errStr)
			return nil, im_error_proto.ErrCode_db_write_err, fmt.Errorf(errStr)
		}
	}
	return dbFriends, 0, nil
}

func (s *Service) GetFriendsList(ctx context.Context, req *api_friends.GetFriendsListReq, res *api_friends.GetFriendsListRes) error {

	dbFriends, errCode, err := s.InitFriendsList(req.ClientInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 获取好友的信息

	if len(dbFriends.Friends.IDs) != 0 {
		getUsersInfoServiceReq := &api_account.GetUsersInfoServiceReq{
			ClientInfo: req.ClientInfo,
			UserIDs:    dbFriends.Friends.IDs,
		}

		getUsersInfoServiceRes, err := api_account.GetUsersInfoService(context.Background(), getUsersInfoServiceReq)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] Get Users[%v] Info Service, but insert db is err:%v", req.ClientInfo.UserID, dbFriends.Friends.IDs, err)
			im_log.Error(errStr)
			res.ErrCode = getUsersInfoServiceRes.ErrCode
			return fmt.Errorf(errStr)
		}

		res.Data = &im_home_proto.GetFriendsListRes{
			List: getUsersInfoServiceRes.Datas,
		}
	} else {
		// 还没有好友
		res.Data = &im_home_proto.GetFriendsListRes{
			List: []*im_home_proto.UserInfo{},
		}
	}

	return nil
}

func (s *Service) ApplyFriends(ctx context.Context, req *api_friends.ApplyFriendsReq, res *api_friends.ApplyFriendsRes) error {
	// 判断添加的这个id是否存在
	userInfoReq := &api_account.GetUserInfoServiceReq{
		ClientInfo: req.ClientInfo,
		UserID:     req.Data.ApplyFriendsID,
	}
	userInfoRes, err := api_account.GetUserInfoService(context.Background(), userInfoReq)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get User[%v] Info is err:%v", req.ClientInfo.UserID, req.Data.ApplyFriendsID, err)
		im_log.Error(errStr)
		res.ErrCode = userInfoRes.ErrCode
		return fmt.Errorf(errStr)
	}

	// 自己
	dbSelfFriends, errCode, err := s.InitFriendsList(req.ClientInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 判断是不是好友
	for _, fID := range dbSelfFriends.Friends.IDs {
		if req.Data.ApplyFriendsID == fID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.ApplyFriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
	}

	// 好友
	dbOtherFriends, errCode, err := s.InitFriendsList(req.Data.ApplyFriendsID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but user[%v] init list is err:%v", req.ClientInfo.UserID, req.Data.ApplyFriendsID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 修改数据库数据
	dbSelfFriends.Friends.ApplyFriendIDs = append(dbSelfFriends.Friends.ApplyFriendIDs, req.Data.ApplyFriendsID)
	dbOtherFriends.Friends.ReceiveApplyIDs = append(dbSelfFriends.Friends.ReceiveApplyIDs, req.ClientInfo.UserID)
	err = s.Dao.UpdateFriends(req.ClientInfo.UserID, dbSelfFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but update db is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}
	err = s.Dao.UpdateFriends(req.Data.ApplyFriendsID, dbOtherFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but user[%v] update db is err:%v", req.ClientInfo.UserID, req.Data.ApplyFriendsID, err)
		im_log.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}

	// 转发给对方
	eventApplyFriend := &event.ApplyFriend{
		UserInfo: req.ClientInfo,
		Message: &im_home_proto.ApplyFriendsToReceiver{
			SenderID:   req.ClientInfo.UserID,
			ReceiverID: req.Data.ApplyFriendsID,
		},
	}
	err = event.Publish(event.Event_ApplyFriend, eventApplyFriend)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but publish is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	res.Data = &im_home_proto.ApplyFriendsRes{
		FriendInfo: userInfoRes.Data,
	}

	return nil
}

func (s *Service) AgreeFriendApply(ctx context.Context, req *api_friends.AgreeFriendApplyReq, res *api_friends.AgreeFriendApplyRes) error {

	dbSelfFriends, errCode, err := s.InitFriendsList(req.ClientInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Agree Friend Apply, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	for _, id := range dbSelfFriends.Friends.IDs {
		if id == req.Data.FriendsID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] agree apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.FriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
	}

	dbOtherFriends, errCode, err := s.InitFriendsList(req.Data.FriendsID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Agree Friend[%v] Apply, but init list is err:%v", req.ClientInfo.UserID, req.Data.FriendsID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	dbSelfFriends.Friends.IDs = append(dbSelfFriends.Friends.IDs, req.Data.FriendsID)
	dbOtherFriends.Friends.IDs = append(dbOtherFriends.Friends.IDs, req.ClientInfo.UserID)

	// 删除申请列表和接收列表中对应的id
	for i := 0; i < len(dbSelfFriends.Friends.ReceiveApplyIDs); i++ {
		if dbSelfFriends.Friends.ReceiveApplyIDs[i] == req.Data.FriendsID {
			dbSelfFriends.Friends.ReceiveApplyIDs = append(dbSelfFriends.Friends.ReceiveApplyIDs[:i], dbSelfFriends.Friends.ReceiveApplyIDs[i+1:]...)
		}
	}

	for i := 0; i < len(dbOtherFriends.Friends.ApplyFriendIDs); i++ {
		if dbOtherFriends.Friends.ApplyFriendIDs[i] == req.ClientInfo.UserID {
			dbOtherFriends.Friends.ApplyFriendIDs = append(dbOtherFriends.Friends.ApplyFriendIDs[:i], dbOtherFriends.Friends.ApplyFriendIDs[i+1:]...)
		}
	}

	// 更新数据库
	err = s.Dao.UpdateFriends(req.ClientInfo.UserID, dbSelfFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] update friend, but db is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}

	err = s.Dao.UpdateFriends(req.Data.FriendsID, dbOtherFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] update friend[%v], but db is err:%v", req.ClientInfo.UserID, req.Data.FriendsID, err)
		im_log.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}

	// 转发给对方
	eventAgreeApplyFriend := &event.AgreeApplyFriend{
		UserInfo: req.ClientInfo,
		Message: &im_home_proto.AgreeApplyFriendsToReceiver{
			SenderID:   req.ClientInfo.UserID,
			ReceiverID: req.Data.FriendsID,
		},
	}
	err = event.Publish(event.Event_AgreeApplyFriend, eventAgreeApplyFriend)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but publish is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}
	res.Data = &im_home_proto.AgreeFriendApplyRes{
		FriendsID: req.Data.FriendsID,
	}
	return nil
}

func (s *Service) DelFriends(ctx context.Context, req *api_friends.DelFriendsReq, res *api_friends.DelFriendsRes) error {
	return nil
}
