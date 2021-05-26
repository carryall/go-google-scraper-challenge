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

var _ = Describe("Forms/SessionForm", func() {
	Describe("#Valid", func() {
		Context("given session form with valid params", func() {
			It("does NOT add error to validation", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := forms.SessionForm{
					Email:    user.Email,
					Password: password,
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given session form with INVALID params", func() {
			Context("given user email does NOT exist", func() {
				It("adds an error to validation", func() {
					form := forms.SessionForm{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.SignInFail))
				})
			})

			Context("given email is INVALID", func() {
				It("adds an error to validation", func() {
					FabricateUser(faker.Email(), faker.Password())
					form := forms.SessionForm{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.SignInFail))
				})
			})

			Context("given password is INVALID", func() {
				It("adds an error to validation", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					form := forms.SessionForm{
						Email:    user.Email,
						Password: faker.Password(),
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Password"))
					Expect(formValidation.Errors[0].Message).To(Equal(constants.SignInFail))
				})
			})
		})
	})

	Describe("#Save", func() {
		Context("given session form with valid params", func() {
			It("returns user with NO error", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := forms.SessionForm{
					Email:    user.Email,
					Password: password,
				}

				currentUser, errors := form.Save()

				Expect(len(errors)).To(BeZero())
				Expect(currentUser.Id).To(Equal(user.Id))
			})
		})

		Context("given session form with INVALID params", func() {
			Context("given email is NOT valid", func() {
				It("returns an email is invalid error", func() {
					form := forms.SessionForm{
						Email:    "not an email",
						Password: faker.Password(),
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
					Expect(user).To(BeNil())
				})
			})

			Context("given NO email", func() {
				It("returns an invalid email or password error", func() {
					form := forms.SessionForm{
						Email:    "",
						Password: faker.Password(),
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
					Expect(user).To(BeNil())
				})
			})

			Context("given NO password", func() {
				It("returns an invalid email or password error", func() {
					form := forms.SessionForm{
						Email:    faker.Email(),
						Password: "",
					}

					user, errors := form.Save()

					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
					Expect(user).To(BeNil())
				})
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users"})
	})
})
