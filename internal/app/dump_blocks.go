package app

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) dumpBlocks() error {
	blocks, err := a.db.GetBlocks(false)
	if err != nil {
		return errors.Wrap(err, "failed to dump blocks")
	}
	for _, block := range blocks {
		printFullBlock(&block)
		fmt.Println()
	}
	return nil
}

func printFullBlock(block *models.Block) {
	printBlock(block)
	fmt.Printf("Block logical presence: %t\n", block.LogicalPresent)
}
