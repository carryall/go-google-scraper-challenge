package models_test

import (
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("#CreateUser", func() {
		Context("given user with valid params", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:          faker.Email(),
					HashedPassword: faker.Password(),
				}
				userID, err := models.CreateUser(&user)
				if err != nil {
					Fail("Failed to add user: " + err.Error())
				}

				Expect(userID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := models.User{
					Email:          faker.Email(),
					HashedPassword: faker.Password(),
				}
				_, err := models.CreateUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given user with INVALID params", func() {
			Context("given email that already exist in database", func() {
				It("returns an error", func() {
					password := faker.Password()
					existingUser := FabricateUser(faker.Email(), password)

					user := models.User{
						Email:          existingUser.Email,
						HashedPassword: password,
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(HavePrefix(`ERROR: duplicate key value violates unique constraint "users_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetUserByID", func() {
		Context("given user id exist in the system", func() {
			It("returns user with given id", func() {
				existUser := FabricateUser(faker.Email(), faker.Password())

				user, err := models.GetUserByID(existUser.ID)
				if err != nil {
					Fail("Failed to get user with ID")
				}

				Expect(user.Email).To(Equal(existUser.Email))
				Expect(user.HashedPassword).To(Equal(existUser.HashedPassword))
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns false", func() {
				user, err := models.GetUserByID(999)

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(user).To(BeNil())
			})
		})
	})

	Describe("#UserEmailAlreadyExist", func() {
		Context("given user email exist in the system", func() {
			It("returns true", func() {
				user := FabricateUser(faker.Email(), faker.Password())

				userExist := models.UserEmailAlreadyExist(user.Email)

				Expect(userExist).To(BeTrue())
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns false", func() {
				userExist := models.UserEmailAlreadyExist(faker.Email())

				Expect(userExist).To(BeFalse())
			})
		})
	})

	Describe("#GetUserByEmail", func() {
		Context("given user email exist in the system", func() {
			It("returns the user", func() {
				existUser := FabricateUser(faker.Email(), faker.Password())

				user, err := models.GetUserByEmail(existUser.Email)
				if err != nil {
					Fail("Failed to find user with given email")
				}

				Expect(user.ID).To(Equal(existUser.ID))
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns error", func() {
				user, err := models.GetUserByEmail(faker.Email())

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(user.ID).To(Equal(int64(0)))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
