package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Authentication", func() {
	Describe("#HashPassword", func() {
		Context("given a valid password", func() {
			It("returns the hashed password", func() {
				hashedPassword, err := helpers.HashPassword("password")
				if err != nil {
					Fail("Failed to hash password")
				}

				err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte("password"))

				Expect(err).To(BeNil())
			})
		})
	})
})
