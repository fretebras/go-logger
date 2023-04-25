package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Cria um arquivo .config.env com valores para as variáveis de ambiente
	err := ioutil.WriteFile(".config.env", []byte("APP=myapp\nENVIRONMENT=myenv\nVERSION=1.0"), 0644)
	assert.NoError(t, err)

	// Defina as variáveis de ambiente temporariamente
	os.Setenv("APP", "myapp")
	os.Setenv("ENVIRONMENT", "myenv")
	os.Setenv("VERSION", "1.0")

	// Chame a função NewConfig() e verifique os valores
	config := NewConfig()
	assert.Equal(t, "myapp", config.App)
	assert.Equal(t, "myenv", config.Environment)
	assert.Equal(t, "1.0", config.Version)

	// Restaure as variáveis de ambiente para os valores originais
	os.Unsetenv("APP")
	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("VERSION")

	// Remove o arquivo .config.env criado
	err = os.Remove(".config.env")
	assert.NoError(t, err)
}

func TestNewConfig_PanicsOnLoadError(t *testing.T) {
	assert.Panics(t, func() {
		os.Setenv("APP", "test-app")
		os.Setenv("ENVIRONMENT", "test-environment")
		os.Setenv("VERSION", "test-version")

		NewConfig()
	}, "NewConfig() should panic on load error")

	os.Unsetenv("APP")
	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("VERSION")
}

func TestNewConfigPanicWhenAppNotRead(t *testing.T) {
	err := ioutil.WriteFile(".config.env", []byte("ENVIRONMENT=myenv"), 0644)
	assert.NoError(t, err)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("NewConfig() should have panicked when APP variable not read")
		}
	}()
	NewConfig()

	os.Unsetenv("TEST")
	err = os.Remove(".config.env")
	assert.NoError(t, err)
}

func TestNewConfigPanicWhenEnvironmentNotRead(t *testing.T) {
	err := ioutil.WriteFile(".config.env", []byte("APP=myapp\nVERSION=1.0"), 0644)
	assert.NoError(t, err)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("NewConfig() should have panicked when ENVIRONMENT variable not read")
		}
	}()
	NewConfig()
	os.Unsetenv("APP")
	os.Unsetenv("VERSION")

	err = os.Remove(".config.env")
	assert.NoError(t, err)
}

func TestNewConfigPanicWhenVersionNotRead(t *testing.T) {
	err := ioutil.WriteFile(".config.env", []byte("APP=myapp\nENVIRONMENT=myenv"), 0644)
	assert.NoError(t, err)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("NewConfig() should have panicked when VERSION variable not read")
		}
	}()
	NewConfig()

	os.Unsetenv("APP")
	os.Unsetenv("ENVIRONMENT")

	err = os.Remove(".config.env")
	assert.NoError(t, err)
}
