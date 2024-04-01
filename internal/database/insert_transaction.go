package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

var ErrTransactionWithSuchIdAlreadyExists = errors.New("transaction with such an id already exists")

func (db *Database) InsertTransaction(tx *models.Transaction) error {
	index, err := db.transactionsIndex.GetIndexById(tx.ID)
	if index != nil && index.Exists {
		return ErrTransactionWithSuchIdAlreadyExists
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
	offset, useGarbage, err := utils.CalculateOffset(db.transactionsGarbage, db.cfg.TransactionsFile())
	if err := utils.Write(encoded, offset, db.cfg.TransactionsFile()); err != nil {
		return errors.Wrap(err, "failed to write transaction")
	}
	newIndex := &models.Index{
		ID:     tx.ID,
		Offset: offset,
		Exists: true,
	}
	if useGarbage {
		db.transactionsGarbage = db.transactionsGarbage[:len(db.transactionsGarbage)-1]
		garbageIndex, err := db.transactionsIndex.GetIndexByOffset(offset)
		if err != nil {
			return err
		}
		err = db.transactionsIndex.UpdateIndex(garbageIndex.ID, newIndex)
		if err != nil {
			return errors.Wrap(err, "failed to update transactions index")
		}
	} else {
		db.transactionsIndex.Add(newIndex)
	}
	return nil
}
