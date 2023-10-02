package interfaces

type Users struct {
	Email			string		`json:"email"`
	DeviceToken		[]string	`json:"device_token"`
}

type Notification struct {
	User			Users	`json:"users"`
	Title           string  `json:"title"`
	Message			string	`json:"message"`
}