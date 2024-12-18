package http

import (
	
	"time"

	"github.com/gofiber/fiber/v3"
)

func (s *server) SetupRouter() {
	s.webSetUp()
	s.authSetUp()
	s.documentSetUp()
}

func (s *server) webSetUp() {
	s.app.Get("/", s.HomeWeb)
	s.app.Get("/about-us", s.AboutWeb)
	s.app.Get("/contacts", s.ContactsWeb)
	s.app.Get("/products", s.ProductsListWeb)
	s.app.Get("/product-single", s.ProductSingleWeb)
	s.app.Get("/add-product", s.AddProductWeb)
		
	s.app.Get("/ping", func(c fiber.Ctx) error{
		return c.SendString("Pong")
	})
	s.app.Get("/login", s.LoginWeb, s.RateLimiter(5, time.Minute))
	//s.app.Get("/dashboard", s.DashboardWeb, s.authMiddleware)

}

func (s *server) authSetUp() {
	route := s.app.Group("/auth")
	route.Post("/login", s.Login, s.RateLimiter(5, time.Minute), s.LoginValidation)
	route.Post("/register", s.Register, s.RateLimiter(5, time.Minute), s.RegisterValidation)
}

func (s *server) documentSetUp() {
	route := s.app.Group("/documenter")
	route.Post("/main", s.CreateMainDocument)
	route.Post("/sub", s.CreateSubDocument)
	route.Post("/content", s.CreateContentDocument)
	route.Get("/all", s.GetAllDocuments)
	route.Get("/all-join", s.GetAllDocumentsByJoin)
}
