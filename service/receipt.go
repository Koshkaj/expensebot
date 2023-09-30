package service

import (
	"github.com/koshkaj/expensebot/bot"
	"github.com/koshkaj/expensebot/db"
	"github.com/koshkaj/expensebot/store"
)

// You could have an interface for services but in this case I implemented without it
type UploadService struct {
	GoogleProcessor *bot.GoogleProcessor
	DB              db.Databaser
	Store           store.Storer
}

func NewUploadService(gp *bot.GoogleProcessor, db db.Databaser, store store.Storer) *UploadService {
	return &UploadService{
		GoogleProcessor: gp,
		DB:              db,
		Store:           store,
	}
}
