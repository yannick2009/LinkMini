package utils

import "math/rand"

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const HashLength = 7

func GenerateHash() string {
	hash := make([]byte, HashLength)

	for i := range hash {
		randI := rand.Intn(len(Charset))
		hash[i] = Charset[randI]
	}

	result := string(hash)
	return string(result)
}
