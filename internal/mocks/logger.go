package mocks

import (
	"bytes"
	"log/slog"
)

type MockLogger struct {
	buff 	*bytes.Buffer
	Logger 	*slog.Logger
}

func NewMockLogger() MockLogger {
	var buff bytes.Buffer
	return MockLogger{
		buff: &buff,
		Logger: slog.New(slog.NewJSONHandler(&buff, nil)),
	}
}

func (mockLogger *MockLogger) ResetBuffer() {
	mockLogger.buff.Reset()
}

func (mockLogger MockLogger) GetBufferValue() string {
	return mockLogger.buff.String()
}