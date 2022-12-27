package view

import (
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

type IdgenObject struct {
	idGenerator *idgen.DefaultIdGenerator
}

func NewIdgenObject() *IdgenObject {
	options := idgen.NewIdGeneratorOptions(1)
	idGenerator := idgen.NewDefaultIdGenerator(options)
	return &IdgenObject{idGenerator: idGenerator}
}

func (d *IdgenObject) NextId() int64 {
	return d.idGenerator.NewLong()
}

func (d *IdgenObject) ExtractTime(id int64) time.Time {
	return d.idGenerator.ExtractTime(id)
}
