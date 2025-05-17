package ident

import (
	"crypto/sha256"
	"math/rand/v2"

	"github.com/google/uuid"
)

func RandomNumericIDBetween(min, max int) int {
	return min + rand.IntN(max-min)
}

// V4 UUID
func RamdomUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

// V5 UUID
// the id is always the same for the same namespace and seed value
func DeterministicUUID(namespace, seed string) string {
	sha256namespace := sha256.Sum256([]byte(namespace))

	namespaceUUID := uuid.NewSHA1(uuid.NameSpaceOID, []byte(sha256namespace[:32]))

	deterministicUUID := uuid.NewSHA1(namespaceUUID, []byte(seed))
    return deterministicUUID.String()

}
