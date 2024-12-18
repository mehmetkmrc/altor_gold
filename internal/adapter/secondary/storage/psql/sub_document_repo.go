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
	_                        documenter.SubDocumentRepositoryPort = (*SubDocumentRepository)(nil)
	SubDocumentRepositorySet                                      = wire.NewSet(NewSubDocumentRepository)
)

type SubDocumentRepository struct {
	dbPool *pgxpool.Pool
}

func NewSubDocumentRepository(em db.EngineMaker) documenter.SubDocumentRepositoryPort {
	return &SubDocumentRepository{
		dbPool: em.GetDB(),
	}
}

func (q *SubDocumentRepository) Insert(ctx context.Context, documentModel *model.SubDocument) (*model.SubDocument, error) {
	query := `INSERT INTO doc_sub (id, main_id, sub_title, product_code, sub_message, asset, status, date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, main_id, sub_title, product_code, sub_message, asset, status, date`
	queryRow := q.dbPool.QueryRow(ctx, query, documentModel.ID, documentModel.MainID, documentModel.SubTitle, documentModel.ProductCode, documentModel.SubMessage, documentModel.Asset, documentModel.Status, documentModel.Date)
	err := queryRow.Scan(&documentModel.ID, &documentModel.MainID, &documentModel.SubTitle, &documentModel.ProductCode, &documentModel.SubMessage, &documentModel.Asset, &documentModel.Status, &documentModel.Date)
	if err != nil {
		return nil, err
	}
	return documentModel, nil
}

func (q *SubDocumentRepository) GetAll(ctx context.Context) ([]*model.SubDocument, error) {
	query := `SELECT * FROM doc_sub`
	queryRows, err := q.dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var documents []*model.SubDocument
	for queryRows.Next() {
		document := new(model.SubDocument)
		err = queryRows.Scan(&document.ID, &document.MainID, &document.ProductCode, &document.SubMessage, &document.Asset, &document.Position, &document.Status, &document.Date)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}
