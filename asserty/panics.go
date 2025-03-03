package asserty

import "reflect"

func AssertTrue(value bool) {
	if !value {
		panic("Assertion failed")
	}
}

func AssertValue(value interface{}, expected interface{}) {
	if !reflect.DeepEqual(value, expected) {
		panic("Assertion failed")
	}
}

func AssertNoError(err error) {
	if err != nil {
		panic("Assertion NoError failed" + err.Error())
	}
}
