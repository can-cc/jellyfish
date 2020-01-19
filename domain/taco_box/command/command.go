package command

type CreateTacoBoxCommand struct {
	CreatorId string
	Name      string `json:"name"`
	Icon      string `json:"icon"`
}

type UpdateTacoCommand struct {
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	OperationUserID string
	TacoBoxID       string
}
