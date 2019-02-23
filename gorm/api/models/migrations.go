package models

func AutoMigrations() {
  db := Connect()
  defer db.Close()
  db.Debug().DropTableIfExists(&Feedback{},&Post{}, &User{})
  db.Debug().AutoMigrate(&User{}, &Post{}, &Feedback{})
  db.Debug().Model(&Post{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
  db.Debug().Model(&Feedback{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
  db.Debug().Model(&Feedback{}).AddForeignKey("post_id", "posts(id)", "cascade", "cascade")
}
