package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"
	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
	"github.com/jfavo/integra-partners-assessment-backend/internal/controllers"
	ipErrors "github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
	"github.com/jfavo/integra-partners-assessment-backend/internal/mocks"
	"github.com/jfavo/integra-partners-assessment-backend/internal/models"
	"github.com/jfavo/integra-partners-assessment-backend/internal/response"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func createTestRequest(method string, url string, body interface{}) *http.Request {
	var reader *bytes.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		reader = bytes.NewReader(b)
		return httptest.NewRequest(method, url, reader)
	}
	
	return httptest.NewRequest(method, url, nil)
}

var _ = Describe("UserController", Ordered, func() {

	var (
		mockCtrl 		*gomock.Controller
		mockRepo 		*mocks.MockIRepo
		e 		 		*echo.Echo

		req *http.Request
		rec *httptest.ResponseRecorder
		ctx echo.Context
	)

	BeforeAll(func() {
		mockLogger := mocks.NewMockLogger()
		logging.Logger = mockLogger.Logger
	})
	
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockIRepo(mockCtrl)
		e = echo.New()

		rec = httptest.NewRecorder()
	})

	Describe("GetAllUsers", func() {

		BeforeEach(func() {
			req = createTestRequest(http.MethodGet, "/users", nil)
			ctx = e.NewContext(req, rec)
		})

		It("should return all users in data store", func() {
			expected := constants.TestUsers

			mockRepo.EXPECT().GetAllUsers().Return(expected, ipErrors.ErrorCode(0), nil)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.GetAllUsers(ctx)

			b, _ := json.Marshal(response.Success(expected))
			
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})

		It("should return error when DB throws error", func() {
			expectedErr := errors.New("DB had an error!")
			expectedCode := ipErrors.UsersRepoGetAllUsersDBQueryFail
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			mockRepo.EXPECT().GetAllUsers().Return([]models.User{}, expectedCode, expectedErr)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.GetAllUsers(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})
	})

	Describe("CreateUser", func() {

		It("should create new user successfully", func() {
			expected := constants.TestUsers[0]
			req = createTestRequest(http.MethodPost, "/users", expected)
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().CreateUser(expected).Return(&expected, ipErrors.ErrorCode(0), nil)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.CreateUser(ctx)

			b, _ := json.Marshal(response.Success(expected))
			
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})

		It("should fail to bind request body", func() {
			expectedCode := ipErrors.UsersControllerUserFailedToBindBody
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPost, "/users", `{"invalid":"type"}`)
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)

			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.CreateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})

		It("should fail if user with username already exists in DB", func() {
			expectedCode := ipErrors.UsersRepoUserDuplicateUsername
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPost, "/users", constants.TestUsers[0])
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().CreateUser(constants.TestUsers[0]).Return(nil, expectedCode, errors.New("Failed!"))
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.CreateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusConflict))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})

		It("should fail with DB error", func() {
			expectedCode := ipErrors.UsersRepoCreateUserDBQueryFail
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPost, "/users", constants.TestUsers[0])
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().CreateUser(constants.TestUsers[0]).Return(nil, expectedCode, errors.New("Failed!"))
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.CreateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})
	})

	Describe("UpdateUser", func() {

		It("should update user successfully", func() {
			expected := constants.TestUsers[0]
			req = createTestRequest(http.MethodPut, "/users", expected)
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().UpdateUser(expected).Return(&expected, ipErrors.ErrorCode(0), nil)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.UpdateUser(ctx)

			b, _ := json.Marshal(response.Success(expected))
			
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 

		It("should fail if user_id is not passed to body", func() {
			input := constants.TestUsers[0]
			input.UserId = 0
			expectedCode := ipErrors.UsersRepoUpdateInvalidUserId
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPut, "/users", input)
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			userController := &controllers.UserController{}
			userController.UpdateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 

		It("should fail if user with email already exists in DB", func() {
			expectedCode := ipErrors.UsersRepoUserDuplicateEmail
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPost, "/users", constants.TestUsers[0])
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().CreateUser(constants.TestUsers[0]).Return(nil, expectedCode, errors.New("Failed!"))
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.CreateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusConflict))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		})

		It("should fail when DB returns an error", func() {
			input := constants.TestUsers[0]

			expectedCode := ipErrors.UsersRepoUpdateUserDBQueryFail
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			req = createTestRequest(http.MethodPut, "/users", input)
			req.Header.Add("Content-Type", "application/json")
			ctx = e.NewContext(req, rec)
			
			mockRepo.EXPECT().UpdateUser(input).Return(nil, expectedCode, errors.New("DB error occurred!"))
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.UpdateUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 
	})

	Describe("DeleteUser", func() {
		var inputId int

		BeforeEach(func() {
			inputId = 1

			req = createTestRequest(http.MethodDelete, "/users/:userId", nil)
			ctx = e.NewContext(req, rec)
			ctx.SetParamNames("userId")
			ctx.SetParamValues(fmt.Sprintf("%d", inputId))
		})

		It("should delete user successfully", func() {
			mockRepo.EXPECT().DeleteUser(inputId).Return(true, ipErrors.ErrorCode(0), nil)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.DeleteUser(ctx)

			b, _ := json.Marshal(response.Success(inputId))
			
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 

		It("should return NotFound if user with Id does not exist", func() {	
			mockRepo.EXPECT().DeleteUser(inputId).Return(false, ipErrors.ErrorCode(0), nil)
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.DeleteUser(ctx)

			b, _ := json.Marshal(response.Success(nil))
			
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 

		It("should return error when DB returns an error", func() {	
			expectedCode := ipErrors.UsersRepoDeleteUserDBQueryFail
			expectedMsg := ipErrors.GetErrorMessage(expectedCode)

			mockRepo.EXPECT().DeleteUser(inputId).Return(false, expectedCode, errors.New("DB error occurred!"))
			userController := &controllers.UserController{
				Repo: mockRepo,
			}
			userController.DeleteUser(ctx)

			b, _ := json.Marshal(response.Failure(expectedCode, expectedMsg))
			
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(strings.ReplaceAll(rec.Body.String(), "\n", "")).To(Equal(string(b)))
		}) 
	})
})
