package app

import (
	"context"
	"sync"

	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/http"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/user"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type App struct {
	rw *sync.RWMutex
	wg *sync.WaitGroup
	Cfg *config.Container
	AMQPConn *amqp.Connection
	HTTP	 http.ServerMaker
	PSQL 	 db.EngineMaker
	TokenMaker auth.TokenMaker
	UserRepo user.UserRepositoryPort
	UserService user.UserServicePort
}

func New(
	rw *sync.RWMutex,
	wg *sync.WaitGroup,
	cfg *config.Container,
	amqpConn *amqp.Connection,
	http http.ServerMaker,
	psql db.EngineMaker,
	token auth.TokenMaker,
	userRepo user.UserRepositoryPort,
	userService user.UserServicePort,
) *App {
	return &App{
		rw: rw,
		wg: wg,
		Cfg: cfg,
		HTTP: http,
		AMQPConn: amqpConn,
		PSQL: psql,
		TokenMaker: token,
		UserRepo: userRepo,
		UserService: userService,
	}
}
func (a *App) Run(ctx context.Context){
	zap.S().Info("Runner!")
}
