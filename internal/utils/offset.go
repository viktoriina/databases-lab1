package utils

import "os"

func CalculateOffset(garbage []int64, filename string) (int64, bool, error) {
	var offset int64
	const invalidAddress = -1
	useGarbage := len(garbage) != 0
	if useGarbage {
		offset = garbage[len(garbage)-1]
	} else {
		fileInfo, err := os.Stat(filename)
		if err != nil {
			return invalidAddress, false, err
		}
		offset = fileInfo.Size()
	}
	return offset, useGarbage, nil
}
