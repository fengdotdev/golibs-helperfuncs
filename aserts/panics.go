package asserty

func AssertTrue(value bool) {
	if !value {
		panic("Assertion failed")
	}
}

func AssertValue(value interface{}, expected interface{}) {
	if value != expected {
		panic("Assertion failed")
	}
}
