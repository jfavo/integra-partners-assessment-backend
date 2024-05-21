package mocks

import (
	"bytes"
	"log/slog"
)

type MockLogger struct {
	buff 	*bytes.Buffer
	Logger 	*slog.Logger
}

// NewMockLogger creates and returns a new MockLogger for testing
func NewMockLogger() MockLogger {
	var buff bytes.Buffer
	return MockLogger{
		buff: &buff,
		Logger: slog.New(slog.NewJSONHandler(&buff, nil)),
	}
}

// ResetBuffer will reset the buffer associated in the mock logger to empty
func (mockLogger *MockLogger) ResetBuffer() {
	mockLogger.buff.Reset()
}

// GetBufferValue returns the string representation of the bytes in the
// buffer of the mock logger.
func (mockLogger MockLogger) GetBufferValue() string {
	return mockLogger.buff.String()
}