package interfaces

type VerifToken struct {
	Username					string 		`json:"username"`
	Email           			string  	`json:"email"`
	GoogleID					string 		`json:"google_id"`
	Role						string		`json:"role"`
	ActivationCode				string 		`json:"email_verification_token"`
	ResetPasswordToken			string		`json:"reset_password_token"`
}