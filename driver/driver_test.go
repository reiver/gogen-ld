package gendriver

import (
	"testing"
)

func TestDriverSimpleDriver(t *testing.T) {

	var driver Driver = SimpleDriver{} // THIS SINGLE LINE OF CODE IS WHAT WE REALLY CARE ABOUT IN THIS TEST.

	if nil == driver {
		t.Errorf("This should never happen.")
	}
}
