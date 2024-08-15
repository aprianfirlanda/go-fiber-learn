package app

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadEnvironment()

	assert.NotEqual(t, "", viper.GetString("DATABASE_USER"))
	assert.NotEqual(t, "", viper.GetString("DATABASE_PASSWORD"))
	assert.NotEqual(t, "", viper.GetString("DATABASE_NAME"))
	assert.NotEqual(t, "", viper.GetString("DATABASE_HOST"))
	assert.NotEqual(t, 0, viper.GetInt("DATABASE_PORT"))
	assert.NotEqual(t, "", viper.GetString("DATABASE_SSLMODE"))
}
