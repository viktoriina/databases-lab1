package database

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (db *Database) GetBlocks(logical bool) ([]models.Block, error) {
	var blocks []models.Block
	for _, index := range db.blocksIndex.Table {
		block, err := db.GetBlock(index.ID, logical)
		if err != nil {
			return nil, errors.Wrap(err,
				fmt.Sprintf("failed to get block with id %d", index.ID),
			)
		}
		blocks = append(blocks, *block)
	}
	return blocks, nil
}
