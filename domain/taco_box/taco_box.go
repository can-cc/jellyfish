package taco_box

import "time"

type TacoBoxID string

type TacoBox struct {
	ID        TacoBoxID `json:"id"`
	Name      string    `json:"name"`
	CreatorID string    `json:"creatorID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *TacoBox) IsNew() bool {
	return t.ID == ""
}

func PointerBoxIdIfEmptyStr(str string) *TacoBoxID {
	if str == "" {
		return nil
	}
	id := TacoBoxID(str)
	return &id
}
