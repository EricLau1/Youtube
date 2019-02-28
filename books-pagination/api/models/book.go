package models

type Book struct {
  Id uint64 `gorm:"primary_key;auto_increment" json:"id"`
  Title string `gorm:"size:100;not null" json:"title"`
  Rating float32 `gorm:"type:double;default:0" json:"rating"`
}

func GetBooks() []Book {
  db := Connect()
  defer db.Close()
  var books []Book
  db.Order("id asc").Find(&books)
  return books
}

func GetBookById(id uint64) Book {
  db := Connect()
  defer db.Close()
  var book Book
  db.Where("id = ?", id).Find(&book)
  return book
}

func NewBook(book Book) error {
  db := Connect()
  defer db.Close()
  return db.Create(&book).Error
}

func UpdateBook(book Book) (int64, error) {
  db := Connect()
  defer db.Close()
  rs := db.Model(&book).Where("id = ?", book.Id).UpdateColumns(
    map[string]interface{}{
      "title": book.Title,
      "rating": book.Rating,
    },
  )
  return rs.RowsAffected, rs.Error
}

func DeleteBook(id uint64) (int64, error) {
  db := Connect()
  defer db.Close()
  rs := db.Where("id = ?", id).Delete(&Book{})
  return rs.RowsAffected, rs.Error
}

func CountBooks() int {
  db := Connect()
  defer db.Close()
  var count int
  db.Model(&Book{}).Count(&count)
  return count
}

func PaginateBooks(begin, limit int) []Book {
  db := Connect()
  defer db.Close()
  var books []Book
  db.Offset(begin).Limit(limit).Find(&books)
  return books
}
