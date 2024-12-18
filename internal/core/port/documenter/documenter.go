package documenter

import (
	"context"

	"github.com/mehmetkmrc/ator_gold/internal/core/domain/model"
)

type MainDocumentRepositoryPort interface {
	Insert(ctx context.Context, document *model.MainDocument) (*model.MainDocument, error)
	GetAll(ctx context.Context)([]*model.MainDocument, error)
	GetAllDocumentsByJoin(ctx context.Context)([]*model.MainDocument, error)
	//Update(ctx context.Context, document *model.MainDocument) (*model.MainDocument, error)
	//GetOneWithTitle(ctx context.Context, title string) (model.MainDocument, error)
	//Delete(ctx context.Context, title string) error
	//DeleteAll(ctx context.Context) error
}

type SubDocumentRepositoryPort interface {
	Insert(ctx context.Context, document *model.SubDocument)(*model.SubDocument, error)
	GetAll(ctx context.Context)([]*model.SubDocument, error)
	//Update(ctx context.Context, document *model.SubDocument) (*model.SubDocument, error)
	//GetOneWithTitle(ctx context.Context, title string) (model.SubDocument, error)
	//Delete(ctx context.Context, title string) error
	//DeleteAll(ctx context.Context) error
}

type ContentDocumentRepositoryPort interface {
	Insert(ctx context.Context, document *model.ContentDocument)(*model.ContentDocument, error)
	GetAll(ctx context.Context)([]*model.ContentDocument, error)
	//Update(ctx context.Context, document *model.ContentDocument) (*model.ContentDocument, error)
	//GetOneWithTitle(ctx context.Context, title string) (model.ContentDocument, error)
	//Delete(ctx context.Context, title string) error
	//DeleteAll(ctx context.Context) error
}

type DocumentServicePort interface {
	AddMainDocument(ctx context.Context, document *model.MainDocument)(*model.MainDocument, error)
	AddSubDocument(ctx context.Context, document *model.SubDocument)(*model.SubDocument, error)
	AddContentDocument(ctx context.Context, document *model.ContentDocument)(*model.ContentDocument, error)
	GetAllDocuments(ctx context.Context)([]*model.MainDocument, error)
	GetAllDocumentsWithMainDocument(ctx context.Context)([]*model.MainDocument, error)
	//UpdateMainDocument(ctx context.Context, document *model.MainDocument) (*model.MainDocument, error)
	//GetMainDocument(ctx context.Context, title string) (model.MainDocument, error)
	//DeleteMainDocument(ctx context.Context, title string) error
	//DeleteAllMains(ctx context.Context) error
}