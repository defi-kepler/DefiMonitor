package event_manage

import (
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
	"sync"
)

var sMap sync.Map

func CreateEventManageFactory() *eventManage {

	return &eventManage{}
}

type eventManage struct {
}

func (e *eventManage) Set(key string, keyFunc func(args ...interface{})) bool {

	if _, exists := e.Get(key); exists == false {
		sMap.Store(key, keyFunc)
		return true
	} else {
		variable.ZapLog.Info(my_errors.ErrorsFuncEventAlreadyExists + " , key :" + key)
	}
	return false
}

func (e *eventManage) Get(key string) (interface{}, bool) {
	if value, exists := sMap.Load(key); exists {
		return value, exists
	}
	return nil, false
}

func (e *eventManage) Call(key string, args ...interface{}) {
	if valueInterface, exists := e.Get(key); exists {
		if fn, ok := valueInterface.(func(args ...interface{})); ok {
			fn(args...)
		} else {
			variable.ZapLog.Error(my_errors.ErrorsFuncEventNotCall + ", key:" + key + ", ")
		}

	} else {
		variable.ZapLog.Error(my_errors.ErrorsFuncEventNotRegister + ", key:" + key)
	}
}

func (e *eventManage) Delete(key string) {
	sMap.Delete(key)
}

func (e *eventManage) FuzzyCall(keyPre string) {

	sMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPre) {
				e.Call(keyName)
			}
		}
		return true
	})
}
