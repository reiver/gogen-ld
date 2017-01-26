package gendriver

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
	errNilWriter     = errors.New("Nil Writer")
	errNotFound      = errors.New("Not Found")
)
