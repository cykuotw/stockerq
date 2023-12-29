package stock_price

import (
	"fmt"
	apperror "stocker-hf-data/web/app/app-error"
	"stocker-hf-data/web/app/model"

	"github.com/google/uuid"
)

// UpdateStockPrice returns the rows that affected
// by update statement
func UpdateStockPrice(price StockPrice) (int64, uuid.UUID, *apperror.ModelError) {
	if !price.isValid() {
		return 0, uuid.UUID{}, apperror.NewModelError(apperror.ErrInputPriceNotValid)
	}

	db := model.GetDB()

	id, err := uuid.Parse(price.Uuid)

	if err != nil {
		return 0, uuid.UUID{}, apperror.NewModelError(err)
	}

	tx, err := db.Begin()

	selectStatment := fmt.Sprintf(`
		SELECT COUNT(*) FROM stock_price 
		WHERE uuid = '%s'
		FOR UPDATE;
	`, id.String())
	rows, err := db.Query(selectStatment)

	if err != nil {
		tx.Rollback()
		return 0, uuid.UUID{}, apperror.NewModelError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		err := rows.Scan(&count)
		if err != nil {
			break
		}
		if count < 1 {
			tx.Rollback()
			return 0, uuid.UUID{}, apperror.NewModelError(
				fmt.Errorf("%s: %w", id.String(), apperror.ErrIdNotExist))
		}
	}

	updateStatment := fmt.Sprintf(`
	UPDATE stock_price 
	SET company_id = '%s', 
		update_date = '%s', 
		price_date = '%s', 
		open = %d, 
		close = %d, 
		high = %d, 
		low = %d,
		price_change = %d, 
		change_percent = %d, 
		volume = %d, 
		amount = %d
	WHERE uuid = '%s';
	`, price.CompanyID,
		price.UpdateDate.Format("2006-01-02 15:04:05"),
		price.PriceDate.Format("2006-01-02"),
		price.Open, price.Close, price.High, price.Low,
		price.PriceChange, price.ChangePercent,
		price.Volume, price.Amount, id.String())
	stmt, err := tx.Prepare(updateStatment)
	if err != nil {
		tx.Rollback()
		return 0, uuid.UUID{}, apperror.NewModelError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		tx.Rollback()
		return 0, uuid.UUID{}, apperror.NewModelError(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		tx.Rollback()
		return 0, uuid.UUID{}, apperror.NewModelError(err)
	}

	tx.Commit()

	return rowsAffected, id, nil
}
