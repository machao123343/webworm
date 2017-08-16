package middleware

import (
	"sync"
	"math"
)

//ID生成器的接口、
type IDGenerator interface {
	GetUint32() uint32 //获得一个uint32类型的ID
}

//创建ID生成器
func NewIdGenerator() IDGenerator {
	return  &cyclicIdGenerator{}
}

//id生成器的实现类型

type cyclicIdGenerator struct {
	sn uint32 //当前的ID
	ended bool //
	mutex sync.Mutex //互斥锁（保证生成ID和网页下载器的并发安全）
}

func (gen *cyclicIdGenerator) GetUint32() uint32 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.ended {
		defer func() {
			gen.ended = false
		}()//活用defer的机制
		gen.sn = 0
		return gen.sn
	}
	id := gen.sn
	if id < math.MaxUint32 {
		gen.sn++
	} else {
		gen.ended = true
	}
	return id
}

//ID生成器的接口类型2
type IdGenerator2 interface{
	GetUint64() uint64 //获取一个uint类型的ID
}

//创建ID生成器2
func NewIdGenerator2() IdGenerator2 {
	return &cyclicIdGenerator2{}
}

type cyclicIdGenerator2 struct {
	base   cyclicIdGenerator//基本的ID生成器
	cycleCount uint64 //基于uint32类型的取值范围的周期计数
}

func (gen *cyclicIdGenerator2) GetUint64() uint64 {
	var id64 uint64
	if gen.cycleCount%2 == 1 {
		id64 += math.MaxUint32
	}
	id32 := gen.base.GetUint32()//接收者的实现
	if id32 == math.MaxUint32 {
		gen.cycleCount++
	}
	id64 += uint64(id32)
	return id64
}

