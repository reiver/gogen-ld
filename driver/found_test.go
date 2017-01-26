package gendriver

import (
	"testing"
)

func TestFoundComplainer(t *testing.T) {

	var complainer FoundComplainer = internalFoundComplainer{} // THIS SINGLE LINE OF CODE IS WHAT WE REALLY CARE ABOUT IN THIS TEST.

	if nil == complainer {
		t.Errorf("This should never happen.")
	}
}
