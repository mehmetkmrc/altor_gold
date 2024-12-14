package http

import (
	"github.com/gofiber/fiber/v3"
	
)

func (s *server) LoginWeb(c fiber.Ctx) error {
	path := "login"
	return c.Render(path, fiber.Map{
		"Title": "Login",
	})
}


func (s *server) HomeWeb(c fiber.Ctx) error {
	path := "home"
	return c.Render(path, fiber.Map{
		"Title": "Ator Gold - Anasayfa",
	})
}

func (s *server) AboutWeb(c fiber.Ctx) error {
	path := "about-us"
	return c.Render(path, fiber.Map{
		"Title": "Hakkımızda",
	})
}

func (s *server) ContactsWeb(c fiber.Ctx) error {
	path := "contacts"
	return c.Render(path, fiber.Map{
		"Title": "İletişim",
	})
}

func (s *server) BlogSingleWeb(c fiber.Ctx) error {
	path := "blog-single"
	return c.Render(path, fiber.Map{
		"Title": "Tek Haberler",
	})
}
func (s *server) BlogsWeb(c fiber.Ctx) error {
	path := "blogs"
	return c.Render(path, fiber.Map{
		"Title": "Haberler",
	})
}
func (s *server) ListingSingle(c fiber.Ctx) error {
	path := "listing-single"
	return c.Render(path, fiber.Map{
		"Title": "Daire",
	})
}

func (s *server) ListingWeb(c fiber.Ctx) error {
	path := "listing"
	return c.Render(path, fiber.Map{
		"Title": "Daireler",
	})
}
func (s *server) ProjectWeb(c fiber.Ctx) error {
	path := "products"
	return c.Render(path, fiber.Map{
		"Title": "Projeler",
	})
}

func (s *server) DashboardWeb(c fiber.Ctx) error {

	//user_ID := c.Params("user_id")
	
	path := "dashboard"
	return c.Render(path, fiber.Map{
		"Title": "Dashboard",
	})
}
