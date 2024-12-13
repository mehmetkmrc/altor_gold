package user

import (
	"context"

	"github.com/mehmetkmrc/ator_gold/internal/core/domain/aggregate"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"
)

type UserRepositoryPort interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string)(*entity.User, error)
	GetUserPassword(ctx context.Context, email string)(string, error)
	Create(ctx context.Context, user *entity.User) error
}

type UserServicePort interface{
	Login(ctx context.Context, email, password string)(*aggregate.UserAccess, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	Register(ctx context.Context, name, surname, email, phone, password string) (*entity.User, error)
}