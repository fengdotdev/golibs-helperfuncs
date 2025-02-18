package secret

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

// Generates a cryptografic secure salt of 1024 bytes
func GenerateSalt1024() ([]byte, error) {
	return GenerateSalt(1024)
}

// Generates a cryptografic secure salt of 2048 bytes
func GenerateSalt2048() ([]byte, error) {
	return GenerateSalt(2048)
}

// Generates a cryptografic secure salt of 4096 bytes
func GenerateSalt4096() ([]byte, error) {
	return GenerateSalt(4096)
}

func GenerateSalt8192() ([]byte, error) {
	return GenerateSalt(8192)
}
