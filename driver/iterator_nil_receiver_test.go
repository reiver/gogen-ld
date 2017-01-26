package gendriver

import (
	"testing"
)

func TestInternalIteratorCopyFromNilReceiver(t *testing.T) {

	iterator := (*internalIterator)(nil)

	if expected, actual := errNilReceiver, iterator.copyFrom(nil); expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := errNilReceiver, iterator.copyFrom( map[string]Driver{} ); expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}

func TestInternalIteratorErrNilReceiver(t *testing.T) {

	var iterator Iterator = (*internalIterator)(nil)

	if expected, actual := errNilReceiver, iterator.Err(); expected != actual {
		t.Errorf("Expected (%T) %v, but actually got (%T) %v.", expected, expected, actual, actual)
		return
	}
}

func TestInternalIteratorNextNilReceiver(t *testing.T) {

	var iterator Iterator = (*internalIterator)(nil)

	if expected, actual := false, iterator.Next(); expected != actual {
		t.Errorf("Expected %t, but actually got %t.", expected, actual)
		return
	}
}
