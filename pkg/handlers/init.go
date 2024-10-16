package handlers

import "github.com/thanpawatpiti/exec-go-loan/pkg"

type InitHandler struct {
	Cmd pkg.CmdInterface
}

func NewInitHandler(cmd pkg.CmdInterface) *InitHandler {
	return &InitHandler{Cmd: cmd}
}
