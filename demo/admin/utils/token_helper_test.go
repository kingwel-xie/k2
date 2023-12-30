package utils

import (
	"testing"
)

// GenerateToken returns a unique token based on the provided some key
func TestToken(t *testing.T) {
	token := GenerateToken()
	t.Log(token)

	if !CompareToken(token) {
		t.Fatal("fail to compare")
	}

	token2 := GenerateToken()

	if token == token2 {
		t.Fatal("duplicate token is generated")
	}
}
