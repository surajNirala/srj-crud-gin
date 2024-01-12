package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
)

// "github.com/golang-jwt/jwt"
// "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID        uint    `gorm:"column:id" json:"id"`
	FirstName string  `gorm:"size:255;column:first_name" json:"first_name"`
	LastName  string  `gorm:"size:255;column:last_name" json:"last_name"`
	Email     string  `gorm:"size:255;column:email;uniqueIndex;not null" json:"email"`
	Password  string  `gorm:"size:255;column:password" json:"password"`
	Phone     string  `gorm:"size:255;column:phone" json:"phone"`
	Image     *string `gorm:"size:255;column:image" json:"image"`
	Age       uint    `gorm:"default:18;column:age" json:"age"`
	Status    bool    `gorm:"default:true;column:status" json:"status"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
