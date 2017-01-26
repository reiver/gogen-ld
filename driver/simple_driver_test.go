package gendriver

import (
	"bytes"

	"testing"
)

func TestSimpleDriver(t *testing.T) {

	tests := []struct{
		Driver   Driver
		Context  Context
		Expected string
	}{
		{
			Driver: SimpleDriver{
				TextTemplate: `package {{.Pkg}}`+"\n"+`type {{.Name}} {{.Type}}`+"\n",
			},
			Context: Context{
				Pkg: "apple",
				Name:    "MyInt64",
				Type:    "int64",
			},
			Expected: "package apple"+"\n"+"type MyInt64 int64"+"\n",
		},
		{
			Driver: SimpleDriver{
				TextTemplate: `package {{.Pkg}}`+"\n"+`type {{.Name}} struct {`+"\n"+"\t"+`value {{.Type}}`+"\n"+`}`+"\n",
			},
			Context: Context{
				Pkg: "banana",
				Name:    "SuperString",
				Type:    "string",
			},
			Expected: "package banana"+"\n"+"type SuperString struct {"+"\n"+"\tvalue string"+"\n"+"}"+"\n",
		},
		{
			Driver: SimpleDriver{
				TextTemplate:
`package {{.Pkg}}

import (
	"fmt"
)

type {{.Name}} {{.Type}}

func (receiver {{.Name}}) String() string {
	return fmt.Sprintf("{{.Name}}(%v)", receiver)
}
`,
			},
			Context: Context{
				Pkg: "cherry",
				Name:    "Nada",
				Type:    "struct{}",
			},
			Expected:
`package cherry

import (
	"fmt"
)

type Nada struct{}

func (receiver Nada) String() string {
	return fmt.Sprintf("Nada(%v)", receiver)
}
`,
		},
	}


	for testNumber, test := range tests {

		writerTo, err := test.Driver.WriterTo(test.Context)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var buffer bytes.Buffer

		n, err := writerTo.WriteTo(&buffer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if expected, actual := int64(len(test.Expected)), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("EXPECTED String: %q", test.Expected)
			t.Errorf("ACTUAL String:   %q", buffer.String())
			continue
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("EXPECTED: %q", expected)
			t.Errorf("ACTUAL:   %q", actual)
			continue
		}
	}
}
