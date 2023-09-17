package lib_test

import (
	"nonelandBackendInterview/internal/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken_Validate_Success(t *testing.T) {
	jwtObj, err := lib.NewJwtObj("8888")
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtObj)

	tokenString, err := lib.CreateToken(*jwtObj)
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)

	valiObj, err := lib.ValidateToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, jwtObj, valiObj)
}
