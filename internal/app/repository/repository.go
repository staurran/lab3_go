package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"lab3/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetAllProducts() ([]ds.Goods, error) {
	var products []ds.Goods
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *Repository) GetProductByID(id uint) (*ds.Goods, error) {
	product := &ds.Goods{}
	err := r.db.First(product, "id = ?", id).Error // find product with code D42
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) CreateProduct(product *ds.Goods) error {
	err := r.db.Create(product).Error
	return err
}

func (r *Repository) ChangeProduct(id uint, new_price uint) error {
	err := r.db.Model(&ds.Goods{}).Where("id = ?", id).Update("price", new_price).Error
	return err
}

func (r *Repository) DeleteProduct(id uint) error {
	err := r.db.Delete(&ds.Goods{}, "id = ?", id).Error
	return err
}
