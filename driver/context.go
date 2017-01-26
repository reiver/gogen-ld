package gendriver

// Context represents the data that is passed to a Driver's WriterTo method.
//
// Since the io.WriterTo returned from the Driver's WriterTo is meant to create
// a  single Go source code file, Context contains information that is relevant
// to that.
//
// More specifically, Context's `Pkg` field is the name of the package for
// the generated Go source code file.
//
// So, if `Pkg` had a value of "something", then the package declaration for
// the generated Go source code file should be:
//
//	package something
//
// Context's `Type` field is the name of the underlying type for the generated Go
// source code file.
//
// And Context's `Name` field is the name of the new type for the generated Go
// source code file.
//
// So (and note that this is just an example, and actual drivers could do this very
// very differently), if `Type` had a value of "int64" and `Name` had a value of
// "MyInt64", then you could get something like:
//
//	type MyInt64 struct {
//		// ...
//		value int64
//		// ...
//	}
type Context struct {
	Imports []string
	Pkg       string
	Type      string
	Name      string
}
