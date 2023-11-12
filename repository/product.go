package repository

import (
	"github.com/jinzhu/gorm"
	"wareHouse/dao"
)

type ProductRepository interface {
	GetAll() ([]dao.Product, error)
	GetByID(id uint) (dao.Product, error)
	Save(product dao.Product) error
	Delete(id uint) error
	Update(product dao.Product) error
	UpdateValues(columnMap map[string]interface{}, conditionMap map[string]interface{}) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAll() ([]dao.Product, error) {
	var products []dao.Product
	if err := r.db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) GetByID(id uint) (dao.Product, error) {
	var product dao.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Save(product dao.Product) error {
	err := r.db.Save(&product).Error
	return err
}

func (r *productRepository) Delete(id uint) error {
	var product dao.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}
	err := r.db.Delete(&product).Error
	return err
}

func (r *productRepository) Update(product dao.Product) error {
	err := r.db.Save(&product).Error
	return err
}

func (r *productRepository) UpdateValues(columnMap map[string]interface{}, conditionMap map[string]interface{}) error {
	err := r.db.Model(&dao.Product{}).Where(conditionMap).Updates(columnMap).Error
	return err
}
