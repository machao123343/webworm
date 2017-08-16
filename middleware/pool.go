package middleware

import (
	"reflect"
	"fmt"
	"errors"
)

//实体的接口类型
type Entity interface {
	Id() uint32//ID的获取方法
}

type Pool interface{
	Take() (Entity, error)   //取出实体
	Return(entity Entity) error//归还实体
	Total() uint32            //实体池的容量
	Used() uint32          //实体池中已被使用的实体的容量
}

//创建实体池
func NewPool(
	total uint32,
	entityType reflect.Type,
	genEntity func() Entity) (bool,  error) {
	if total == 0 {
		errMsg :=
			fmt.Sprintf("The pool can not be initailized ! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	size := int(total)
	container := make(chan Entity, size)
	idContainner := make(map[uint32]bool)
	for  i := 0; i < size; i++ {
	}

}



