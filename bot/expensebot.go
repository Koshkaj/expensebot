package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	documentai "cloud.google.com/go/documentai/apiv1"
	documentaipb "cloud.google.com/go/documentai/apiv1/documentaipb"
	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/types"
	"google.golang.org/api/option"
)

type GoogleProcessor struct {
	*config.GoogleProcessorConfig
	client *documentai.DocumentProcessorClient
}

func NewGoogleProcessor(cfg *config.GoogleProcessorConfig) *GoogleProcessor {
	ctx := context.Background()
	client, err := documentai.NewDocumentProcessorClient(ctx, option.WithEndpoint(cfg.Endpoint))
	if err != nil {
		log.Fatal(err)
	}
	return &GoogleProcessor{
		GoogleProcessorConfig: cfg,
		client:                client,
	}
}

func (gp *GoogleProcessor) getProcessorName() string {
	return fmt.Sprintf("projects/%s/locations/%s/processors/%s", gp.ProjectID, gp.Location, gp.ProcessorID)
}

func (gp *GoogleProcessor) createProcessRequest(ctx context.Context, file *types.File) (*documentaipb.ProcessRequest, error) {
	r := &documentaipb.ProcessRequest{
		SkipHumanReview: true,
		Name:            gp.getProcessorName(),
		Source: &documentaipb.ProcessRequest_RawDocument{
			RawDocument: &documentaipb.RawDocument{
				MimeType: file.MimeType,
				Content:  file.Data,
			},
		},
	}

	return r, nil
}

func (gp *GoogleProcessor) Process(fileDocument *types.File) error {
	ctx := context.Background()

	client, err := documentai.NewDocumentProcessorClient(ctx)
	defer client.Close()
	if err != nil {
		return err
	}
	request, err := gp.createProcessRequest(ctx, fileDocument)
	if err != nil {
		return err
	}
	response, err := client.ProcessDocument(ctx, request)
	if err != nil {
		log.Print("Error processing document: ", err)
		return err
	}
	document := response.GetDocument()
	jsoned, err := json.Marshal(document)
	if err != nil {
		log.Print("Error marshaling document: ", err)
		return err
	}
	fmt.Println(jsoned)

	return nil
}
