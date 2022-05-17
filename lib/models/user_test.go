package models_test

import (
	"go-google-scraper-challenge/lib/models"
	"go-google-scraper-challenge/test"

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
					existingUser := test.FabricateUser(faker.Email(), password)

					user := models.User{
						Email:          existingUser.Email,
						HashedPassword: password,
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(Equal(`ERROR: pq: duplicate key value violates unique constraint "user_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetUserById", func() {
		Context("given user id exist in the system", func() {
			It("returns user with given id", func() {
				existUser := test.FabricateUser(faker.Email(), faker.Password())

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

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(user).To(BeNil())
			})
		})
	})

	Describe("#UserEmailAlreadyExist", func() {
		Context("given user email exist in the system", func() {
			It("returns true", func() {
				user := test.FabricateUser(faker.Email(), faker.Password())

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
				existUser := test.FabricateUser(faker.Email(), faker.Password())

				user, err := models.GetUserByEmail(existUser.Email)
				if err != nil {
					Fail("Failed to find user with given email")
				}

				Expect(user.Id).To(Equal(existUser.Id))
			})
		})

		Context("given user email does NOT exist in the system", func() {
			It("returns error", func() {
				user, err := models.GetUserByEmail(faker.Email())

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(user.Id).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		test.CleanupDatabase([]string{"users"})
	})
})
