package database

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (db *Database) GetTransactionsPerBlock(blockId int64, logical bool) ([]models.Transaction, error) {
	var transactions []models.Transaction
	for _, index := range db.transactionsIndex.Table {
		tx, err := db.GetTransaction(index.ID, logical)
		if err != nil {
			return nil, errors.Wrap(err,
				fmt.Sprintf("failed to get tx with id %d", index.ID),
			)
		}
		if tx.BlockID == blockId {
			transactions = append(transactions, *tx)
		}
	}
	return transactions, nil
}
