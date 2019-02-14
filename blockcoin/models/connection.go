package models

import (
  "database/sql"
  "fmt"
  "log"
  "go-api/config"
  "os"
  _ "github.com/lib/pq"
)

var configs = config.LoadConfigs()

const DEV = false

func Connect() *sql.DB {
  URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", configs.Database.User, configs.Database.Pass, 
    configs.Database.Name, "disable")
  var db *sql.DB
  var err error
  if DEV { 
    db, err = sql.Open("postgres", URL)
  } else {
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
  }
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return db
}

func TestConnection() {
  con := Connect()
  defer con.Close()
  err := con.Ping()
  if err != nil {
    fmt.Errorf("%s", err.Error())
    return
  }
  fmt.Println("Database connected!")
}
