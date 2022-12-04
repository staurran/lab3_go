package repository

import (
	"errors"
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
	user.Role = user_db.Role
	return nil
}

func (r *Repository) CheckLogin(login string) error {
	err := r.db.First(&ds.Users{}).Where("login = ?", login).Error
	if err != nil {
		return nil
	}
	err1 := errors.New("math: square root of negative number")
	return err1
}

func (r *Repository) GetUserByID(id uint) (*ds.Users, error) {
	user := &ds.Users{}
	err := r.db.First(user, "id_user = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetIdByLogin(login string) (uint, error) {
	user := &ds.Users{}
	err := r.db.First(user, "login = ?", login).Error
	if err != nil {
		return 0, err
	}
	return user.Id_user, nil
}

func (r *Repository) CreateBasketRow(basket_row *ds.Basket) error {
	err := r.db.Create(basket_row).Error
	return err
}

func (r *Repository) GetBasket(id_user uint) ([]ds.Basket, error) {
	var basket []ds.Basket
	result := r.db.Find(&basket, "id_user = ?", id_user)
	if result.Error != nil {
		return nil, result.Error
	}
	return basket, nil
}
func (r *Repository) DeleteBasketRow(basket_row *ds.Basket) error {
	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.Basket{}, "id_row = ?", basket_row.Id_row).Error
	return err
}

func (r *Repository) ChangeQuantity(basket_row *ds.Basket, quantity int) error {
	err := r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_good, "id_user = ?", basket_row.Id_user).Take(&basket_row).Error
	if err != nil {
		return err
	}
	err = r.db.Model(&ds.Basket{}).Where("id_good = ?", basket_row.Id_row).Update("quantity", quantity).Error
	return err
}
