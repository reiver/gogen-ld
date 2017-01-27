package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "{{.Name}}_{{.Name}}_test.go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

{{if or (eq .Type "bool") (eq .Type "float32") (eq .Type "float64") (eq .Type "int8") (eq .Type "int16") (eq .Type "int32") (eq .Type "int64") (eq .Type "string") (eq .Type "time.Time") (eq .Type "uint8") (eq .Type "uint16") (eq .Type "uint32") (eq .Type "uint64") }}

// THIS FILE IS GENERATED; DO NOT EDIT

import (
{{range .Imports}}
	"{{.}}"
{{end}}
	"testing"
)

func Test{{.Name}}{{.Name}}(t *testing.T) {

	{
		var x {{.Name}}

		casted, err := x.{{.Name}}()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: (%T) %v; got value: %v", err, err, casted)
			return
		}
		if _, ok := err.(NotLoadedComplainer); !ok {
			t.Errorf("Expected error to fit \"ld.NotLoadedComplainer\", but actually didn't.")
			return
		}
	}

	{
		{{if eq .Type "bool"}}
		expected := {{.Type}}(true)
		{{else if or (eq .Type "float32") (eq .Type "float64") }}
		expected := {{.Type}}(2.2)
		{{else if or (eq .Type "int8") (eq .Type "int16") (eq .Type "int32") (eq .Type "int64") }}
		expected := {{.Type}}(-3)
		{{else if eq .Type "string"}}
		expected := {{.Type}}("FOUR")
		{{else if eq .Type "time.Time"}}
		expected := {{.Type}}(time.Date(2005, time.May, 15, 23, 0, 0, 0, time.UTC))
		{{else if or (eq .Type "uint8") (eq .Type "uint16") (eq .Type "uint32") (eq .Type "uint64") }}
		expected := {{.Type}}(3)
		{{else}}
		panic("THIS SHOULD NOT HAVE BEEN GENERATED!")
		{{end}}

		var x {{.Name}} = {{.Name}}Some(expected)

		actual, err := x.{{.Name}}()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected != actual {
			t.Errorf("Expected %v, but actually got %v.", expected, actual)
			return
		}
	}
}
{{end}}
`}

	registry := gendriver.Registry
	if nil == registry {
		panic(errNilRegistry)
	}

	if err := registry.Register(driver.NamePattern, driver); nil != err {
		panic(err)
	}
}
