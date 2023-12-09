package dbhandler

import (
	"github.com/fonteeboa/go-smtp-hub/pkg/domain"

	"gorm.io/gorm"
)

func InsertSMTPConfig(db *gorm.DB, smtpConfig domain.SMTPConfig) error {
	return db.Create(&smtpConfig).Error
}

func GetDataGorm(db *gorm.DB) ([]*domain.SMTPConfig, error) {
	var configs []*domain.SMTPConfig
	err := db.Find(&configs).Error
	if err != nil {
		return nil, err
	}
	return configs, nil
}
