package config_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jfavo/integra-partners-assessment-backend/internal/config"
	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
)

var _ = Describe("Config", func() {

	Describe("New", func() {

		AfterEach(func() {
			os.Clearenv()
		})

		It("should create a new config object with environment variables", func() {
			os.Setenv("POSTGRES_HOSTNAME", "postgres")
			os.Setenv("PORT", "80")
			os.Setenv("POSTGRES_MAX_IDLE_CONNS", "5")

			config := config.New()

			Expect(config.Server.Port).To(Equal("80"))
			Expect(config.Database.Host).To(Equal("postgres"))
			Expect(config.Database.MaxIdleConnections).To(Equal(5))
		})

		It("should use all defaults when environment variables are not set", func() {
			config := config.New()

			Expect(config.Server.Port).To(Equal(constants.ServerPortDefault))
			Expect(config.Database.Name).To(Equal(constants.DBNameDefault))
			Expect(config.Database.SSLMode).To(Equal(constants.DBSSLModeDefault))
		})
	})
})
