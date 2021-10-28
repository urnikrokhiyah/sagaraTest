package database

import (
	"errors"
	"test/config"
	"test/middlewares"
	"test/models"
)

func RegisterUser(inputUser models.User) (interface{}, error) {

	err := validateRegistration(inputUser)
	if err != nil {
		return nil, err
	}

	tx := config.Db.Create(&inputUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return inputUser, nil
}

func LoginUser(loginData models.LoginUser) (interface{}, int, error) {
	user := models.User{}
	tx := config.Db.Where("users.email = ? && users.password = ?", loginData.Email, loginData.Password).Find(&user)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		token, _ := middlewares.CreateToken(int(user.ID), user.Role)
		loginResponse := models.LoginResponse{
			Email: user.Email,
			Token: token,
		}
		return loginResponse, 1, nil
	}
	return nil, 0, nil
}

func validateRegistration(inputData models.User) error {

	switch {
	case inputData.Name == "":
		return errors.New("name must be exist!")
	case inputData.Email == "":
		return errors.New("email must be exist!")
	case inputData.Password == "":
		return errors.New("password must be exist!")
	}

	return nil
}
