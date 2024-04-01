package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

func (db *Database) DeleteBlock(id int64) (errClose error) {
	if err := db.deleteTransactionsPerBlock(id); err != nil {
		return err
	}
	block, err := db.GetBlock(id, false)
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}
	block.LogicalPresent = false
	encoded, err := utils.EncodeBinary(block)
	if err != nil {
		return errors.Wrap(err, "failed to encode block")
	}
	index, _ := db.blocksIndex.GetIndexById(id)
	if err := utils.Write(encoded, index.Offset, db.cfg.BlocksFile()); err != nil {
		return errors.Wrap(err, "failed to write block")
	}
	db.blocksGarbage = append(db.blocksGarbage, index.Offset)
	index.Exists = false
	return db.blocksIndex.UpdateIndex(index.ID, index)
}

func (db *Database) deleteTransactionsPerBlock(blockID int64) error {
	transactions, err := db.GetTransactionsPerBlock(blockID, true)
	if err != nil {
		return errors.Wrap(err, "failed to get transactions per block")
	}
	for _, tx := range transactions {
		if err := db.DeleteTransaction(tx.ID); err != nil {
			return errors.Wrap(err, "failed to delete transaction")
		}
	}
	return nil
}
