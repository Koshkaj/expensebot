package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/service"
	"github.com/koshkaj/expensebot/types"
	"github.com/koshkaj/expensebot/util"
	"github.com/labstack/echo/v4"
)

func HandleUploadDocument(svc *service.UploadService) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("receipt")
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid file")
		}
		src, err := file.Open()
		if err != nil {
			return c.String(http.StatusUnprocessableEntity, "failed to open file")
		}
		defer src.Close()
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, src); err != nil {
			return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("failed to copy file: %s", err.Error()))
		}
		mimeType := http.DetectContentType(buf.Bytes())
		if !util.IsValidMimeType(mimeType) {
			return c.String(http.StatusBadRequest, "invalid file type")
		}

		uuid := uuid.New().String()
		document := &types.Document{
			Id:       uuid,
			Filename: file.Filename,
			MimeType: mimeType,
		}
		fileName := fmt.Sprintf("%s.%s", uuid, util.GetFileExtension(mimeType))
		fileDocument := &types.File{
			Id:        uuid,
			Name:      fileName,
			Extension: util.GetFileExtension(mimeType),
			MimeType:  mimeType,
			Data:      make([]byte, buf.Len()),
		}
		copy(fileDocument.Data, buf.Bytes())
		ctx := context.Background()
		if err := svc.Save(ctx, fileName, buf); err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}
		jsoned, err := svc.Process(ctx, fileDocument)
		if err != nil || jsoned == nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}
		document.FileDataJSON = fmt.Sprintf("%s.json", uuid)

		if err := svc.Save(ctx, document.FileDataJSON, bytes.NewBuffer(jsoned)); err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}

		if err := svc.CreateDocument(ctx, document); err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}
		return c.JSON(200, document)
	}
}

func HandleGetDocument(svc *service.UploadService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		document, err := svc.GetDocument(ctx, uuid.MustParse(c.Param("id")))
		if err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}
		return c.JSON(200, document)
	}
}
