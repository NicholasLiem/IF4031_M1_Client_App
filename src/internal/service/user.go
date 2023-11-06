package service

import (
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user datastruct.User) *utils.HttpError
	UpdateUser(requestedUserID uint, user dto.UpdateUserDTO, issuerID uint) (*datastruct.UserResponse, *utils.HttpError)
	DeleteUser(requestedUserID, issuerID uint) (*datastruct.User, *utils.HttpError)
	GetUser(requestedUserID, userID uint) (*datastruct.UserResponse, *utils.HttpError)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) CreateUser(user datastruct.User) *utils.HttpError {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	// Check unique constraint
	isUnique, err := u.dao.NewUserQuery().IsEmailUnique(user.Email, nil)
	if err != nil {
		return &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !isUnique {
		return &utils.HttpError{
			Message:    "email is already exists",
			StatusCode: http.StatusConflict,
		}
	}

	user.Password = string(hashedPassword)
	_, err = u.dao.NewUserQuery().CreateUser(user)

	return &utils.HttpError{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}

func (u *userService) UpdateUser(requestedUserID uint, user dto.UpdateUserDTO, issuerID uint) (*datastruct.UserResponse, *utils.HttpError) {
	var userBySession *datastruct.User
	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role == datastruct.ADMIN {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			return nil, &utils.HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		// Check unique constraint
		isUnique, err := u.dao.NewUserQuery().IsEmailUnique(user.Email, &requestedUserID)
		if err != nil {
			return nil, &utils.HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		if !isUnique {
			return nil, &utils.HttpError{
				Message:    "email is already exists",
				StatusCode: http.StatusConflict,
			}
		}

		user.Password = string(hashedPassword)
		updatedUser, err := u.dao.NewUserQuery().UpdateUser(requestedUserID, user)

		if err != nil {
			return nil, &utils.HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		responseData := datastruct.UserResponse{
			ID:        updatedUser.ID,
			FirstName: updatedUser.FirstName,
			LastName:  updatedUser.LastName,
			Email:     updatedUser.Email,
			Role:      updatedUser.Role,
		}

		return &responseData, nil
	}

	return nil, &utils.HttpError{
		Message:    "unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func (u *userService) DeleteUser(requestedUserID, issuerID uint) (*datastruct.User, *utils.HttpError) {
	var userBySession *datastruct.User
	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	if userBySession.Role == datastruct.ADMIN {
		userData, err := u.dao.NewUserQuery().DeleteUser(requestedUserID)
		return userData, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil, &utils.HttpError{
		Message:    "unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func (u *userService) GetUser(requestedUserID uint, issuerID uint) (*datastruct.UserResponse, *utils.HttpError) {
	var userBySession *datastruct.User

	userBySession, err := u.dao.NewUserQuery().GetUser(issuerID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}

	userByRequest, err := u.dao.NewUserQuery().GetUser(requestedUserID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    "requested user doesn't exist",
			StatusCode: http.StatusNotFound,
		}
	}

	if userByRequest.ID == userBySession.ID || userBySession.Role == datastruct.ADMIN {
		responseData := datastruct.UserResponse{
			ID:        userByRequest.ID,
			FirstName: userByRequest.FirstName,
			LastName:  userByRequest.LastName,
			Email:     userByRequest.Email,
			Role:      userByRequest.Role,
		}

		return &responseData, nil
	} else {
		return nil, &utils.HttpError{
			Message:    "unauthorized",
			StatusCode: http.StatusUnauthorized,
		}
	}
}
