package service

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"

	"github.com/fonteeboa/go-smtp-hub/pkg/domain"

	"github.com/fonteeboa/go-smtp-hub/internal/database/dbhandler"
	"github.com/fonteeboa/go-smtp-hub/internal/database/mongodb"
)

// getSmtpConfig retrieves the SMTP configuration from the database.
//
// It returns a slice of SMTPConfig pointers and an error.
func getSmtpConfig() ([]*domain.SMTPConfig, error) {
	var config []*domain.SMTPConfig
	gormDB, mongoClient, err := dbhandler.GetConnection()
	if err != nil {
		return config, err
	}

	if gormDB == nil && mongoClient == nil {
		return config, errors.New("no.database.connection")
	}

	if gormDB != nil {
		config, err = dbhandler.GetDataGorm(gormDB)
		if err != nil {
			return config, err
		}
	}

	if mongoClient != nil {
		config, err = mongodb.GetDataMongo(mongoClient)
		if err != nil {
			return config, err
		}
	}

	return config, nil

}

// SetupInitAuth returns an smtp.Auth implementation using plain authentication.
//
// It takes the following parameters:
// - email: the email address to use for authentication
// - password: the password to use for authentication
// - smtpHost: the SMTP server host
//
// It returns an smtp.Auth object.
func SetupInitAuth(email, password, smtpHost string) smtp.Auth {
	return smtp.PlainAuth("", email, password, smtpHost)
}

// BuildSMTPAddress returns the SMTP address string.
//
// It takes in the SMTP host and port as parameters and returns a formatted string
// in the format "host:port".
//
// Parameters:
// - smtpHost: the SMTP host as a string.
// - smtpPort: the SMTP port as an integer.
//
// Returns:
// - The SMTP address as a formatted string.
func BuildSMTPAddress(smtpHost string, smtpPort int) string {
	return fmt.Sprintf("%s:%d", smtpHost, smtpPort)
}

// checkEnvironment checks the environment and returns a boolean value.
//
// This function does not take any parameters.
// It returns a boolean value indicating whether the "DATABASE_TYPE" environment variable is set or not.
func CheckEnvironment() bool {
	insertDB := os.Getenv("DATABASE_TYPE")
	return insertDB != ""
}

// MountAndSendMail sends an email using the provided SMTP configuration.
//
// It takes in two parameters:
//   - to: a slice of strings representing the recipients of the email.
//   - message: a byte array containing the content of the email.
//
// It returns a boolean value indicating whether the email was sent successfully, and an error if any.
func MountAndSendMail(to []string, message []byte) (bool, error) {

	dbConfig := CheckEnvironment()

	if !dbConfig {
		return false, errors.New("no.database.config")
	}

	var smtpConfig *domain.SMTPConfig
	smtpConfigs, err := getSmtpConfig()
	if err != nil {
		return false, err
	}
	// Verificando se há dados retornados
	if len(smtpConfigs) > 0 {
		return false, errors.New("no.smtp.config")
	} else {
		smtpConfig = smtpConfigs[0]
	}

	auth := SetupInitAuth(smtpConfig.Email, smtpConfig.Password, smtpConfig.Host)
	smtpAddr := BuildSMTPAddress(smtpConfig.Host, smtpConfig.Port)
	send, err := ExecSendMail(smtpAddr, auth, smtpConfig.Email, to, message)
	if err != nil {
		return send, err
	}
	return send, nil
}

func ExecSendMail(addr string, auth smtp.Auth, from string, to []string, msg []byte) (bool, error) {
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}
