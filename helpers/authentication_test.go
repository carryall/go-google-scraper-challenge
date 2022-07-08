package helpers_test

import (
	"go-google-scraper-challenge/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("Authentication", func() {
	Describe("#HashPassword", func() {
		Context("given a valid password", func() {
			It("returns the hashed password", func() {
				password := faker.Password()
				hashedPassword, err := helpers.HashPassword(password)
				if err != nil {
					Fail("Failed to hash password")
				}

				err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CompareHashWithPassword", func() {
		Context("given a valid hashed password and password", func() {
			It("returns true", func() {
				password := faker.Password()
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
				if err != nil {
					Fail("Failed to hash password")
				}
				result := helpers.CompareHashWithPassword(string(hashedPassword), password)

				Expect(result).To(BeTrue())
			})
		})

		Context("given an INVALID hashed password and password", func() {
			It("returns false", func() {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(faker.Password()), bcrypt.MinCost)
				if err != nil {
					Fail("Failed to hash password")
				}
				result := helpers.CompareHashWithPassword(string(hashedPassword), faker.Password())

				Expect(result).To(BeFalse())
			})
		})
	})
})
