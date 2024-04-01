package index

import (
	"encoding/binary"
	"errors"
	"os"
	"sort"

	"github.com/viktoriina/databases-lab1/internal/models"
)

var ErrIndexDoesNotExist = errors.New("index does not exist")

func (in *Index) readIndexTable() (errClose error) {
	indexFile, err := os.OpenFile(in.filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		if err := indexFile.Close(); err != nil {
			errClose = err
		}
	}()
	fileInfo, err := indexFile.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	var numElements int
	elementSize := binary.Size(models.Index{})
	if fileSize%int64(elementSize) == 0 {
		numElements = int(fileSize) / elementSize
	} else {
		return errors.New("file size is not a multiple of element size")
	}

	defer in.mu.Unlock()
	in.mu.Lock()

	in.Table = make([]models.Index, numElements)
	if err := binary.Read(indexFile, binary.BigEndian, in.Table); err != nil {
		return err
	}

	return nil
}

func (in *Index) WriteIndexTable() (errClose error) {
	indexFile, err := os.OpenFile(in.filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		if err := indexFile.Close(); err != nil {
			errClose = err
		}
	}()
	defer in.mu.Unlock()
	in.mu.Lock()

	if err := binary.Write(indexFile, binary.BigEndian, in.Table); err != nil {
		return err
	}
	return nil
}

func (in *Index) SortIndexTable() {
	defer in.mu.Unlock()
	in.mu.Lock()

	sort.Slice(in.Table, func(i, j int) bool {
		return in.Table[i].ID < in.Table[j].ID
	})
}

func (in *Index) Add(index *models.Index) {
	defer in.mu.Unlock()
	in.mu.Lock()

	in.Table = append(in.Table, *index)
}

func (in *Index) GetIndexById(id int64) (*models.Index, error) {
	defer in.mu.Unlock()
	in.mu.Lock()

	for _, index := range in.Table {
		if index.ID == id {
			return &index, nil
		}
	}
	return nil, ErrIndexDoesNotExist
}

func (in *Index) GetIndexByOffset(offset int64) (*models.Index, error) {
	defer in.mu.Unlock()
	in.mu.Lock()

	for _, index := range in.Table {
		if index.Offset == offset {
			return &index, nil
		}
	}
	return nil, ErrIndexDoesNotExist
}

func (in *Index) UpdateIndex(id int64, newIndex *models.Index) error {
	defer in.mu.Unlock()
	in.mu.Lock()

	for i, index := range in.Table {
		if index.ID == id {
			in.Table[i] = *newIndex
			return nil
		}
	}
	return ErrIndexDoesNotExist
}
