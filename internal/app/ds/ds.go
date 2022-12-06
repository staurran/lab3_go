package ds

import (
	"github.com/dgrijalva/jwt-go"
	"lab3/internal/app/role"
)

type Goods struct {
	Id_good     uint   `sql:"type:uuid;primary_key;default:" json:"Id_good" gorm:"primarykey"`
	Type        string `json:"type"`
	Company     string `json:"company"`
	Color       string `json:"color"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Users struct {
	Id_user  uint      `sql:"type:uuid;primary_key;default:" json:"Id_user" gorm:"primarykey"`
	Role     role.Role `json:"type"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

type Basket struct {
	Id_row   uint `sql:"type:uuid;primary_key;default:" json:"Id_row" gorm:"primarykey"`
	Id_user  uint `json:"id_user"`
	Id_good  uint `json:"id_good"`
	Quantity int  `json:"quantity"`
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
	Role   role.Role
}

type Orders struct {
	Id_order uint   `sql:"type:uuid;primary_key;default:" json:"Id_order" gorm:"primarykey"`
	Date     string `json:"date"`
	Status   uint   `json:"status"`
	Id_user  uint   `json:"id_User"`
}

type GoodOrder struct {
	Id_row   uint `sql:"type:uuid;primary_key;default:" json:"Id_row" gorm:"primarykey"`
	Id_good  uint `json:"id_good"`
	Quantity int  `json:"quantity"`
	Id_order uint `json:"id_order"`
}

type Statuses struct {
	Id_status   uint   `sql:"type:uuid;primary_key;default:" json:"Id_row" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
