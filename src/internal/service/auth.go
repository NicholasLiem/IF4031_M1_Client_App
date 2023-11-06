package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(loginDTO dto.LoginDTO) (*datastruct.User, *jwt.JWTToken, error)
	SignUp(model datastruct.User) (*datastruct.User, error)
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

func (a *authService) SignIn(loginDTO dto.LoginDTO) (*datastruct.User, *jwt.JWTToken, error) {
	password, err := a.dao.NewUserQuery().GetUserPasswordByEmail(loginDTO.Email)
	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(loginDTO.Password))
	if err != nil {
		return nil, nil, fmt.Errorf("passwords dont match %v", err)
	} else {
		userData, err := a.dao.NewUserQuery().GetUserByEmail(loginDTO.Email)
		if err != nil {
			return nil, nil, err
		}

		jwtToken, err := jwt.CreateJWT(strconv.Itoa(int(userData.ID)), userData.Email, string(userData.Role))
		if err != nil {
			return nil, nil, err
		}

		return userData, &jwtToken, nil
	}
}

func (a *authService) SignUp(model datastruct.User) (*datastruct.User, error) {

	if !utils.IsEmailValid(model.Email) {
		return nil, errors.New("email is not valid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	model.Password = string(hashedPassword)

	userData, err := a.dao.NewUserQuery().CreateUser(model)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
