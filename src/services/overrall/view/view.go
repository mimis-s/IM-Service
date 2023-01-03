package view

import (
	"sync"
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

var singletonMutex sync.Mutex

// 每一个实例的workid都是不一样的
// 默认workid最大值为2^WorkerIdBitLength - 1
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
	// 刚开始为一个月的时间差
	options.BaseTime = time.Now().Add(-time.Second*2592000).UnixNano() / 1e6
	options.WorkerIdBitLength = 4
	options.SeqBitLength = 3
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
