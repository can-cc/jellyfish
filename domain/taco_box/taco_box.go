package taco_box

import "time"

type TacoBoxID string

type TacoBox struct {
	ID        TacoBoxID `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatorID string    `json:"creatorID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (t *TacoBox) IsNew() bool {
	return t.ID == ""
}
