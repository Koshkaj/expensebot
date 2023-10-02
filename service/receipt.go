package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/bot"
	"github.com/koshkaj/expensebot/db"
	"github.com/koshkaj/expensebot/store"
	"github.com/koshkaj/expensebot/types"
)

// You could have an interface for services but in this case I implemented without it
type UploadService struct {
	googleProcessor *bot.GoogleProcessor
	db              db.Databaser
	store           store.Storer
}

func NewUploadService(gp *bot.GoogleProcessor, db db.Databaser, store store.Storer) *UploadService {
	return &UploadService{
		googleProcessor: gp,
		db:              db,
		store:           store,
	}
}

func (s *UploadService) GetDocument(ctx context.Context, id uuid.UUID) (types.Document, error) {
	return s.db.Get(id)
}

func (s *UploadService) CreateDocument(ctx context.Context, document *types.Document) error {
	return s.db.Create(document)
}

func (s *UploadService) Save(ctx context.Context, fileName string, file io.Reader) error {
	return s.store.Save(fileName, file)
}

func (s *UploadService) Get(ctx context.Context, fileName string) (io.Reader, error) {
	return s.store.Get(fileName)
}

func (s *UploadService) Process(ctx context.Context, file *types.File) ([]byte, error) {
	return s.googleProcessor.Process(file)
}
