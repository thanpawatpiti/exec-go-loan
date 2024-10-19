package cmd

import (
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

	helpers "github.com/thanpawatpiti/exec-go-loan/pkg/helpers"
)

func (c *InitCmd) Login(body usecase.RequestLogin) (*usecase.ResponseLogin, error) {
	// Call the repository
	userId, err := c.Model.Login(body)
	if err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, refreshToken, err := helpers.GenerateToken(*userId)
	if err != nil {
		return nil, err
	}

	return &usecase.ResponseLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (c *InitCmd) Refresh(body usecase.RequestRefresh) (*usecase.ResponseRefresh, error) {
	// Validate the refresh token
	claims, err := helpers.ValidateToken(body.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Generate new access token
	userID := uint(claims["user_id"].(float64))
	newAccessToken, newRefreshToken, err := helpers.GenerateToken(userID)
	if err != nil {
		return nil, err
	}

	// Return new tokens
	return &usecase.ResponseRefresh{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
