package psql

import (
	"context"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/model"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
)

var (
	_ 	documenter.ContentDocumentRepositoryPort = (*ContentDocumentRepository)(nil)
	ContentRepositorySet				= wire.NewSet(NewContentDocumentRepository)
)

type ContentDocumentRepository struct {
	dbPool *pgxpool.Pool
}

func NewContentDocumentRepository(em db.EngineMaker) documenter.ContentDocumentRepositoryPort{
	return &ContentDocumentRepository{
		dbPool: em.GetDB(),
	}
}

func (q *ContentDocumentRepository) Insert(ctx context.Context, documentModel *model.ContentDocument) (*model.ContentDocument, error) {
	query := `INSERT INTO doc_content (id, sub_id, about_collection, jewellery_care, status, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, sub_id, about_collection, jewellery_care, status, date`
	queryRow := q.dbPool.QueryRow(ctx, query, documentModel.ID, documentModel.SubID, documentModel.ColText, documentModel.JewCare, documentModel.Status, documentModel.Date)
	err := queryRow.Scan(&documentModel.ID, &documentModel.SubID, &documentModel.ColText, &documentModel.JewCare, &documentModel.Status, &documentModel.Date)
	if err != nil {
		return nil, err
	}
	return documentModel, nil
}

func (q *ContentDocumentRepository) GetAll(ctx context.Context) ([]*model.ContentDocument, error) {
	query := `SELECT * FROM doc_content`
	queryRows, err := q.dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var documents []*model.ContentDocument
	for queryRows.Next() {
		document := new(model.ContentDocument)
		err = queryRows.Scan(&document.ID, &document.SubID, &document.ColText, &document.JewCare, &document.Position, &document.Status, &document.Date)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}
