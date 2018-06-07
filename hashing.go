package htpasswd

import (
	"os/exec"

	"golang.org/x/crypto/bcrypt"
)

func hashSha(password string) string {
	// s := sha512.New()
	// s.Write([]byte(password))
	// passwordSum := []byte(s.Sum(nil))
	// return base64.StdEncoding.EncodeToString(passwordSum)
	args := []string{"-m", "sha-512", password}
	output, _ := exec.Command("mkpasswd", args).Output()
	return string(output)
}

func hashBcrypt(password string) (hash string, err error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	return string(passwordBytes), nil
}
