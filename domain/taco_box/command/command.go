package command

type CreateTacoBoxCommand struct {
	CreatorId string
	Name      string `json:"name"`
}

type UpdateTacoCommand struct {
	Name            string `json:"name"`
	OperationUserID string
	TacoBoxID       string
}
