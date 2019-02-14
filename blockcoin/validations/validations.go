package validations

import (
  "github.com/badoux/checkmail"
)

func IsEmpty(param string) bool {
  if param == "" {
    return true
  }
  return false
}

func IsEmail(email string) bool {
  err := checkmail.ValidateFormat(email)
  if err != nil {
    return false
  }
  return true
}
