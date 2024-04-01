package utils

import (
	"os"
	"path"
)

func EnsureFileExists(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil
	}
	if os.IsNotExist(err) {
		dir := path.Dir(filename)
		if err := ensureDirExists(dir); err != nil {
			return err
		}
		if _, err := os.Create(filename); err != nil {
			return err
		}
	}
	return nil
}

func ensureDirExists(dir string) error {
	_, err := os.Stat(dir)
	if err != nil && !os.IsNotExist(err) {
		return nil
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
