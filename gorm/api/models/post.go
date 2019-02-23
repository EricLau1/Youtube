package models

import (
  "time"
)

type Post struct {
  Id          uint32     `gorm:"primary_key;auto_increment" json:"id"`
  Description string     `gorm:"type:varchar(255)" json:"description"`
  ImageUrl    string     `gorm:"type:varchar(255)" json:"image_url"`
  Subtitle    string     `gorm:"type:varchar(100)" json:"subtitle"`
  UserId      uint32     `gorm:"not null" json:"user_id"`
  User        User       `json:"user"`
  CreatedAt   time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
  UpdatedAt   time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
  Feedbacks   []Feedback `gorm:"ForeignKey:PostId" json:"feedbacks"`
}

func NewPost(post Post) error {
  db := Connect()
  defer db.Close()
  return db.Create(&post).Error
}

func GetPosts() []Post {
  db := Connect()
  defer db.Close()
  var posts []Post
  db.Order("id asc").Find(&posts)
  for i, _ := range posts {
    db.Model(&posts[i]).Related(&posts[i].User)
    posts[i].Feedbacks = GetFeedbacksByPost(posts[i])
  }
  return posts
}

func UpdatePost(post Post) (int64, error) {
  db := Connect()
  defer db.Close()
  rs := db.Model(&post).Where("id = ?", post.Id).UpdateColumns(
    map[string]interface{}{
      "description": post.Description,
      "image_url": post.ImageUrl,
      "subtitle": post.Subtitle,
    },
  )
  return rs.RowsAffected, rs.Error
}
