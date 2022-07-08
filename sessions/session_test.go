package sessions_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	"go-google-scraper-challenge/bootstrap"
	"go-google-scraper-challenge/sessions"
	. "go-google-scraper-challenge/test"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sessions", func() {
	engine := gin.Default()
	engine = bootstrap.SetupSession(engine)

	Describe("#GetCurrentUser", func() {
		engine.GET("/test-get-current-user", func(ctx *gin.Context) {
			returnUser := sessions.GetCurrentUser(ctx)

			if returnUser != nil {
				ctx.String(http.StatusOK, fmt.Sprint(returnUser.ID))
			} else {
				ctx.String(http.StatusNotFound, "")
			}
		})

		Context("given session have the user ID", func() {
			It("returns the user with the ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				cookie := FabricateAuthUserCookie(user.ID)
				responseRecorder := httptest.NewRecorder()
				request, err := http.NewRequest("GET", "/test-get-current-user", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				if err != nil {
					Fail("Fail to test the session " + err.Error())
				}
				engine.ServeHTTP(responseRecorder, request)
				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal(fmt.Sprint(user.ID)))
			})

			Context("given user ID does NOT exist in the database", func() {
				It("returns nil", func() {
					cookie := FabricateAuthUserCookie(999)
					responseRecorder := httptest.NewRecorder()
					request, err := http.NewRequest("GET", "/test-get-current-user", nil)
					request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
					if err != nil {
						Fail("Fail to test the session " + err.Error())
					}
					engine.ServeHTTP(responseRecorder, request)
					response := responseRecorder.Result()
					responseBody := GetResponseBody(response)

					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
					Expect(responseBody).To(BeEmpty())
				})
			})
		})

		Context("given session does NOT have the user ID", func() {
			It("returns nil", func() {
				responseRecorder := httptest.NewRecorder()
				request, err := http.NewRequest("GET", "/test-get-current-user", nil)
				if err != nil {
					Fail("Fail to test the session " + err.Error())
				}
				engine.ServeHTTP(responseRecorder, request)
				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				Expect(responseBody).To(BeEmpty())
			})
		})
	})

	Describe("#SetCurrentUser", func() {
		engine.GET("/test-set-current-user", func(ctx *gin.Context) {
			userIDStr, _ := ctx.GetQuery("userID")
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				ctx.String(http.StatusBadRequest, userIDStr)
			}

			sessions.SetCurrentUser(ctx, int64(userID))

			ctx.String(http.StatusOK, "")
		})

		Context("given a user ID", func() {
			It("sets the current user ID to session", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				responseRecorder := httptest.NewRecorder()
				request, err := http.NewRequest("GET", fmt.Sprintf("/test-set-current-user?userID=%d", user.ID), nil)
				if err != nil {
					Fail("Fail to test the session " + err.Error())
				}
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				encodedSession := ""
				for _, cookie := range response.Cookies() {
					if cookie.Name == "google_scraper_session" {
						encodedSession = cookie.Value
					}
				}
				decodedCookie := DecodeCookieString(encodedSession)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(decodedCookie[sessions.CurrentUserKey]).To(Equal(fmt.Sprint(user.ID)))
			})
		})
	})

	Describe("#SetFlash", func() {
		engine.GET("/test-set-flash", func(ctx *gin.Context) {
			flashType, _ := ctx.GetQuery("type")
			flashMessage, _ := ctx.GetQuery("message")

			sessions.SetFlash(ctx, flashType, flashMessage)

			ctx.String(http.StatusOK, "")
		})

		Context("given flashes", func() {
			It("sets flashes to the session", func() {
				flashType := sessions.FlashTypeError
				flashMessage := "ERRORMSG"
				responseRecorder := httptest.NewRecorder()
				request, err := http.NewRequest("GET", fmt.Sprintf("/test-set-flash?type=%s&message=%s", flashType, flashMessage), nil)
				if err != nil {
					Fail("Fail to test the session " + err.Error())
				}
				engine.ServeHTTP(responseRecorder, request)

				response := responseRecorder.Result()
				encodedSession := ""
				for _, cookie := range response.Cookies() {
					if cookie.Name == "google_scraper_session" {
						encodedSession = cookie.Value
					}
				}
				decodedCookie := DecodeCookieString(encodedSession)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(decodedCookie).To(HaveKey(flashType))
				Expect(decodedCookie[sessions.FlashTypeError]).To(ConsistOf(flashMessage))
			})
		})
	})

	Describe("#GetFlash", func() {
		engine.GET("/test-get-flash", func(ctx *gin.Context) {
			flashes := sessions.GetFlash(ctx)

			ctx.JSON(http.StatusOK, flashes)
		})

		Context("given the session with flashes", func() {
			It("returns the flashes", func() {
				flashes := map[string]interface{}{}
				flashes[sessions.FlashTypeError] = []interface{}{"Error Message"}
				flashes[sessions.FlashTypeInfo] = []interface{}{"Info Message"}
				flashes[sessions.FlashTypeSuccess] = []interface{}{"Success Message"}
				expectedFlashes, _ := json.Marshal(flashes)
				cookie := FabricateCookieWithFlashes(flashes)

				responseRecorder := httptest.NewRecorder()
				request, err := http.NewRequest("GET", "/test-get-flash", nil)
				request.Header.Set("Cookie", cookie.Name+"="+cookie.Value)
				if err != nil {
					Fail("Fail to test the session " + err.Error())
				}
				engine.ServeHTTP(responseRecorder, request)
				response := responseRecorder.Result()
				responseBody := GetResponseBody(response)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(responseBody).To(Equal(string(expectedFlashes)))
			})
		})
	})

	AfterEach(func() {
		CleanupDatabase([]string{"users"})
	})
})
