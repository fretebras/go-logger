package logger

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	logger := New()
	logger.SetLogService("test-service")
	var buf bytes.Buffer
	logger.logger = logger.logger.Output(&buf)
	logger.Debug("debug message", "context1")
	assert.Contains(t, buf.String(), "debug message")
	assert.Contains(t, buf.String(), "test-service")
	assert.Contains(t, buf.String(), "context1")
}

func TestLogger_Info(t *testing.T) {
	logger := New()
	logger.SetLogService("test-service-info")
	var buf bytes.Buffer
	logger.logger = logger.logger.Output(&buf)
	logger.Info("info message", "context-info")
	assert.Contains(t, buf.String(), "info message")
	assert.Contains(t, buf.String(), "test-service-info")
	assert.Contains(t, buf.String(), "context-info")
}

func TestLogger_Warning(t *testing.T) {
	logger := New()
	logger.SetLogService("test-service-warning")
	var buf bytes.Buffer
	logger.logger = logger.logger.Output(&buf)
	logger.Info("warning message", "context-warning")
	assert.Contains(t, buf.String(), "warning message")
	assert.Contains(t, buf.String(), "test-service-warning")
	assert.Contains(t, buf.String(), "context-warning")
}

func TestLogger_Error(t *testing.T) {
	logger := New()
	logger.SetLogService("test-service-error")
	var buf bytes.Buffer
	logger.logger = logger.logger.Output(&buf)
	err := errors.New("test error log")
	logger.Error(err, "context-error")
	assert.Contains(t, buf.String(), err.Error())
	assert.Contains(t, buf.String(), "test-service-error")
	assert.Contains(t, buf.String(), "context-error")
}

func TestLogger_Critical(t *testing.T) {
	logger := New()
	logger.SetLogService("test-service-critical")
	var buf bytes.Buffer
	logger.logger = logger.logger.Output(&buf)
	err := errors.New("test critical log")
	logger.Error(err, "context-critical")
	assert.Contains(t, buf.String(), err.Error())
	assert.Contains(t, buf.String(), "test-service-critical")
	assert.Contains(t, buf.String(), "context-critical")
}
