package taco

import (
	"strings"
	"time"
)

type Status string

const (
	Doing Status = "Doing"
	Done  Status = "Done"
)

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

type ListTacoFilter struct {
	Statues []Status
	BoxId   *string
	Type    *Type
}

const (
	Task Type = "Task"
)

type Taco struct {
	Id        string     `json:"id"`
	CreatorId string     `json:"creatorId"`
	Content   string     `json:"content"`
	Detail    *string    `json:"detail"`
	Status    Status     `json:"status"`
	Type      Type       `json:"type"`
	Deadline  *time.Time `json:"deadline"`
	BoxId     *string    `json:"boxId"`
	Order     float64    `json:"order"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdateAt  *time.Time `json:"updatedAt"`
}

func (t *Taco) IsNew() bool {
	return t.Id == ""
}

func IndexOfSlice(slice []Taco, itemId string) int {
	for i := range slice {
		if slice[i].Id == itemId {
			return i
		}
	}
	return -1
}

func SliceRemove(tacos []Taco, index int) {
	tacos[index] = tacos[len(tacos)-1]
	tacos[len(tacos)-1] = Taco{}
	tacos = tacos[:len(tacos)-1]
}
