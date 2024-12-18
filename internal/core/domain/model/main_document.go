package model

import (
	"time"

	"github.com/google/uuid"
)

type MainDocument struct {
	ID uuid.UUID	`json:"id"`
	MainTitle string `json:"title"`
	Status uint8 `json:"status"`
	Position uint8 `json:"position"`
	Date time.Time `json:"date"`
	SubDocuments []*SubDocument `json:"sub_documents"`
}