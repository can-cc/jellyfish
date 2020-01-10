package factory

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
)

func NewTacoFromCommand(command *command.CreateTacoCommand, userID string) *taco.Taco {
	taco := &taco.Taco{
		CreatorID: userID,
		Content:   command.Content,
		Detail:    command.Detail,
		Status:    taco.Doing,
		Type:      taco.Task,
		Deadline:  command.Deadline,
	}
	return taco
}
