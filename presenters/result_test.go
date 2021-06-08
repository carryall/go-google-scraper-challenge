package presenters_test

import (
	"go-google-scraper-challenge/helpers"
	"go-google-scraper-challenge/models"
	"go-google-scraper-challenge/presenters"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Result", func() {
	Describe("#PrepareResultSet", func() {
		Context("given a blank list", func() {
			It("returns nil", func() {
				var results []*models.Result

				Expect(presenters.PrepareResultSet(results)).To(BeNil())
			})
		})

		Context("given a result list with half of configured `PaginationPerPage` length", func() {
			It("returns a set of all results", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				var results []*models.Result
				perPage := helpers.GetPaginationPerPage()
				for i := 0; i < perPage/2; i++ {
					results = append(results, FabricateResult(user))
				}

				resultSet := presenters.PrepareResultSet(results)

				Expect(resultSet).To(HaveLen(1))
				Expect(resultSet[0]).To(HaveLen(perPage/2))
			})
		})

		Context("given a result list with more than half of configured `PaginationPerPage` length", func() {
			It("returns two sets of result divided by half of configured `PaginationPerPage`", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				var results []*models.Result
				perPage := helpers.GetPaginationPerPage()
				for i := 0; i < (perPage / 2) + 5; i++ {
					results = append(results, FabricateResult(user))
				}

				resultSet := presenters.PrepareResultSet(results)

				Expect(resultSet).To(HaveLen(2))
				Expect(resultSet[0]).To(HaveLen(perPage / 2))
				Expect(resultSet[1]).To(HaveLen(5))
			})
		})
	})

	Describe("#GetResultPresenter", func() {
		Context("given a valid result", func() {
			It("returns a result presenter", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				result := FabricateResult(user)
				adLink1 := FabricateAdLinkWithParams(result, models.AdLinkPositionTop)
				adLink2 := FabricateAdLinkWithParams(result, models.AdLinkPositionBottom)
				link := FabricateLink(result)
				expectedAdLinkCollection := map[string][]string{
					models.AdLinkPositionTop: {adLink1.Link},
					models.AdLinkPositionBottom: {adLink2.Link},
					models.AdLinkPositionSide: nil,
				}
				result, err := models.GetResultByIdWithRelations(result.Id)
				if err != nil {
					Fail("Failed to get result by ID")
				}

				presenter := presenters.GetResultPresenter(result)

				Expect(presenter.Id).To(Equal(result.Id))
				Expect(presenter.Keyword).To(Equal(result.Keyword))
				Expect(presenter.Status).To(Equal(result.Status))
				Expect(presenter.PageCache).To(Equal(result.PageCache))
				Expect(presenter.TotalLinkCount).To(Equal(3))
				Expect(presenter.AdLinkCount).To(Equal(2))
				Expect(presenter.NonAdLinkCount).To(Equal(1))
				Expect(presenter.AdLinks).To(Equal(expectedAdLinkCollection))
				Expect(presenter.Links).To(Equal([]*models.Link{link}))
			})
		})

		Context("given no result", func() {
			It("returns nil", func() {
				presenter := presenters.GetResultPresenter(nil)

				Expect(presenter).To(BeNil())
			})
		})
	})
})
