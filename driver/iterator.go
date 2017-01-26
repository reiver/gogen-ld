package gendriver

import (
	"sync"
)

// Iterator is an iterator on Driver.
type Iterator interface {
	Datum() (string, Driver, error)
	Err() error
	Next() bool
}

type internalIterator struct {
	data []struct{
		Name   string
		Driver Driver
	}
	err   error
	indexPlusOne int
	mutex sync.RWMutex
}

func (receiver internalIterator) index() int {
	return receiver.indexPlusOne-1
}

func (receiver *internalIterator) copyFrom(m map[string]Driver) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil != receiver.data {
		return errInternalError
	}

	var data []struct{
		Name   string
		Driver Driver
	}

	for k,v := range m {
		datum := struct{
			Name   string
			Driver Driver
		}{
			Name:   k,
			Driver: v,
		}

		data = append(data, datum)
	}

	receiver.data = data

	return nil
}

func (receiver internalIterator) Datum() (string, Driver, error) {

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	if err := receiver.err; nil != err {
		return "", nil, err
	}

	data := receiver.data
	if nil == data {
		return "", nil, errInternalError
	}

	index := receiver.index()
	if len(data) <= index {
		return "", nil, errNotFound
	}

	datum := data[index]

	return datum.Name, datum.Driver, nil
}

func (receiver *internalIterator) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	return receiver.err
}

func (receiver *internalIterator) Next() bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if err := receiver.err; nil != err {
		return false
	}

	if data, index := receiver.data, receiver.index(); len(data) <= (1+index) {
		return false
	}

	receiver.indexPlusOne++


	return true
}
