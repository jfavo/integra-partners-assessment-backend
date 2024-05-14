package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
	"github.com/jmoiron/sqlx"
)

// CreateRepoWithMockedDBDriver returns a Repo instance with a sqlmock DB instance
// attached to perform unit testing with.
func CreateRepoWithMockedDBDriver() (database.IRepo, sqlmock.Sqlmock, func()) {
	var repo database.IRepo
	var dbMock sqlmock.Sqlmock
	var db *sql.DB

	// Create our DB driver mock to append to our mock repo
	db, dbMock, _ = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new Repo instance
	r := database.CreateDefault()
	r.DB = sqlxDB
	repo = r
	
	return repo, dbMock, func() { db.Close() }
}