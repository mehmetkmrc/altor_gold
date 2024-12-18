package converter

import (
	"time"

	"github.com/google/uuid"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/entity"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/model"
	"github.com/mehmetkmrc/ator_gold/internal/dto"
)

func MainDocumentCreateRequestToModel(req *dto.MainDocumentCreateRequest)(*model.MainDocument, error) {
	mainDocument := new(model.MainDocument)
	mainID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	mainDocument = &model.MainDocument{
		ID: mainID,
		MainTitle: req.MainTitle,
		Date: time.Now(),
	}
	return mainDocument, nil
}

func SubDocumentCreateRequestToModel(req *dto.SubDocumentCreateRequest)(*model.SubDocument, error) {
	subDocument := new(model.SubDocument)
	ID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	mainID, err := uuid.Parse(req.MainID)
	if err != nil{
		return nil, err
	}
	subDocument = &model.SubDocument{
		ID: ID,
		MainID: mainID,
		SubTitle: req.SubTitle,
		ProductCode: req.ProductCode,
		SubMessage: req.SubMessage,
		Asset: req.Asset,
		Date: time.Now(),
	}
	return subDocument, nil
}

func ContentDocumentCreateRequestToModel(req *dto.ContentDocumentCreateRequest)(*model.ContentDocument, error){
	document := new(model.ContentDocument)
	ID, err := uuid.NewV7()
	if err != nil{
		return nil, err
	}
	subID, err := uuid.Parse(req.SubID)
	if err != nil {
		return nil, err
	}
	document = &model.ContentDocument{
		ID: ID,
		SubID: subID,
		ColText: req.ColText,
		JewCare: req.JewCare,
		Date: time.Now(),
	}
	return document, nil
}

func GetUserModelToDto(userData *entity.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		UserID:    userData.UserID,
		Name:      userData.Name,
		Surname:   userData.Surname,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}
}
