package psql

import (
	"context"

	"time"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/user"
	"golang.org/x/crypto/bcrypt"
)

var UserRepoSet = wire.NewSet(NewUserRepository)

type(
	UserRepository struct{
		db *pgxpool.Pool
	}
)

func NewUserRepository(db db.EngineMaker) user.UserRepositoryPort {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error){
	userQuery := struct {
		UserID		  string
		Name	  	  string
		Surname   	  string
		Email	  	  string
		Password  	  string
		CreatedAt 	  time.Time
	}{}
	query := `
	SELECT CAST(user_id AS VARCHAR(64)) as ID, 
       first_name, 
       last_name, 
       email, 
       password, 
       created_at 
	FROM Users 
	WHERE Email = $1 
  		AND password IS NOT NULL 
  		AND email IS NOT NULL;
	`
	err := r.db.QueryRow(ctx, query, email).Scan(&userQuery.UserID, &userQuery.Name, &userQuery.Surname, &userQuery.Email, &userQuery.Password, &userQuery.CreatedAt)
	if err != nil{
		return nil, err
	}

	userData := &entity.User{
		UserID:    userQuery.UserID,
		Name:      userQuery.Name,
		Surname:   userQuery.Surname,
		Email:     userQuery.Email,
		Password:  userQuery.Password,
		CreatedAt: userQuery.CreatedAt,
	}
	return userData, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	userQuery := struct {
		UserID        string
		Name      	  string
		Surname   	  string
		Email     	  string
		Password  	  string
		CreatedAt 	  time.Time
	}{}
	query := `SELECT CAST(user_id AS VARCHAR(64)) as UserID, 
       first_name, 
       last_name, 
       email, 
       password, 
       created_at 
	FROM Users 
	WHERE user_id = $1 
  		AND password IS NOT NULL 
  		AND email IS NOT NULL;
	`
	err := r.db.QueryRow(ctx, query, id).Scan(&userQuery.UserID, &userQuery.Name, &userQuery.Surname, &userQuery.Email, &userQuery.Password, &userQuery.CreatedAt)
	if err != nil {
		return nil, err
	}

	userData := &entity.User{
		UserID:    userQuery.UserID,
		Name:      userQuery.Name,
		Surname:   userQuery.Surname,
		Email:     userQuery.Email,
		Password:  userQuery.Password,
		CreatedAt: userQuery.CreatedAt,
	}
	return userData, nil
}

func (r *UserRepository) GetUserPassword(ctx context.Context, email string) (string, error) {
	var password string
	query := `SELECT password 
	FROM users 
	WHERE email = $1;
	`
	err := r.db.QueryRow(ctx, query, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        return err
    }
	user.CreatedAt = time.Now()
	user.Password = string(hashedPassword)

	query := `
	INSERT INTO users (user_id, first_name, last_name, email, phone, password, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`
	_, err = r.db.Exec(ctx, query, user.UserID, user.Name, user.Surname, user.Email, user.Phone, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
