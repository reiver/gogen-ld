package main

import (
	"github.com/reiver/gogen-ld/driver"
)

func init() {

	driver := gendriver.SimpleDriver{
		NamePattern: "notloadedcomplainer.go",
		TextTemplate:
`// DO NOT EDIT
//
// MACHINE GENERATED BY THE FOLLOWING COMMAND
// gogen-ld --pkg={{.Pkg}} --type={{.Type}} --name={{.Name}}
package {{.Pkg}}

// THIS FILE IS GENERATED; DO NOT EDIT

type NotLoadedComplainer interface {
	error
	NotLoadedComplainer()
}

// THIS FILE IS GENERATED; DO NOT EDIT

type internalNotLoadedComplainer struct{}

// THIS FILE IS GENERATED; DO NOT EDIT

func (receiver internalNotLoadedComplainer) Error() string {
	return "Not Loaded"
}

// THIS FILE IS GENERATED; DO NOT EDIT

func (internalNotLoadedComplainer) NotLoadedComplainer() {
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
