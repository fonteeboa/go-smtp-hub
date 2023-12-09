package pkg

import (
	"github.com/fonteeboa/go-smtp-hub/internal/service"
)

// SendMail sends an email to the specified recipients.
//
// Parameters:
// - to: an array of email addresses to send the email to.
// - message: the content of the email message as a byte array.
//
// Returns:
// - a boolean indicating whether the email was successfully sent.
// - an error if there was a problem sending the email.
func SendMail(to []string, message []byte) (bool, error) {
	send, err := service.MountAndSendMail(to, message)
	if send {
		return send, nil
	} else {
		return send, err
	}
}

// SendMailCustom sends a custom email using the provided SMTP host, port, email credentials, and message details.
//
// Parameters:
// - smtpHost: the SMTP server host address.
// - smtpPort: the SMTP server port number.
// - email: the email address used for authentication.
// - password: the password associated with the email address.
// - from: the sender's email address.
// - to: a list of recipient email addresses.
// - message: the email message content as a byte array.
//
// Returns:
// - bool: a boolean value indicating whether the email was successfully sent.
// - error: an error object if there was an issue during the email sending process.
func SendMailCustom(smtpHost string, smtpPort int, email string, password string, from string, to []string, message []byte) (bool, error) {
	auth := service.SetupInitAuth(email, password, smtpHost)
	smtpAddr := service.BuildSMTPAddress(smtpHost, smtpPort)
	send, err := service.ExecSendMail(smtpAddr, auth, email, to, message)
	if err != nil {
		return send, err
	}
	return send, nil
}

// SaveSMTPConfig saves the SMTP configuration.
//
// Host is the SMTP server host.
// Port is the port number for the SMTP server.
// Email is the email address for authentication.
// Password is the password for authentication.
// Returns an error if there was a problem saving the configuration.
func SaveSMTPConfig(Host string, Port int, Email string, Password string) error {

	err := service.SaveConfig(Host, Port, Email, Password)

	if err != nil {
		return err
	}

	return nil
}
