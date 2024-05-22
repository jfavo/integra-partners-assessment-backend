package database

import (
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
	ipErrors "github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
	"github.com/jfavo/integra-partners-assessment-backend/internal/models"
)

// GetAllUsers fetchs all user entries from the DB.
//
// Returns a slice of Users.
// Returns an error and error code if creating the SQL query or querying DB fails.
// If error is returned, an error code associated with it will be returned as well.
func (r ServiceRepo) GetAllUsers() ([]models.User, ipErrors.ErrorCode, error) {
	users := []models.User{}

	rows, err := r.psql.
		Select("*").
		From(constants.UsersTableName).
		RunWith(r.DB).
		Query()

	if err != nil {
		return users, ipErrors.UsersRepoGetAllUsersDBQueryFail, err
	}

	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserId,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
			&user.UserStatus,
			&user.Department); err != nil {
			logging.Error("GetAllUsers", "failed to scan user data", err)
		}

		users = append(users, user)
	}

	return users, 0, err
}

// CreateUser adds a new user entry into the DB.
//
// Returns the created User if successful.
// Returns an error and error code if creating the SQL query or querying DB fails.
// If error is returned, an error code associated with it will be returned as well.
func (r ServiceRepo) CreateUser(user models.User) (*models.User, ipErrors.ErrorCode, error) {
	returnedUser := new(models.User)

	err := r.psql.
		Insert(constants.UsersTableName).
		Columns("user_name", "first_name", "last_name", "email", "user_status", "department").
		Values(user.Username, user.Firstname, user.Lastname, user.Email, user.UserStatus, user.Department).
		Suffix("RETURNING *").
		RunWith(r.DB).
		QueryRow().
		Scan(&returnedUser.UserId,
			&returnedUser.Username,
			&returnedUser.Firstname,
			&returnedUser.Lastname,
			&returnedUser.Email,
			&returnedUser.UserStatus,
			&returnedUser.Department)

	if err != nil {
		// Check duplicate username/email err and return the appropriate error
		if valid, errCode := checkUserDBError(err); valid {
			return nil, errCode, err
		}

		return nil, ipErrors.UsersRepoCreateUserDBQueryFail, err
	}

	return returnedUser, 0, nil
}

// UpdateUser updates an existing user entry in the DB.
//
// Returns the updated User if successful.
// Returns an error and error code if creating the SQL query or querying DB fails.
// If error is returned, an error code associated with it will be returned as well.
func (r ServiceRepo) UpdateUser(user models.User) (*models.User, ipErrors.ErrorCode, error) {
	returnedUser := new(models.User)

	// Creates our set statements
	setMap := createUpdateSetMap(user)

	err := r.psql.Update(constants.UsersTableName).
		SetMap(setMap).
		Where("user_id = ?", user.UserId).
		Suffix("RETURNING *").
		RunWith(r.DB).
		QueryRow().
		Scan(&returnedUser.UserId,
			&returnedUser.Username,
			&returnedUser.Firstname,
			&returnedUser.Lastname,
			&returnedUser.Email,
			&returnedUser.UserStatus,
			&returnedUser.Department)

	if err != nil {
		// Check duplicate username/email err and return the appropriate error
		if valid, errCode := checkUserDBError(err); valid {
			return nil, errCode, err
		}

		return nil, ipErrors.UsersRepoUpdateUserDBQueryFail, err
	}

	return returnedUser, 0, nil
}

// DeleteUser remove user entry in the DB with the associated id.
//
// Returns true if the user was successfully removed.
// Returns an error and error code if creating the SQL query or querying DB fails.
// If error is returned, an error code associated with it will be returned as well.
func (r ServiceRepo) DeleteUser(userId int) (bool, ipErrors.ErrorCode, error) {
	res, err := r.psql.Delete(constants.UsersTableName).
		Where("user_id = ?", userId).
		Suffix("RETURNING user_id").
		RunWith(r.DB).
		Exec()

	if err != nil {
		return false, ipErrors.UsersRepoDeleteUserDBQueryFail, err
	}

	rows, _ := res.RowsAffected()
	return rows > 0, 0, nil
}

// checkUserDBError checks to see if error from the DB is specific
// to invalid data and returns the appropriate error codes for them
//
// Returns true and the appropriate error code if successful
// Otherwise, returns false with 0 as the code
func checkUserDBError(err error) (bool, ipErrors.ErrorCode) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		var errCode ipErrors.ErrorCode

		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			if strings.Contains(err.Error(), "user_name") {
				errCode = ipErrors.UsersRepoUserDuplicateUsername
			} else if strings.Contains(err.Error(), "email") {
				errCode = ipErrors.UsersRepoUserDuplicateEmail
			}
		case pgerrcode.InvalidTextRepresentation:
			if strings.Contains(err.Error(), "user_status") {
				errCode = ipErrors.UsersRepoUserInvalidUserStatus
			}
		}

		if errCode != 0 {
			return true, errCode
		}
	}

	return false, 0
}

// createUpdateSetMap creates and returns a squirrel.Eq{} with the
// Set statements for our user update query.
func createUpdateSetMap(user models.User) squirrel.Eq {
	setMap := squirrel.Eq{}

	if user.Username != "" {
		setMap["user_name"] = user.Username
	}

	if user.Firstname != "" {
		setMap["first_name"] = user.Firstname
	}

	if user.Lastname != "" {
		setMap["last_name"] = user.Lastname
	}

	if user.Email != "" {
		setMap["email"] = user.Email
	}

	if user.UserStatus != "" {
		setMap["user_status"] = user.UserStatus
	}

	setMap["department"] = user.Department

	return setMap
}
