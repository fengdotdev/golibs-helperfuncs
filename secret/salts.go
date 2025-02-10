package secret

import (
	"crypto/rand"
	"errors"
	"sync"
)

// Generates a cryptografic secure salt of the given size
func GenerateSalt(size int) ([]byte, error) {
	if size < 1 {
		return nil, errors.New("size must be greater than 0")
	}
	salt := make([]byte, size)
	_, err := rand.Read(salt) // Use crypto/rand to generate cryptographically secure bytes
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GenerateSaltParallel(size, workers int) ([]byte, error) {
	if workers < 1 {
		return nil, errors.New("workers must be greater than 0")
	}
	if size < 1 {
		return nil, errors.New("size must be greater than 0")
	}
	if workers > size {
		// Alternatively, you could allow more workers than bytes,
		// but some might just generate 0 bytes. For safety:
		workers = size
	}

	// Distribute `size` as evenly as possible across workers.
	// e.g. if size=10, workers=3 => sizes: 4,3,3
	base := size / workers
	remainder := size % workers

	// Prepare array to hold each worker's salt
	salts := make([][]byte, workers)

	var wg sync.WaitGroup
	wg.Add(workers)

	cancel := make(chan struct{})
	var once sync.Once
	var retErr error

	// Start goroutines
	offset := 0
	for i := 0; i < workers; i++ {
		// Determine how many bytes this worker should generate
		chunkSize := base
		if i < remainder {
			chunkSize++
		}

		go func(idx, sz, start int) {
			defer wg.Done()

			// If an error has already occurred, we can skip work or short-circuit:
			select {
			case <-cancel:
				// Another goroutine has failed
				return
			default:
				// Continue
			}

			salt, err := GenerateSalt(sz)
			if err != nil {
				once.Do(func() {
					retErr = err
					close(cancel)
				})
				return
			}
			salts[idx] = salt

		}(i, chunkSize, offset)

		offset += chunkSize
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check if any goroutine signaled an error
	select {
	case <-cancel:
		return nil, errors.New("error generating salt: " + retErr.Error())
	default:
		// No errors, continue
	}

	// Combine all slices into a single slice of length `size`
	salt := make([]byte, size)
	pos := 0
	for i := 0; i < workers; i++ {
		copy(salt[pos:], salts[i])
		pos += len(salts[i])
	}

	return salt, nil
}

// Generates a cryptografic secure salt of 16 bytes
func GenerateSalt16() ([]byte, error) {
	return GenerateSalt(16)
}

// Generates a cryptografic secure salt of 32 bytes
func GenerateSalt32() ([]byte, error) {
	return GenerateSalt(32)
}

// Generates a cryptografic secure salt of 64 bytes
func GenerateSalt64() ([]byte, error) {
	return GenerateSalt(64)
}

// Generates a cryptografic secure salt of 128 bytes
func GenerateSalt128() ([]byte, error) {
	return GenerateSalt(128)
}

// Generates a cryptografic secure salt of 256 bytes
func GenerateSalt256() ([]byte, error) {
	return GenerateSalt(256)
}

// Generates a cryptografic secure salt of 512 bytes
func GenerateSalt512() ([]byte, error) {
	return GenerateSalt(512)
}

func GenerateSalt512Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(512, workers)
}

// Generates a cryptografic secure salt of 1024 bytes
func GenerateSalt1024() ([]byte, error) {
	return GenerateSalt(1024)
}

func GenerateSalt1024Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(1024, workers)
}

// Generates a cryptografic secure salt of 2048 bytes
func GenerateSalt2048() ([]byte, error) {
	return GenerateSalt(2048)
}

func GenerateSalt2048Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(2048, workers)
}

// Generates a cryptografic secure salt of 4096 bytes
func GenerateSalt4096() ([]byte, error) {
	return GenerateSalt(4096)
}

func GenerateSalt4096Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(4096, workers)
}

func GenerateSalt8192() ([]byte, error) {
	return GenerateSalt(8192)
}

func GenerateSalt8192Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(8192, workers)
}
