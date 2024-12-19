package http

import (
	"encoding/base64"
	"strings"

	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)


var year = time.Now().Year()
func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}


func (s *server) LoginWeb(c fiber.Ctx) error {
	path := "login"
	return c.Render(path, fiber.Map{
		"Title": "Login",
	},"layouts/back-layout")
}


func (s *server) HomeWeb(c fiber.Ctx) error {
	path := "home"
	return c.Render(path, fiber.Map{
		"Title": "Ator Gold - Anasayfa",
	},"layouts/main")
}

func (s *server) AboutWeb(c fiber.Ctx) error {
	path := "about-us"
	return c.Render(path, fiber.Map{
		"Title": "Hakkımızda",
	},"layouts/main")
}

func (s *server) ContactsWeb(c fiber.Ctx) error {
	path := "contacts"
	return c.Render(path, fiber.Map{
		"Title": "İletişim",
	},"layouts/main")
}


func (s *server) ProductsListWeb(c fiber.Ctx) error {
	path := "product-list"
	return c.Render(path, fiber.Map{
		"Title": "Ürün listesi",
	},"layouts/main")
}



func (s *server) AddProductWeb(c fiber.Ctx) error {
	path := "add-product"
	return c.Render(path, fiber.Map{
		"Title": "Ürün Ekle",
	},"layouts/main")
}





func (s *server) ProductSingleWeb(c fiber.Ctx) error {
	mainIDStr := c.Params("main_id")
	mainID, err := uuid.Parse(mainIDStr)
	if err != nil{
		return s.errorResponse(c, "invalid main_id", err, nil, fiber.StatusBadRequest)
	}

	documents, err := s.documentService.GetAllDocumentsWithMainDocument(c.Context())
	if err != nil {
		return s.errorResponse(c, "error while trying to get all documents", err, nil, fiber.StatusBadRequest)
	}

	allDocuments, err := s.documentService.GetAllDocumentsWithMainDocument(c.Context())
	if err != nil {
		return s.errorResponse(c, "error while trying to get all documents", err, nil, fiber.StatusBadRequest)
	}
	var mainDocs []interface{}
	var title string // Title için değişken
	

	for _, document := range allDocuments {
		mainDoc := map[string]interface{}{
			"id": document.ID,
			"title": document.MainTitle,
		}
		mainDocs = append(mainDocs, mainDoc)
	}

	var filteredDocuments []interface{}
	//Sadece belirtilen main_id'ye ait belgeleri filtrele
	for _, document := range documents{
		if document.ID == mainID {
			title = strings.ToUpper(document.MainTitle) // Title'ı burada al
			mainDoc := map[string]interface{}{
				"id":		document.ID,
				"title":	document.MainTitle,
				"status":	document.Status,
				"position":	document.Position,
				"date":		document.Date,
			}

			var subDocs []interface{}
			for _, subDoc := range document.SubDocuments{
				var assets []string
				for _, asset := range subDoc.Asset {
					encodedAsset := "data:" + http.DetectContentType(asset) + ";base64," + encodeBase64(asset)
					assets = append(assets, encodedAsset)
				}

				subDocument := map[string]interface{}{
					"id":	subDoc.ID,
					"main_id": subDoc.MainID,
					"sub_title":	subDoc.SubTitle,
					"product_code":	subDoc.ProductCode,
					"sub_message": subDoc.SubMessage,
					"asset": assets,
					"status": subDoc.Status,
					"date":	subDoc.Date,
				}

				var contentDocs []interface{}
				for _, contentDoc := range subDoc.ContentDocuments {
					contentDocument := map[string]interface{}{
						"id":	contentDoc.ID,
						"sub_id": contentDoc.SubID, 
						"about_collection": contentDoc.ColText,
						"jewellery_care": contentDoc.JewCare,
						"position": contentDoc.Position,
						"status": contentDoc.Status,
						"date": contentDoc.Date,
					}
					contentDocs = append(contentDocs, contentDocument)
				}
				subDocument["ContentDocuments"] = contentDocs
                subDocs = append(subDocs, subDocument)

			}
			mainDoc["SubDocuments"] = subDocs
				filteredDocuments = append(filteredDocuments, mainDoc)
				break // Sadece ilgili main_id yi al
		}
	}
	path := "product-single"
	return c.Render(path, fiber.Map{
		"PageTitle": "Tekli Ürün",
		"Title":	title,
		"Year":		year,
		"FilteredDocuments": filteredDocuments,
		"AllDocuments": mainDocs,
	},"layouts/main")
	
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
