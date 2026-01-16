package hasher

type PasswordHasher interface {
	Hash(pass []byte) ([]byte, error)
	ComparePassword(hashPass, pass []byte) error
}
