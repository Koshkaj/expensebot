package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	*config.MongoConfig
	Client *mongo.Client
	Db     *mongo.Database
}

func (m *MongoDatabase) Get(id uuid.UUID) (types.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Db.Collection(m.COLLECTION)

	filter := bson.M{"_id": id.String()}

	var result types.Document

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return types.Document{}, fmt.Errorf("document not found")
		}
		return types.Document{}, err
	}

	return result, nil
}

func (m *MongoDatabase) Create(doc *types.Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Db.Collection(m.COLLECTION)

	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	return nil
}

func NewMongoDatabase(client *mongo.Client, cfg *config.MongoConfig) Databaser {
	db := client.Database(cfg.DB_NAME)
	return &MongoDatabase{Client: client, MongoConfig: cfg, Db: db}
}

func InitMongo(ctx context.Context, cfg *config.MongoConfig) Databaser {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(cfg.URI).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient := NewMongoDatabase(client, cfg)
	return mongoClient
}
