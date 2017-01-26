package gendriver

import (
	"testing"
)

func TestInternalIteratorCopyFrom(t *testing.T) {

	tests := []struct{
		Map map[string]Driver
	}{
		{
			Map: nil,
		},
		{
			Map:  map[string]Driver{},
		},



		{
			Map:  map[string]Driver{
				"apple": SimpleDriver{TextTemplate:"Apple"},
			},
		},
		{
			Map:  map[string]Driver{
				"apple":  SimpleDriver{TextTemplate:"Apple"},
				"banana": SimpleDriver{TextTemplate:"Banana"},
			},
		},
		{
			Map:  map[string]Driver{
				"apple":  SimpleDriver{TextTemplate:"Apple"},
				"banana": SimpleDriver{TextTemplate:"Banana"},
				"cherry": SimpleDriver{TextTemplate:"Cherry"},
			},
		},
		{
			Map:  map[string]Driver{
				"apple":  SimpleDriver{TextTemplate:"Apple"},
				"banana": SimpleDriver{TextTemplate:"Banana"},
				"cherry": SimpleDriver{TextTemplate:"Cherry"},
				"date":   SimpleDriver{TextTemplate:"Date"},
			},
		},
	}


	for testNumber, test := range tests {
		var iterator internalIterator

		if data := iterator.data; nil != data {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, data)
			continue
		}

		if err := iterator.copyFrom(test.Map); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Map), len(iterator.data); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for datumNumber, datum := range iterator.data {
			key := datum.Name

			expectedDriver, ok := test.Map[key]
			if !ok {
				t.Errorf("For test #%d and datum #%d, expected key %q to exist, but actually didn't.", testNumber, datumNumber, key)
				t.Errorf("Actual keys:...")
				for _, d := range iterator.data {
					t.Errorf("\tKey: %q", d.Name)
				}
				continue
			}

			if expected, actual := expectedDriver, datum.Driver; expected != actual {
				t.Errorf("For test #%d and datum #%d, expected %#v, but actually got %#v.", testNumber, datumNumber, expected, actual)
				continue
			}
		}
	}
}
