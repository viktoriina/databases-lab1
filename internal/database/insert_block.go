package database

import (
	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

var ErrBlockWithSuchIdAlreadyExists = errors.New("block with such an id already exists")

func (db *Database) InsertBlock(block *models.Block) error {
	index, err := db.blocksIndex.GetIndexById(block.ID)
	if index != nil && index.Exists {
		return ErrBlockWithSuchIdAlreadyExists
	}
	block.LogicalPresent = true
	encoded, err := utils.EncodeBinary(block)
	if err != nil {
		return errors.Wrap(err, "failed to encode block")
	}
	offset, useGarbage, err := utils.CalculateOffset(db.blocksGarbage, db.cfg.BlocksFile())
	if err := utils.Write(encoded, offset, db.cfg.BlocksFile()); err != nil {
		return errors.Wrap(err, "failed to write block")
	}
	newIndex := &models.Index{
		ID:     block.ID,
		Offset: offset,
		Exists: true,
	}
	if useGarbage {
		db.blocksGarbage = db.blocksGarbage[:len(db.blocksGarbage)-1]
		garbageIndex, err := db.blocksIndex.GetIndexByOffset(offset)
		if err != nil {
			return err
		}
		err = db.blocksIndex.UpdateIndex(garbageIndex.ID, newIndex)
		if err != nil {
			return errors.Wrap(err, "failed to update index")
		}
	} else {
		db.blocksIndex.Add(newIndex)
	}
	db.blocksIndex.SortIndexTable()
	return nil
}
