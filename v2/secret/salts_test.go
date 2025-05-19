package secret_test

import (
	"testing"

	"github.com/fengdotdev/golibs-helperfuncs/v2/secret"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestSalts_GenerateSalt(t *testing.T) {
	lowerBound := 1
	upperBound := 100000
	times := 100000

	current := lowerBound
	for i := 0; i < times; i++ {
		if current > upperBound {
			break
		}
		size := current
		salt, err := secret.GenerateSalt(size)
		assert.NoError(t, err)
		assert.Equal(t, size, len(salt))
		current++

	}

}

func TestSalts_GenerateSaltParallels(t *testing.T) {

	lowerBound := 512
	upperBound := 1000000
	times := 10000000

	current := lowerBound
	for i := 0; i < times; i++ {
		if current > upperBound {
			break
		}
		size := current
		workers := 10
		salt, err := secret.GenerateSaltParallel(size, workers)
		assert.NoError(t, err)
		assert.Equal(t, size, len(salt))
		current = current + 32

	}

}

func TestSalts_GenerateSalt16(t *testing.T) {
	salt, err := secret.GenerateSalt16()
	assert.NoError(t, err)
	assert.Equal(t, 16, len(salt))
}

func TestSalts_GenerateSalt32(t *testing.T) {
	salt, err := secret.GenerateSalt32()
	assert.NoError(t, err)
	assert.Equal(t, 32, len(salt))
}

func TestSalts_GenerateSalt64(t *testing.T) {
	salt, err := secret.GenerateSalt64()
	assert.NoError(t, err)
	assert.Equal(t, 64, len(salt))
}

func TestSalts_GenerateSalt128(t *testing.T) {
	salt, err := secret.GenerateSalt128()
	assert.NoError(t, err)
	assert.Equal(t, 128, len(salt))
}

func TestSalts_GenerateSalt256(t *testing.T) {
	salt, err := secret.GenerateSalt256()
	assert.NoError(t, err)
	assert.Equal(t, 256, len(salt))
}

func TestSalts_GenerateSalt512(t *testing.T) {
	salt, err := secret.GenerateSalt512()
	assert.NoError(t, err)
	assert.Equal(t, 512, len(salt))
}

func TestSalts_GenerateSalt512Parallel(t *testing.T) {
	workers := 10
	salt, err := secret.GenerateSalt512Parallel(workers)
	assert.NoError(t, err)
	assert.Equal(t, 512, len(salt))
}

func TestSalts_GenerateSalt1024(t *testing.T) {
	salt, err := secret.GenerateSalt1024()
	assert.NoError(t, err)
	assert.Equal(t, 1024, len(salt))
}

func TestSalts_GenerateSalt1024Parallel(t *testing.T) {
	workers := 10
	salt, err := secret.GenerateSalt1024Parallel(workers)
	assert.NoError(t, err)
	assert.Equal(t, 1024, len(salt))
}

func TestSalts_GenerateSalt2048(t *testing.T) {
	salt, err := secret.GenerateSalt2048()
	assert.NoError(t, err)
	assert.Equal(t, 2048, len(salt))
}

func TestSalts_GenerateSalt2048Parallel(t *testing.T) {
	workers := 10
	salt, err := secret.GenerateSalt2048Parallel(workers)
	assert.NoError(t, err)
	assert.Equal(t, 2048, len(salt))
}

func TestSalts_GenerateSalt4096(t *testing.T) {
	salt, err := secret.GenerateSalt4096()
	assert.NoError(t, err)
	assert.Equal(t, 4096, len(salt))
}

func TestSalts_GenerateSalt4096Parallel(t *testing.T) {

	workers := 10
	salt, err := secret.GenerateSalt4096Parallel(workers)
	assert.NoError(t, err)
	assert.Equal(t, 4096, len(salt))
}

func TestSalts_GenerateSalt8192(t *testing.T) {
	salt, err := secret.GenerateSalt8192()
	assert.NoError(t, err)
	assert.Equal(t, 8192, len(salt))
}

func TestSalts_GenerateSalt8192Parallel(t *testing.T) {
	workers := 10
	salt, err := secret.GenerateSalt8192Parallel(workers)
	assert.NoError(t, err)
	assert.Equal(t, 8192, len(salt))
}
