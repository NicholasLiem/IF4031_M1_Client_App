package service

import (
	"net/http"
	"strconv"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(loginDTO dto.LoginDTO) (*datastruct.User, *jwt.JWTToken, *utils.HttpError)
	SignUp(model datastruct.User) (*datastruct.User, *utils.HttpError)
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

func (a *authService) SignIn(loginDTO dto.LoginDTO) (*datastruct.User, *jwt.JWTToken, *utils.HttpError) {
	password, err := a.dao.NewUserQuery().GetUserPasswordByEmail(loginDTO.Email)
	if err != nil {
		return nil, nil, &utils.HttpError{
			Message:    "invalid credentials",
			StatusCode: http.StatusUnauthorized,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(loginDTO.Password))
	if err != nil {
		return nil, nil, &utils.HttpError{
			Message:    "invalid credentials",
			StatusCode: http.StatusUnauthorized,
		}
	} else {
		userData, err := a.dao.NewUserQuery().GetUserByEmail(loginDTO.Email)
		if err != nil {
			return nil, nil, &utils.HttpError{
				Message:    "invalid credentials",
				StatusCode: http.StatusUnauthorized,
			}
		}

		jwtToken, err := jwt.CreateJWT(strconv.Itoa(int(userData.ID)), userData.Email, string(userData.Role))
		if err != nil {
			return nil, nil, &utils.HttpError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		return userData, &jwtToken, nil
	}
}

func (a *authService) SignUp(model datastruct.User) (*datastruct.User, *utils.HttpError) {

	if !utils.IsEmailValid(model.Email) {
		return nil, &utils.HttpError{
			Message:    "email is not valid",
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.MinCost)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	model.Password = string(hashedPassword)

	// Check unique constraint
	isUnique, err := a.dao.NewUserQuery().IsEmailUnique(model.Email, nil)

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

	userData, err := a.dao.NewUserQuery().CreateUser(model)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return userData, nil
}
