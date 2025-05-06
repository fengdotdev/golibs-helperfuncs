package chile

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const (
	RutLowerBound = 1000000
	RutUpperBound = 99999999
)

// ValidateRUT validates a chilean RUT (Rol Único Tributario)
// RUT is a unique identifier for people and companies in Chile
// It is composed by a number and a verification digit example: 12345678-9
// The verification digit is calculated using the modulo 11 algorithm
// The RUT can be validated by checking the verification digit
// The RUT can be formatted with dots and dashes example: 12.345.678-9

// GenerateRut generates a random RUT (Rol Único Tributario) for testing purposes
// The RUT is composed by a number and a verification digit example: 12345678-9
// The verification digit is calculated using the modulo 11 algorithm
// The RUT can be validated by checking the verification digit
func GenerateRut() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	rutNum := rand.Intn(RutUpperBound-RutLowerBound) + RutLowerBound
	validationDigit := RutCalculateValidationDigit(rutNum)
	return strconv.Itoa(rutNum) + "-" + validationDigit
}

// the firstDigit and secondDigit are the first two digits of the RUT example someone born in 1989 could be 16
func GenerateRutFrom2First(firstDigit, secondDigit int) string {

	if firstDigit < 1 || firstDigit > 9 || secondDigit < 0 || secondDigit > 9 {
		return ""
	}
	rand.Seed(uint64(time.Now().UnixNano()))
	rutNum := rand.Intn(1000000) + (firstDigit*10+secondDigit)*1000000
	validationDigit := RutCalculateValidationDigit(rutNum)
	return strconv.Itoa(rutNum) + "-" + validationDigit
}

func ValidateRut(rut string) bool {
	cleanRut := RutRemoveDashesAndDots(rut)
	cleanRut = RutKUpperCase(cleanRut)

	match, _ := regexp.MatchString(`^[0-9]{7,8}[0-9K]$`, cleanRut)
	if !match {
		return false
	}

	rutNum, err := strconv.Atoi(cleanRut[:len(cleanRut)-1])
	if err != nil {
		return false
	}

	validationDigit := strings.ToUpper(cleanRut[len(cleanRut)-1:])
	expected := RutCalculateValidationDigit(rutNum)

	return validationDigit == expected
}

func RutRemoveDashesAndDots(rut string) string {
	rut = strings.ReplaceAll(rut, ".", "")
	rut = strings.ReplaceAll(rut, "-", "")
	return rut
}

func RutKUpperCase(rut string) string {
	return strings.ToUpper(rut)
}

// use K instead of k (uppercase)
func RutCalculateValidationDigit(rut int) string {
	adder := 0
	multiplier := 2

	// modulo 11 algorithm
	for rut > 0 {
		digit := rut % 10
		adder += digit * multiplier
		multiplier++
		if multiplier > 7 {
			multiplier = 2
		}
		rut = rut / 10
	}

	remaninder := adder % 11
	validationDigit := 11 - remaninder

	switch validationDigit {
	case 11:
		return "0"
	case 10:
		return "K"
	default:
		return strconv.Itoa(validationDigit)
	}
}
