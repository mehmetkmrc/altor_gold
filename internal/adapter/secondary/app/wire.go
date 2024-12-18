//go:build wireinject
// +build wireinject

package app

import (
	"context"
	http_adapter "github.com/mehmetkmrc/ator_gold/internal/adapter/primary/http"
	rabbitmq "github.com/mehmetkmrc/ator_gold/internal/adapter/primary/rabbit"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/storage/memcache"	
	
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/auth/paseto"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	psql "github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/storage/psql"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/http"
	
	"github.com/mehmetkmrc/ator_gold/internal/core/port/user"
	"github.com/mehmetkmrc/ator_gold/internal/core/service"
	"sync"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func InitApp(
	ctx context.Context,
	wg *sync.WaitGroup,
	rw *sync.RWMutex,
	cfg *config.Container,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		httpServerFunc,
		memcache.NewMemcacheSet,
		memcache.MemcacheTTLSet,
		paseto.PasetoSet,
		psql.MainRepositorySet,
		psql.SubDocumentRepositorySet,
		psql.ContentRepositorySet,
		psql.UserRepoSet,
		service.DocumentServiceSet,
		service.UserServiceSet,
	))
}

func dbEngineFunc(
	ctx context.Context,
	Cfg *config.Container,
) (db.EngineMaker, func(), error) {
	psqlDb := psql.NewDB(Cfg)
	err := psqlDb.Start(ctx)
	if err != nil {
		zap.S().Fatal("failed to start db:", err)
	}

	return psqlDb, func() { psqlDb.Close(ctx) }, nil
}

func httpServerFunc(
	ctx context.Context,
	Cfg *config.Container,
	user user.UserServicePort,
	token auth.TokenMaker,
	docService documenter.DocumentServicePort,
) (http.ServerMaker, func(), error) {
	httpServer := http_adapter.NewHTTPServer(ctx, Cfg, user, token, docService)
	err:= httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}

func rabbitMQFunc(
	ctx context.Context,
	Cfg *config.Container,
) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(Cfg)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
