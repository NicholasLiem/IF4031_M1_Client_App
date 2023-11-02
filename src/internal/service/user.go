package service

import (
	"errors"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user datastruct.UserModel) error
	UpdateUser(user dto.UpdateUserDTO, issuerID uint) (*datastruct.UserModel, error)
	DeleteUser(requestedUserID, issuerID uint) (*datastruct.UserModel, error)
	GetUser(requestedUserID, userID uint) (*datastruct.UserModel, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) CreateUser(user datastruct.UserModel) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	_, err = u.dao.NewUserQuery().CreateUser(user)
	return err
}

func (u *userService) UpdateUser(user dto.UpdateUserDTO, issuerID uint) (*datastruct.UserModel, error) {
	var userBySession *datastruct.UserModel
	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role == datastruct.ADMIN {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
		updatedUser, err := u.dao.NewUserQuery().UpdateUser(user)
		return updatedUser, err
	}

	return nil, errors.New("unauthorized access")
}

func (u *userService) DeleteUser(requestedUserID, issuerID uint) (*datastruct.UserModel, error) {
	var userBySession *datastruct.UserModel
	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	if userBySession.Role == datastruct.ADMIN {
		userData, err := u.dao.NewUserQuery().DeleteUser(requestedUserID)
		return userData, err
	}

	return nil, errors.New("unauthorized access")
}

func (u *userService) GetUser(requestedUserID uint, issuerID uint) (*datastruct.UserModel, error) {
	var userBySession *datastruct.UserModel

	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, errors.New("user isn't authorized")
	}

	userByRequest, err := u.dao.NewUserQuery().GetUser(requestedUserID)
	if err != nil {
		return nil, errors.New("requested user doesn't exist")
	}

	if userByRequest.ID == userBySession.ID || userBySession.Role == datastruct.ADMIN {
		return userByRequest, nil
	} else {
		return &datastruct.UserModel{
			FirstName: userByRequest.FirstName,
			LastName:  userByRequest.LastName,
			Email:     userByRequest.Email,
		}, nil
	}
}
