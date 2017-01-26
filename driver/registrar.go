package gendriver

import (
	"sync"
)

type Registrar interface {
	Iterator() (Iterator, error)
	Register(string, Driver) error
}

type internalRegistrar struct {
	m map[string]Driver
	mutex sync.RWMutex
}

func (receiver *internalRegistrar) Iterator() (Iterator, error) {

	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	m := receiver.m
	if nil == m {
		 var iterator Iterator = new(internalIterator)
		return iterator, nil
	}

	var it internalIterator
	if err := it.copyFrom(m); nil != err {
		return nil, errInternalError
	}

	var iterator Iterator = &it

	return iterator, nil
}

func (receiver *internalRegistrar) Register(name string, driver Driver) error {

	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.m {
		receiver.m = map[string]Driver{}
	}

	if _, ok := receiver.m[name]; ok {
		return &internalFoundComplainer{name}
	}

	receiver.m[name] = driver

	return nil
}
