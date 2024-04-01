package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

func (db *Database) DeleteTransaction(id int64) error {
	tx, err := db.GetTransaction(id, false)
	if err != nil {
		return errors.Wrap(err, "failed to get transaction")
	}
	tx.LogicalPresent = false
	encoded, err := utils.EncodeBinary(tx)
	if err != nil {
		return errors.Wrap(err, "failed to encode transaction")
	}
	index, _ := db.transactionsIndex.GetIndexById(id)
	if err := utils.Write(encoded, index.Offset, db.cfg.TransactionsFile()); err != nil {
		return errors.Wrap(err, "failed to write transaction")
	}
	db.transactionsGarbage = append(db.transactionsGarbage, index.Offset)
	index.Exists = false
	return db.transactionsIndex.UpdateIndex(index.ID, index)
}
