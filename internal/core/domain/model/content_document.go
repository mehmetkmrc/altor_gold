package model

import (
	"github.com/google/uuid"
	"time"
)

type ContentDocument struct {
	ID       uuid.UUID `json:"id"`
	SubID    uuid.UUID `json:"sub_id"`
	ColText  string    `json:"about_collection"`
	JewCare string     `json:"jewellery_care"`
	Position uint8     `json:"position"`
	Status   uint8     `json:"status"`
	Date     time.Time `json:"date"`
}
