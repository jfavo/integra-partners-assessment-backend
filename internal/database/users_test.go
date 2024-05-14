package database_test

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
	ipErrors "github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/mocks"
	"github.com/jfavo/integra-partners-assessment-backend/internal/models"
)

var _ = Describe("Users", Ordered, func() {
	var repo database.IRepo
	var dbMock sqlmock.Sqlmock
	var closeFunc func()
	var testUser models.User

	BeforeAll(func() {
		repo, dbMock, closeFunc = mocks.CreateRepoWithMockedDBDriver()
	})

	BeforeEach(func() {
		testUser = constants.TestUsers[0]
	})

	AfterAll(func() {
		closeFunc()
	})

	Describe("GetAllUsers", func() {
		It("should return a list of users", func() {

			rows := sqlmock.NewRows([]string{"user_id", "user_name", "first_name", "last_name", "email", "user_status", "department"}).
				AddRow("1", "testUser", "test", "user", "test@user.com", "A", "sales").
				AddRow("2", "testUser2", "test", "user2", "test2@user.com", "I", "accounting")

			dbMock.ExpectQuery("SELECT * FROM integra_partners.users").
				WillReturnRows(rows)

			users, errCode, err := repo.GetAllUsers()

			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(err).To(BeNil())
			Expect(len(users)).To(Equal(2))
			// Spot check some values
			Expect(users[0].Email).To(Equal("test@user.com"))
			Expect(users[1].UserStatus).To(Equal("I"))
		})

		It("should return error when db query fails", func() {
			expectedErr := errors.New("DB query failed!")

			dbMock.ExpectQuery("SELECT * FROM integra_partners.users").
				WillReturnError(expectedErr)

			users, errCode, err := repo.GetAllUsers()

			Expect(errCode).To(Equal(ipErrors.UsersRepoGetAllUsersDBQueryFail))
			Expect(err).To(Equal(expectedErr))
			Expect(len(users)).To(Equal(0))
		})
	})

	Describe("CreateUser", func() {
		insertQuery := fmt.Sprintf(
			"INSERT INTO %s (user_name,first_name,last_name,email,user_status,department) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *",
			constants.UsersTableName)

		It("should successfully create a new user", func() {
			rows := sqlmock.NewRows([]string{"user_id", "user_name", "first_name", "last_name", "email", "user_status", "department"}).
				AddRow("1", "testUser", "test", "user", "test@user.com", "A", "sales")

			dbMock.ExpectQuery(insertQuery).
				WillReturnRows(rows)

			user, errCode, err := repo.CreateUser(testUser)

			Expect(err).To(BeNil())
			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(user).To(Equal(&testUser))
		})

		It("should fail due to duplicate username", func() {
			expectedErr := &pgconn.PgError{
				Code:    pgerrcode.UniqueViolation,
				Message: "duplicate key value violates unique constraint \"users_user_name_idx\"",
			}

			dbMock.ExpectQuery(insertQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.CreateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserDuplicateUsername))
			Expect(user).To(BeNil())
		})

		It("should fail due to duplicate email", func() {
			expectedErr := &pgconn.PgError{Code: pgerrcode.UniqueViolation, Message: "duplicate key value violates unique constraint \"users_email_idx\""}

			dbMock.ExpectQuery(insertQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.CreateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserDuplicateEmail))
			Expect(user).To(BeNil())
		})

		It("should fail due to invalid user_status value", func() {
			expectedErr := &pgconn.PgError{
				Code:    pgerrcode.InvalidTextRepresentation,
				Message: "invalid input value for enum integra_partners.user_status: \"e\" (SQLSTATE 22P02)",
			}

			dbMock.ExpectQuery(insertQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.CreateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserInvalidUserStatus))
			Expect(user).To(BeNil())
		})

		It("should fail due to DB error", func() {
			expectedErr := errors.New("DB encountered and error!")

			dbMock.ExpectQuery(insertQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.CreateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoCreateUserDBQueryFail))
			Expect(user).To(BeNil())
		})
	})

	Describe("UpdateUser", func() {
		fullUpdateQuery := fmt.Sprintf(
			"UPDATE %s SET department = $1, email = $2, first_name = $3, last_name = $4, user_name = $5, user_status = $6 WHERE user_id = $7 RETURNING *",
			constants.UsersTableName)

		It("should successfully update a user", func() {
			rows := sqlmock.NewRows([]string{"user_id", "user_name", "first_name", "last_name", "email", "user_status", "department"}).
				AddRow("1", "testUser", "test", "user", "test@user.com", "A", "sales")

			dbMock.ExpectQuery(fullUpdateQuery).
				WillReturnRows(rows)

			user, errCode, err := repo.UpdateUser(testUser)

			Expect(err).To(BeNil())
			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(user).To(Equal(&testUser))
		})

		It("should successfully update a few user fields", func() {
			rows := sqlmock.NewRows([]string{"user_id", "user_name", "first_name", "last_name", "email", "user_status", "department"}).
				AddRow("1", "testUserChange", "test", "user", "test@user.com", "A", "warehouse")

			updateUser := models.User{UserId: 1, Username: "testUserChange", Department: "warehouse"}
			partialUpdateQuery := fmt.Sprintf(
				"UPDATE %s SET department = $1, user_name = $2 WHERE user_id = $3 RETURNING *",
				constants.UsersTableName)

			dbMock.ExpectQuery(partialUpdateQuery).
				WillReturnRows(rows)

			user, errCode, err := repo.UpdateUser(updateUser)

			testUser.Username = updateUser.Username
			testUser.Department = updateUser.Department

			Expect(err).To(BeNil())
			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(user).To(Equal(&testUser))
		})

		It("should fail due to duplicate username", func() {
			expectedErr := &pgconn.PgError{
				Code:    pgerrcode.UniqueViolation,
				Message: "duplicate key value violates unique constraint \"users_user_name_idx\"",
			}

			dbMock.ExpectQuery(fullUpdateQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.UpdateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserDuplicateUsername))
			Expect(user).To(BeNil())
		})

		It("should fail due to duplicate email", func() {
			expectedErr := &pgconn.PgError{Code: pgerrcode.UniqueViolation, Message: "duplicate key value violates unique constraint \"users_email_idx\""}

			dbMock.ExpectQuery(fullUpdateQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.UpdateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserDuplicateEmail))
			Expect(user).To(BeNil())
		})

		It("should fail due to invalid user_status value", func() {
			expectedErr := &pgconn.PgError{
				Code:    pgerrcode.InvalidTextRepresentation,
				Message: "invalid input value for enum integra_partners.user_status: \"e\" (SQLSTATE 22P02)",
			}

			dbMock.ExpectQuery(fullUpdateQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.UpdateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUserInvalidUserStatus))
			Expect(user).To(BeNil())
		})

		It("should fail due to DB error", func() {
			expectedErr := errors.New("DB encountered and error!")

			dbMock.ExpectQuery(fullUpdateQuery).
				WillReturnError(expectedErr)

			user, errCode, err := repo.UpdateUser(testUser)

			Expect(err).To(Equal(expectedErr))
			Expect(errCode).To(Equal(ipErrors.UsersRepoUpdateUserDBQueryFail))
			Expect(user).To(BeNil())
		})
	})

	Describe("DeleteUser", func() {
		deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 RETURNING user_id", constants.UsersTableName)

		It("should successfully delete user", func() {
			dbMock.ExpectExec(deleteQuery).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 1))

			deleted, errCode, err := repo.DeleteUser(testUser.UserId)

			Expect(err).To(BeNil())
			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(deleted).To(Equal(true))
		})

		It("should return nil user if they do not exist", func() {
			dbMock.ExpectExec(deleteQuery).
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 0))

			deleted, errCode, err := repo.DeleteUser(testUser.UserId)

			Expect(err).To(BeNil())
			Expect(errCode).To(Equal(ipErrors.ErrorCode(0)))
			Expect(deleted).To(Equal(false))
		})

		It("should return error if DB throws error", func() {
			err := errors.New("DB threw an error!")

			dbMock.ExpectExec(deleteQuery).
				WithArgs(1).
				WillReturnError(err)

			deleted, errCode, err := repo.DeleteUser(testUser.UserId)

			Expect(err).To(Equal(err))
			Expect(errCode).To(Equal(ipErrors.UsersRepoDeleteUserDBQueryFail))
			Expect(deleted).To(Equal(false))
		})
	})
})
