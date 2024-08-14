package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}
