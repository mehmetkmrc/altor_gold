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
