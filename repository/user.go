package repository

import (
	"wareHouse/dao"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAll() ([]dao.User, error)
	GetByID(id uint) (dao.User, error)
	Save(user dao.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll() ([]dao.User, error) {
	var users []dao.User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint) (dao.User, error) {
	var user dao.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Save(user dao.User) error {
	err := r.db.Save(&user).Error
	return err
}
