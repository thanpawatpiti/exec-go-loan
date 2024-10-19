package pkg

import "github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

type ModelInterface interface {
	Login(body usecase.RequestLogin) (*uint, error)
	GetProfile(userID uint) (*usecase.ResponseProfile, error)
}
