package usecase

type ResponseProfile struct {
	FistName string                   `json:"first_name"`
	LastName string                   `json:"last_name"`
	Accounts []ResponseProfileAccount `json:"accounts"`
}

type ResponseProfileAccount struct {
	AccountType  string `json:"account_type"`
	AccountValue string `json:"account_value"`
}
