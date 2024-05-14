package controllers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/labstack/echo/v4"

	"github.com/jfavo/integra-partners-assessment-backend/internal/controllers"
	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
)

var _ = Describe("Controller", func() {

	Describe("Initialize", Ordered, func() {
		var repo database.Repo
		var e *echo.Echo

		BeforeAll(func() {
			repo = database.CreateDefault()
			e = echo.New()
		})

		It("should create new user controller", func() {
			controllers.Initialize[controllers.UserController](&repo, e)

			Expect(len(e.Routes())).To(Equal(4))
		})
	})
})
