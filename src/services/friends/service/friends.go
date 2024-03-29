package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/golang_tools/zlog"
)

/*
	好友列表只会给出简略信息, 详细信息存储在客户端, 客户端登录时检查比对好友信息
	网页客户端只能实时获取
*/

func (s *Service) InitFriendsList(userID int64) (*dbmodel.Friends, im_error_proto.ErrCode, error) {
	dbFriends, find, err := s.Dao.GetFriends(userID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but db is err:%v", userID, err)
		zlog.Error(errStr)
		return nil, im_error_proto.ErrCode_db_read_err, fmt.Errorf(errStr)
	}

	if !find {
		// 找不到说明是新账号, 插入一个
		dbFriends = &dbmodel.Friends{
			UserId: userID,
			Friends: dbmodel.TBJsonField_Friends{
				IDs:             make([]int64, 0),
				ApplyFriendIDs:  make([]int64, 0),
				ReceiveApplyIDs: make([]int64, 0),
			},
		}
		err = s.Dao.InsertFriends(userID, dbFriends)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get friends list, but insert db is err:%v", userID, err)
			zlog.Error(errStr)
			return nil, im_error_proto.ErrCode_db_write_err, fmt.Errorf(errStr)
		}
	}
	return dbFriends, 0, nil
}

func (s *Service) GetFriendsStatusList(ctx context.Context, userInfo *im_home_proto.UserInfo) (*im_home_proto.FriendsStatusList, im_error_proto.ErrCode, error) {

	friendsStatusList := &im_home_proto.FriendsStatusList{}

	dbFriends, errCode, err := s.InitFriendsList(userInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but init list is err:%v", userInfo.UserID, err)
		zlog.Error(errStr)
		return friendsStatusList, errCode, fmt.Errorf(errStr)
	}

	friendsStatusList.FriendsStatusList = make([]*im_home_proto.FriendInfoStatus, 0, len(dbFriends.Friends.IDs))
	friendsStatusList.SendFriendApplyList = make([]*im_home_proto.FriendInfoStatus, 0, len(dbFriends.Friends.ApplyFriendIDs))
	friendsStatusList.ReceiveFriendApplyList = make([]*im_home_proto.FriendInfoStatus, 0, len(dbFriends.Friends.ReceiveApplyIDs))

	// 获取好友的信息
	userIDs := append(dbFriends.Friends.IDs, dbFriends.Friends.ApplyFriendIDs...)
	userIDs = append(userIDs, dbFriends.Friends.ReceiveApplyIDs...)

	if len(userIDs) != 0 {
		getUsersInfoServiceReq := &api_account.GetUsersInfoServiceReq{
			ClientInfo: &im_home_proto.ClientOnlineInfo{
				UserID: userInfo.UserID,
			},
			UserIDs: userIDs,
		}

		getUsersInfoServiceRes, err := api_account.GetUsersInfoService(context.Background(), getUsersInfoServiceReq)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] Get Users[%v] Info Service is err:%v", userInfo.UserID, userIDs, err)
			zlog.Error(errStr)
			return friendsStatusList, getUsersInfoServiceRes.ErrCode, fmt.Errorf(errStr)
		}
		findFriend := func(id int64, array []int64) bool {
			for _, a := range array {
				if id == a {
					return true
				}
			}
			return false
		}
		for _, friend := range getUsersInfoServiceRes.Datas {
			if findFriend(friend.UserID, dbFriends.Friends.IDs) {
				friendsStatusList.FriendsStatusList = append(friendsStatusList.FriendsStatusList,
					&im_home_proto.FriendInfoStatus{
						Friend:       friend,
						IsUpdateHead: true,
					})
				continue
			}
			if findFriend(friend.UserID, dbFriends.Friends.ApplyFriendIDs) {
				friendsStatusList.SendFriendApplyList = append(friendsStatusList.SendFriendApplyList,
					&im_home_proto.FriendInfoStatus{
						Friend:       friend,
						IsUpdateHead: true,
					})
				continue
			}
			if findFriend(friend.UserID, dbFriends.Friends.ReceiveApplyIDs) {
				friendsStatusList.ReceiveFriendApplyList = append(friendsStatusList.ReceiveFriendApplyList,
					&im_home_proto.FriendInfoStatus{
						Friend:       friend,
						IsUpdateHead: true,
					})
				continue
			}
		}
	}

	return friendsStatusList, 0, nil
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
		zlog.Error(errStr)
		res.ErrCode = userInfoRes.ErrCode
		return fmt.Errorf(errStr)
	}

	// 自己
	dbSelfFriends, errCode, err := s.InitFriendsList(req.ClientInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but init list is err:%v", req.ClientInfo.UserID, err)
		zlog.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 查询重复
	findRepeat := func(id int64, friendIDs []int64) bool {
		for _, i := range friendIDs {
			if i == id {
				return true
			}
		}
		return false
	}
	if findRepeat(req.Data.ApplyFriendsID, dbSelfFriends.Friends.ApplyFriendIDs) {
		errStr := fmt.Sprintf("user[%v] apply friend, but [%v] already apply friend", req.ClientInfo.UserID, req.Data.ApplyFriendsID)
		res.ErrCode = im_error_proto.ErrCode_friends_user_already_apply_friend
		zlog.Error(errStr)
		return fmt.Errorf(errStr)
	}

	// 判断是不是好友
	for _, fID := range dbSelfFriends.Friends.IDs {
		if req.Data.ApplyFriendsID == fID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.ApplyFriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			zlog.Error(errStr)
			return fmt.Errorf(errStr)
		}
	}

	// 好友
	dbOtherFriends, errCode, err := s.InitFriendsList(req.Data.ApplyFriendsID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but user[%v] init list is err:%v", req.ClientInfo.UserID, req.Data.ApplyFriendsID, err)
		zlog.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 修改数据库数据
	dbSelfFriends.Friends.ApplyFriendIDs = append(dbSelfFriends.Friends.ApplyFriendIDs, req.Data.ApplyFriendsID)
	dbOtherFriends.Friends.ReceiveApplyIDs = append(dbOtherFriends.Friends.ReceiveApplyIDs, req.ClientInfo.UserID)

	err = s.Dao.UpdateFriends(req.ClientInfo.UserID, dbSelfFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but update db is err:%v", req.ClientInfo.UserID, err)
		zlog.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}
	err = s.Dao.UpdateFriends(req.Data.ApplyFriendsID, dbOtherFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Apply Friends, but user[%v] update db is err:%v", req.ClientInfo.UserID, req.Data.ApplyFriendsID, err)
		zlog.Error(errStr)
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
		zlog.Error(errStr)
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
		zlog.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	for _, id := range dbSelfFriends.Friends.IDs {
		if id == req.Data.FriendsID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] agree apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.FriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			zlog.Error(errStr)
			return fmt.Errorf(errStr)
		}
	}

	dbOtherFriends, errCode, err := s.InitFriendsList(req.Data.FriendsID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] Agree Friend[%v] Apply, but init list is err:%v", req.ClientInfo.UserID, req.Data.FriendsID, err)
		zlog.Error(errStr)
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
		zlog.Error(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		return fmt.Errorf(errStr)
	}

	err = s.Dao.UpdateFriends(req.Data.FriendsID, dbOtherFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] update friend[%v], but db is err:%v", req.ClientInfo.UserID, req.Data.FriendsID, err)
		zlog.Error(errStr)
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
		zlog.Error(errStr)
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
