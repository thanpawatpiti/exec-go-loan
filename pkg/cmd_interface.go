package pkg

import "github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

type CmdInterface interface {
	Login(body usecase.RequestLogin) (*usecase.ResponseLogin, error)
	Refresh(body usecase.RequestRefresh) (*usecase.ResponseRefresh, error)
}
