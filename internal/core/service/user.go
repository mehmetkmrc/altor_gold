package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/aggregate"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/user"
)

var (
	_              user.UserServicePort = (*UserService)(nil)
	UserServiceSet                      = wire.NewSet(NewUserService)
)

type UserService struct {
	userRepo user.UserRepositoryPort
	token    auth.TokenMaker
}

func NewUserService(userRepo user.UserRepositoryPort, token auth.TokenMaker) user.UserServicePort {
	return &UserService{
		userRepo,
		token,
	}
}

func (us *UserService) Login(ctx context.Context, email, password string) (*aggregate.UserAccess, error) {
	userPassword, err := us.userRepo.GetUserPassword(ctx, email)
	if err != nil {
		return nil, err
	}
	if strings.Compare(password, userPassword) != 0 {
		return nil, errors.New("Şifreler eşleşmiyor")
	}
	userModel, err := us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	accessToken, publicKey, accessPayload, err := us.token.CreateToken(userModel.UserID)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshPublicKey, err := us.token.CreateRefreshToken(accessPayload)
	if err != nil {
		return nil, err
	}

	sessionModel := aggregate.NewUserAccess(userModel, accessToken, publicKey, refreshToken, refreshPublicKey)

	return sessionModel, nil
}

func (us *UserService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	userModel, err := us.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (us *UserService) Register(ctx context.Context, first_name, last_name, email, phone, password string)(*entity.User, error) {
	newUser := &entity.User{
		UserID: uuid.New().String(),
		Name: first_name,
		Surname: last_name,
		Email: email,
		Phone: phone,
		Password: password,
	}
	err := us.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}