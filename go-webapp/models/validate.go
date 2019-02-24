package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
  "fmt"
)

var (
	ErrRequiredFirstName = errors.New("Nome requerido")
	ErrRequiredLastName = errors.New("Sobrenome requerido")
	ErrRequiredEmail = errors.New("Email requerido")
	ErrInvalidEmail = errors.New("Email inválido")
	ErrRequiredPassword = errors.New("Password requerido")
	ErrMaxLimit = errors.New("Ultrapassou o limite de caracteres")
	ErrEmailTaken = errors.New("Email já está em uso")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

func Max(attr string, lim int) bool {
	if len(attr) <= lim {
		return true
	}
	return false
}

func ValidateLimitFields(user User) (User, error) {
	if !Max(user.FirstName, 15) || !Max(user.LastName, 20) || !Max(user.Email, 40) || !Max(user.Password, 100) {
		return User{}, ErrMaxLimit
	}
	return user, nil
}

func VerifyEmail(email string) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := "select count(email) from users where email = $1" 
	rs, err := con.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()
	var count int64
	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	if count > 0 {
		return false, ErrEmailTaken
	}
	return true, nil
}

func ValidateNewUser(user User) (User, error) {
	user, err := ValidateLimitFields(user)
	if err != nil {
		return user, err
	}
	user.FirstName = Trim(user.FirstName)
	user.LastName = Trim(user.LastName)
	user.Email = Trim(strings.ToLower(user.Email))
	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}
	_, err = VerifyEmail(user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func Count(table string) (int64, error) {
  con := Connect()
  defer con.Close()
  sql := fmt.Sprintf("select count(*) from %s", table)
  var count int64
  err := con.QueryRow(sql).Scan(&count)
  if err != nil {
    return 0, err
  }
  return count, nil
}
