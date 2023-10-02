package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/types"
	"github.com/stretchr/testify/assert"
)

func TestMemoryDatabase_Create(t *testing.T) {
	db := NewMemoryDatabase()

	sampleDocumentID := uuid.New()
	sampleDocument := &types.Document{
		Id: sampleDocumentID.String(),
	}

	t.Run("Create Existing Document", func(t *testing.T) {
		err := db.Create(sampleDocument)
		assert.NoError(t, err)

		err = db.Create(sampleDocument)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "already exists")
	})

	t.Run("Create New Document", func(t *testing.T) {
		newDocumentID := uuid.New()
		newDocument := &types.Document{
			Id: newDocumentID.String(),
		}

		err := db.Create(newDocument)
		assert.NoError(t, err)

		retrievedDocument, err := db.Get(newDocumentID)
		assert.NoError(t, err)
		assert.Equal(t, *newDocument, retrievedDocument)
	})
}

func TestMemoryDatabase_Get(t *testing.T) {
	db := NewMemoryDatabase()

	sampleDocumentID := uuid.New()
	sampleDocument := &types.Document{
		Id: sampleDocumentID.String(),
	}

	t.Run("Get Existing Document", func(t *testing.T) {
		err := db.Create(sampleDocument)
		assert.NoError(t, err)

		retrievedDocument, err := db.Get(sampleDocumentID)
		assert.NoError(t, err)
		assert.Equal(t, *sampleDocument, retrievedDocument)
	})

	t.Run("Get Non-Existing Document", func(t *testing.T) {
		nonExistentID := uuid.New()
		_, err := db.Get(nonExistentID)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}
