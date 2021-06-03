package helpers

import (
	"encoding/csv"
	"io"
	"mime/multipart"
	"strings"
)

func GetFileType(fileHeader *multipart.FileHeader) string {
	return fileHeader.Header.Get("content-Type")
}

func GetFileContent(csvFile multipart.File) ([]string, error) {
	reader := csv.NewReader(csvFile)
	var allContent []string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}

		content := strings.TrimSpace(row[0])
		if len(content) > 0 {
			allContent = append(allContent, content)
		}
	}

	return allContent, nil
}
