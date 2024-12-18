package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/model"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
)

var (
	_ 				  documenter.DocumentServicePort = (*DocumentService)(nil)
	DocumentServiceSet			= wire.NewSet(NewDocumentService)
)

type DocumentService struct {
	mainDocumentRepo documenter.MainDocumentRepositoryPort
	subDocumentRepo documenter.SubDocumentRepositoryPort
	contentDocumentRepo documenter.ContentDocumentRepositoryPort
}

func NewDocumentService(
	mainDocumentRepo documenter.MainDocumentRepositoryPort,
	subDocumentRepo documenter.SubDocumentRepositoryPort,
	contentDocumentRepo documenter.ContentDocumentRepositoryPort) documenter.DocumentServicePort {
	return &DocumentService{
		mainDocumentRepo: mainDocumentRepo,
		subDocumentRepo: subDocumentRepo,
		contentDocumentRepo: contentDocumentRepo,
	}
}

func (s *DocumentService) AddMainDocument(ctx context.Context, document *model.MainDocument)(*model.MainDocument, error) {
	return s.mainDocumentRepo.Insert(ctx, document)
}

func (s *DocumentService) AddSubDocument(ctx context.Context, document *model.SubDocument)(*model.SubDocument, error) {
	return s.subDocumentRepo.Insert(ctx, document)
}

func (s *DocumentService) AddContentDocument(ctx context.Context, document *model.ContentDocument)(*model.ContentDocument, error) {
	return s.contentDocumentRepo.Insert(ctx, document)
}

func (s *DocumentService) GetAllDocuments(ctx context.Context)([]*model.MainDocument, error) {
	mainDocuments, err := s.mainDocumentRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	subDocuments, err := s.subDocumentRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	contentDocuments, err := s.contentDocumentRepo.GetAll(ctx)
	if err != nil{
		return nil, err
	}
	subDocMap := make(map[uuid.UUID]*model.SubDocument)
	for _, subDocument := range subDocuments {
		subDocMap[subDocument.ID] = subDocument
		subDocument.ContentDocuments = []*model.ContentDocument{}
	}

	for _, contentDocument := range contentDocuments {
		if subDocument, ok := subDocMap[contentDocument.SubID]; ok {
			subDocument.ContentDocuments = append(subDocument.ContentDocuments, contentDocument)
		}
	}

	mainDocMap := make(map[uuid.UUID]*model.MainDocument)
	for _, mainDocument := range mainDocuments {
		mainDocMap[mainDocument.ID] = mainDocument
		mainDocument.SubDocuments = []*model.SubDocument{}
	}

	for _, subDocument := range subDocuments {
		if mainDocument, ok := mainDocMap[subDocument.MainID]; ok {
			mainDocument.SubDocuments = append(mainDocument.SubDocuments, subDocument)
		}
	}

	return mainDocuments, nil

}

func (s *DocumentService) GetAllDocumentsWithMainDocument(ctx context.Context) ([]*model.MainDocument, error) {
	return s.mainDocumentRepo.GetAllDocumentsByJoin(ctx)
}
