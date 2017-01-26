package gendriver

import (
	"io"
)

// Driver represents a driver. A driver has a simple WriterTo method that
// returns a io.WriterTo. The returned io.WriterTo is meant to create a
// single Go source code file.
//
// (Conceptually, there would be many drivers, each that would create
// a single Go source code file. But collectively, they would create a
// set of Go source code files.)
type Driver interface {
	WriterTo(context Context) (io.WriterTo, error)
}
