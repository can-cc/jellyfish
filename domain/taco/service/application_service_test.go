package service

import (
	"testing"

	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/stretchr/testify/assert"
)

func TestSortTacos(t *testing.T) {
	tacos := []taco.Taco{
		taco.Taco{
			Id: "1",
		},
		taco.Taco{
			Id: "2",
		},
		taco.Taco{
			Id: "3",
		},
	}
	tacos = SortTacos(tacos, 0, 1)
	assert.Equal(t, tacos[0].Id, "2")
	assert.Equal(t, tacos[1].Id, "1")
	assert.Equal(t, tacos[2].Id, "3")
}
