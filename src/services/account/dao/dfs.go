package dao

import (
	"strconv"

	"github.com/mimis-s/IM-Service/src/common/common_client"
)

/*
	文件的上传和下载
*/

func (d *Dao) upLoadDFS(path, fileName string, payload []byte) error {
	err := d.DFSHandler.PutObject(path, fileName, payload)
	return err
}

func (d *Dao) downLoadDFS(path, fileName string) ([]byte, error) {
	return d.DFSHandler.GetObject(path, fileName)
}

// 用户头像上传和下载
func (d *Dao) UpLoadUserHead(userID int64, payload []byte) error {
	return d.upLoadDFS(common_client.ENUM_DFS_PATH_UserHead, strconv.FormatInt(userID, 10), payload)
}

func (d *Dao) DownLoadUserHead(userID int64) ([]byte, error) {
	return d.downLoadDFS(common_client.ENUM_DFS_PATH_UserHead, strconv.FormatInt(userID, 10))
}
