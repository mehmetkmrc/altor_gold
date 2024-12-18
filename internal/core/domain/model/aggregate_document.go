package model

type AggregateDocument struct {
	MainDocument     *MainDocument    `json:"main_document"`
	SubDocuments     *SubDocument     `json:"sub_document"`
	ContentDocuments *ContentDocument `json:"content_document"`
}
