package tests

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo"
)

func GetMultipartFromFile(filePath string) (multipart.File, *multipart.FileHeader) {
	httpHeader, body := CreateRequestInfoFormFile(filePath)

	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		ginkgo.Fail("Failed to create request form file: " + err.Error())
	}
	req.Header = httpHeader

	file, header, err := req.FormFile("file")
	if err != nil {
		ginkgo.Fail("Failed to get multipart from request: " + err.Error())
	}

	return file, header
}

func CreateRequestInfoFormFile(filePath string) (http.Header, *bytes.Buffer) {
	file, err := os.Open(filePath)
	if err != nil {
		ginkgo.Fail("Failed to open file: " + err.Error())
	}
	defer file.Close()

	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)
	writer := createPart(multipartWriter, filepath.Base(filePath))

	_, err = io.Copy(writer, file)
	if err != nil {
		ginkgo.Fail("Failed to copy file to FormFile: " + err.Error())
	}

	err = multipartWriter.Close()
	if err != nil {
		ginkgo.Fail("Failed to close writer: " + err.Error())
	}

	headers := http.Header{}
	headers.Set("Content-Type", multipartWriter.FormDataContentType())

	return headers, body
}

func createPart(multipartWriter *multipart.Writer, fileName string) io.Writer {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", fileName))
	h.Set("Content-Type", mockFileType(fileName))

	writer, err := multipartWriter.CreatePart(h)
	if err != nil {
		ginkgo.Fail("Failed to create part:" + err.Error())
	}

	return writer
}

func mockFileType(fileName string) string {
	switch filepath.Ext(fileName) {
	case ".csv":
		return "text/csv"
	case ".txt":
		return "text/txt"
	}
	return ""
}
