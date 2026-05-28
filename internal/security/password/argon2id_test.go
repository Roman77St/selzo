package password

import "testing"

func TestArgon2IDHasher(t *testing.T) {
	hasher := NewArgon2IDHasher()

	password := "super-secure-password"

	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}

	ok, err := hasher.Verify(password, hash)
	if err != nil {
		t.Fatalf("verify password: %v", err)
	}

	if !ok {
		t.Fatal("expected password to match hash")
	}
}