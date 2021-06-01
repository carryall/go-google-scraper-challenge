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

	Describe("#GetFirstPendingResult", func() {
		Context("given at least one pending result", func() {
			It("returns the oldest pending result", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResultWithParams(user, "keyword", models.ResultStatusFailed)
				FabricateResultWithParams(user, "keyword", models.ResultStatusCompleted)
				pendingResult := FabricateResultWithParams(user, "keyword", models.ResultStatusPending)
				FabricateResultWithParams(user, "keyword", models.ResultStatusPending)
				FabricateResultWithParams(user, "keyword", models.ResultStatusProcessing)

				result, err := models.GetFirstPendingResult()
				if err != nil {
					Fail("Failed to get first pending result")
				}

				Expect(result.Id).To(Equal(pendingResult.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResultWithParams(user, "keyword", models.ResultStatusPending)

				_, err := models.GetFirstPendingResult()

				Expect(err).To(BeNil())
			})
		})

		Context("given NO pending result", func() {
			It("returns an error", func() {
				result, err := models.GetFirstPendingResult()

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("#GetPaginatedResultsByUserId", func() {
		Context("given valid params", func() {
			Context("given a valid user id", func() {
				Context("given no limit and offset", func() {
					It("returns all user results", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						otherUser := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						otherUserResult := FabricateResult(otherUser)

						results, err := models.GetPaginatedResultsByUserId(user.Id, 0,0)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIds []int64
						for _, r := range results {
							resultIds = append(resultIds, r.Id)
						}

						Expect(resultIds).NotTo(ContainElement(otherUserResult.Id))
						Expect(resultIds).To(ConsistOf(result2.Id, result1.Id))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						otherUser := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(otherUser)

						_, err := models.GetPaginatedResultsByUserId(user.Id, 0, 0)
						Expect(err).To(BeNil())
					})
				})

				Context("given a limit but no offset", func() {
					Context("given a positive limit", func() {
						It("returns the latest user results with the given limit", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, 2,0)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result3.Id, result2.Id))
							Expect(resultIds).NotTo(ConsistOf(result1.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, 2,0)
							Expect(err).To(BeNil())
						})
					})

					Context("given a negative limit", func() {
						It("returns the latest user results with no limit", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, -1,0)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result3.Id, result2.Id, result1.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, -1,0)
							Expect(err).To(BeNil())
						})
					})
				})

				Context("given no limit but an offset", func() {
					Context("given a positive offset", func() {
						It("returns the latest user results with the given offset", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, 0,1)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result2.Id, result1.Id))
							Expect(resultIds).NotTo(ConsistOf(result3.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, 0,1)
							Expect(err).To(BeNil())
						})
					})

					Context("given a negative offset", func() {
						It("returns the latest user results with no offset", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, 0,-1)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result3.Id, result2.Id, result1.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, 0,-1)
							Expect(err).To(BeNil())
						})
					})
				})

				Context("given a limit and an offset", func() {
					Context("given a positive limit and offset", func() {
						It("returns the latest user results with the given limit and offset", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)
							result4 := FabricateResult(user)
							result5 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, 2,1)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result4.Id, result3.Id))
							Expect(resultIds).NotTo(ConsistOf(result1.Id, result2.Id, result5.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, 2,1)
							Expect(err).To(BeNil())
						})
					})

					Context("given a negative limit and offset", func() {
						It("returns the user results with no limit and no offset", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							result1 := FabricateResult(user)
							result2 := FabricateResult(user)
							result3 := FabricateResult(user)
							result4 := FabricateResult(user)

							results, err := models.GetPaginatedResultsByUserId(user.Id, -1,-1)
							if err != nil {
								Fail("Failed to get results with User Id")
							}

							var resultIds []int64
							for _, r := range results {
								resultIds = append(resultIds, r.Id)
							}

							Expect(resultIds).To(ConsistOf(result4.Id, result3.Id, result2.Id, result1.Id))
						})

						It("returns NO error", func() {
							user := FabricateUser(faker.Email(), faker.Password())
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)
							FabricateResult(user)

							_, err := models.GetPaginatedResultsByUserId(user.Id, -1,-1)
							Expect(err).To(BeNil())
						})
					})
				})
			})
		})

		Context("given invalid params", func() {
			Context("given an invalid user id", func() {
				It("returns an empty list", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					results, err := models.GetPaginatedResultsByUserId(999, 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					Expect(results).To(BeEmpty())
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					_, err := models.GetPaginatedResultsByUserId(999, 0, 0)
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

	Describe("#Process", func() {
		It("update result status to processing", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Process()
			if err != nil {
				Fail("Failed to process result")
			}

			Expect(result.Status).To(Equal(models.ResultStatusProcessing))
		})

		It("returns NO error", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Process()
			Expect(err).To(BeNil())
		})
	})

	Describe("#Complete", func() {
		It("update result status to completed", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Complete()
			if err != nil {
				Fail("Failed to compleete result")
			}

			Expect(result.Status).To(Equal(models.ResultStatusCompleted))
		})

		It("returns NO error", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Complete()
			Expect(err).To(BeNil())
		})
	})

	Describe("#Fail", func() {
		It("update result status to failed", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Fail()
			if err != nil {
				Fail("Failed to compleete result")
			}

			Expect(result.Status).To(Equal(models.ResultStatusFailed))
		})

		It("returns NO error", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			result := FabricateResult(user)

			err := result.Fail()
			Expect(err).To(BeNil())
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase([]string{"users", "results"})
	})
})
