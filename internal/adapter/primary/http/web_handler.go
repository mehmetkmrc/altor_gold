package http

import (
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v3"
)


var year = time.Now().Year()
func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}


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

func (s *server) ProductsListWeb(c fiber.Ctx) error {
	path := "product-list"
	return c.Render(path, fiber.Map{
		"Title": "Ürün listesi",
	})
}

func (s *server) ProductSingleWeb(c fiber.Ctx) error {
	path := "product-single"
	return c.Render(path, fiber.Map{
		"Title": "Tekli Ürün",
	})
}

func (s *server) AddProductWeb(c fiber.Ctx) error {
	path := "add-product"
	return c.Render(path, fiber.Map{
		"Title": "Ürün Ekle",
	})
}


func (s *server) uploadHandler(c fiber.Ctx) error {
    allDocuments, err := s.documentService.GetAllDocumentsWithMainDocument(c.Context())
    if err != nil {
        return s.errorResponse(c, "error while trying to get all documents", err, nil, fiber.StatusBadRequest)
    }
    var mainDocs []interface{}

    for _, document := range allDocuments {
        mainDoc := map[string]interface{}{
            "id":    document.ID,
            "title": document.MainTitle,
        }
        mainDocs = append(mainDocs, mainDoc)
    }
	return c.Render("upload", fiber.Map{
		"PageTitle": "Upload Page",
		"Title":     "Welcome to Otovinn App!",
		"Year":      year,
        "AllDocuments": mainDocs,
	}, "layouts/main")
}
