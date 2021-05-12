package models_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/test/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Result", func() {
	Describe("CreateResult", func() {
		Context("given result with valid params", func() {
			It("returns the user ID", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				resultID, err := models.CreateResult(&result)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(resultID).To(BeNumerically(">", 0))
			})

			It("sets default result status to pending", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := &models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				resultID, err := models.CreateResult(result)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				result, err = models.GetResultById(resultID)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(resultID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				_, err := models.CreateResult(&result)

				Expect(err).To(BeNil())
			})
		})

		Context("given result with INVALID params", func() {

		})
	})

	Describe("#GetResultById", func() {
		Context("given result id exist in the system", func() {
			It("returns result with given id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				existResult := FabricateResult(user)
				result, err := models.GetResultById(existResult.Id)
				if err != nil {
					Fail("Failed to get user with ID")
				}

				Expect(result.Keyword).To(Equal(existResult.Keyword))
				Expect(result.User.Id).To(Equal(user.Id))
			})
		})

		Context("given result id does NOT exist in the system", func() {
			It("returns false", func() {
				result, err := models.GetResultById(999)

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(result).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("results")
		initializers.CleanupDatabase("adwords")
	})
})
