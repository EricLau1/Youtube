package utils

import (
  "crypto/md5"
  "fmt"
  "golang.org/x/crypto/bcrypt"
)

func Md5(str string) string {
  return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Bcrypt(password string) ([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func IsPassword(hashedPassword, password string) error {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
