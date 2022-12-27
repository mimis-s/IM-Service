package view

import (
	"sync"
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

var singletonMutex sync.Mutex

// 每一个实例的workid都是不一样的
// 默认workid最大值为2^6-1 -> 63, 意思就是最大可以创建63个雪花算法对象实例, 对于一个服务器来说已经够用了
var workerId = uint16(1)

type IdgenObject struct {
	idGenerator *idgen.DefaultIdGenerator
}

func NewIdgenObject() *IdgenObject {
	singletonMutex.Lock()
	defer func() {
		singletonMutex.Unlock()
	}()
	options := idgen.NewIdGeneratorOptions(workerId)
	idGenerator := idgen.NewDefaultIdGenerator(options)
	workerId++
	return &IdgenObject{idGenerator: idGenerator}
}

func (d *IdgenObject) NextId() int64 {
	return d.idGenerator.NewLong()
}

func (d *IdgenObject) ExtractTime(id int64) time.Time {
	return d.idGenerator.ExtractTime(id)
}
