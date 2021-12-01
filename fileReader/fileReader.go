package fileReader

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
)

func ReadStringLines(path string) ([]string, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadIntLines(path string) ([]int, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, number)
	}
	return lines, scanner.Err()
}

func openFile(filePath string) (*os.File, error) {
	_, callerFileName, _, _ := runtime.Caller(2)
	fileRealPath := path.Join(path.Dir(callerFileName), filePath)
	file, err := os.Open(fileRealPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}