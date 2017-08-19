package analyzer

import (
	"reflect"
	"fmt"
	"errors"

	mdw "webworm/middleware"

)

//todo：学会封装池的建立方法

//生成分析器的函数类型
type GenAnalyzer func() Analyzer

//分析器池的接口类型
type AnalyzerPool interface {
	Take() (Analyzer, error)
	Return(analyzer Analyzer) error
	Total() uint32
	Used() uint32
}

type myAnalyzerPool struct {
	pool mdw.Pool   //实体池
	etype reflect.Type   //实体池内的类型
}

func NewAnalyzerPool(
	total uint32,
	gen GenAnalyzer) (AnalyzerPool, error) {
	etype := reflect.TypeOf(gen())
	genEntity := func() mdw.Entity {
		return gen()
	}
	pool, err := mdw.NewPool(total, etype, genEntity)
	if err != nil {
		return nil, err
	}
	dlpool := &myAnalyzerPool{pool: pool, etype: etype}
	return dlpool, nil
}

func (spdpool *myAnalyzerPool) Take() (Analyzer, error) {
	entity, err := spdpool.pool.Take()
	if err != nil {
		return nil ,err
	}

	analyzer, ok := entity.(Analyzer)
	if !ok {
		errMsg := fmt.Sprintf("The type of entity is NOT %s!\n", spdpool.etype)
		panic(errors.New(errMsg))
	}

	return analyzer, nil
}

func (spdpool *myAnalyzerPool) Return(analyzer Analyzer) error {
	return spdpool.pool.Return(analyzer)
}

func (spdpool *myAnalyzerPool) Total() uint32 {
	return spdpool.pool.Total()
}

func (spdpool *myAnalyzerPool) Used() uint32 {
	return spdpool.Used()
}

//新建池生成对象， 分析吃接口作为新建池对象的内嵌方法用以外部调用。



