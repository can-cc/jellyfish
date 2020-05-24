package command

type CreateTacoBoxCommand struct {
	CreatorId string
	Name      string `json:"name" validate:"required"`
}

type UpdateTacoCommand struct {
	Name            string `json:"name"`
	OperationUserID string
	TacoBoxID       string `validate:"required"`
}
