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
})
