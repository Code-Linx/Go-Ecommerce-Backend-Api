package auth

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func ComparePasswords(hahsed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hahsed), plain)
	return err == nil
}
