package taco

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestParseStatues(t *testing.T) {
	assert.Equal(t, ParseStatues("Doing, Done"), []Status{"Doing", "Done"})
	assert.Equal(t, ParseStatues("Doing"), []Status{"Doing"})
}
