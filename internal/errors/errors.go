package errors

import (
	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
)

// Enum to store the error codes to send to the client for
// easier diagnosis of issues
type ErrorCode int

const (
	DBRepoFailedToInitialize ErrorCode = iota + 10000

	UsersRepoGetAllUsersDBQueryFail
	UsersRepoCreateUserDBQueryFail
	UsersRepoUserDuplicateUsername
	UsersRepoUserDuplicateEmail
	UsersRepoUserInvalidUserStatus
	UsersRepoUpdateUserDBQueryFail
	UsersRepoUpdateInvalidUserId
	UsersRepoDeleteUserDBQueryFail

	UsersControllerUserFailedToBindBody
	UsersControllerInvalidUserIdParam
)

var mappedErrors = map[ErrorCode]string{
	// DB creation errors
	DBRepoFailedToInitialize: constants.ErrDBRepoFailedToInitializeMessage,

	// User repo errors
	UsersRepoGetAllUsersDBQueryFail: constants.ErrUsersRepoGetAllUsersDBQueryFailMessage,
	UsersRepoCreateUserDBQueryFail:  constants.ErrUsersRepoCreateUserDBQueryFailMessage,
	UsersRepoUserDuplicateUsername:  constants.ErrUsersRepoUserDuplicateUserNameMessage,
	UsersRepoUserDuplicateEmail:     constants.ErrUsersRepoUserDuplicateEmailMessage,
	UsersRepoUserInvalidUserStatus:  constants.ErrUsersRepoInvalidUserStatusMessage,
	UsersRepoUpdateUserDBQueryFail:  constants.ErrUsersRepoUpdateUserDBQueryFailMessage,
	UsersRepoUpdateInvalidUserId:    constants.ErrUsersRepoUpdateInvalidUserIdMessage,
	UsersRepoDeleteUserDBQueryFail:  constants.ErrUsersRepoDeleteUserDBQueryFailMessage,

	// User controller errors
	UsersControllerUserFailedToBindBody: constants.ErrUsersControllerUserFailedToBindBodyFailMessage,
	UsersControllerInvalidUserIdParam:   constants.ErrUsersControllerInvalidUserIdParamMessage,
}

// GetErrorMessage returns the error message for the specified code
//
// Returns an empty string if it does not exist
func GetErrorMessage(code ErrorCode) string {
	return mappedErrors[code]
}

// GetErrorCode returns the code for the specified error message
//
// Returns 0 if the message was not found
func GetErrorCode(message string) ErrorCode {
	for code, msg := range mappedErrors {
		if message == msg {
			return code
		}
	}

	return 0
}
