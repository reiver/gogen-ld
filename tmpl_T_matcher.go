package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "{{.Name}}_matcher.go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

// THIS FILE IS GENERATED; DO NOT EDIT

type {{.Name}}Matcher interface {
	{{.Name}}Matcher()
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
