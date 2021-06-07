package models_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Result", func() {
	Describe("#CreateResult", func() {
		Context("given result with valid params", func() {
			It("returns the result ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
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
				user := FabricateUser(faker.Email(), faker.Password())
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

				Expect(result.Status).To(Equal(models.ResultStatusPending))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := &models.Result{
					User: user,
					Keyword: "valid keyword",
				}
				_, err := models.CreateResult(result)

				Expect(err).To(BeNil())
			})

			Context("given result with status", func() {
				It("sets status to given status", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := &models.Result{
						User: user,
						Keyword: "valid keyword",
						Status: models.ResultStatusProcessing,
					}
					resultID, err := models.CreateResult(result)
					if err != nil {
						Fail("Failed to add result: " + err.Error())
					}

					result, err = models.GetResultById(resultID)
					if err != nil {
						Fail("Failed to add result: " + err.Error())
					}

					Expect(result.Status).To(Equal(models.ResultStatusProcessing))
				})
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
				user := FabricateUser(faker.Email(), faker.Password())
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

	Describe("#GetOldestPendingResult()", func() {
		Context("given at least one pending result", func() {
			It("returns the oldest pending result", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResultWithParams(user, "keyword", models.ResultStatusFailed)
				FabricateResultWithParams(user, "keyword", models.ResultStatusCompleted)
				pendingResult := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)
				FabricateResultWithParams(user, "keyword", models.ResultStatusPending)
				FabricateResultWithParams(user, "keyword", models.ResultStatusProcessing)

				result, err := models.GetOldestPendingResult()
				if err != nil {
					Fail("Failed to get first pending result")
				}

				Expect(result.Id).To(Equal(pendingResult.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				_, err := models.GetOldestPendingResult()

				Expect(err).To(BeNil())
			})
		})

		Context("given NO pending result", func() {
			It("returns an error", func() {
				result, err := models.GetOldestPendingResult()

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("#GetResultsBy", func() {
		Context("given valid params", func() {
			Context("given a valid limit", func() {
				Context("given limit is > 0", func() {
					It("returns the results with the given limit", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						var query = map[string]interface{}{
							"limit": 2,
						}
						results, err := models.GetResultsBy(query)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIds []int64
						for _, r := range results {
							resultIds = append(resultIds, r.Id)
						}

						Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
						Expect(resultIds).NotTo(ConsistOf(result3.Id))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						var query = map[string]interface{}{
							"limit": 2,
						}
						_, err := models.GetResultsBy(query)
						Expect(err).To(BeNil())
					})
				})

				Context("given a 0 limit", func() {
					It("returns the results with no limit", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						var query = map[string]interface{}{
							"limit": 0,
						}
						results, err := models.GetResultsBy(query)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIds []int64
						for _, r := range results {
							resultIds = append(resultIds, r.Id)
						}

						Expect(resultIds).To(ConsistOf(result1.Id, result2.Id, result3.Id))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						var query = map[string]interface{}{
							"limit": 0,
						}
						_, err := models.GetResultsBy(query)
						Expect(err).To(BeNil())
					})
				})
			})

			Context("given a valid offset", func() {
				Context("given offset is > 0", func() {
					It("returns the results with the given offset", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						var query = map[string]interface{}{
							"offset": 1,
						}
						results, err := models.GetResultsBy(query)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIds []int64
						for _, r := range results {
							resultIds = append(resultIds, r.Id)
						}

						Expect(resultIds).To(ConsistOf(result2.Id, result3.Id))
						Expect(resultIds).NotTo(ConsistOf(result1.Id))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						var query = map[string]interface{}{
							"offset": 0,
						}
						_, err := models.GetResultsBy(query)
						Expect(err).To(BeNil())
					})
				})

				Context("given a 0 offset", func() {
					It("returns the results with no offset", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						var query = map[string]interface{}{
							"offset": 0,
						}
						results, err := models.GetResultsBy(query)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIds []int64
						for _, r := range results {
							resultIds = append(resultIds, r.Id)
						}

						Expect(resultIds).To(ConsistOf(result1.Id, result2.Id, result3.Id))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						var query = map[string]interface{}{
							"offset": 0,
						}
						_, err := models.GetResultsBy(query)
						Expect(err).To(BeNil())
					})
				})
			})

			Context("given a valid user id", func() {
				It("returns all user results", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)
					otherUserResult := FabricateResult(otherUser)

					var query = map[string]interface{}{
						"user_id": user.Id,
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					Expect(resultIds).NotTo(ContainElement(otherUserResult.Id))
					Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)
					FabricateResult(otherUser)

					var query = map[string]interface{}{
						"user_id": user.Id,
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})

			Context("given a valid user id", func() {
				It("returns all user results", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)
					otherUserResult := FabricateResult(otherUser)

					var query = map[string]interface{}{
						"user_id": user.Id,
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					Expect(resultIds).NotTo(ContainElement(otherUserResult.Id))
					Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)
					FabricateResult(otherUser)

					var query = map[string]interface{}{
						"user_id": user.Id,
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})

			Context("given a valid keyword query", func() {
				It("returns the results that match the query", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResultWithParams(user, "search for Keyword 1", models.ResultStatusPending)
					result2 := FabricateResultWithParams(user, "keyword 2", models.ResultStatusPending)
					result3 := FabricateResultWithParams(user, "some other result", models.ResultStatusPending)

					var query = map[string]interface{}{
						"keyword__icontains": "keyword",
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					Expect(resultIds).NotTo(ContainElement(result3.Id))
					Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResultWithParams(user, "search for Keyword 1", models.ResultStatusPending)
					FabricateResultWithParams(user, "keyword 2", models.ResultStatusPending)
					FabricateResultWithParams(user, "some other result", models.ResultStatusPending)

					var query = map[string]interface{}{
						"keyword__icontains": "keyword",
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})

			Context("given a valid order", func() {
				It("returns the results order by the given order", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)

					var query = map[string]interface{}{
						"order": "-id",
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					var expectedResultIds = []int64{ result2.Id, result1.Id }
					Expect(resultIds).To(Equal(expectedResultIds))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					var query = map[string]interface{}{
						"order": "-id",
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})
		})

		Context("given invalid params", func() {
			Context("given a negative limit", func() {
				It("returns the results with no limit", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)

					var query = map[string]interface{}{
						"limit": -1,
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					var query = map[string]interface{}{
						"limit": -1,
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})

			Context("given a negative offset", func() {
				It("returns the results with no offset", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)

					var query = map[string]interface{}{
						"offset": -1,
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIds []int64
					for _, r := range results {
						resultIds = append(resultIds, r.Id)
					}

					Expect(resultIds).To(ConsistOf(result1.Id, result2.Id))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					var query = map[string]interface{}{
						"offset": -1,
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})

			Context("given an invalid user id", func() {
				It("returns an empty list", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					var query = map[string]interface{}{
						"user_id": 999,
					}
					results, err := models.GetResultsBy(query)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					Expect(results).To(BeEmpty())
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					var query = map[string]interface{}{
						"user_id": 999,
					}
					_, err := models.GetResultsBy(query)
					Expect(err).To(BeNil())
				})
			})
		})
	})

	Describe("#CountResultsByUserId", func() {
		Context("given a valid user id", func() {
			It("returns the correct number of user results", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)
				FabricateResult(otherUser)

				count, err := models.CountResultsByUserId(user.Id)
				if err != nil {
					Fail("Failed to count results with User Id")
				}

				Expect(count).To(BeEquivalentTo(2))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)
				FabricateResult(otherUser)

				_, err := models.CountResultsByUserId(user.Id)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid user id", func() {
			It("returns a 0", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)

				count, err := models.CountResultsByUserId(999)
				if err != nil {
					Fail("Failed to count results with User Id")
				}

				Expect(count).To(BeZero())
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)

				_, err := models.CountResultsByUserId(999)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#UpdateResultById", func() {
		Context("given result id exist in the system", func() {
			It("updates the result with given id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				existResult := FabricateResult(user)
				existResult.Status = models.ResultStatusPending

				err := models.UpdateResultById(existResult)
				if err != nil {
					Fail("Failed to update result with ID")
				}

				result, err := models.GetResultById(existResult.Id)
				if err != nil {
					Fail("Failed to get result with ID")
				}

				Expect(result.Keyword).To(Equal(existResult.Keyword))
				Expect(result.Status).To(Equal(models.ResultStatusPending))
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

	Describe("#UpdateResultStatus", func() {
		Context("given a valid status", func() {
			It("updates result status to the given status", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				err := models.UpdateResultStatus(result, models.ResultStatusCompleted)
				if err != nil {
					Fail("Failed to update result status")
				}

				Expect(result.Status).To(Equal(models.ResultStatusCompleted))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				err := models.UpdateResultStatus(result, models.ResultStatusCompleted)
				Expect(err).To(BeNil())
			})
		})

		Context("given an INVALID status", func() {
			It("does NOT update the result status", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				err := models.UpdateResultStatus(result, "invalid status")
				Expect(err).NotTo(BeNil())
				Expect(result.Status).To(Equal(models.ResultStatusPending))
			})

			It("returns the error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				err := models.UpdateResultStatus(result, "invalid status")
				Expect(err.Error()).To(Equal("Invalid result status"))
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users", "results"})
	})
})
