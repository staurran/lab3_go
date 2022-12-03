package ds

type Goods struct {
	Id_good  uint   `sql:"type:uuid;primary_key;default:" json:"Id_good" gorm:"primarykey"`
	Type     string `json:"type"`
	Company  string `json:"company"`
	Color    string `json:"color"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
}

type Users struct {
	Id_user  uint   `sql:"type:uuid;primary_key;default:" json:"Id_user" gorm:"primarykey"`
	Type     string `json:"type"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Basket struct {
	Id_row  uint `sql:"type:uuid;primary_key;default:" json:"Id_row" gorm:"primarykey"`
	Id_user uint `sql:"type:uuid;foreign_key;" json:"id_user" gorm:"foreignkey"`
	Id_good uint `sql:"type:uuid;foreign_key;" json:"id_good" gorm:"foreignkey"`
}
