package errors_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
)

var _ = Describe("Errors", func() {

	Describe("GetErrorMessage", func() {
		It("should return message associated with the code", func() {
			Expect(
				errors.GetErrorMessage(errors.DBRepoFailedToInitialize)).
				To(Equal(constants.ErrDBRepoFailedToInitializeMessage))
			Expect(
				errors.GetErrorMessage(errors.UsersControllerInvalidUserIdParam)).
				To(Equal(constants.ErrUsersControllerInvalidUserIdParamMessage))
		})
	})

	Describe("GetErrorCode", func() {
		It("should return code for associated error message", func() {
			Expect(
				errors.GetErrorCode(constants.ErrUsersRepoCreateUserDBQueryFailMessage)).
				To(Equal(errors.UsersRepoCreateUserDBQueryFail))
			Expect(
				errors.GetErrorCode(constants.ErrUsersRepoUpdateInvalidUserIdMessage)).
				To(Equal(errors.UsersRepoUpdateInvalidUserId))
		})
	})
})
