package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testName   = "Leeroy Jenkins"
	testText   = "<SAMPLE TEXT>"
	testFlavor = "At least I have Chicken"
	testImage  = "<SAMPLE URL>"
)

func TestDataStructure(t *testing.T) {
	testCard := NewCard(testName, testText, testFlavor, testImage)

	assert.Equal(t, testCard.Name(), testName)
	assert.Equal(t, testCard.Text(), testText)
	assert.Equal(t, testCard.Flavor(), testFlavor)
	assert.Equal(t, testCard.Image(), testImage)
}
