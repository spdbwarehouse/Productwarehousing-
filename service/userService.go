package service

import (
	"wareHouse/dao"
	"wareHouse/repository"
)

type Service interface {
	RegisterUser(user *dao.User) error
	GetAllUsers() ([]dao.User, error)
	GetUserByID(id uint) (dao.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return &userService{repo}
}

func (u *userService) RegisterUser(user *dao.User) error {
	return u.userRepo.Save(*user)
}

func (u *userService) GetAllUsers() ([]dao.User, error) {
	return u.userRepo.GetAll()
}

func (u *userService) GetUserByID(id uint) (dao.User, error) {
	return u.userRepo.GetByID(id)
}
