package index

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/viktoriina/databases-lab1/internal/models"
	"github.com/viktoriina/databases-lab1/internal/utils"
)

type Index struct {
	filename string
	Table    []models.Index
	mu       sync.Mutex
}

func NewIndex(filename string) (*Index, error) {
	index := &Index{filename: filename}
	return index, index.init()
}

func (in *Index) init() error {
	if err := utils.EnsureFileExists(in.filename); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to ensure %s existence", in.filename))
	}
	if err := in.readIndexTable(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read %s index table", in.filename))
	}
	return nil
}
