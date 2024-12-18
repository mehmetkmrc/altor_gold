package http

import (
	"context"
	"errors"
	"fmt"
	std_http "net/http"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/auth"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/http"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/user"
	"go.uber.org/zap"
)

const (
	viewPath   = "../../client/templates"
	publicPath = "../../client/public"
	renderType = ".html"
)

var (
	_ http.ServerMaker = (*server)(nil)
)

type(
	server struct{
		ctx 	context.Context
		cfg		*config.Container
		app 	*fiber.App
		cfgFiber *fiber.Config
		userService user.UserServicePort
		tokenService auth.TokenMaker
		documentService documenter.DocumentServicePort
	}
)

func NewHTTPServer(
	ctx context.Context,
	cfg *config.Container,
	userService user.UserServicePort,
	tokenService auth.TokenMaker,
	documentService documenter.DocumentServicePort,
) http.ServerMaker{
	return &server{
		ctx: ctx,
		cfg: cfg,
		userService: userService,
		tokenService: tokenService,
		documentService: documentService,
	}
}
func (s *server) Start(ctx context.Context) error{
	engine := html.New(viewPath, renderType)

	app := fiber.New(fiber.Config{
		ReadTimeout: time.Minute * time.Duration(s.cfg.Settings.ServerReadTimeout),
		StrictRouting: false,
		CaseSensitive: true,
		BodyLimit: 4 * 1024 * 1024,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		AppName: "ator_gold",
		Immutable: true,
		Views: engine,
		ViewsLayout: "layouts/main",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			var e *fiber.Error
			if errors.As(err, &e) {
				if e.Code == fiber.StatusNotFound {
					return c.Render("404", fiber.Map{
						"Title" : "Page Not Found",
					})
				}
				return c.Status(e.Code).Render("error", fiber.Map{
					"Title": "Error",
					"Message": e.Message,
				})
			}
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"Title" : "Internal Server Error",
				"Message": "An unexpected error occured.",
			})
		},
	})
	app.Use(static.New(publicPath))
	//app.Static('/', publicPath)
	s.app = app
	fiberConnURL := fmt.Sprintf("%s:%d", s.cfg.HTTP.Host, s.cfg.HTTP.Port)

	go func() {
		zap.S().Info("Starting HTTP server on ", fiberConnURL)
		if err := s.app.Listen(fiberConnURL); err != nil{
			if errors.Is(err, std_http.ErrServerClosed){
				return 
			}
			zap.S().Fatal("server listen error: %w", err)
		}
	}()
	err := s.HTTPMiddleware()
	if err != nil{
		zap.S().Fatal("middleware error:", err)
	}
	s.SetupRouter()
	zap.S().Info("Router setup successfully")
	return nil
}

func (s *server) Close(ctx context.Context) error{
	zap.S().Info("HTTP-Server Context is done. Shutting down server...")
	if err := s.app.ShutdownWithContext(ctx); err != nil{
		zap.S().Info("server shutdown error: %w", err)
		return err
	}
	return nil
}