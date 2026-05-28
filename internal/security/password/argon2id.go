package password

import "github.com/alexedwards/argon2id"

type Argon2IDHasher struct{}

func NewArgon2IDHasher() *Argon2IDHasher {
	return &Argon2IDHasher{}
}

func (h *Argon2IDHasher) Hash(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func (h *Argon2IDHasher) Verify(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}