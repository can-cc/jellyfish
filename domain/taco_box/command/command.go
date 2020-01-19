package command

type CreateTacoBoxCommand struct {
	CreatorId string
	Name      string `json:"name"`
	Icon      string `json:"icon"`
}
