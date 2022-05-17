package models_test

import (
	"go-google-scraper-challenge/lib/models"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Link", func() {
	Describe("#CreateLink", func() {
		Context("given link with valid params", func() {
			It("returns the link ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				link := &models.Link{
					Result: result,
				}
				linkID, err := models.CreateLink(link)
				if err != nil {
					Fail("Failed to add result: " + err.Error())
				}

				Expect(linkID).To(BeNumerically(">", 0))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				link := &models.Link{
					Result: result,
				}
				_, err := models.CreateLink(link)

				Expect(err).To(BeNil())
			})
		})

		Context("given link with INVALID params", func() {
			Context("given NO user and keyword", func() {
				It("returns an error", func() {
					link := &models.Link{}

					linkID, err := models.CreateLink(link)

					Expect(err.Error()).To(HavePrefix("ERROR: insert or update on table \"links\" violates foreign key constraint \"links_result_id_fkey\""))
					Expect(linkID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetLinkById", func() {
		Context("given link id exist in the system", func() {
			It("returns link with given id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				existLink := FabricateLink(result)
				link, err := models.GetLinkById(existLink.Id)
				if err != nil {
					Fail("Failed to get link with ID")
				}

				Expect(link.Link).To(Equal(existLink.Link))
				Expect(link.ResultId).To(Equal(result.Id))
			})
		})

		Context("given link id does NOT exist in the system", func() {
			It("returns false", func() {
				link, err := models.GetLinkById(999)

				Expect(err.Error()).To(ContainSubstring("record not found"))
				Expect(link).To(BeNil())
			})
		})
	})

	Describe("#GetLinksByResultId", func() {
		Context("given a valid result id", func() {
			It("returns links with the given result id", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				link1 := FabricateLink(result)
				link2 := FabricateLink(result)
				otherLink := FabricateLink(otherResult)

				links, err := models.GetLinksByResultId(result.Id)
				if err != nil {
					Fail("Failed to get links with Result Id")
				}

				var LinkIds []int64
				for _, a := range links {
					LinkIds = append(LinkIds, a.Id)
				}

				Expect(LinkIds).NotTo(ContainElement(otherLink.Id))
				Expect(LinkIds).To(ConsistOf(link1.Id, link2.Id))
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				otherUser := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				otherResult := FabricateResult(otherUser)
				FabricateLink(result)
				FabricateLink(result)
				FabricateLink(otherResult)

				_, err := models.GetLinksByResultId(result.Id)
				Expect(err).To(BeNil())
			})
		})

		Context("given an invalid result id", func() {
			It("returns an empty list", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				FabricateLink(result)
				FabricateLink(result)

				results, err := models.GetLinksByResultId(999)
				if err != nil {
					Fail("Failed to get results with User Id")
				}

				Expect(results).To(BeEmpty())
			})

			It("returns NO error", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				FabricateLink(result)
				FabricateLink(result)

				_, err := models.GetLinksByResultId(999)
				Expect(err).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users", "results", "links"})
	})
})
