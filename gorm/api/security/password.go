package security

import (
  "golang.org/x/crypto/bcrypt"
  "fmt"
)

func Hash(password string) (string, error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%s", hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
