package dao

import (
	"strconv"

	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/golang_tools/lib"
)

/*
	聊天图片和文件的上传下载
*/

func (d *Dao) upLoadDFS(path, fileName string, payload []byte) error {
	err := d.dfsHandler.PutObject(path, fileName, payload)
	return err
}

func (d *Dao) downLoadDFS(path, fileName string) ([]byte, error) {
	return d.dfsHandler.GetObject(path, fileName)
}

// 聊天图片上传下载
func (d *Dao) UpLoadChatImage(user_1 int64, user_2 int64, messageID int64, index int, payload []byte) error {
	userFrist := lib.MinInt64(user_1, user_2)
	userSecond := lib.MaxInt64(user_1, user_2)
	name := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10) + "-" + strconv.FormatInt(messageID, 10) + "-" + strconv.Itoa(index)

	return d.upLoadDFS(common_client.ENUM_DFS_PATH_ChatImage, name, payload)
}

func (d *Dao) DownLoadChatImage(user_1 int64, user_2 int64, messageID int64, index int) ([]byte, error) {
	userFrist := lib.MinInt64(user_1, user_2)
	userSecond := lib.MaxInt64(user_1, user_2)
	name := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10) + "-" + strconv.FormatInt(messageID, 10) + "-" + strconv.Itoa(index)
	return d.downLoadDFS(common_client.ENUM_DFS_PATH_ChatImage, name)
}

// 聊天文件上传下载
func (d *Dao) UpLoadChatFile(user_1 int64, user_2 int64, messageID int64, index int, payload []byte) error {

	userFrist := lib.MinInt64(user_1, user_2)
	userSecond := lib.MaxInt64(user_1, user_2)
	name := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10) + "-" + strconv.FormatInt(messageID, 10) + "-" + strconv.Itoa(index)

	return d.upLoadDFS(common_client.ENUM_DFS_PATH_ChatFile, name, payload)
}

func (d *Dao) DownLoadChatFile(user_1 int64, user_2 int64, messageID int64, index int) ([]byte, error) {
	userFrist := lib.MinInt64(user_1, user_2)
	userSecond := lib.MaxInt64(user_1, user_2)
	name := strconv.FormatInt(userFrist, 10) + "-" + strconv.FormatInt(userSecond, 10) + "-" + strconv.FormatInt(messageID, 10) + "-" + strconv.Itoa(index)

	return d.downLoadDFS(common_client.ENUM_DFS_PATH_ChatFile, name)
}
