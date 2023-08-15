package interfaces

type APIKey struct {
	Label			string	`json:"label"`
	ApiKey			string	`json:"api_key"`
	ActivationToken	int		`json:"activation_token"`
}