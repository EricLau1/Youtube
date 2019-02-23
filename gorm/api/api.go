package api

import (
  _ "gorm/api/models"
)

func Run() {
  //models.AutoMigrations()
  listen(9000)
}
