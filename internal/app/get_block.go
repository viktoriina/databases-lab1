package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
)

func (a *App) getBlock() error {
	id, err := enterBlockId()
	if err != nil {
		return errors.Wrap(err, "invalid block id")
	}
	block, err := a.db.GetBlock(id, true)
	if err != nil {
		return err
	}
	fmt.Println()
	printBlock(block)
	return nil
}

func enterBlockId() (int64, error) {
	fmt.Print("Enter block id: ")
	reader := bufio.NewReader(os.Stdin)
	idString, _ := reader.ReadString('\n')
	idString = strings.Replace(idString, "\n", "", replacementLimit)
	id, err := strconv.Atoi(idString)
	const invalidId = -1
	if err != nil || id < 1 {
		return invalidId, err
	}
	return int64(id), nil
}

func printBlock(block *models.Block) {
	fmt.Println("Block id:", block.ID)
	fmt.Println("Block timestamp:", block.Timestamp)
	fmt.Println("Block nonce:", block.Nonce)
}
