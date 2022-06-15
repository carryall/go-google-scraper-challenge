package models_test

import (
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

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
					User:    user,
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
					User:    user,
					Keyword: "valid keyword",
				}
				resultID, err := models.CreateResult(result)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				result, err = models.GetResultByID(resultID, []string{})
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(result.Status).To(Equal(models.ResultStatusPending))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := &models.Result{
					User:    user,
					Keyword: "valid keyword",
				}
				_, err := models.CreateResult(result)

				Expect(err).To(BeNil())
			})

			Context("given result with status", func() {
				It("sets status to given status", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result := &models.Result{
						User:    user,
						Keyword: "valid keyword",
						Status:  models.ResultStatusProcessing,
					}
					resultID, err := models.CreateResult(result)
					if err != nil {
						Fail("Failed to add result: " + err.Error())
					}

					result, err = models.GetResultByID(resultID, []string{})
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

					Expect(err.Error()).To(HavePrefix("ERROR: insert or update on table \"results\" violates foreign key constraint \"results_user_id_fkey\""))
					Expect(resultID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#CreateResults", func() {
		Context("given results with valid params", func() {
			It("returns the result IDs", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				results := []models.Result{
					{
						User:    user,
						Keyword: "valid keyword",
					}, {
						User:    user,
						Keyword: "valid keyword",
					},
				}

				resultIDs, err := models.CreateResults(&results)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(resultIDs).To(HaveLen(2))
				for i := range resultIDs {
					Expect(resultIDs[i]).To(BeNumerically(">", 0))
				}
			})

			It("sets default result status to pending", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				results := &[]models.Result{
					{
						User:    user,
						Keyword: "valid keyword",
					}, {
						User:    user,
						Keyword: "valid keyword",
					},
				}
				resultIDs, err := models.CreateResults(results)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				results, err = models.GetResultsByIDs(resultIDs)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				resultList := *results
				for i := range resultList {
					Expect(resultList[i].Status).To(Equal(models.ResultStatusPending))
				}
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				results := []models.Result{
					{
						User:    user,
						Keyword: "valid keyword",
					}, {
						User:    user,
						Keyword: "valid keyword",
					},
				}
				_, err := models.CreateResults(&results)

				Expect(err).To(BeNil())
			})

			Context("given results with status", func() {
				It("sets status to given status", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					results := &[]models.Result{
						{
							User:    user,
							Keyword: "valid keyword",
							Status:  models.ResultStatusProcessing,
						}, {
							User:    user,
							Keyword: "valid keyword",
							Status:  models.ResultStatusProcessing,
						},
					}
					resultIDs, err := models.CreateResults(results)
					if err != nil {
						Fail("Failed to add result: " + err.Error())
					}

					results, err = models.GetResultsByIDs(resultIDs)
					if err != nil {
						Fail("Failed to add result: " + err.Error())
					}

					resultList := *results
					for i := range resultList {
						Expect(resultList[i].Status).To(Equal(models.ResultStatusProcessing))
					}
				})
			})
		})

		Context("given one result with INVALID params", func() {
			Context("given NO user and keyword", func() {
				It("returns an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					results := &[]models.Result{
						{}, {
							User:    user,
							Keyword: "valid keyword",
							Status:  models.ResultStatusProcessing,
						},
					}
					resultIDs, err := models.CreateResults(results)

					Expect(err.Error()).To(HavePrefix("ERROR: insert or update on table \"results\" violates foreign key constraint \"results_user_id_fkey\""))
					Expect(resultIDs).To(Equal([]int64{}))
				})
			})
		})
	})

	Describe("#GetResultByID", func() {
		Context("given result id exist in the system", func() {
			It("returns result with given id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				existResult := FabricateResult(user)
				result, err := models.GetResultByID(existResult.ID, []string{})
				if err != nil {
					Fail("Failed to get result with ID")
				}

				Expect(result.Keyword).To(Equal(existResult.Keyword))
				Expect(result.UserID).To(Equal(user.ID))
			})
		})

		Context("given result id does NOT exist in the system", func() {
			It("returns the error", func() {
				result, err := models.GetResultByID(999, []string{})

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(result).To(BeNil())
			})
		})
	})

	Describe("#GetResultsByIDs", func() {
		Context("given result IDs exist in the system", func() {
			It("returns results with the given ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				result2 := FabricateResult(user)
				expectedResults := []*models.Result{result, result2}

				results, err := models.GetResultsByIDs([]int64{result.ID, result2.ID})

				Expect(err).To(BeNil())
				for i, result := range *results {
					Expect(result.Keyword).To(Equal(expectedResults[i].Keyword))
					Expect(result.UserID).To(Equal(user.ID))
				}
			})
		})

		Context("given at least one result ID that exist in the system", func() {
			It("returns the existing result", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				existResult := FabricateResult(user)

				results, err := models.GetResultsByIDs([]int64{existResult.ID, 999})

				Expect(err).To(BeNil())
				Expect(*results).To(HaveLen(1))
				for _, result := range *results {
					Expect(result.Keyword).To(Equal(existResult.Keyword))
					Expect(result.UserID).To(Equal(user.ID))
				}
			})
		})

		Context("given NO result ID exist in the system", func() {
			It("returns the error", func() {
				results, err := models.GetResultsByIDs([]int64{888, 999})

				Expect(*results).To(HaveLen(0))
				Expect(err.Error()).To(ContainSubstring("record not found"))
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

				Expect(result.ID).To(Equal(pendingResult.ID))
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

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(result.ID).To(Equal((int64(0))))
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

						results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 2)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIDs []int64
						for _, r := range results {
							resultIDs = append(resultIDs, r.ID)
						}

						Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
						Expect(resultIDs).NotTo(ConsistOf(result3.ID))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 2)
						Expect(err).To(BeNil())
					})
				})

				Context("given a 0 limit", func() {
					It("returns the results with no limit", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 0)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIDs []int64
						for _, r := range results {
							resultIDs = append(resultIDs, r.ID)
						}

						Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID, result3.ID))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 0)
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

						results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 1, 0)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIDs []int64
						for _, r := range results {
							resultIDs = append(resultIDs, r.ID)
						}

						Expect(resultIDs).To(ConsistOf(result2.ID, result3.ID))
						Expect(resultIDs).NotTo(ConsistOf(result1.ID))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 1, 0)
						Expect(err).To(BeNil())
					})
				})

				Context("given a 0 offset", func() {
					It("returns the results with no offset", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						result1 := FabricateResult(user)
						result2 := FabricateResult(user)
						result3 := FabricateResult(user)

						results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 0)
						if err != nil {
							Fail("Failed to get results with User Id")
						}

						var resultIDs []int64
						for _, r := range results {
							resultIDs = append(resultIDs, r.ID)
						}

						Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID, result3.ID))
					})

					It("returns NO error", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(user)

						_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, 0)
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

					query := map[string]interface{}{
						"user_id": user.ID,
					}
					results, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).NotTo(ContainElement(otherUserResult.ID))
					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)
					FabricateResult(otherUser)

					query := map[string]interface{}{
						"user_id": user.ID,
					}
					_, err := models.GetResultsBy(query, []string{}, "", 0, 0)
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

					query := map[string]interface{}{
						"user_id": user.ID,
					}
					results, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).NotTo(ContainElement(otherUserResult.ID))
					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					otherUser := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)
					FabricateResult(otherUser)

					query := map[string]interface{}{
						"user_id": user.ID,
					}
					_, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					Expect(err).To(BeNil())
				})
			})

			Context("given a valid keyword query", func() {
				It("returns the results that match the query", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResultWithParams(user, "search for Keyword 1", models.ResultStatusPending)
					result2 := FabricateResultWithParams(user, "keyword 2", models.ResultStatusPending)
					result3 := FabricateResultWithParams(user, "some other result", models.ResultStatusPending)

					query := map[string]interface{}{
						"keyword": "keyword",
					}
					results, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).NotTo(ContainElement(result3.ID))
					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResultWithParams(user, "search for Keyword 1", models.ResultStatusPending)
					FabricateResultWithParams(user, "keyword 2", models.ResultStatusPending)
					FabricateResultWithParams(user, "some other result", models.ResultStatusPending)

					query := map[string]interface{}{
						"keyword": "keyword",
					}
					_, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					Expect(err).To(BeNil())
				})
			})

			Context("given a valid order", func() {
				It("returns the results order by the given order", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)

					results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "-id", 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					var expectedResultIDs = []int64{result2.ID, result1.ID}
					Expect(resultIDs).To(Equal(expectedResultIDs))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "-id", 0, 0)
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

					results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", -1, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", -1, 0)
					Expect(err).To(BeNil())
				})
			})

			Context("given a negative offset", func() {
				It("returns the results with no offset", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)

					results, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, -1)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					_, err := models.GetResultsBy(map[string]interface{}{}, []string{}, "", 0, -1)
					Expect(err).To(BeNil())
				})
			})

			Context("given an invalid user id", func() {
				It("returns an empty list", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					query := map[string]interface{}{
						"user_id": 999,
					}
					results, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					if err != nil {
						Fail("Failed to get results with User Id")
					}

					Expect(results).To(BeEmpty())
				})

				It("returns NO error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					FabricateResult(user)
					FabricateResult(user)

					query := map[string]interface{}{
						"user_id": 999,
					}
					_, err := models.GetResultsBy(query, []string{}, "", 0, 0)
					Expect(err).To(BeNil())
				})
			})
		})
	})

	Describe("#GetUserResults", func() {
		Context("given valid params", func() {
			Context("given user ID with results", func() {
				It("returns a list of results that belongs to the user", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					anotherUser := FabricateUser(faker.Email(), faker.Password())
					result1 := FabricateResult(user)
					result2 := FabricateResult(user)
					result3 := FabricateResult(anotherUser)

					results, err := models.GetUserResults(user.ID, []string{})

					Expect(err).To(BeNil())

					var resultIDs []int64
					for _, r := range results {
						resultIDs = append(resultIDs, r.ID)
					}

					Expect(resultIDs).To(ConsistOf(result1.ID, result2.ID))
					Expect(resultIDs).NotTo(ConsistOf(result3.ID))
				})

				Context("given an array of preload relations", func() {
					It("returns an array of results with relations", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						anotherUser := FabricateUser(faker.Email(), faker.Password())
						FabricateResult(user)
						FabricateResult(user)
						FabricateResult(anotherUser)

						results, err := models.GetUserResults(user.ID, []string{"User"})

						Expect(err).To(BeNil())

						for _, r := range results {
							Expect(r.User).ToNot(BeNil())
							Expect(r.User.ID).To(Equal(user.ID))
						}
					})
				})
			})

			Context("given user ID without results", func() {
				It("returns a blank array of result", func() {
					user := FabricateUser(faker.Email(), faker.Password())

					results, err := models.GetUserResults(user.ID, []string{})

					Expect(err).To(BeNil())
					Expect(results).To(HaveLen(0))
				})
			})
		})

		Context("given invalid params", func() {
			Context("given an INVALID user ID", func() {
				It("returns a blank array of result", func() {
					results, err := models.GetUserResults(999, []string{})

					Expect(err).To(BeNil())
					Expect(results).To(HaveLen(0))
				})
			})
		})
	})

	Describe("#CountResultsBy", func() {
		Context("given a valid user id", func() {
			It("returns the correct number of user results", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)
				FabricateResult(otherUser)

				query := map[string]interface{}{
					"user_id": user.ID,
				}
				count, err := models.CountResultsBy(query, []string{}, "", 0, 0)
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

				query := map[string]interface{}{
					"user_id": user.ID,
				}
				_, err := models.CountResultsBy(query, []string{}, "", 0, 0)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid user id", func() {
			It("returns a 0", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)

				query := map[string]interface{}{
					"user_id": 999,
				}
				count, err := models.CountResultsBy(query, []string{}, "", 0, 0)
				if err != nil {
					Fail("Failed to count results with User Id")
				}

				Expect(count).To(BeZero())
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				FabricateResult(user)
				FabricateResult(user)

				query := map[string]interface{}{
					"user_id": 999,
				}
				_, err := models.CountResultsBy(query, []string{}, "", 0, 0)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#UpdateResult", func() {
		Context("given result id exist in the system", func() {
			It("updates the result with given id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				existResult := FabricateResult(user)
				existResult.Status = models.ResultStatusPending

				err := models.UpdateResult(existResult)
				if err != nil {
					Fail("Failed to update result with ID")
				}

				result, err := models.GetResultByID(existResult.ID, []string{})
				if err != nil {
					Fail("Failed to get result with ID")
				}

				Expect(result.Keyword).To(Equal(existResult.Keyword))
				Expect(result.Status).To(Equal(models.ResultStatusPending))
			})
		})

		Context("given result id does NOT exist in the system", func() {
			It("returns error", func() {
				result := &models.Result{Base: models.Base{ID: 999}}
				err := models.UpdateResult(result)

				Expect(err.Error()).To(ContainSubstring("record not found"))
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
		CleanupDatabase([]string{"users", "results"})
	})
})
