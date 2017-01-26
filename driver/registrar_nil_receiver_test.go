package gendriver

import (
	"testing"
)

func TestInternalRegistrarIteratorNilReceiver(t *testing.T) {

	var registrar Registrar = (*internalRegistrar)(nil)

	iterator, err := registrar.Iterator()

	if nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}
	if nil != iterator {
		t.Errorf("Expected nil iterator, but did not actually get that: (%T) %v", iterator, iterator)
		return
	}

	if expected, actual := errNilReceiver, err; expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}

func TestInternalRegistrarRegisterNilReceiver(t *testing.T) {

	var registrar Registrar = (*internalRegistrar)(nil)

	{
		err := registrar.Register("", nil)
		if nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
			return
		}
	}

	{
		err := registrar.Register("something", SimpleDriver{})
		if nil == err {
			t.Errorf("Expected an error, but did not actually got one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
			return
		}
	}
}
