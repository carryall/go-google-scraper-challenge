package models_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("#CreateUser", func() {
		Context("given user with valid params", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:          "dev@nimblehq.co",
					HashedPassword: "password",
				}
				userID, err := models.CreateUser(&user)
				if err != nil {
					Fail("Failed to add user: " + err.Error())
				}

				Expect(userID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := models.User{
					Email:          "dev@nimblehq.co",
					HashedPassword: "password",
				}
				_, err := models.CreateUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given user with INVALID params", func() {
			Context("given email that already exist in database", func() {
				It("returns an error", func() {
					FabricateUser("dev@nimblehq.co", "password")

					user := models.User{
						Email:          "dev@nimblehq.co",
						HashedPassword: "password",
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "users_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetUserById", func() {
		Context("given user id exist in the system", func() {
			It("returns user with given id", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				user, err := models.GetUserById(existUser.Id)
				if err != nil {
					Fail("Failed to get user with ID")
				}

				Expect(user.Email).To(Equal(existUser.Email))
				Expect(user.HashedPassword).To(Equal(existUser.HashedPassword))
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns false", func() {
				user, err := models.GetUserById(999)

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(user).To(BeNil())
			})
		})
	})

	Describe("#UserEmailAlreadyExist", func() {
		Context("given user email exist in the system", func() {
			It("returns true", func() {
				FabricateUser("dev@nimblehq.co", "password")

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

	Describe("#GetUserByEmail", func() {
		Context("given user email exist in the system", func() {
			It("returns the user", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				user, err := models.GetUserByEmail("dev@nimblehq.co")
				if err != nil {
					Fail("Failed to find user with given email")
				}

				Expect(user.Id).To(Equal(existUser.Id))
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns error", func() {
				user, err := models.GetUserByEmail("dev@nimblehq.co")

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(user).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
	})
})
