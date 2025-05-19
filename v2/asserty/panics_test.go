package asserty_test

import (
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/v2/asserty"
	"github.com/fengdotdev/golibs-testing/assert"
)

// testing objects
type obj struct {
	Name string
}

type obj2 struct {
	name string
}

func NewOnobj2(name string) *obj2 {
	return &obj2{
		name: name,
	}
}

func TestAsertyPanic_True(t *testing.T) {

	t.Run("Should panic when false", func(t *testing.T) {
		didPanic := false

		func() {

			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertTrue(false) // should panic
		}()

		assert.True(t, didPanic)
	})

	t.Run("Should panic not panic", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertTrue(true) // should not panic
		}()

		assert.False(t, didPanic)
	})

}

func TestAsertyPanic_Value(t *testing.T) {

	t.Run("Should panic when values are different", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertValue(obj{Name: "test"}, obj{Name: "test2"}) // should panic
		}()

		assert.True(t, didPanic)
	})



	t.Run("Should panic when values are different2", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertValue(2,3) // should panic
		}()

		assert.True(t, didPanic)
	})

	t.Run("Should not panic ", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertValue(obj{Name: "test"}, obj{Name: "test"}) // should not panic
		}()

		assert.False(t, didPanic)
	})

	t.Run("Should not panic2 ", func(t *testing.T) {
		didPanic := false

		func() {
			defer func() {
				if r := recover(); r != nil {
					didPanic = true
				}
			}()

			asserty.AssertValue(1,1) // should not panic
		}()

		assert.False(t, didPanic)
	})

}
