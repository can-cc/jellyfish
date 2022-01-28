package factory

import (
	"jellyfish/domain/taco"
	"jellyfish/domain/taco/command"
)

func NewTacoFromCreateCommand(command *command.CreateTacoCommand, userID string) *taco.Taco {
	t := &taco.Taco{
		CreatorId: userID,
		Content:   command.Content,
		Detail:    command.Detail,
		Status:    taco.Doing,
		BoxId:     command.BoxId,
		Type:      taco.Task,
		Order:     command.Order,
		Deadline:  command.Deadline,
	}
	return t
}
