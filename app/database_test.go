package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	LoadEnvironment("test")
	db := OpenConnection()
	assert.NotNil(t, db)
}
