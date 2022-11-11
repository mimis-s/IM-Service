package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
)

/*
	好友列表只会给出简略信息, 详细信息存储在客户端, 客户端登录时检查比对好友信息
	网页客户端只能实时获取
*/

func (s *Service) InitFriendsList(userInfo *im_home_proto.ClientOnlineInfo) (*dbmodel.Friends, im_error_proto.ErrCode, error) {
	dbFriends, find, err := s.Dao.GetFriends(userInfo.UserID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but db is err:%v", userInfo.UserID, err)
		im_log.Error(errStr)
		return nil, im_error_proto.ErrCode_db_read_err, fmt.Errorf(errStr)
	}

	if !find {
		// 找不到说明是新账号, 插入一个
		dbFriends = &dbmodel.Friends{
			UserId: userInfo.UserID,
			Friends: dbmodel.TBJsonField_Friends{
				IDs: make([]int64, 0),
			},
		}
		err = s.Dao.InsertFriends(userInfo, dbFriends)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get friends list, but insert db is err:%v", userInfo.UserID, err)
			im_log.Error(errStr)
			return nil, im_error_proto.ErrCode_db_write_err, fmt.Errorf(errStr)
		}
	}
	return dbFriends, 0, nil
}

func (s *Service) GetFriendsList(ctx context.Context, req *api_friends.GetFriendsListReq, res *api_friends.GetFriendsListRes) error {

	dbFriends, errCode, err := s.InitFriendsList(req.ClientInfo)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	listFriends := make([]*im_home_proto.GetFriendsListRes_BriedFriends, 0, len(dbFriends.Friends.IDs))

	for _, id := range dbFriends.Friends.IDs {
		listFriends = append(listFriends, &im_home_proto.GetFriendsListRes_BriedFriends{
			UserID: id,
		})
	}

	res.Data = &im_home_proto.GetFriendsListRes{
		List: listFriends,
	}

	return nil
}

func (s *Service) ApplyFriends(ctx context.Context, req *api_friends.ApplyFriendsReq, res *api_friends.ApplyFriendsRes) error {
	dbFriends, errCode, err := s.InitFriendsList(req.ClientInfo)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	// 判断是不是好友
	for _, fID := range dbFriends.Friends.IDs {
		if req.Data.ApplyFriendsID == fID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.ApplyFriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
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
		ApplyFriendsID: req.Data.ApplyFriendsID,
	}

	return nil
}

func (s *Service) AgreeFriendApply(ctx context.Context, req *api_friends.AgreeFriendApplyReq, res *api_friends.AgreeFriendApplyRes) error {

	dbFriends, errCode, err := s.InitFriendsList(req.ClientInfo)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friends list, but init list is err:%v", req.ClientInfo.UserID, err)
		im_log.Error(errStr)
		res.ErrCode = errCode
		return fmt.Errorf(errStr)
	}

	for _, id := range dbFriends.Friends.IDs {
		if id == req.Data.FriendsID {
			// 已经是好友
			errStr := fmt.Sprintf("user[%v] agree apply friend, but [%v] already friend", req.ClientInfo.UserID, req.Data.FriendsID)
			res.ErrCode = im_error_proto.ErrCode_friends_user_already_be_friend
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
	}

	dbFriends.Friends.IDs = append(dbFriends.Friends.IDs, req.Data.FriendsID)

	// 更新数据库
	err = s.Dao.UpdateFriends(req.ClientInfo, dbFriends)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] update friends list, but db is err:%v", req.ClientInfo.UserID, err)
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
	return nil
}

func (s *Service) DelFriends(ctx context.Context, req *api_friends.DelFriendsReq, res *api_friends.DelFriendsRes) error {
	return nil
}
