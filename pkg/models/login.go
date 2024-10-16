package models

import "github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

func (m *Init) Login(body usecase.RequestLogin) (*uint, error) {
	userId := uint(1)

	return &userId, nil
}
