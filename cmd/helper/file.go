package helper

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFolder(sourceFolder string, destinationFolder string) error {
	err := filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		destinationPath := filepath.Join(destinationFolder, relativePath)

		if info.IsDir() {
			err = os.MkdirAll(destinationPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			destinationFile, err := os.Create(destinationPath)
			if err != nil {
				return err
			}
			defer destinationFile.Close()

			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
