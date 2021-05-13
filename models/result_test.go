package models_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/models/results"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Result", func() {
	Describe("CreateResult", func() {
		Context("given result with valid params", func() {
			It("returns the result ID", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := &models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				resultID, err := models.CreateResult(result)
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

				Expect(result.Status).To(Equal("pending"))
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := &models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				_, err := models.CreateResult(result)

				Expect(err).To(BeNil())
			})
		})

		Context("given result with INVALID params", func() {
			Context("given NO user and keyword", func() {
				It("returns an error", func() {
					result := &models.Result{}
					resultID, err := models.CreateResult(result)

					Expect(err.Error()).To(Equal("field `go-google-scraper-challenge/models.Result.User` cannot be NULL"))
					Expect(resultID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetResultById", func() {
		Context("given result id exist in the system", func() {
			It("returns result with given id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				existResult := FabricateResult(user)
				result, err := models.GetResultById(existResult.Id)
				if err != nil {
					Fail("Failed to get result with ID")
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

	Describe("#GetResultsByUserId", func() {
		Context("given a valid user id", func() {
			It("returns results with the given user id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				otherUser := FabricateUser("dev2@nimblehq.co", "password")
				userResult1 := FabricateResult(user)
				userResult2 := FabricateResult(user)
				otherUserResult := FabricateResult(otherUser)

				results, err := models.GetResultsByUserId(user.Id)
				if err != nil {
					Fail("Failed to get results with User Id")
				}

				resultIds := []int64{}
				for _, r := range results {
					resultIds = append(resultIds, r.Id)
				}

				Expect(resultIds).NotTo(ContainElement(otherUserResult.Id))
				Expect(resultIds).To(ConsistOf(userResult1.Id, userResult2.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				otherUser := FabricateUser("dev2@nimblehq.co", "password")
				FabricateResult(user)
				FabricateResult(user)
				FabricateResult(otherUser)

				_, err := models.GetResultsByUserId(user.Id)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid user id", func() {
			It("returns an empty list", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				FabricateResult(user)
				FabricateResult(user)

				results, err := models.GetResultsByUserId(999)
				if err != nil {
					Fail("Failed to get results with User Id")
				}

				Expect(results).To(BeEmpty())
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				FabricateResult(user)
				FabricateResult(user)

				_, err := models.GetResultsByUserId(999)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#UpdateResultById", func() {
		Context("given result id exist in the system", func() {
			It("updates the result with given id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				existResult := FabricateResult(user)
				existResult.Status = results.Processing

				err := models.UpdateResultById(existResult)
				if err != nil {
					Fail("Failed to update result with ID")
				}

				result, err := models.GetResultById(existResult.Id)
				if err != nil {
					Fail("Failed to get result with ID")
				}

				Expect(result.Keyword).To(Equal(existResult.Keyword))
				Expect(result.Status).To(Equal(results.Processing))
			})
		})

		Context("given result id does NOT exist in the system", func() {
			It("returns error", func() {
				result := &models.Result{Base: models.Base{Id: 999}}
				err := models.UpdateResultById(result)

				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
		initializers.CleanupDatabase("results")
	})
})
