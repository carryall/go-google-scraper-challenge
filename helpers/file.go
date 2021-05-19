package helpers

import (
	"encoding/csv"
	"io"
	"mime/multipart"
)

func GetFileType(fileHeader *multipart.FileHeader) string {
	return fileHeader.Header.Get("content-Type")
}

func GetFileContent(csvFile multipart.File) ([]string, error) {
	reader := csv.NewReader(csvFile)
	var content []string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}

		content = append(content, row[0])
	}

	return content, nil
}
