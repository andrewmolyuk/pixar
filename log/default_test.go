package log_test

import (
	"github.com/andrewmolyuk/pixar/exitor"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"os"
	"testing"
)

// Ensure exitorMock implements IExitor interface
var _ exitor.IExitor = (*exitorMock)(nil)

type exitorMock struct {
	mock.Mock
}

func (m *exitorMock) Exit(code int) {
	m.Called(code)
}

func catchStdOut() (*os.File, *os.File, *os.File) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return stdout, r, w
}

func releaseStdOut(stdout *os.File, r *os.File, w *os.File) string {
	os.Stdout = stdout
	err := io.Closer.Close(w)
	if err != nil {
		panic(err)
	}
	output, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func TestNew_NotNil(t *testing.T) {
	// Arrange & Act
	l := log.New(true, nil, true, nil)

	// Assert
	assert.NotNil(t, l, "New logger should not be nil")
}

func TestDefault_NotNil(t *testing.T) {
	// Arrange & Act
	l := log.Default()

	// Assert
	assert.NotNil(t, l, "Default logger should not be nil")
}

func TestLogger_Debug_WithDebug(t *testing.T) {
	// Arrange
	l := log.New(true, nil, true, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Debug("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [DEBUG] (default_test.go:66) test\x1b[0m\n", output[23:], "logger should print debug message")
}

func TestLogger_Debug_WithoutDebug(t *testing.T) {
	// Arrange
	l := log.New(false, nil, true, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Debug("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, "", output, "logger should not print debug message")
}

func TestLogger_Secrets(t *testing.T) {
	// Arrange
	l := log.New(false, []string{"secret1", "secret2"}, true, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Info("%s%s", "secret1", "secret2")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [INFO] **********\x1b[0m\n", output[28:], "logger should print stars instead of secret message")
}

func TestLogger_Info(t *testing.T) {
	// Arrange
	l := log.New(false, []string{"secret1", "secret2"}, false, nil)
	stdout, r, w := catchStdOut()

	// Act
	l.Info("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [INFO] test\n", output[23:], "logger should print info message")
}

func TestLogger_Warn(t *testing.T) {
	// Arrange
	l := log.Default()
	stdout, r, w := catchStdOut()

	// Act
	l.Warn("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [WARN] test\x1b[0m\n", string(output)[28:], "logger should print warn message")
}

func TestLogger_Error(t *testing.T) {
	// Arrange
	e := &exitorMock{}
	e.On("Exit", 1).Return()

	l := log.New(true, []string{"secret1", "secret2"}, true, e)
	stdout, r, w := catchStdOut()

	// Act
	l.Error("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [ERROR] test\x1b[0m\n", string(output)[28:], "logger should print fatal message")
}

func TestDefaultLogger_Debug(t *testing.T) {
	// Arrange
	log.InitDefault(true, nil, false, nil)
	stdout, r, w := catchStdOut()

	// Act
	log.Debug("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [DEBUG] (default.go:119) test\n", output[23:], "logger should print debug message")
}

func TestDefaultLogger_Info(t *testing.T) {
	// Arrange
	log.InitDefault(true, nil, true, nil)
	stdout, r, w := catchStdOut()

	// Act
	log.Info("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [INFO] test\x1b[0m\n", output[28:], "logger should print info message")
}

func TestDefaultLogger_Warn(t *testing.T) {
	// Arrange
	log.InitDefault(false, nil, true, nil)
	stdout, r, w := catchStdOut()

	// Act
	log.Warn("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [WARN] test\x1b[0m\n", string(output)[28:], "logger should print warn message")
}

func TestDefaultLogger_Error(t *testing.T) {
	// Arrange
	e := &exitorMock{}
	e.On("Exit", 1).Return()

	log.InitDefault(true, []string{"secret1", "secret2"}, true, e)
	stdout, r, w := catchStdOut()

	// Act
	log.Error("test")

	// Assert
	output := releaseStdOut(stdout, r, w)
	assert.Equal(t, " [ERROR] test\x1b[0m\n", string(output)[28:], "logger should print fatal message")
}
