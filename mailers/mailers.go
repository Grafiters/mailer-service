package mailers

import (
	"os"
	"log"
	"bytes"
	"strconv"
	"text/template"
	"mailer/interfaces"

	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
)

func SendMail(record interfaces.Record, msg interface{}){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	senderEmail := os.Getenv("SMTP_SENDER_EMAIL")
	senderName := os.Getenv("SMTP_SENDER_NAME")

	templateName := record.Tag+"_template.html"

	var tmpl *template.Template

	tmpl, err = template.ParseFiles("templates/"+templateName)
	if err != nil {
		log.Printf("Failed to load %v template: %v\n", record.Tag, err)
		return
	}

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Printf("Failed to convert SMTP port to integer: %v\n", err)
		return
	}

	mailer := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	m := mail.NewMessage()

	if record.Email == "" || record.Subject == ""{
		log.Printf("Invalid record: %+v\n", record)
		return
	}

	m.SetHeader("From", senderName+" <"+senderEmail+">")
	m.SetHeader("To", record.Email)
	m.SetHeader("Subject", record.Subject)

	body := &bytes.Buffer{}
	if err := tmpl.Execute(body, msg); err != nil {
		log.Printf("Failed to execute template: %v\n", err)
		return
	}
	
	m.SetBody("text/html", body.String())
	if err := mailer.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v\n", err)
	}else{
		log.Printf("Sent email Successfully")
	}
}