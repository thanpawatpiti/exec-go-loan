package usecase

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
