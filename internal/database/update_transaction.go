package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

func (db *Database) UpdateTransaction(tx *models.Transaction) error {
	index, err := db.transactionsIndex.GetIndexById(tx.ID)
	if index == nil || !index.Exists {
		return ErrTransactionIsNotPresent
	}
	blockIndex, err := db.blocksIndex.GetIndexById(tx.BlockID)
	if blockIndex == nil || !blockIndex.Exists {
		return ErrBlockIsNotPresent
	}
	tx.LogicalPresent = true
	encoded, err := utils.EncodeBinary(tx)
	if err != nil {
		return errors.Wrap(err, "failed to encode transaction")
	}
	if err := utils.Write(encoded, index.Offset, db.cfg.TransactionsFile()); err != nil {
		return errors.Wrap(err, "failed to write transaction")
	}
	return nil
}
