package utils

import(
  "golang.org/x/crypto/bcrypt"
)

func Hash(password string)([]byte, error){

  cost := bcrypt.DefaultCost

  return bcrypt.GenerateFromPassword([]byte(password), cost)

}


func VerifyPassword(hash, password []byte) error {

  return bcrypt.CompareHashAndPassword(hash, password)

}
