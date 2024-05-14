package logging

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
)

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Error prints an error log with specific formatting.
//
// Format will be "{fromFunc}: {message}. Error: {err.Error()}".
// The Error will be omitted if it is nil.
func Error(fromFunc string, message string, err error) {
	logMsg := fmt.Sprintf("%s: %s", fromFunc, message)
	if err != nil {
		logMsg = fmt.Sprintf("%s. Error: %s", logMsg, err.Error())
	}

	Logger.Error(logMsg)
}

// Error prints an error log with specific formatting.
//
// Format will be "Code: {code}. {message} Error: {err.Error()}".
// The Error will be omitted if it is nil.
func ErrorWithCode(code errors.ErrorCode, message string, err error) {
	logMsg := fmt.Sprintf("Code: %d. %s", code, message)
	if err != nil {
		logMsg = fmt.Sprintf("%s. Error: %s", logMsg, err.Error())
	}
	
	Logger.Error(logMsg)
}