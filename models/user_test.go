package models_test

import (
	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("#CreateUser", func() {
		Context("given user with valid params", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				}
				userID, err := models.CreateUser(&user)
				if err != nil {
					Fail("Failed to add user: " + err.Error())
				}

				Expect(userID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				}
				_, err := models.CreateUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given user with INVALID params", func() {
			Context("given email that already exist in database", func() {
				It("returns an error", func() {
					FabricateUser(&models.User{
						Email:             "dev@nimblehq.co",
						EncryptedPassword: "password",
					})

					user := models.User{
						Email:             "dev@nimblehq.co",
						EncryptedPassword: "password",
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "user_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#UserEmailAlreadyExist", func() {
		Context("given user email exist in the system", func() {
			It("returns true", func() {
				FabricateUser(&models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				})

				userExist := models.UserEmailAlreadyExist("dev@nimblehq.co")

				Expect(userExist).To(BeTrue())
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns false", func() {
				userExist := models.UserEmailAlreadyExist("dev@nimblehq.co")

				Expect(userExist).To(BeFalse())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
