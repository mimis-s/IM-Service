package seralize

import (
	"hash/crc32"
	"reflect"

	"github.com/golang/protobuf/proto"
)

/*
	序列化和反序列化消息
*/

type Message proto.Message

func Marshal(m Message) ([]byte, error) {
	return proto.Marshal(m)
}

func Unmarshal(buf []byte, m Message) error {
	return proto.Unmarshal(buf, m)
}

/*
	用消息结构生成ID
*/
func GetMsgIDByName(msgName string) uint32 {
	return crc32.ChecksumIEEE([]byte(msgName))
}

func GetMsgIdByStruct(msgStruct interface{}) uint32 {
	return GetMsgIDByName(reflect.TypeOf(msgStruct).Name())
}

func GetMsgNameByStruct(msgStruct interface{}) string {
	return reflect.TypeOf(msgStruct).Name()
}
