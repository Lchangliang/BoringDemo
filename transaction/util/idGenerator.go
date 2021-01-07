package util

import (
	"sync"
	"sync/atomic"
	"strconv"
)

type IdGenerator interface {
	GenerateTransactionId() string
}

type atomicIdGenerator struct {
	i int64
}

func (generator *atomicIdGenerator) GenerateTransactionId() string {
	i := atomic.AddInt64(&generator.i, 1)
	return strconv.FormatInt(i, 10)
}

var once sync.Once
var generator IdGenerator

func GetIdGenerator() IdGenerator {
	once.Do(func() {
		generator = &atomicIdGenerator{0}
	})
	return generator
}