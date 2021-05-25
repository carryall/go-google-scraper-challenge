package helpers_test

import (
	"go-google-scraper-challenge/models"
	. "go-google-scraper-challenge/tests/helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-google-scraper-challenge/helpers"
)

var _ = Describe("Result", func() {
	Describe("#PrepareResultSet", func() {
		Context("given a blank list", func() {
			It("returns nil", func() {
				var results []*models.Result

				Expect(helpers.PrepareResultSet(results)).To(BeNil())
			})
		})

		Context("given a list of 10 results", func() {
			It("returns a kebab case string", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				var results []*models.Result
				for i := 0; i < 10; i++ {
					results = append(results, FabricateResult(user))
				}

				resultSet := helpers.PrepareResultSet(results)

				Expect(resultSet).To(HaveLen(1))
				Expect(resultSet[0]).To(HaveLen(10))
			})
		})

		Context("given a list of more than 10 results", func() {
			It("returns two sets of result", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				var results []*models.Result
				for i := 0; i < 15; i++ {
					results = append(results, FabricateResult(user))
				}

				resultSet := helpers.PrepareResultSet(results)

				Expect(resultSet).To(HaveLen(2))
				Expect(resultSet[0]).To(HaveLen(10))
				Expect(resultSet[1]).To(HaveLen(5))
			})
		})
	})
})
