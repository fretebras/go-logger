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
