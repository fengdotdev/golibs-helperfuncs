package chile_test

import (
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/chile"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestValidateRut(t *testing.T) {
	validRut := "5.126.663-3"
	assert.True(t, chile.ValidateRut(validRut))


	invalidRut := "1.234.567-8"
	assert.False(t, chile.ValidateRut(invalidRut))

}
