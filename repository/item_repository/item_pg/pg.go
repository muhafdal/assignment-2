package item_pg

import (
	"database/sql"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository"
)

type itemPG struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) item_repository.Repository {
	return &itemPG{
		db: db,
	}
}

func (itemPG *itemPG) GetItemsByCodes(itemCodes []string) ([]entity.Item, errs.Error) {
	query := itemPG.generateGetItemsByCodesQuery(len(itemCodes))

	itemCodeDbArg := []any{}

	for _, itemCode := range itemCodes {
		itemCodeDbArg = append(itemCodeDbArg, itemCode)
	}

	rows, err := itemPG.db.Query(query, itemCodeDbArg...)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	items := []entity.Item{}

	for rows.Next() {

		item := entity.Item{}

		err = rows.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.OrderId, &item.CreatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
