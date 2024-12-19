package model

import (
	"time"

	"github.com/google/uuid"
)

type SubDocument struct {
	ID  	uuid.UUID	`json:"id"`
	MainID 	uuid.UUID	`json:"main_id"`
	SubTitle string 	`json:"sub_title"`
	ProductCode string 	`json:"product_code"`
	SubMessage string 	`json:"sub_message"`//Ürün Özellikleri
	Asset [][]byte		`json:"asset"`
	Position uint8		`json:"position"`
	Status	uint8 		`json:"status"`
	Date 	time.Time 	`json:"date"`
	ContentDocuments []*ContentDocument `json:"content_documents"`
}