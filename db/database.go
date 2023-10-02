package db

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/types"
)

type Databaser interface {
	Get(uuid.UUID) (types.Document, error)
	Create(*types.Document) error
}

func CreateDatabase(dbType string) (Databaser, error) {
	switch dbType {
	case "memory":
		return NewMemoryDatabase(), nil
	case "mongo":
		mongoConf := &config.MongoConfig{
			URI:        os.Getenv("MONGO_URI"),
			DB_NAME:    os.Getenv("MONGO_DB_NAME"),
			COLLECTION: os.Getenv("MONGO_COLLECTION"),
		}
		return InitMongo(context.Background(), mongoConf), nil
	default:
		return nil, fmt.Errorf("invalid  type")

	}
}
