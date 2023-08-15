package interfaces

type Login struct {
	Email           string  `json:"email"`
	Device			string	`json:"device,omitempty"`
	LoginTime		string	`json:"login_time,omitempty"`
}