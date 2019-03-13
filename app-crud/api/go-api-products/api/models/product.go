package models

import "time"

type Product struct {
	Id          uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Description string     `gorm:"size:255;not null" json:"description"`
	Price       float32    `gorm:"type:decimal(10,2);not null" json:"price"`
	Quantity    int        `gorm:"default:0" json:"quantity"`
	Status      string     `gorm:"type:enum('0', '1');default:'1'" json:"status"`
	CreatedAt   *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

func NewProduct(product Product) (interface{}, error) {
	db := Connect()
	defer db.Close()
	product.StatusVerify()
	rs := db.Create(&product)
	return rs.Value, rs.Error
}

func (p *Product) StatusVerify() string {
	if p.Quantity > 0 {
		p.Status = "1"
	} else {
		p.Status = "0"
	}
	return p.Status
}

func GetProducts() []Product {
	db := Connect()
	defer db.Close()
	var products []Product
	db.Order("id asc").Find(&products)
	return products
}

func GetProductById(id uint64) Product {
	db := Connect()
	defer db.Close()
	var product Product
	db.Where("id = ?", id).Find(&product)
	return product
}

func UpdateProduct(product Product) (interface{}, error) {
	db := Connect()
	defer db.Close()
	product.StatusVerify()
	rs := db.Model(&product).Where("id = ?", product.Id).UpdateColumns(
		map[string]interface{}{
			"description": product.Description,
			"price": product.Price,
			"quantity": product.Quantity,
			"status": product.Status,
			"updated_at": time.Now(),
		},
	)
	return rs.Value, rs.Error
}

func DeleteProduct(id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", id).Delete(&Product{})
	return rs.RowsAffected, rs.Error
}
