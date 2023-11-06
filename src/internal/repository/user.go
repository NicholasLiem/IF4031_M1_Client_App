package repository

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"gorm.io/gorm"
)

type UserQuery interface {
	CreateUser(user datastruct.User) (*datastruct.User, error)
	UpdateUser(userID uint, user dto.UpdateUserDTO) (*datastruct.User, error)
	DeleteUser(userID uint) (*datastruct.User, error)
	GetUser(userID uint) (*datastruct.User, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetUserByEmail(email string) (*datastruct.User, error)
}

type userQuery struct {
	pgdb *gorm.DB
}

func NewUserQuery(pgdb *gorm.DB) UserQuery {
	return &userQuery{
		pgdb: pgdb,
	}
}

func (u *userQuery) CreateUser(user datastruct.User) (*datastruct.User, error) {
	newUser := datastruct.User{
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

func (u *userQuery) UpdateUser(userID uint, user dto.UpdateUserDTO) (*datastruct.User, error) {
	err := u.pgdb.Model(datastruct.User{}).Where("id = ?", userID).Updates(user).Error

	var updatedUser datastruct.User
	err = u.pgdb.Where("id = ?", userID).First(&updatedUser).Error
	if err != nil {
		return nil, err
	}

	return &updatedUser, err
}

func (u *userQuery) DeleteUser(userID uint) (*datastruct.User, error) {
	var userData datastruct.User
	err := u.pgdb.Model(datastruct.User{}).Where("id = ?", userID).First(&userData).Error
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

func (u *userQuery) GetUser(userID uint) (*datastruct.User, error) {
	var userData datastruct.User
	err := u.pgdb.Where("id = ?", userID).First(&userData).Error
	return &userData, err
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	var password string
	err := u.pgdb.Model(&datastruct.User{}).Where("email = ?", email).Select("password").Scan(&password).Error
	return &password, err
}

func (u *userQuery) GetUserByEmail(email string) (*datastruct.User, error) {
	var userData datastruct.User
	err := u.pgdb.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}

	return &userData, err
}
