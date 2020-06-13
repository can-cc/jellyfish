package factory

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
)

func NewTacoFromCreateCommand(command *command.CreateTacoCommand, userID string) *taco.Taco {
	t := &taco.Taco{
		CreatorId: userID,
		Content:   command.Content,
		Detail:    command.Detail,
		Status:    taco.Doing,
		BoxId:     command.BoxId,
		Type:      taco.Task,
		Deadline:  command.Deadline,
	}
	return t
}
