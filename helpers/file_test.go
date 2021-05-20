package helpers_test

import (
	"mime/multipart"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
	. "go-google-scraper-challenge/tests/helpers"
)

var _ = Describe("File", func() {
	Describe("#GetFileType", func() {
		Context("given a CSV file header", func() {
			It("returns csv file type", func() {
				csvMIMEHeader := CreateMIMEHaader("valid.csv")
				fileHeader := multipart.FileHeader{
					Filename: "valid.csv",
					Header: csvMIMEHeader,
					Size: 0,
				}
				fileType := helpers.GetFileType(&fileHeader)

				Expect(fileType).To(Equal("text/csv"))
			})
		})

		Context("given a text file header", func() {
			It("returns text file type", func() {
				csvMIMEHeader := CreateMIMEHaader("text.txt")
				fileHeader := multipart.FileHeader{
					Filename: "text.txt",
					Header: csvMIMEHeader,
					Size: 0,
				}
				fileType := helpers.GetFileType(&fileHeader)

				Expect(fileType).To(Equal("text/txt"))
			})
		})
	})

	Describe("#GetFileContent", func() {
		Context("given a CSV file", func() {
			It("returns CSV file content", func() {
				file, _ := GetMultipartFromFile("tests/fixtures/files/valid.csv")

				content, err := helpers.GetFileContent(file)
				if err != nil {
					Fail("Failed to get file content: " + err.Error())
				}

				Expect(content).To(Equal([]string{"cloud computing service", "crypto currency"}))
			})

			It("does NOT return error", func() {
				file, _ := GetMultipartFromFile("tests/fixtures/files/valid.csv")

				_, err := helpers.GetFileContent(file)

				Expect(err).To(BeNil())
			})
		})

		Context("given a blank text file", func() {
			It("returns an empty array", func() {
				file, _ := GetMultipartFromFile("tests/fixtures/files/text.txt")

				content, err := helpers.GetFileContent(file)
				if err != nil {
					Fail("Failed to get file content: " + err.Error())
				}

				Expect(content).To(BeEmpty())
			})

			It("does NOT return error", func() {
				file, _ := GetMultipartFromFile("tests/fixtures/files/text.txt")

				_, err := helpers.GetFileContent(file)

				Expect(err).To(BeNil())
			})
		})
	})
})
