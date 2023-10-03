package server

import (
	"log"
	"os"

	"github.com/koshkaj/expensebot/bot"
	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/db"
	"github.com/koshkaj/expensebot/handlers"
	"github.com/koshkaj/expensebot/service"
	"github.com/koshkaj/expensebot/store"
	"github.com/koshkaj/expensebot/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*config.Config
	*echo.Echo
	db db.Databaser
}

func (s *Server) handlePing(c echo.Context) error {
	return c.String(200, "pong")
}

func NewServer(e *echo.Echo, c *config.Config, m db.Databaser) *Server {
	return &Server{Echo: e, Config: c, db: m}
}

func InitServer() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	util.LoadDotenv()
	serverConfig := &config.Config{
		DbType:        os.Getenv("DB_TYPE"),
		StoreType:     os.Getenv("STORE_TYPE"),
		ProcessorType: os.Getenv("PROCESSOR_TYPE"),
		Server: config.ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
	gp, err := bot.CreateProcessor(serverConfig.ProcessorType)
	if err != nil {
		log.Fatalf("error creating a processor: %v", err)
	}
	database, err := db.CreateDatabase(serverConfig.DbType)
	// defer database.close() if we have such method
	if err != nil {
		log.Fatalf("error creating a database: %v", err)
	}
	store, err := store.CreateFileStore(serverConfig.StoreType)
	// defer store.close() if we have such method
	if err != nil {
		log.Fatalf("error creating a store: %v", err)
	}
	if serverConfig.StoreType == "local" {
		// Create a directory to store files
		directoryName := os.Getenv("STORE_DIRECTORY_NAME")
		os.Mkdir(directoryName, 0755)
	}

	uploadService := service.NewUploadService(gp, database, store)

	server := NewServer(e, serverConfig, database)
	indexGroup := server.Group("/")
	{
		indexGroup.Add("GET", "documents/:id", handlers.HandleGetDocument(uploadService))
		indexGroup.Add("POST", "documents", handlers.HandleUploadDocument(uploadService))

		indexGroup.Add("GET", "ping", server.handlePing)
	}
	return server.Start(serverConfig.Server.Port)
}
