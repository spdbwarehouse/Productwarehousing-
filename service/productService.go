package service

import (
	"wareHouse/dao"
	"wareHouse/repository"
)

type ProductService interface {
	GetAllProducts() ([]dao.Product, error)
	GetProductByID(id uint) (dao.Product, error)
	Save(product dao.Product) error
	Delete(id uint) error
	Update(product dao.Product) error
	UpdateValues(columnMap map[string]interface{}, conditionMap map[string]interface{}) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (u *productService) GetAllProducts() ([]dao.Product, error) {
	return u.productRepo.GetAll()
}

func (u *productService) GetProductByID(id uint) (dao.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *productService) Save(product dao.Product) error {
	return u.productRepo.Save(product)
}

func (u *productService) Delete(id uint) error {
	return u.productRepo.Delete(id)
}

func (u *productService) Update(product dao.Product) error {
	return u.productRepo.Update(product)
}

func (u *productService) UpdateValues(columnMap map[string]interface{}, conditionMap map[string]interface{}) error {
	return u.productRepo.UpdateValues(columnMap, conditionMap)
}
