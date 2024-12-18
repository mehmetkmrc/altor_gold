package http

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/ator_gold/internal/converter"
	"github.com/mehmetkmrc/ator_gold/internal/dto"
	"go.uber.org/zap"
)

func (s *server) CreateMainDocument(c fiber.Ctx) error {
	reqBody := new(dto.MainDocumentCreateRequest)
	body := c.Body()
	if err := json.Unmarshal(body, reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	documentModel, err := converter.MainDocumentCreateRequestToModel(reqBody)
	if err != nil {
		return s.errorResponse(c, "error while trying to convert document create request to model", err, nil, fiber.StatusBadRequest)
	}
	document, err := s.documentService.AddMainDocument(c.Context(), documentModel)
	if err != nil {
		return s.errorResponse(c, "error while trying to create document", err, nil, fiber.StatusBadRequest)
	}

	zap.S().Info("Document Created Successfully! Document:", document)
	return s.successResponse(c, documentModel.ID, "document created successfully",  fiber.StatusOK)
}

func (s *server) CreateSubDocument(c fiber.Ctx) error {
	reqBody := new(dto.SubDocumentCreateRequest)
	body := c.Body()
	if err := json.Unmarshal(body, reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	documentModel, err := converter.SubDocumentCreateRequestToModel(reqBody)
	if err != nil {
		return s.errorResponse(c, "error while trying to convert document create request to model", err, nil, fiber.StatusBadRequest)
	}
	document, err := s.documentService.AddSubDocument(c.Context(), documentModel)
	if err != nil {
		return s.errorResponse(c, "error while trying to create document", err, nil, fiber.StatusBadRequest)
	}

	zap.S().Info("Document Created Successfully! Document:", document)
	return s.successResponse(c, documentModel.ID, "document created successfully", fiber.StatusOK)
}

func (s *server) CreateContentDocument(c fiber.Ctx) error {
	reqBody := new(dto.ContentDocumentCreateRequest)
	body := c.Body()
	if err := json.Unmarshal(body, reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	documentModel, err := converter.ContentDocumentCreateRequestToModel(reqBody)
	if err != nil {
		return s.errorResponse(c, "error while trying to convert document create request to model", err, nil, fiber.StatusBadRequest)
	}
	document, err := s.documentService.AddContentDocument(c.Context(), documentModel)
	if err != nil {
		return s.errorResponse(c, "error while trying to create document", err, nil, fiber.StatusBadRequest)
	}

	zap.S().Info("Document Created Successfully! Document:", document)
	return s.successResponse(c, nil, "document created successfully", fiber.StatusOK)
}

func (s *server) GetAllDocuments(c fiber.Ctx) error {
	documents, err := s.documentService.GetAllDocuments(c.Context())
	if err != nil {
		return s.errorResponse(c, "error while trying to get all documents", err, nil, fiber.StatusBadRequest)
	}

	return s.successResponse(c, documents, "documents fetched successfully", fiber.StatusOK)
}

func (s *server) GetAllDocumentsByJoin(c fiber.Ctx) error {
	documents, err := s.documentService.GetAllDocumentsWithMainDocument(c.Context())
	if err != nil {
		return s.errorResponse(c, "error while trying to get all documents", err, nil, fiber.StatusBadRequest)
	}

	return s.successResponse(c, documents, "documents fetched successfully", fiber.StatusOK)
}
