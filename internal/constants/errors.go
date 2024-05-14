// package constants provides constant values for
package constants

const (
	// Contains all returned error messages for the clients
	ErrDBRepoFailedToInitializeMessage = "failed to initialize DB"

	ErrUsersRepoGetAllUsersDBQueryFailMessage = "failed to get users from records"
	ErrUsersRepoCreateUserDBQueryFailMessage  = "failed to create user in records"
	ErrUsersRepoUserDuplicateUserNameMessage  = "user with username already exists"
	ErrUsersRepoUserDuplicateEmailMessage     = "user with email already exists"
	ErrUsersRepoInvalidUserStatusMessage      = "input for user_status is invalid"
	ErrUsersRepoUpdateUserDBQueryFailMessage  = "failed to update user in records"
	ErrUsersRepoUpdateInvalidUserIdMessage    = "user Id is required to update the user"
	ErrUsersRepoDeleteUserDBQueryFailMessage  = "failed to delete user from records"

	ErrUsersControllerUserFailedToBindBodyFailMessage = "user input body is invalid"
	ErrUsersControllerInvalidUserIdParamMessage       = "user id passed as URL param is invalid"
)
