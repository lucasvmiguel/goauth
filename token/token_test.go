package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {

	token1 := Generate()
	token2 := Generate()
	token3 := Generate()

	assert.NotEqual(t, token1, token2, "they should not be equal")
	assert.NotEqual(t, token1, token3, "they should not be equal")
	assert.NotEqual(t, token2, token3, "they should not be equal")
}
