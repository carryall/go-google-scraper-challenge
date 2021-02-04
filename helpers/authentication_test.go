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

	Describe("#CompareHashWithPassword", func() {
		Context("given a valid hashed password and password", func() {
			It("returns true", func() {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
				if err != nil {
					Fail("Failed to hash password")
				}
				result := helpers.CompareHashWithPassword(string(hashedPassword), "password")

				Expect(result).To(BeTrue())
			})
		})

		Context("given an INVALID hashed password and password", func() {
			It("returns false", func() {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte("not the password"), bcrypt.MinCost)
				if err != nil {
					Fail("Failed to hash password")
				}
				result := helpers.CompareHashWithPassword(string(hashedPassword), "password")

				Expect(result).To(BeFalse())
			})
		})
	})
})
