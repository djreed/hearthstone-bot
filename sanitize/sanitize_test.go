package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftTickReplace(t *testing.T) {
	ideal := "witch's apprentice"
	actual := "witch‘s apprentice"

	assert.Equal(t, SanitizeString(actual), ideal)
}

func TestRightTickReplace(t *testing.T) {
	ideal := "witch's apprentice"
	actual := "witch’s apprentice"

	assert.Equal(t, SanitizeString(actual), ideal)
}

func TestTicksAreDistinct(t *testing.T) {
	leftTick := "‘"  // U+2018
	rightTick := "’" // U+2019

	assert.NotEqual(t, leftTick, rightTick)
	assert.Equal(t, SanitizeString(leftTick), SanitizeString(rightTick))
}
