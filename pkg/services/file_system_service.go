package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileSystemService struct {
}

func NewFileSystemService() *FileSystemService {
	return &FileSystemService{}
}

func (r FileSystemService) Walk(sourceDir string, pattern string, name string) error {
	fileInfo, err := os.Stat(sourceDir)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a dir", fileInfo.Name())
	}

	gitDirFolder := filepath.Join(sourceDir, string(os.PathSeparator), ".git/")
	gitDirInfo, err := os.Stat(gitDirFolder)
	if err != nil {
		return err
	}
	if gitDirInfo.IsDir() {
		rErr := os.RemoveAll(gitDirFolder)
		if rErr != nil {
			return rErr
		}
	}

	err = filepath.Walk(sourceDir, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			file, fileErr := os.Open(path)
			if fileErr != nil {
				return fileErr
			}
			bytes, rErr := os.ReadFile(file.Name())
			if rErr != nil {
				return rErr
			}
			replaced := strings.ReplaceAll(string(bytes), pattern, name)
			wErr := os.WriteFile(path, []byte(replaced), 0)
			if wErr != nil {
				return wErr
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if err = os.Mkdir(name, 0777); err != nil {
		fmt.Println(err)
	}

	return nil
}
