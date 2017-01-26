package gendriver

import (
	"testing"
)

func TestInternalRegistrar(t *testing.T) {

	tests := []struct{
		Drivers []struct{
			Name   string
			Driver Driver
		}
	}{
		{
			Drivers: []struct{
				Name   string
				Driver Driver
			}{},
		},



		{
			Drivers: []struct{
				Name   string
				Driver Driver
			}{
				{
					Name:   "apple",
					Driver: SimpleDriver{TextTemplate:"Apple"},
				},
			},
		},
		{
			Drivers: []struct{
				Name   string
				Driver Driver
			}{
				{
					Name:   "apple",
					Driver: SimpleDriver{TextTemplate:"Apple"},
				},
				{
					Name:   "banana",
					Driver: SimpleDriver{TextTemplate:"Banana"},
				},
			},
		},
		{
			Drivers: []struct{
				Name   string
				Driver Driver
			}{
				{
					Name:   "apple",
					Driver: SimpleDriver{TextTemplate:"Apple"},
				},
				{
					Name:   "banana",
					Driver: SimpleDriver{TextTemplate:"Banana"},
				},
				{
					Name:   "cherry",
					Driver: SimpleDriver{TextTemplate:"Cherry"},
				},
			},
		},
	}


	TestLoop: for testNumber, test := range tests {
		var registry Registrar = new(internalRegistrar)

		if nil == registry {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, registry)
			continue
		}

		for driverNumber, driverData := range test.Drivers {
			if err := registry.Register(driverData.Name, driverData.Driver); nil != err {
				t.Errorf("For test #%d and driver #%d, did not expect an error, but actually got: (%T) %v", testNumber, driverNumber, err, err)
				continue
			}
		}


		iterator, err := registry.Iterator()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got: (%T) %v", testNumber, err, err)
			continue
		}
		if nil == iterator {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, iterator)
			continue
		}

		names   := []string{}
		drivers := []Driver{}
		actualIterations := 0

		for iterator.Next() {
			actualIterations++

			iterationNumber := actualIterations-1

			name, driver, err := iterator.Datum()
			if nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue TestLoop
			}

			names   = append(names,   name)
			drivers = append(drivers, driver)
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Drivers), actualIterations; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("Names:...")
			for nameNumber, name := range names {
				t.Errorf("\tname [%d] -> %q", nameNumber, name)
			}
			t.Errorf("Drivers:...")
			for driverNumber, driver := range drivers {
				t.Errorf("\tdriver [%d] -> %#v", driverNumber, driver)
			}
			continue
		}

		for driverNumber, driverData := range test.Drivers {
			expectedName := driverData.Name
			found := false
			foundIndex := -1
			for nameIndex, actualName := range names {
				if expected, actual := expectedName, actualName; expected == actual {
					found = true
					foundIndex = nameIndex
				}
			}
			if !found {
				t.Errorf("For test #%d and driver #%d, could not find driver %q.", testNumber, driverNumber, expectedName)
				continue
			}

			if expected, actual := driverData.Driver, drivers[foundIndex]; expected != actual {
				t.Errorf("For test #%d and driver #%d, expected %#v, but actually got %#v.", testNumber, driverNumber, expected, actual)
				continue
			}
		}
	}
}
