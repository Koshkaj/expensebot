package server

import (
	"fmt"
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

func InitServer() *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	util.LoadDotenv()
	serverConfig := &config.Config{
		DbType:    os.Getenv("DB_TYPE"),
		StoreType: os.Getenv("STORE_TYPE"),
		Server: config.ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
	botConfig := &config.GoogleProcessorConfig{
		Location:    os.Getenv("GCP_LOCATION"),
		ProjectID:   os.Getenv("GCP_PROJECT_ID"),
		ProcessorID: os.Getenv("GCP_PROCESSOR_ID"),
		Endpoint:    fmt.Sprintf("%s-documentai.googleapis.com:443", os.Getenv("GCP_LOCATION")),
	}
	gp := bot.NewGoogleProcessor(botConfig)
	database, err := db.CreateDatabase(serverConfig.DbType)
	if err != nil {
		log.Fatal(err)
	}
	store, err := store.CreateFileStore(serverConfig.StoreType)
	if err != nil {
		log.Fatal(err)
	}

	uploadService := service.NewUploadService(gp, database, store)

	server := NewServer(e, serverConfig, database)
	indexGroup := server.Group("/")
	{
		indexGroup.Add("GET", "documents/:id", handlers.HandleGetDocument(uploadService))
		indexGroup.Add("POST", "documents", handlers.HandleUploadDocument(uploadService))

		indexGroup.Add("GET", "ping", server.handlePing)
	}
	return server
}
