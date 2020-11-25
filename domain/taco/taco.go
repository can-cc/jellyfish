package taco

import (
	"sort"
	"strings"
	"time"

	"github.com/fwchen/jellyfish/domain/taco_box"
)

type Status string

const (
	Doing Status = "Doing"
	Done  Status = "Done"
)

type TacoFilter struct {
	Important bool
	Scheduled bool
	BoxId     *taco_box.TacoBoxID // TODO rename boxId
	Statues   []Status
	Type      *Type
}

func ParseStatues(str string) (statues []Status) {
	if str == "" {
		return
	}
	ss := strings.Split(str, ",")
	for _, s := range ss {
		statues = append(statues, Status(strings.Trim(s, " ")))
	}
	return
}

type Type string

const (
	Task Type = "Task"
)

type Taco struct {
	Id          string     `json:"id"`
	CreatorId   string     `json:"creatorId"`
	Content     string     `json:"content"`
	Detail      *string    `json:"detail"`
	Status      Status     `json:"status"`
	Type        Type       `json:"type"`
	Deadline    *time.Time `json:"deadline"`
	BoxId       *string    `json:"boxId"`
	Order       float64    `json:"order"`
	IsImportant bool       `json:"isImportant"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdateAt    *time.Time `json:"updatedAt"`
}

func (t *Taco) IsNew() bool {
	return t.Id == ""
}

func IndexOfTacos(slice []Taco, itemId string) int {
	for i := range slice {
		if slice[i].Id == itemId {
			return i
		}
	}
	return -1
}

type ByOrder []Taco

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

func SortTacosByOrder(tacos []Taco) []Taco {
	sort.Sort(ByOrder(tacos))
	return tacos
}

func SliceRemove(tacos []Taco, index int) []Taco {
	return append(tacos[:index], tacos[index+1:]...)
}

func InsertInTacos(tacos []Taco, taco Taco, index int) []Taco {
	if len(tacos) == index { // nil or empty slice or after last element
		return append(tacos, taco)
	}
	tacos = append(tacos[:index+1], tacos[index:]...)
	tacos[index] = taco
	return tacos
}
