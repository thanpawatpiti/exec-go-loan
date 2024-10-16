package cmd

import "github.com/thanpawatpiti/exec-go-loan/pkg"

type InitCmd struct {
	Model pkg.ModelInterface
}

func NewInitCmd(model pkg.ModelInterface) pkg.CmdInterface {
	return &InitCmd{Model: model}
}
