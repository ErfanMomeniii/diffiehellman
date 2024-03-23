package diffiehellman

import (
	"math/rand"
	"time"
)

const defaultPrime = 10009

type DiffieHellman struct {
	primeNumber   int64
	primitiveRoot int64
	privateKey    int64
	publicKey     int64
}

func isPrime(number int64) bool {
	var i int64
	for i = 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func isPrimitive(candidate int64, primeNumber int64) bool {
	m := make(map[int64]int64)

	var i int64
	modP := candidate
	for i = 1; i < primeNumber; i++ {
		_, ok := m[modP]
		if ok {
			return false
		}
		m[modP] = 1
		modP *= candidate
		modP %= primeNumber
	}

	return true
}

func findPrimitiveRoot(primeNumber int64) int64 {
	var i int64
	for i = 2; i < primeNumber; i++ {
		if isPrimitive(i, primeNumber) {
			return i
		}
	}

	return i
}

func findPrivateKey(primeNumber int64) int64 {
	rand.Seed(time.Now().UnixNano())

	return int64(rand.Intn(int(primeNumber)-1) + 1)
}

func findPublicKey(privateKey int64, primitiveRoot int64, primeNumber int64) int64 {
	base := primitiveRoot
	result := base

	var i int64
	for i = 0; i < privateKey; i++ {
		result *= base
		result %= primeNumber
	}

	return result
}

func generateKeysFromPrimeNumber(primeNumber int64) (primitiveRoot, privateKey, publicKey int64) {
	primitiveRoot = findPrimitiveRoot(primeNumber)
	privateKey = findPrivateKey(primeNumber)
	publicKey = findPublicKey(privateKey, primitiveRoot, primeNumber)

	return
}
func (dh *DiffieHellman) SetPrimeNumber(primeNumber int64) error {
	if isPrime(primeNumber) {
		dh.primeNumber = primeNumber
		dh.primitiveRoot, dh.privateKey, dh.publicKey = generateKeysFromPrimeNumber(dh.primeNumber)
	}

	return nil
}

func (dh *DiffieHellman) GetPrivateKey() int64 {
	return dh.privateKey
}

func (dh *DiffieHellman) GetPublicKey() int64 {
	return dh.publicKey
}

func (dh *DiffieHellman) GenerateTransportKey(publicKey int64) int64 {
	base := publicKey
	result := base

	var i int64
	for i = 0; i < dh.privateKey; i++ {
		result *= base
		result %= dh.primeNumber
	}

	return result
}

func New(primeNumber ...int64) (*DiffieHellman, error) {
	dh := &DiffieHellman{}

	if len(primeNumber) == 0 {
		dh.primeNumber = defaultPrime
	} else {
		prime := primeNumber[0]
		if isPrime(prime) {
			dh.primeNumber = prime
		} else {
			dh.primeNumber = defaultPrime
		}
	}

	dh.primitiveRoot, dh.privateKey, dh.publicKey = generateKeysFromPrimeNumber(dh.primeNumber)

	return dh, nil
}
