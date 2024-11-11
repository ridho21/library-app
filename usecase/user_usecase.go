package usecase

import (
	"errors"
	"test-ordent/model"
	"test-ordent/model/dto/request"
	"test-ordent/model/dto/response"
	"test-ordent/repository"
	"test-ordent/utils/common"
	"test-ordent/utils/encryption"
	"time"

	"github.com/google/uuid"
)

type UserUsecase interface {
	FindUserById(id string) (model.User, error)
	CreateUser(payload model.User) (response.RegisterResponseDto, error)
	LoginUser(paylod request.LoginRequestDto) (response.LoginResponseDto, error)
}

type userUsecase struct {
	repo     repository.UserRepository
	jwtToken common.JwtToken
}

func (u *userUsecase) FindUserById(id string) (model.User, error) {
	return model.User{}, nil
}

func (u *userUsecase) CreateUser(payload model.User) (response.RegisterResponseDto, error) {
	payload.Id = uuid.NewString()
	payload.CreatedAt = time.Now()
	hasPass, err := encryption.HashPassword(payload.Password)
	payload.Password = hasPass
	if err != nil {
		return response.RegisterResponseDto{}, err
	}
	user, err := u.repo.Create(payload)
	if err != nil {
		return response.RegisterResponseDto{}, err
	}
	return response.RegisterResponseDto{
		Id:          user.Id,
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (u *userUsecase) LoginUser(payload request.LoginRequestDto) (response.LoginResponseDto, error) {

	user, err := u.repo.GetByEmail(payload.Email)
	if err != nil {
		return response.LoginResponseDto{}, err
	}

	isValid := encryption.CheckPassword(payload.Password, user.Password)
	if !isValid {
		return response.LoginResponseDto{}, errors.New("Password is invalid")
	}

	accessToken, err := u.jwtToken.GenerateTokenJwt(user)
	if err != nil {
		return response.LoginResponseDto{}, err
	}
	return response.LoginResponseDto{
		AccessToken: accessToken,
		UserId:      user.Id,
	}, nil
}

func NewUserUsecase(
	repo repository.UserRepository,
	jwtToken common.JwtToken,
) UserUsecase {
	return &userUsecase{
		repo:     repo,
		jwtToken: jwtToken,
	}
}
