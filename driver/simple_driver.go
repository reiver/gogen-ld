package gendriver

import (
	"bytes"
	"io"
	"text/template"
)

// SimpleDriver is a helper, that makes creating a driver easier, by providing
// a simple way of creating a driver that used Go's "text/template" package.
//
// Example Usage
//
//	var driver Driver = SimpleDriver{
//		NamePattern: "slice_{{.Name}}.go",
//		TextTemplate:
//	`package {{.Package}}
//	
//	import (
//		"fmt"
//	)
//	
//	type {{.Name}} []{{.Type}}
//	
//	func (receiver {{.Name}}) String() string {
//		return fmt.Sprintf("{{.Name}}(%v)", receiver)
//	}
//	`
//	}
type SimpleDriver struct {
	NamePattern  string
	TextTemplate string
	tmpl *template.Template
}

func (receiver *SimpleDriver) init() error {

	if nil == receiver {
		return errNilReceiver
	}

	if nil != receiver.tmpl {
		return nil
	}

	tmpl := template.New(receiver.NamePattern)
	if nil == tmpl {
		return errInternalError
	}

	{
		var err error

		tmpl, err = tmpl.Parse(receiver.TextTemplate)
		if nil != err {
			return err
		}
		if nil == tmpl {
			return errInternalError
		}
	}

	receiver.tmpl = tmpl

	return nil

}

func (receiver SimpleDriver) WriterTo(ctx Context) (io.WriterTo, error) {

	if err := receiver.init(); nil != err {
		return nil, errInternalError
	}

	tmpl := receiver.tmpl
	if nil == tmpl {
		return nil, errInternalError
	}

	var writerTo io.WriterTo = internalSimpleDriverWriterTo{
		ctx:  ctx,
		tmpl: tmpl,
	}

	return writerTo, nil
}

type internalSimpleDriverWriterTo struct {
	ctx  Context
	tmpl *template.Template
}

func (receiver internalSimpleDriverWriterTo) WriteTo(w io.Writer) (int64, error) {

	if nil == w {
		return 0, errNilWriter
	}

	tmpl := receiver.tmpl
	if nil == tmpl {
		return 0, errInternalError
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, receiver.ctx); nil != err {
		return 0, err
	}

	n, err := w.Write(buffer.Bytes())

	n64 := int64(n)

	return n64, err
}
