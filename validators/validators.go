package validators

import (
	"errors"
	"log"
	"regexp"
)

// example of a validator function
// this function should return true if the input is valid, and false otherwise
// it should also return an error, as the reason for the failure

// considering that some valitators can do the same
// validation in different ways, we can have multiple functions
// that do the same thing, but with different names at the end example _byletter
// may do a benchmark to see which one is faster or comment the reason for prefer one over the other like some use of a standard

// also considering make a default validator function an keep the other ones

// default validator function for foo
func FooValidator(foo string) (bool, error) {
	return FooValidator_byEqual(foo)
}

// this is a simple example of a validator function for foo
// it checks if the input is equal to "foo"
func FooValidator_byEqual(foo string) (bool, error) {
	log.Println("FooValidator is not for production use")

	if foo == "foo" {
		return true, nil
	}
	return false, errors.New("foo is not valid")
}

// this is a simple example of a validator function for foo
// it checks if the input is equal to "foo" comparing by letter
func FooValidator_byletter(foo string) (bool, error) {

	log.Println("FooValidator_byletter is not for production use")

	if len(foo) == 3 && foo[0] == 'f' && foo[1] == 'o' && foo[2] == 'o' {
		return true, nil
	}

	return false, errors.New("foo is not valid")
}

// this is a simple example of a validator function for foo
// it checks if the input is equal to "foo" using regex
func FooValidator_byRegex(s string) (bool, error) {

	log.Println("FooValidator_byRegex is not for production use")
	match, err := regexp.MatchString(`^foo$`, s)
	if err != nil {
		return false, err
	}
	return match, nil
}
