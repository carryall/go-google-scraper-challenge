package models_test

import (
	. "go-google-scraper-challenge/helpers/test"
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("#AddUser", func() {
		Context("given user with valid params", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				}
				userID, err := models.AddUser(&user)
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
				_, err := models.AddUser(&user)

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
					userID, err := models.AddUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "user_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetUserByEmail", func() {
		Context("given user email exist in system", func() {
			It("returns the existing user ID", func() {
				existingUserID := FabricateUser(&models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				})

				user, err := models.GetUserByEmail("dev@nimblehq.co")
				if err != nil {
					Fail("Failed to get user by email: " + err.Error())
				}

				Expect(user.Id).To(Equal(existingUserID))
			})

			It("does NOT return errors", func() {
				FabricateUser(&models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "password",
				})

				_, err := models.GetUserByEmail("dev@nimblehq.co")

				Expect(err).To(BeNil())
			})
		})

		Context("given user email does NOT exist in system", func() {
			It("returns error", func() {
				userID, err := models.GetUserByEmail("dev@nimblehq.co")

				Expect(userID).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("user")
	})
})
