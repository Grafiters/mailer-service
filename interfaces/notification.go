package interfaces

type Notification struct {
	DeviceToken		[]string	`json:"device_token"`
	Title           string  	`json:"title"`
	Message			string		`json:"message"`
	Channel			string		`json:"channel,omitempty"`
	Time			string		`json:"time,omitempty"`
	Event			string		`json:"event,omitempty"`
	ChannelName		string		`json:"channel_name,omitempty"`
}