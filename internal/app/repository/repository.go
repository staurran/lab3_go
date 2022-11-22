package repository

import (
	"golang.org/x/crypto/bcrypt"
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
	err := r.db.First(product, "id_good = ?", id).Error // find product with code D42
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
	err := r.db.Model(&ds.Goods{}).Where("id_good = ?", id).Update("price", new_price).Error
	return err
}

func (r *Repository) DeleteProduct(id uint) error {
	err := r.db.First(&ds.Goods{}, "id_good = ?", id).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.Goods{}, "id_good = ?", id).Error
	return err
}

//Users

func (r *Repository) CreateUser(user *ds.Users) error {
	err := r.db.Create(user).Error
	return err
}

func (r *Repository) LoginCheck(user *ds.Users) error {
	user_db := ds.Users{}
	err := r.db.Model(&ds.Users{}).Where("login = ?", user.Login).Take(&user_db).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	user.Id_user = user_db.Id_user
	return nil
}

func (r *Repository) GetUserByID(id uint) (*ds.Users, error) {
	user := &ds.Users{}
	err := r.db.First(user, "id_user = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
