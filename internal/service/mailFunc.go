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

// ExecSendMail sends an email using the provided SMTP server address, authentication, sender and recipient addresses, and message.
//
// Parameters:
// - addr: the address of the SMTP server to use.
// - auth: the authentication mechanism to use when connecting to the SMTP server.
// - from: the sender's email address.
// - to: a list of recipient email addresses.
// - msg: the message to be sent.
//
// Returns:
// - bool: a boolean value indicating whether the email was sent successfully or not.
// - error: an error object if there was an error sending the email.
func ExecSendMail(addr string, auth smtp.Auth, from string, to []string, msg []byte) (bool, error) {
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SaveLog saves the SMTP configuration to the database.
//
// It takes a domain.SMTPConfig parameter and returns an error.
func SaveConfigDb(config domain.SMTPConfig) error {

	dbConfig := CheckEnvironment()
	if !dbConfig {
		return errors.New("no.database.config")
	}

	gormDB, mongoClient, err := dbhandler.GetConnection()

	if err != nil {
		return err
	}

	if gormDB == nil && mongoClient == nil {
		return errors.New("no.database.connection")
	}

	if gormDB != nil {
		errGorm := dbhandler.InsertSMTPConfig(gormDB, config)
		if errGorm != nil {
			return errGorm
		}
	}

	if mongoClient != nil {
		errMongo := mongodb.InsertSMTPConfig(mongoClient, config)
		if errMongo != nil {
			return errMongo
		}
	}

	return nil

}

// NewSMTPConfig creates a new SMTPConfig object.
//
// Parameters:
// - Host: the host of the SMTP server (string).
// - Port: the port of the SMTP server (int).
// - Email: the email address to be used for authentication (string).
// - Password: the password to be used for authentication (string).
//
// Return type:
// - SMTPConfig: the newly created SMTPConfig object.
func newSMTPConfig(Host string, Port int, Email string, Password string) domain.SMTPConfig {
	return domain.SMTPConfig{
		Host:     Host,
		Port:     Port,
		Email:    Email,
		Password: Password,
	}
}

// SaveConfig saves the SMTP configuration to the database.
//
// Parameters:
// - Host: the host of the SMTP server.
// - Port: the port number of the SMTP server.
// - Email: the email address to use for authentication.
// - Password: the password to use for authentication.
//
// Returns:
// - error: an error if there was a problem saving the configuration.
func SaveConfig(Host string, Port int, Email string, Password string) error {

	config := newSMTPConfig(Host, Port, Email, Password)

	err := SaveConfigDb(config)
	if err != nil {

		return err
	}

	return nil

}
