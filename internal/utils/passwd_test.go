package utils

import (
	"log"
	"testing"
)

func TestDifferentPasswd(t *testing.T) {
	encodedHash, _ := HashPasswd("password123", &DefaultArgonParams)

	match, err := ComparePasswordAndHash("pa$$word", encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	if match {
		t.Error("Different passwords should not match")
	}
}

func TestPasswd(t *testing.T) {
	encodedHash, _ := HashPasswd("password123", &DefaultArgonParams)

	// Use a different password...
	match, err := ComparePasswordAndHash("password123", encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	if !match {
		t.Error("Different passwords should not match")
	}
}
