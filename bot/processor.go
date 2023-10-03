package bot

import (
	"fmt"
	"os"

	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/types"
)

type Processor interface {
	Process(fileDocument *types.File) ([]byte, error)
}

func CreateProcessor(processorType string) (Processor, error) {
	switch processorType {
	case "google":
		cfg := &config.GoogleProcessorConfig{
			Location:        os.Getenv("GCP_LOCATION"),
			ProjectID:       os.Getenv("GCP_PROJECT_ID"),
			ProcessorID:     os.Getenv("GCP_PROCESSOR_ID"),
			CredentialsFile: os.Getenv("CREDENTIALS_FILE_PATH"),
			Endpoint:        fmt.Sprintf("%s-documentai.googleapis.com:443", os.Getenv("GCP_LOCATION")),
		}
		return NewGoogleProcessor(cfg), nil
	default:
		return nil, fmt.Errorf("invalid processor type")

	}
}
