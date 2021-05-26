package forms_test

import (
	"go-google-scraper-challenge/constants"
	"go-google-scraper-challenge/forms"
	"go-google-scraper-challenge/initializers"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/beego/beego/v2/core/validation"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forms/RegistrationForm", func() {
	Describe("#Valid", func() {
		Context("given registration form with valid params", func() {
			It("does NOT add error to validation", func() {
				password := faker.Password()
				form := forms.RegistrationForm{
					Email:                faker.Email(),
					Password:             password,
					PasswordConfirmation: password,
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("given email that already registered", func() {
				It("adds duplicate email error to validation", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)

					form := forms.RegistrationForm{
						Email:                user.Email,
						Password:             password,
						PasswordConfirmation: password,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.UserAlreadyExist))
				})
			})

			Context("given password confirmation is NOT match with the password", func() {
				It("adds a mismatch password confirmation error to validation", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             faker.Password(),
						PasswordConfirmation: "does not match the password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("PasswordConfirmation"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.PasswordConfirmNotMatch))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given registration form with valid params", func() {
			It("returns a user ID", func() {
				password := faker.Password()
				form := forms.RegistrationForm{
					Email:                faker.Email(),
					Password:             password,
					PasswordConfirmation: password,
				}

				userID, err := form.Save()
				if err != nil {
					Fail("Failed to save form")
				}

				Expect(userID).NotTo(BeNil())
			})

			It("returns NO error", func() {
				password := faker.Password()
				form := forms.RegistrationForm{
					Email:                faker.Email(),
					Password:             password,
					PasswordConfirmation: password,
				}

				_, err := form.Save()

				Expect(err).To(BeNil())
			})
		})

		Context("given registration form with INVALID params", func() {
			Context("given email that already registered", func() {
				It("returns a duplicate email error", func() {
					user := FabricateUser(faker.Email(), faker.Password())

					password := faker.Password()
					form := forms.RegistrationForm{
						Email:                user.Email,
						Password:             password,
						PasswordConfirmation: password,
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal(constants.UserAlreadyExist))
				})
			})

			Context("given NO email", func() {
				It("returns an invalid email error", func() {
					password := faker.Password()
					form := forms.RegistrationForm{
						Email:                "",
						Password:             password,
						PasswordConfirmation: password,
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given an INVALID email", func() {
				It("returns an invalid email error", func() {
					password := faker.Password()
					form := forms.RegistrationForm{
						Email:                "invalid",
						Password:             password,
						PasswordConfirmation: password,
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given NO password", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             "",
						PasswordConfirmation: faker.Password(),
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given password length is less than 6", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             "1234",
						PasswordConfirmation: faker.Password(),
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given NO password confirmation", func() {
				It("returns an invalid password confirmation error", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             faker.Password(),
						PasswordConfirmation: "",
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("PasswordConfirmation can not be empty"))
				})
			})

			Context("given password confirmation is length less than 6", func() {
				It("returns an invalid password error", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             faker.Password(),
						PasswordConfirmation: "1234",
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal("PasswordConfirmation minimum size is 6"))
				})
			})

			Context("given password confirmation is NOT match with the password", func() {
				It("returns a mismatch password confirmation error", func() {
					form := forms.RegistrationForm{
						Email:                faker.Email(),
						Password:             faker.Password(),
						PasswordConfirmation: "does not match the password",
					}

					userID, err := form.Save()

					Expect(userID).To(BeNil())
					Expect(err.Error()).To(Equal(constants.PasswordConfirmNotMatch))
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users"})
	})
})
