package main

import (
	"github.com/reiver/gogen-ld/driver"
	"github.com/reiver/go-tmpl"

	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var flags struct {
		Imports string
		Pkg     string
		Type    string
		Name    string
	}
	var imports []string
	{
		flag.StringVar(&flags.Imports, "imports", "",     "Extra imports. Ex: --imports=")
		flag.StringVar(&flags.Pkg,     "pkg",     "main", "Package name. Ex: --pkg=main")
		flag.StringVar(&flags.Type,    "type",    "",     "Type. Ex: --type=int64")
		flag.StringVar(&flags.Name,    "name",    "",     "Name of generated type. Ex: --name=Int64")

		flag.Parse()

		if "" == flags.Name {
			fmt.Fprintf(os.Stderr, "ERROR: \"name\" may not be empty.\n")
			return
		}
		if "" == flags.Type {
			fmt.Fprintf(os.Stderr, "ERROR: \"type\" may not be empty.\n")
			return
		}

		if "" == flags.Pkg {
			flags.Pkg = "main"
		}

		if "" != flags.Imports {
			imports = strings.Split(flags.Imports, ",")
		}
	}

	registry := gendriver.Registry
	if nil == registry {
		fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could not find registry.\n")
		return
	}

	iterator, err := registry.Iterator()
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Could get an iterator from the registry, because: %v\n", err)
		return
	}
	if nil == iterator {
		fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Received a bad iterator from the registry.\n")
		return
	}

	tmplNumber := 0
	for iterator.Next() {
		tmplNumber++

		name, driver, err := iterator.Datum()
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Problem getting next driver: %v\n", err)
			return
		}
		if nil == driver {
			fmt.Fprintf(os.Stderr, "ERROR: Problem getting next driver: Nil Driver.\n")
			return
		}

		fileNameTemplate := "gen_ld_"+name
		fileName := tmpl.Sprintt(fileNameTemplate, struct{
			Name string
		}{
			Name: strings.ToLower(flags.Name),
		})

		fmt.Fprintf(os.Stdout, "Creating %q... ", fileName)

		writerTo, err := driver.WriterTo(gendriver.Context{
			Imports:    imports,
			Pkg:  flags.Pkg,
			Name: flags.Name,
			Type: flags.Type,
		})
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Problem setting up to write file %q: %v\n", fileName, err)
			return
		}
		if nil == writerTo {
			fmt.Fprintf(os.Stderr, "ERROR: Problem setting up to write file %q: Nil io.WriterTo\n", fileName)
			return
		}

		func() {

			file, err := os.Create(fileName)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR (initially) creating (empty) file %q for template #%d: %v\n", fileName, tmplNumber, err)
			}
			defer file.Close()


			_, err = writerTo.WriteTo(file)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR creating template #%d as file %q: %v\n", tmplNumber, fileName, err)
			}
			fmt.Fprintf(os.Stdout, "Done.\n")
		}()

	}
	if err := iterator.Err(); nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: This should not happen. Received an error when iterating through drivers: %v.\n", err)
		return
	}
}
