package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "{{.Name}}_notloaded_go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

// THIS FILE IS GENERATED; DO NOT EDIT

// {{.Name}}NotLoaded returns an {{.Name}} which does not have a value;
// and in particular the "not loaded" type of a lack of a value.
//
// For example:
//
//	var x {{.Pkg}}.{{.Name}} = {{.Pkg}}.{{.Name}}NotLoaded()
func {{.Name}}NotLoaded() {{.Name}} {
	return {{.Name}}{}
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

