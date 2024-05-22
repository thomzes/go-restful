package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomzes/go-restful/simple"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
