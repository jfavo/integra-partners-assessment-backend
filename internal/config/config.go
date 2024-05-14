package config

import (
	"os"
	"strconv"

	"github.com/jfavo/integra-partners-assessment-backend/internal/constants"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	SSLMode  string
	// Maximum number of idle connections allowed
	MaxIdleConnections int
	// How long in minutes idle connections will stick around
	// before they get terminated
	ConnectionMaxIdleTime int
}

type ServerConfig struct {
	Port string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", constants.ServerPortDefault),
		},
		Database: DatabaseConfig{
			Host:                  getEnv("POSTGRES_HOSTNAME", constants.DBHostDefault),
			Username:              getEnv("POSTGRES_USER", constants.DBUsernameDefault),
			Password:              getEnv("POSTGRES_PASSWORD", constants.DBPasswordDefault),
			Name:                  getEnv("POSTGRES_DB", constants.DBNameDefault),
			Port:                  getEnv("POSTGRES_PORT", constants.DBPortDefault),
			SSLMode:               getEnv("POSTGRES_SSL", constants.DBSSLModeDefault),
			MaxIdleConnections:    getEnvInt("POSTGRES_MAX_IDLE_CONNS", constants.DBMaxIdleConnectionsDefault),
			ConnectionMaxIdleTime: getEnvInt("POSTGRES_CONN_MAX_IDLE_TIME", constants.DBConnectionMaxIdleTime),
		},
	}
}

// GetEnv attempts to retrieve the environment variable for the key param
// If one does not exist, returns the defaultVal
func getEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	return defaultVal
}

// GetEnvInt attempts to retrieve the environment variable for the key param and tries
// to convert it to an integer
// If one does not exist, or it fails to convert to an integer, it returns the defaultVal
func getEnvInt(key string, defaultVal int) int {
	if val, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(val)
		if err == nil {
			return i
		}
		// If there were an error, we log and let it fallthrough to return the default value
		logging.Error("GetEnvInt", "failed to convert environment variable to int", err)
	}

	return defaultVal
}
