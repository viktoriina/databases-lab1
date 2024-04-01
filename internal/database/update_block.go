package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

func (db *Database) UpdateBlock(block *models.Block) error {
	index, err := db.blocksIndex.GetIndexById(block.ID)
	if index == nil || !index.Exists {
		return ErrBlockIsNotPresent
	}
	block.LogicalPresent = true
	encoded, err := utils.EncodeBinary(block)
	if err != nil {
		return errors.Wrap(err, "failed to encode block")
	}
	if err := utils.Write(encoded, index.Offset, db.cfg.BlocksFile()); err != nil {
		return errors.Wrap(err, "failed to write block")
	}
	return nil
}
