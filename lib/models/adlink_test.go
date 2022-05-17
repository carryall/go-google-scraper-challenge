package models_test

import (
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AdLink", func() {
	Describe("#CreateAdLink", func() {
		Context("given ad link with valid params", func() {
			It("returns the ad link ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink := &models.AdLink{
					Result:   result,
					Position: models.AdLinkPositionTop,
					Type:     models.AdLinkTypeLink,
				}
				adLinkID, err := models.CreateAdLink(adLink)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(adLinkID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink := &models.AdLink{
					Result:   result,
					Position: models.AdLinkPositionTop,
					Type:     models.AdLinkTypeLink,
				}
				_, err := models.CreateAdLink(adLink)

				Expect(err).To(BeNil())
			})
		})

		Context("given adLink with INVALID params", func() {
			Context("given NO user and keyword", func() {
				It("returns an error", func() {
					adLink := &models.AdLink{}

					adLinkID, err := models.CreateAdLink(adLink)

					Expect(err.Error()).To(HavePrefix("ERROR: insert or update on table \"ad_links\" violates foreign key constraint \"ad_links_result_id_fkey\""))
					Expect(adLinkID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetAdLinkById", func() {
		Context("given adLink id exist in the system", func() {
			It("returns adLink with given id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				existAdLink := FabricateAdLink(result)
				adLink, err := models.GetAdLinkById(existAdLink.Id)
				if err != nil {
					Fail("Failed to get adLink with ID")
				}

				Expect(adLink.Link).To(Equal(existAdLink.Link))
				Expect(adLink.ResultId).To(Equal(result.Id))
			})
		})

		Context("given adLink id does NOT exist in the system", func() {
			It("returns false", func() {
				adLink, err := models.GetAdLinkById(999)

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(adLink).To(BeNil())
			})
		})
	})

	Describe("#GetAdLinksByResultId", func() {
		Context("given a valid result id", func() {
			It("returns adlinks with the given result id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				adLink1 := FabricateAdLink(result)
				adLink2 := FabricateAdLink(result)
				otherAdLink := FabricateAdLink(otherResult)

				adLinks, err := models.GetAdLinksByResultId(result.Id)
				if err != nil {
					Fail("Failed to get adlinks with Result Id")
				}

				var AdLinkIds []int64
				for _, a := range adLinks {
					AdLinkIds = append(AdLinkIds, a.Id)
				}

				Expect(AdLinkIds).NotTo(ContainElement(otherAdLink.Id))
				Expect(AdLinkIds).To(ConsistOf(adLink1.Id, adLink2.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				FabricateAdLink(result)
				FabricateAdLink(result)
				FabricateAdLink(otherResult)

				_, err := models.GetAdLinksByResultId(result.Id)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid result id", func() {
			It("returns an empty list", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				FabricateAdLink(result)
				FabricateAdLink(result)

				results, err := models.GetAdLinksByResultId(999)
				if err != nil {
					Fail("Failed to get results with User Id")
				}

				Expect(results).To(BeEmpty())
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				FabricateAdLink(result)
				FabricateAdLink(result)

				_, err := models.GetAdLinksByResultId(999)
				Expect(err).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "ad_links"})
	})
})
