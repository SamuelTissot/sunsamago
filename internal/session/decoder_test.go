package session_test

import (
	"testing"

	"sunsamago/internal/session"

	"github.com/stretchr/testify/assert"
)

func TestDecoder_UserID(t *testing.T) {
	userId, err := session.UserID(
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI5ODc2NTQzMjEiLCJzZXNzaW9uSWQiOiIiLCJzZXNzaW9uTm9uY2UiOiIiLCJncm91cElkIjoiMTIzNDUiLCJncm91cFVybFBhdGgiOiJzb21lcGF0aCIsImlhdCI6MTcwNzc1MzQxOCwiZXhwIjoxNzEwMzQ1NDE4fQ.JI3_y4G2OcqgRjrsHo-1rGGpPCJ9nnv0LQr7USXYLMc",
	)

	assert.NoError(t, err)
	assert.Equal(t, "987654321", userId)
}

func TestDecoder_GroupID(t *testing.T) {
	groupID, err := session.GroupID(
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI5ODc2NTQzMjEiLCJzZXNzaW9uSWQiOiIiLCJzZXNzaW9uTm9uY2UiOiIiLCJncm91cElkIjoiMTIzNDUiLCJncm91cFVybFBhdGgiOiJzb21lcGF0aCIsImlhdCI6MTcwNzc1MzQxOCwiZXhwIjoxNzEwMzQ1NDE4fQ.JI3_y4G2OcqgRjrsHo-1rGGpPCJ9nnv0LQr7USXYLMc",
	)

	assert.NoError(t, err)
	assert.Equal(t, "12345", groupID)
}
