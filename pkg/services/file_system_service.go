package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

type FileSystemService struct {
	foldersToDelete []string
}

func NewFileSystemService() *FileSystemService {
	return &FileSystemService{
		foldersToDelete: []string{
			".git/",
		},
	}
}

func (r FileSystemService) WalkDir(sourceDir string, pattern string, name string) error {
	fileInfo, err := os.Stat(sourceDir)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a dir", fileInfo.Name())
	}

	for i := 0; i < len(r.foldersToDelete); i++ {
		err = r.removeDir(sourceDir, r.foldersToDelete[i])
		if err != nil {
			return err
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
			err = r.applyChange(path, file, pattern, name)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if err = os.Mkdir(name, 0777); err != nil {
		return fmt.Errorf("dir %s is not empty", name)
	}

	err = cp.Copy(sourceDir, name)
	if err != nil {
		return err
	}

	return nil
}

func (r FileSystemService) applyChange(path string, file *os.File, pattern string, name string) error {
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}

	replaced := strings.ReplaceAll(string(bytes), pattern, name)
	err = os.WriteFile(path, []byte(replaced), 0)
	if err != nil {
		return err
	}

	return nil
}

func (r FileSystemService) removeDir(sourceDir string, dir string) error {
	dirFolder := filepath.Join(sourceDir, string(os.PathSeparator), dir)
	fileInfo, err := os.Stat(dirFolder)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		err = os.RemoveAll(dirFolder)
		if err != nil {
			return err
		}
	}

	return nil
}
