package models

import (
	"go-webapp/utils"
)

type User struct{
	Id uint64
	FirstName string
	LastName string
	Email string
	Password string
	Status string
}

func NewUser(user User) (bool, error) {
	user, err := ValidateNewUser(user)
	if err != nil {
		return false, err
	}
	con := Connect()
	defer con.Close()
	sql := "insert into users (firstname, lastname, email, password) values ($1, $2, $3, $4)"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	hash, err := utils.Hash(user.Password)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hash)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetUserByEmail(email string) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func GetUsers() ([]User, error) {
  con := Connect()
  defer con.Close()
  sql := "select * from users"
  rs, err := con.Query(sql)
  if err != nil {
    return nil, err
  }
  defer rs.Close()
  var users []User
  for rs.Next() {
    var user User
    err := rs.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status)
    if err != nil {
      return nil, err
    }
    users = append(users, user)
  }
  return users, nil
}
