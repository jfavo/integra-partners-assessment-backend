package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
	"github.com/jfavo/integra-partners-assessment-backend/internal/models"
	"github.com/jfavo/integra-partners-assessment-backend/internal/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Controller
	Repo database.Repo
}

// createDefault will update itself with necessary components 
func (uc UserController) createDefault(repo database.Repo) Controller {
	return &UserController{
		Repo: repo,
	}
}

// registerRoutes will register all controller routes to the Echo instance
func (uc UserController) registerRoutes(e *echo.Echo) Controller {
	e.GET("/users", uc.GetAllUsers)
	e.POST("/users", uc.CreateUser)
	e.PUT("/users", uc.UpdateUser)
	e.DELETE("/users/:userId", uc.DeleteUser)

	return uc
}

// @Summary Returns all users
// @Description Show all available users from data store
// @Tags 	Users
// @Produce json
// @Success 200 {object} response.Response{data=models.User,error_code=nil,error_message=nil}
// @Failure 500 {object} response.Response{data=nil,error_code=int,error_message=string}
// @Router	/users		 [get]
func (uc UserController) GetAllUsers(ctx echo.Context) error {
	users, errCode, err := uc.Repo.GetAllUsers()
	if err != nil {
		logging.ErrorWithCode(errCode, "failed to fetch user data", err)

		return ctx.JSON(http.StatusInternalServerError,
			response.Failure(errCode, errors.GetErrorMessage(errCode)))
	}

	return ctx.JSON(http.StatusOK, response.Success(users))
}

// @Summary Creates a new user
// @Description Creates a new user in the data store. Returns new user when successful
// @Tags 	Users
// @Produce json
// @Param	user body models.User true "User data to be ingested"
// @Success 200 {object} response.Response{data=[]models.User,error_code=nil,error_message=nil}
// @Failure 500 {object} response.Response{data=nil,error_code=int,error_message=string}
// @Router	/users		 [post]
func (uc UserController) CreateUser(ctx echo.Context) error {
	user := models.User{}
	if err := ctx.Bind(&user); err != nil {
		code := errors.UsersControllerUserFailedToBindBody
		errMessage := errors.GetErrorMessage(code)
		logging.ErrorWithCode(code, errMessage, err)

		return ctx.JSON(http.StatusBadRequest,
			response.Failure(code, errMessage))
	}

	newUser, errCode, err := uc.Repo.CreateUser(user)
	if err != nil {
		errMessage := errors.GetErrorMessage(errCode)
		logging.ErrorWithCode(errCode, errMessage, err)

		statusCode := getHttpStatusCodeForErr(errCode)

		return ctx.JSON(statusCode, response.Failure(errCode, errMessage))
	}

	return ctx.JSON(http.StatusOK, response.Success(newUser))
}

// @Summary Updates an existing user
// @Description Updates a new user in the data store. Returns updated user when successful
// @Tags 	Users
// @Produce json
// @Param	user body models.User true "User data to be ingested"
// @Success 200 {object} response.Response{data=[]models.User,error_code=nil,error_message=nil}
// @Failure 500 {object} response.Response{data=nil,error_code=int,error_message=string}
// @Router	/users		 [put]
func (uc UserController) UpdateUser(ctx echo.Context) error {
	user := models.User{}
	if err := ctx.Bind(&user); err != nil {
		code := errors.UsersControllerUserFailedToBindBody
		errMessage := errors.GetErrorMessage(code)
		logging.ErrorWithCode(code, errMessage, err)

		return ctx.JSON(http.StatusBadRequest,
			response.Failure(code, fmt.Sprintf("%s. %s", errMessage, err.Error())))
	}

	// Ensure that the user_id is passed
	if user.UserId == 0 {
		code := errors.UsersRepoUpdateInvalidUserId

		return ctx.JSON(http.StatusBadRequest, response.Failure(code, errors.GetErrorMessage(code)))
	}

	newUser, errCode, err := uc.Repo.UpdateUser(user)
	if err != nil {
		errMessage := errors.GetErrorMessage(errCode)
		logging.ErrorWithCode(errCode, errMessage, err)

		statusCode := getHttpStatusCodeForErr(errCode)

		return ctx.JSON(statusCode, response.Failure(errCode, errMessage))
	}

	return ctx.JSON(http.StatusOK, response.Success(newUser))
}

// @Summary Delete a user by the userId
// @Description Deletes the user from the data store with the associated ID
// @Tags 	Users
// @Produce json
// @Param 	userId path string true "User Id for the user to be removed"
// @Success 200 {object} 			response.Response{data=[]models.User,error_code=nil,error_message=nil}
// @Failure 404 {object} 			response.Response{data=nil,error_code=nil,error_message=nil}
// @Failure 500 {object} 			response.Response{data=nil,error_code=int,error_message=string}
// @Router	/users/{userId}			[delete]
func (uc UserController) DeleteUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("userId"))

	if err != nil {
		code := errors.UsersControllerInvalidUserIdParam
		message := errors.GetErrorMessage(code)
		logging.ErrorWithCode(code, message, err)

		return ctx.JSON(http.StatusBadRequest, response.Failure(code, errors.GetErrorMessage(code)))
	}

	deleted, errCode, err := uc.Repo.DeleteUser(id)
	if err != nil {
		message := errors.GetErrorMessage(errCode)
		logging.ErrorWithCode(errCode, message, err)

		return ctx.JSON(http.StatusInternalServerError, response.Failure(errCode, errors.GetErrorMessage(errCode)))
	}

	// If the DB returns empty, then we relay to the client that the user
	// for this ID was not found.
	if !deleted {
		return ctx.JSON(http.StatusNotFound, response.Success(nil))
	}

	return ctx.JSON(http.StatusOK, response.Success(id))
}

// getHttpStatusCodeForErr returns the http status code for the specified
// errors.ErrorCode.
//
// Will return http.StatusInternalServerError (500) status code if errCode
// does not have a specific status code for it.
func getHttpStatusCodeForErr(errCode errors.ErrorCode) int {
	switch errCode {
	case errors.UsersRepoUserDuplicateEmail:
		fallthrough
	case errors.UsersRepoUserDuplicateUsername:
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}