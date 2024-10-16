package cmd

import (
	"time"

	"github.com/thanpawatpiti/exec-go-loan/config"
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"

	"github.com/golang-jwt/jwt/v4"
)

func (c *InitCmd) Login(body usecase.RequestLogin) (*usecase.ResponseLogin, error) {
	// Call the repository
	userId, err := c.Model.Login(body)
	if err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, refreshToken, err := GenerateToken(*userId)
	if err != nil {
		return nil, err
	}

	return &usecase.ResponseLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GenerateToken creates an access token and a refresh token
func GenerateToken(userID uint) (string, string, error) {
	// Create access token
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // 15 minutes expiration
		"iss":     config.AppName,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAccessToken, err := accessToken.SignedString(config.JwtSecretKey)
	if err != nil {
		return "", "", err
	}

	// Create refresh token
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiration
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshToken.SignedString(config.JwtSecretKey)
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}
