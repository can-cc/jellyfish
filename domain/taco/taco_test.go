package taco

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestParseStatues(t *testing.T) {
	assert.Equal(t, ParseStatues("Doing, Done"), []Status{"Doing", "Done"})
	assert.Equal(t, ParseStatues("Doing"), []Status{"Doing"})
}

func TestSortTacosByOrder(t *testing.T) {
	tacos := []Taco{
		{
			Id:    "1",
			Order: 20,
		},
		{
			Id:    "2",
			Order: 10,
		},
		{
			Id:    "3",
			Order: 30,
		},
	}
	SortTacosByOrder(tacos)
	assert.Equal(t, tacos[0].Id, "2")
	assert.Equal(t, tacos[1].Id, "1")
	assert.Equal(t, tacos[2].Id, "3")

}

func TestSliceRemove(t *testing.T) {
	tacos := []Taco{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
	}
	tacos = SliceRemove(tacos, 0)
	assert.Equal(t, len(tacos), 1)
	assert.Equal(t, tacos[0].Id, "2")

	tacos2 := []Taco{
		{
			Id: "1",
		},
		{
			Id: "2",
		},
	}
	tacos2 = SliceRemove(tacos2, 1)
	assert.Equal(t, len(tacos2), 1)
	assert.Equal(t, tacos2[0].Id, "1")
}

func TestInsertInTacos(t *testing.T) {
	tacos := []Taco{
		{
			Id: "1",
		},
		{
			Id: "3",
		},
	}
	tacos = InsertInTacos(tacos, Taco{Id: "2"}, 1)
	assert.Equal(t, len(tacos), 3)
	assert.Equal(t, tacos[0].Id, "1")
	assert.Equal(t, tacos[1].Id, "2")
	assert.Equal(t, tacos[2].Id, "3")

	tacos2 := []Taco{
		Taco{
			Id: "1",
		},
	}
	tacos2 = InsertInTacos(tacos2, Taco{Id: "2"}, 0)
	assert.Equal(t, len(tacos2), 2)
	assert.Equal(t, tacos2[0].Id, "2")
	assert.Equal(t, tacos2[1].Id, "1")

	tacos3 := []Taco{
		Taco{
			Id: "1",
		},
	}
	tacos3 = InsertInTacos(tacos3, Taco{Id: "2"}, 1)
	assert.Equal(t, len(tacos3), 2)
	assert.Equal(t, tacos3[0].Id, "1")
	assert.Equal(t, tacos3[1].Id, "2")
}
