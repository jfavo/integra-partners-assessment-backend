package response_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/response"
)

var _ = Describe("Response", func() {

	Describe("Success", func() {
		It("Should return Response object with data", func() {
			expected := response.Response{
				Data: "hi",
			}

			Expect(response.Success("hi")).To(Equal(expected))
		})
	})

	Describe("Failure", func() {
		It("Should return Response object with data", func() {
			code := errors.UsersControllerInvalidUserIdParam
			errMsg := errors.GetErrorMessage(errors.UsersControllerInvalidUserIdParam)
			expected := response.Response{
				ErrorCode:    code,
				ErrorMessage: errMsg,
			}

			Expect(response.Failure(code, errMsg)).To(Equal(expected))
		})
	})
})
