package middleware

import (
	"reflect"
	"fmt"
	"errors"
	"Mouse/container"
	"sync"
)

//实体的接口类型
type Entity interface {
	Id() uint32//ID的获取方法
}

//实体池的接口类型
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
	idContainer := make(map[uint32]bool)
	for  i := 0; i < size; i++ {
		newEntity := genEntity()
		if entityType != reflect.TypeOf(newEntity) {
			errMsg :=
				fmt.Sprintf("The type of result of function genEntity() is NOT %s!\n", entityType)
		return nil, errors.New(errMsg)
		}
		container <- newEntity
		idContainer[newEntity.Id()] = true
	}
	pool := &myPool{
		total:    total,
		etype:    entityType,
	}

}

//实体池的实现类型
type myPool struct{
	total   uint32
	etype   reflect.Type //池中实体的类型
	genEntity func() Entity //池中实体的生成函数
	container chan Entity  //实体容器
	idContainer map[uint32]bool //实体ID的容器
	mutex      sync.Mutex    //针对实体ID容器操作的互斥锁
}

func (pool *myPool) Take() (Entity, error) {
	entity, ok := <- pool.container
	if !ok {
		return nil, errors.New("The inner container is invalid")
	}
	pool.mutex.Lock()//锁是针对于对象而言的
	defer pool.mutex.Unlock()
	pool.idContainer[entity.Id()] = false
	return entity, nil
}



