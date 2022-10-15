package ds

type Goods struct {
	Id       uint   `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID" gorm:"primarykey"`
	Type     string `json:"type"`
	Company  string `json:"company"`
	Color    string `json:"color"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
}
