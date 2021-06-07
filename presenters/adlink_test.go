package presenters_test

import (
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/presenters"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AdLink", func() {
	Describe("#GetAdLinkCollection", func() {
		Context("given a list of ad links", func() {
			It("returns link collection map with position", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink1 := FabricateAdLinkWithParams(result, models.AdLinkPositionTop)
				adLink2 := FabricateAdLinkWithParams(result, models.AdLinkPositionBottom)
				adLink3 := FabricateAdLinkWithParams(result, models.AdLinkPositionSide)
				expectedAdLinkCollection := map[string][]string{
					models.AdLinkPositionTop: {adLink1.Link},
					models.AdLinkPositionBottom: {adLink2.Link},
					models.AdLinkPositionSide: {adLink3.Link},
				}

				adLinkCollection := presenters.GetAdLinkCollection([]*models.AdLink{adLink1, adLink2, adLink3})

				Expect(adLinkCollection).To(Equal(expectedAdLinkCollection))
			})
		})

		Context("given a blank list", func() {
			It("returns blank collections map with ad link position", func() {
				expectedAdLinkCollection := map[string][]string{
					models.AdLinkPositionTop: nil,
					models.AdLinkPositionBottom: nil,
					models.AdLinkPositionSide: nil,
				}

				adLinkCollection := presenters.GetAdLinkCollection([]*models.AdLink{})

				Expect(adLinkCollection).To(Equal(expectedAdLinkCollection))
			})
		})
	})
})
