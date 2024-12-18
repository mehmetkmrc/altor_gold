package psql


import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/model"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/documenter"
)

var (
	_                 documenter.MainDocumentRepositoryPort = (*MainDocumentRepository)(nil)
	MainRepositorySet                                       = wire.NewSet(NewMainDocumentRepository)
)

type MainDocumentRepository struct {
	dbPool *pgxpool.Pool
}

func NewMainDocumentRepository(em db.EngineMaker) documenter.MainDocumentRepositoryPort {
	return &MainDocumentRepository{
		dbPool: em.GetDB(),
	}
}
func (q *MainDocumentRepository) Insert(ctx context.Context, documentModel *model.MainDocument) (*model.MainDocument, error) {
	query := `INSERT INTO doc_main (id, main_title, status, date) VALUES ($1, $2, $3, $4) RETURNING id, main_title, status, date`
	queryRow := q.dbPool.QueryRow(ctx, query, documentModel.ID, documentModel.MainTitle, documentModel.Status, documentModel.Date)
	err := queryRow.Scan(&documentModel.ID, &documentModel.MainTitle, &documentModel.Status, &documentModel.Date)
	if err != nil {
		return nil, err
	}
	return documentModel, nil
}

func (q *MainDocumentRepository) GetAll(ctx context.Context) ([]*model.MainDocument, error) {
	query := `SELECT * FROM doc_main`
	queryRows, err := q.dbPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var documents []*model.MainDocument
	for queryRows.Next() {
		document := new(model.MainDocument)
		err = queryRows.Scan(&document.ID, &document.MainTitle, &document.Position, &document.Status, &document.Date)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}

func (q *MainDocumentRepository) GetAllDocumentsByJoin(ctx context.Context) ([]*model.MainDocument, error) {
	rows, err := q.dbPool.Query(ctx, `
        SELECT 
            dm.id as main_id, 
            dm.title, 
            dm.position as main_position, 
            dm.status as main_status, 
            dm.date as main_date,
            ds.id as sub_id,
            ds.main_id as sub_main_id,
            ds.sub_title,
			ds.product_code,
			ds.sub_message, 
            ds.asset as sub_asset, 
            ds.position as sub_position, 
            ds.status as sub_status, 
            ds.date as sub_date,
            dc.id as content_id, 
            dc.sub_id as content_sub_id,
            dc.about_collection,
			dc.jewellery_care, 
            dc.position as content_position, 
            dc.status as content_status, 
            dc.date as content_date
        FROM 
            doc_main dm
        LEFT JOIN 
            doc_sub ds ON dm.id = ds.main_id
        LEFT JOIN 
            doc_content dc ON ds.id = dc.sub_id
        ORDER BY 
            dm.position, ds.position, dc.position;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mainDocMap := make(map[uuid.UUID]*model.MainDocument)
	subDocMap := make(map[uuid.UUID]*model.SubDocument)

	for rows.Next() {
		var mainDocument model.MainDocument
		var subDocument model.SubDocument
		var contentDocument model.ContentDocument

		err := rows.Scan(
			&mainDocument.ID, &mainDocument.MainTitle, &mainDocument.Position, &mainDocument.Status, &mainDocument.Date,
			&subDocument.ID, &subDocument.MainID, &subDocument.SubTitle, &subDocument.ProductCode, &subDocument.SubMessage, &subDocument.Asset, &subDocument.Position, &subDocument.Status, &subDocument.Date,
			&contentDocument.ID, &contentDocument.SubID, &contentDocument.ColText, &contentDocument.JewCare, &contentDocument.Position, &contentDocument.Status, &contentDocument.Date,
		)
		if err != nil {
			return nil, err
		}

		if mainDocMap[mainDocument.ID] == nil {
			mainDocument.SubDocuments = []*model.SubDocument{}
			mainDocMap[mainDocument.ID] = &mainDocument
		}

		if subDocument.ID != uuid.Nil {
			if subDocMap[subDocument.ID] == nil {
				subDocument.ContentDocuments = []*model.ContentDocument{}
				subDocMap[subDocument.ID] = &subDocument
				mainDocMap[mainDocument.ID].SubDocuments = append(mainDocMap[mainDocument.ID].SubDocuments, &subDocument)
			}

			if contentDocument.ID != uuid.Nil {
				subDocMap[subDocument.ID].ContentDocuments = append(subDocMap[subDocument.ID].ContentDocuments, &contentDocument)
			}
		}
	}

	mainDocuments := make([]*model.MainDocument, 0, len(mainDocMap))
	for _, mainDoc := range mainDocMap {
		mainDocuments = append(mainDocuments, mainDoc)
	}

	return mainDocuments, nil
}
