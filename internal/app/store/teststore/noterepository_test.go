package teststore_test

import (
	"context"
	"testing"

	"github.com/jMurad/notes/internal/app/model"
	"github.com/jMurad/notes/internal/app/store"
	"github.com/jMurad/notes/internal/app/store/teststore"
)

func TestNoteRepository_Create(t *testing.T) {
	s := teststore.New()

	n := model.TestNote(t)

	if err := s.Note().Create(context.Background(), n); err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}
	if n == nil {
		t.Errorf("Expected value not to be nil")
	}
}

func TestNoteRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	n1 := model.TestNote(t)
	n2 := model.TestNote(t)

	_, err := s.Note().GetNotes(context.Background(), n1.AuthorID)
	if err != store.ErrRecordNotFound {
		t.Errorf("Error message not equal:\nexpected: %v\nactual  : %v", store.ErrRecordNotFound, err)
	}

	s.Note().Create(context.Background(), n1)
	s.Note().Create(context.Background(), n2)

	u3, err := s.Note().GetNotes(context.Background(), n1.AuthorID)
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}

	if u3 == nil {
		t.Errorf("Expected value not to be nil")
	}
}
