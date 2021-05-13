package models_test

import (
	"go-google-scraper-challenge/initializers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/models/adwords"
	. "go-google-scraper-challenge/tests/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adword", func() {
	Describe("CreateAdword", func() {
		Context("given adword with valid params", func() {
			It("returns the adword ID", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResult(user)
				adword := &models.Adword{
					Result: result,
					Position: adwords.Top,
					Type: adwords.Link,
				}
				adwordID, err := models.CreateAdword(adword)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(adwordID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResult(user)
				adword := &models.Adword{
					Result: result,
					Position: adwords.Top,
					Type: adwords.Link,
				}
				_, err := models.CreateAdword(adword)

				Expect(err).To(BeNil())
			})
		})

		Context("given adword with INVALID params", func() {
			Context("given NO user and keyword", func() {
				It("returns an error", func() {
					adword := &models.Adword{}

					adwordID, err := models.CreateAdword(adword)

					Expect(err.Error()).To(Equal("field `go-google-scraper-challenge/models.Adword.Result` cannot be NULL"))
					Expect(adwordID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetAdwordById", func() {
		Context("given adword id exist in the system", func() {
			It("returns adword with given id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResult(user)
				existAdword := FabricateAdword(result)
				adword, err := models.GetAdwordById(existAdword.Id)
				if err != nil {
					Fail("Failed to get adword with ID")
				}

				Expect(adword.Link).To(Equal(existAdword.Link))
				Expect(adword.Result.Id).To(Equal(result.Id))
			})
		})

		Context("given adword id does NOT exist in the system", func() {
			It("returns false", func() {
				adword, err := models.GetAdwordById(999)

				Expect(err.Error()).To(ContainSubstring("no row found"))
				Expect(adword).To(BeNil())
			})
		})
	})

	Describe("#GetAdwordsByResultId", func() {
		Context("given a valid result id", func() {
			It("returns adwords with the given result id", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				otherUser := FabricateUser("dev2@nimblehq.co", "password")
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				adword1 := FabricateAdword(result)
				adword2 := FabricateAdword(result)
				otherAdword := FabricateAdword(otherResult)

				adwords, err := models.GetAdwordsByResultId(result.Id)
				if err != nil {
					Fail("Failed to get adwords with Result Id")
				}

				AdwordIds := []int64{}
				for _, a := range adwords {
					AdwordIds = append(AdwordIds, a.Id)
				}

				Expect(AdwordIds).NotTo(ContainElement(otherAdword.Id))
				Expect(AdwordIds).To(ConsistOf(adword1.Id, adword2.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				otherUser := FabricateUser("dev2@nimblehq.co", "password")
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				FabricateAdword(result)
				FabricateAdword(result)
				FabricateAdword(otherResult)

				_, err := models.GetAdwordsByResultId(result.Id)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid result id", func() {
			It("returns an empty list", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResult(user)
				FabricateAdword(result)
				FabricateAdword(result)

				results, err := models.GetAdwordsByResultId(999)
				if err != nil {
					Fail("Failed to get results with User Id")
				}

				Expect(results).To(BeEmpty())
			})

			It("returns NO error", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				result := FabricateResult(user)
				FabricateAdword(result)
				FabricateAdword(result)

				_, err := models.GetAdwordsByResultId(999)
				Expect(err).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		initializers.CleanupDatabase("users")
		initializers.CleanupDatabase("results")
		initializers.CleanupDatabase("adwords")
	})
})
