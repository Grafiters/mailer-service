package handlers

import (
	"log"

	"mailer/token"
	"mailer/interfaces"
)

func ClaimAllRecordData(msg string, tag string) (interface{}, error) {
	switch tag {
		case "login":
			var interfaceData interfaces.Login
			record := token.DecodeToken(msg, &interfaceData)
			if record != nil {
				log.Println(record)
				return nil, record
			}

			return interfaceData, nil
		case "send_code":
			var interfaceData interfaces.VerifToken
			record := token.DecodeToken(msg, &interfaceData)
			if record != nil {
				log.Println(record)
				return nil, record
			}

			return interfaceData, nil
		case "reset_password":
			var interfaceData interfaces.VerifToken
			record := token.DecodeToken(msg, &interfaceData)
			if record != nil {
				log.Println(record)
				return nil, record
			}

			return interfaceData, nil
		case "api_key":
			var interfaceData interfaces.APIKey
			record := token.DecodeToken(msg, &interfaceData)
			if record != nil {
				log.Println(record)
				return nil, record
			}

			return interfaceData, nil
		case "notification":
			var interfaceData interfaces.Notification
			record := token.DecodeToken(msg, &interfaceData)
			if record != nil {
				log.Println(record)
				return nil, record
			}

			return interfaceData, nil
		default:
			log.Println("Tag not recognized")
			
			var interfaceData interfaces.Record
			return interfaceData, nil
	}
}