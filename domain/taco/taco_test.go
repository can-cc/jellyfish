package taco

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestParseStatues(t *testing.T) {
	assert.Equal(t, ParseStatues("Doing, Done"), []Status{"Doing", "Done"})
	assert.Equal(t, ParseStatues("Doing"), []Status{"Doing"})
}

func TestSliceRemove(t *testing.T) {
	tacos := []Taco{
		Taco{
			Id: "1",
		},
		Taco{
			Id: "2",
		},
		Taco{
			Id: "3",
		},
	}
	tacos = SliceRemove(tacos, 1)
	assert.Equal(t, len(tacos), 2)
	assert.Equal(t, tacos[0].Id, "1")
	assert.Equal(t, tacos[1].Id, "3")
}
