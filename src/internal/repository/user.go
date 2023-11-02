package repository

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"gorm.io/gorm"
)

type UserQuery interface {
	CreateUser(user datastruct.UserModel) (*datastruct.UserModel, error)
	UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error)
	DeleteUser(userID uint) (*datastruct.UserModel, error)
	GetUser(userID uint) (*datastruct.UserModel, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetUserByEmail(email string) (*datastruct.UserModel, error)
}

type userQuery struct {
	pgdb *gorm.DB
}

func NewUserQuery(pgdb *gorm.DB) UserQuery {
	return &userQuery{
		pgdb: pgdb,
	}
}

func (u *userQuery) CreateUser(user datastruct.UserModel) (*datastruct.UserModel, error) {
	newUser := datastruct.UserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	if err := u.pgdb.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *userQuery) UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error) {
	err := u.pgdb.Model(datastruct.UserModel{}).Where("id = ?", user.UserID).Updates(user).Error

	var updatedUser datastruct.UserModel
	err = u.pgdb.Where("id = ?", user.UserID).First(&updatedUser).Error
	if err != nil {
		return nil, err
	}

	return &updatedUser, err
}

func (u *userQuery) DeleteUser(userID uint) (*datastruct.UserModel, error) {
	var userData datastruct.UserModel
	err := u.pgdb.Model(datastruct.UserModel{}).Where("id = ?", userID).First(&userData).Error
	if err != nil {
		return nil, err
	}

	/**
	Perform hard delete, if you want to soft-delete, delete the Unscoped function
	*/
	err = u.pgdb.Where("id = ?", userID).Delete(&userData).Error
	if err != nil {
		return nil, err
	}

	return &userData, err
}

func (u *userQuery) GetUser(userID uint) (*datastruct.UserModel, error) {
	var userData datastruct.UserModel
	err := u.pgdb.Where("id = ?", userID).First(&userData).Error
	return &userData, err
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	var password string
	err := u.pgdb.Model(&datastruct.UserModel{}).Where("email = ?", email).Select("password").Scan(&password).Error
	return &password, err
}

func (u *userQuery) GetUserByEmail(email string) (*datastruct.UserModel, error) {
	var userData datastruct.UserModel
	err := u.pgdb.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}

	return &userData, err
}
