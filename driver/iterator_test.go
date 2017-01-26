package gendriver

import (
	"testing"
)

func TestIteratorInternalIterator(t *testing.T) {
	var x Iterator = new(internalIterator) // THIS SINGLE LINE OF CODE IS WHAT WE REALLY CARE ABOUT IN THIS TEST.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
