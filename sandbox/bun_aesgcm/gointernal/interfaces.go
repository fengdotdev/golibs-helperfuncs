package gointernal

import (
	traits "github.com/fengdotdev/golibs-traits"
)

type AESGCM_GeneratorInterface interface {
	GenerateKey() ([]byte, error)                    // 32 bytes random key
	GenerateIV() ([]byte, error)                     // 12 bytes random iv or nonce
	GenerateIVFromSeed(seed []byte) ([]byte, error)  // 12 bytes iv or nonce from seed
	GenerateKeyFromSeed(seed []byte) ([]byte, error) // 32 bytes key from seed
}

type AESGCM_HelperInterface interface {
	Getkey() []byte
	GetKey64() string
	GetIV() []byte
	GetIV64() string
}

type AESGCM_ObjectInterface interface {
	traits.JSONTrait
	AESGCM_DataInterface
	AESGCM_CypherInterface
}

type AESGCM_DataInterface interface {
	Data() ([]byte, error)
	IV() ([]byte, error)
	AdditionalData() ([]byte, error)
	//AuthTag() ([]byte, error)
	IV64() string
	Data64() string
	AdditionalData64() string
}

type AESGCM_CypherInterface interface {
	Decrypt(helper AESGCM_HelperInterface) ([]byte, error)
}

// TODO MAKE THIS PART OF TRAITS
type SecretTrait interface {
	GetResource() (any, error) 
	GetResource64() (string, error)
	IsValid() bool
	IsClosed() bool
	Open(password string) error
	Close(password string) error
}
