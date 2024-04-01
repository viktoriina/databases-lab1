package utils

import "os"

func Write(value []byte, offset int64, filename string) (errClose error) {
	file, err := os.OpenFile(filename, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			errClose = err
		}
	}()
	if _, err := file.WriteAt(value, offset); err != nil {
		return err
	}
	return nil
}
