package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "nullable_{{.Name}}_some.go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

{{range .Imports}}
import "{{.}}"
{{end}}

// THIS FILE IS GENERATED; DO NOT EDIT

// Nullable{{.Name}}Some returns a Nullable{{.Name}} with a value.
//
// For example:
//
//	var x {{.Pkg}}.Nullable{{.Name}} = {{.Pkg}}.Nullable{{.Name}}Some(41)
func Nullable{{.Name}}Some(v {{.Type}}) Nullable{{.Name}} {
	return Nullable{{.Name}} {
		value: v,
		loaded: true,
		null: false,
	}
}
`}

	registry := gendriver.Registry
	if nil == registry {
		panic(errNilRegistry)
	}

	if err := registry.Register(driver.NamePattern, driver); nil != err {
		panic(err)
	}
}
