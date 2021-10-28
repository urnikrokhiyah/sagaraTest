package database

import (
	"errors"
	"test/config"
	"test/models"
)

func CreateNewProduct(inputProduct models.Product) (interface{}, error) {
	if inputProduct.Name == "" || inputProduct.Category == "" || inputProduct.Price == 0 {
		return nil, errors.New("field must be exist")
	}

	tx := config.Db.Create(&inputProduct)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return inputProduct, nil
}

func GetListProduct() (interface{}, int, error) {
	listResponse := []models.ListProductResponse{}
	product := []models.Product{}

	tx := config.Db.Model(product).Find(&listResponse)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return listResponse, 1, nil
	}
	return nil, 0, nil
}

func GetProductDetail(productId int) (interface{}, int, error) {
	product := models.Product{}

	tx := config.Db.Find(&product, productId)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return product, 1, nil
	}
	return nil, 0, nil
}

func DeleteProduct(productId int) (interface{}, int, error) {
	product := models.Product{}

	tx := config.Db.Delete(&product, productId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return "deleted", 1, nil
	}

	return nil, 0, nil
}

func UpdateProduct(productId int, newData models.Product) (interface{}, int, error) {
	product := models.Product{}

	tx := config.Db.Find(&product, productId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {

		saveData := config.Db.Model(&product).Updates(newData)
		if saveData.Error != nil {
			return nil, 0, saveData.Error
		}

		return "updated data", 1, nil

	}

	return nil, 0, nil
}

func UpdateImage(productId int, path string) (interface{}, int, error) {
	product := models.Product{}

	tx := config.Db.Find(&product, productId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {

		product.Image = path
		saveData := config.Db.Save(&product)
		if saveData.Error != nil {
			return nil, 0, saveData.Error
		}

		return "updated photo", 1, nil

	}

	return nil, 0, nil
}
