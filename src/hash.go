package main

import (
	"crypto/sha1"
	"encoding/hex"
)

func calculateHash(bytes []byte) string {
	sum := sha1.Sum(bytes)
	return hex.EncodeToString(sum[:])
}
