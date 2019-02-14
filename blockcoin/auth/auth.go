package auth

import (
  "go-api/models"
  "go-api/utils"
  "go-api/config"
  "errors"
  "time"
  jwt "github.com/dgrijalva/jwt-go"
)

var (
  ErrInvalidPassword = errors.New("Senha inválida")
)

type Auth struct {
  User    models.User `json:"user"`
  Token   string      `json:"token"`
  IsValid bool        `json:"is_valid"`
}

var configs = config.LoadConfigs()

func SignIn(user models.User) (Auth, error) {
  password := user.Password
  user, err := models.GetUserByEmail(user.Email)
  if err != nil {
    return Auth{IsValid:false}, err
  }
  err = utils.IsPassword(user.Password, password)  
  if err != nil {
    return Auth{IsValid:false}, ErrInvalidPassword
  }
  token, err := GenerateJWT(user)
  if err != nil {
    return Auth{IsValid:false}, err
  }
  return Auth{user, token, true}, nil
}

func GenerateJWT(user models.User) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["authorized"] = true
  claims["userId"] = user.UID // informe algum dado único do usuário que irá receber o token
  claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // tempo que o token irá expirar
  return token.SignedString(configs.Jwt.SecretKey)
}
