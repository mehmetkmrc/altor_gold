package app

import (
	"context"
	"sync"

	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/cache"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
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
	MemCache cache.Memcache
	MemCacheTTL cache.MemcacheTTL
	TokenMaker auth.TokenMaker
	MainRepo  documenter.MainDocumentRepositoryPort
	SubRepo   documenter.SubDocumentRepositoryPort
	ContentRepo documenter.ContentDocumentRepositoryPort
	UserRepo user.UserRepositoryPort
	DocService documenter.DocumentServicePort
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
	memCache cache.Memcache,
	memcacheTTL cache.MemcacheTTL,
	mainRepo documenter.MainDocumentRepositoryPort,
	subRepo documenter.SubDocumentRepositoryPort,
	contentRepo documenter.ContentDocumentRepositoryPort,
	userRepo user.UserRepositoryPort,
	userService user.UserServicePort,
	docService documenter.DocumentServicePort,
) *App {
	return &App{
		rw: rw,
		wg: wg,
		Cfg: cfg,
		HTTP: http,
		AMQPConn: amqpConn,
		PSQL: psql,
		TokenMaker: token,
		MemCache: memCache,
		MemCacheTTL: memcacheTTL,
		MainRepo: mainRepo,
		SubRepo: subRepo,
		ContentRepo: contentRepo,
		UserRepo: userRepo,
		UserService: userService,
		DocService: docService,
	}
}
func (a *App) Run(ctx context.Context){
	zap.S().Info("Runner!")
}
