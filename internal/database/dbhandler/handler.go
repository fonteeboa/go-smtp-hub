package dbhandler

import (
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	"github.com/fonteeboa/go-smtp-hub/internal/database/mongodb"
	"github.com/fonteeboa/go-smtp-hub/internal/database/mysql"
	"github.com/fonteeboa/go-smtp-hub/internal/database/postgres"
	"github.com/fonteeboa/go-smtp-hub/internal/database/sqlite"
)

// GetConnection retrieves the database connection based on the DATABASE_TYPE environment variable.
//
// It returns the *gorm.DB and *mongo.Client connections and an error.
func GetConnection() (*gorm.DB, *mongo.Client, error) {
	dbType := os.Getenv("DATABASE_TYPE")
	if dbType == "" {
		fmt.Println("DATABASE_TYPE not set")
		return nil, nil, errors.New("database.type.not.specified")
	}

	var gormDB *gorm.DB
	var mongoClient *mongo.Client
	var err error

	switch dbType {
	case "postgres":
		gormDB, err = postgres.Connect()
		if err != nil {
			return nil, nil, err
		}
	case "mysql":
		gormDB, err = mysql.Connect()
		if err != nil {
			return nil, nil, err
		}
	case "sqlite":
		gormDB, err = sqlite.Connect()
		if err != nil {
			return nil, nil, err
		}
	case "mongodb":
		mongoClient, err = mongodb.Connect()
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("unsupported.database.type")
	}

	return gormDB, mongoClient, nil
}
