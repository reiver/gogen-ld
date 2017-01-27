package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "{{.Name}}_value.go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

// THIS FILE IS GENERATED; DO NOT EDIT

import (
{{range .Imports}}
	"{{.}}"
{{end}}
	"fmt"
)

// THIS FILE IS GENERATED; DO NOT EDIT

type {{.Name}}Value {{.Type}}

// THIS FILE IS GENERATED; DO NOT EDIT

func (receiver {{.Name}}Value) String() string {
	var casted interface{} = {{.Type}}(receiver)

	switch t := casted.(type) {
	case interface{ String()string }:
		return t.String()
	case {{.Type}}:
		return fmt.Sprintf("%v", casted)
	default:
		return fmt.Sprintf("%v", casted)
	}
}

// THIS FILE IS GENERATED; DO NOT EDIT

// {{.Name}}Matcher makes {{.Name}}Value fit the {{.Name}}Matcher interface.
func ({{.Name}}Value) {{.Name}}Matcher() {
	// Nothing here.
}

// THIS FILE IS GENERATED; DO NOT EDIT

// Nullable{{.Name}}Matcher makes {{.Name}}Value fit the Nullable{{.Name}}Matcher interface.
func ({{.Name}}Value) Nullable{{.Name}}Matcher() {
	// Nothing here.
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
