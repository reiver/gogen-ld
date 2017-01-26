package gendriver

import (
	"fmt"
)

// An error that fits FoundComplainer could be returned from a Registry.Register() method,
// in a case where a driver has already been registered for that name.
//
// Example Usage
//
//	err := gendriver.Registry.Register(driverName, driver)
//	
//	switch complainer := err.(type) {
//	case FoundComplainer:
//		//@TODO
//	default:
//		//@TODO
//	}
type FoundComplainer interface {
	error
	FoundComplainer()
}

type internalFoundComplainer struct{
	name string
}

func (receiver internalFoundComplainer) Error() string {
	return fmt.Sprintf("Found: %q", receiver.name)
}

func (internalFoundComplainer) FoundComplainer() {
	// Nothing here.
}
