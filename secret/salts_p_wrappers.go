package secret

func GenerateSalt512Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(512, workers)
}

func GenerateSalt1024Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(1024, workers)
}

func GenerateSalt2048Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(2048, workers)
}

func GenerateSalt4096Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(4096, workers)
}

func GenerateSalt8192Parallel(workers int) ([]byte, error) {
	return GenerateSaltParallel(8192, workers)
}
