package database

import (
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/jfavo/integra-partners-assessment-backend/internal/config"
	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
	"github.com/jfavo/integra-partners-assessment-backend/internal/models"
)

type Repo interface {
	GetAllUsers() ([]models.User, errors.ErrorCode, error)
	CreateUser(models.User) (*models.User, errors.ErrorCode, error)
	UpdateUser(models.User) (*models.User, errors.ErrorCode, error)
	DeleteUser(userId int) (bool, errors.ErrorCode, error)
}

type ServiceRepo struct {
	DB 		*sqlx.DB
	psql 	squirrel.StatementBuilderType
}

func CreateDefault() ServiceRepo {
	return ServiceRepo{
		psql: 	squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

// CreateNewRepo returns a new repo instance with an initialized sqlx.DB client
//
// Returns error if either the sqlx.DB client fails to open, or if we cannot
// verify the connection to the DB.
// If error is returned, an error code associated with it will be returned as well.
func CreateNewRepo(dbConfig config.DatabaseConfig) (Repo, error) {
	dsnString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s database=%s sslmode=%s", 
		dbConfig.Username, 
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port, 
		dbConfig.Name, 
		dbConfig.SSLMode)

	db, err := sqlx.Connect("pgx", dsnString)
	if err != nil {
		return nil, err
	}

	// Ensure the connection pool doesn't hold onto
	// idle connections for too long
	db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(dbConfig.ConnectionMaxIdleTime) * time.Minute)

	logging.Logger.Info("Successfully connected to DB")

	repo := CreateDefault()
	repo.DB = db

	return &repo, nil
}

