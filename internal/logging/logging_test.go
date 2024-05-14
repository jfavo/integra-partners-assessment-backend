package logging_test

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ipErrors "github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
	"github.com/jfavo/integra-partners-assessment-backend/internal/mocks"
)

var _ = Describe("Logging", Ordered, func() {
	var mockLogger mocks.MockLogger

	BeforeAll(func() {
		mockLogger = mocks.NewMockLogger()
		logging.Logger = mockLogger.Logger
	})

	BeforeEach(func() {
		mockLogger.ResetBuffer()
	})

	Describe("Error", func() {
		It("should log an error message in a specified format, with error", func() {
			fromFunc := "testFunc"
			message := "test message to use!"
			err := errors.New("a fake error to use!")
			expected := fmt.Sprintf("%s: %s. Error: %s", fromFunc, message, err.Error())

			logging.Error(fromFunc, message, err)
			Expect(strings.Contains(mockLogger.GetBufferValue(), expected)).To(Equal(true))
		})
		It("should log an error message in a specified format, without error", func() {
			fromFunc := "testFunc"
			message := "test message to use!"
			expected := fmt.Sprintf("%s: %s", fromFunc, message)

			logging.Error(fromFunc, message, nil)
			Expect(strings.Contains(mockLogger.GetBufferValue(), expected)).To(Equal(true))
		})
	})

	Describe("ErrorWithCode", func() {
		It("should log an error message with code in a specified format, with error", func() {
			code := ipErrors.UsersControllerInvalidUserIdParam
			message := "test message to use!"
			err := errors.New("a fake error to use!")
			expected := fmt.Sprintf("Code: %d. %s. Error: %s", code, message, err.Error())

			logging.ErrorWithCode(code, message, err)
			Expect(strings.Contains(mockLogger.GetBufferValue(), expected)).To(Equal(true))
		})

		It("should log an error message with code in a specified format, without error", func() {
			code := ipErrors.UsersControllerInvalidUserIdParam
			message := "test message to use!"

			expected := fmt.Sprintf("Code: %d. %s", code, message)

			logging.ErrorWithCode(code, message, nil)
			Expect(strings.Contains(mockLogger.GetBufferValue(), expected)).To(Equal(true))
		})
	})
})
