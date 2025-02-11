package secret

import (
	"errors"
	"sync"
)

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
