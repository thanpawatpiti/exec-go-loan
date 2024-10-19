package cmd

import "github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

func (c *InitCmd) GetProfile(userID uint) (*usecase.ResponseProfile, error) {
	res, err := c.Model.GetProfile(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
