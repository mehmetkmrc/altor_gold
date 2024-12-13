package http

import (
	
	"time"

	"github.com/gofiber/fiber/v3"
)

func (s *server) SetupRouter() {
	s.webSetUp()
	s.authSetUp()
	s.dashboardSetUp()
}

func (s *server) webSetUp() {
	s.app.Get("/", s.HomeWeb)
	s.app.Get("/about", s.AboutWeb)
	s.app.Get("/contacts", s.ContactsWeb)
	s.app.Get("/blog-single", s.BlogSingleWeb)
	s.app.Get("/blogs", s.BlogsWeb)
	s.app.Get("/listing-single", s.ListingSingle)
	s.app.Get("/listing", s.ListingWeb)
	s.app.Get("/projects", s.ProjectWeb)

	
	// s.app.Get("/", func(c fiber.Ctx) error{
	// 	return c.Redirect().To("/login")
	// })
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

func (s *server) dashboardSetUp(){
	route := s.app.Group("/dashboard")
	route.Get("",s.DashboardWeb, s.IsAuthorized, s.GetUserDetail, s.RateLimiter(120, time.Minute), )
}
