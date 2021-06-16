package note

import (
	"errors"
	"fmt"
	"time"
)

type Service interface {
	GetAll() ([]Note, error)
	GetByUserLogin(userID string) ([]Note, error)
	GetByID(ID string) (Note, error)
	SaveNewNote(noteID string, userID string, input NoteInput) (Note, error)
	UpdateNote(ID string, dataUpdate NoteUpdateInput) (Note, error)
	DeleteNote(ID string) (interface{}, error)
	UnDeleteNote(ID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Note, error) {
	notes, err := s.repository.FindAll()

	if err != nil {
		return notes, err
	}

	// formatNotes := FormatNotes(notes)

	return notes, nil
}

func (s *service) GetByUserLogin(userID string) ([]Note, error) {
	notes, err := s.repository.FindAllByUser(userID)

	if err != nil {
		return notes, err
	}

	// formatNotes := FormatNotes(notes)

	return notes, nil
}

func (s *service) GetByID(ID string) (Note, error) {
	note, err := s.repository.FIndByID(ID)

	if note.ID == "" || len(note.ID) <= 1 {
		errResponse := fmt.Sprintf("error note id %s not found", ID)
		return note, errors.New(errResponse)
	}

	if err != nil {
		return note, err
	}

	// formatNote := FormatNote(note)

	return note, nil
}

func (s *service) SaveNewNote(noteID string, userID string, input NoteInput) (Note, error) {

	var newNote = Note{
		ID:        noteID,
		UserID:    userID,
		Title:     input.Title,
		Body:      input.Body,
		Secret:    input.Secret,
		Type:      input.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	note, err := s.repository.Create(newNote)

	if err != nil {
		return note, err
	}

	// formatNote := FormatNote(note)

	return note, nil
}

func (s *service) UpdateNote(ID string, dataInput NoteUpdateInput) (Note, error) {
	var dataUpdate = map[string]interface{}{}

	note, err := s.repository.FIndByID(ID)

	if err != nil {
		return note, err
	}

	if note.ID == "" || len(note.ID) <= 1 {
		errResponse := fmt.Sprintf("error note id %s not found", ID)
		return note, errors.New(errResponse)
	}

	if dataInput.Body != "" || len(dataInput.Body) > 0 {
		dataUpdate["body"] = dataInput.Body
	}

	if dataInput.Title != "" || len(dataInput.Title) > 0 {
		dataUpdate["title"] = dataInput.Title
	}

	if dataInput.Secret != "" || len(dataInput.Secret) > 0 {
		dataUpdate["secret"] = dataInput.Secret
	}

	if dataInput.Type != "" || len(dataInput.Type) > 0 {
		dataUpdate["type"] = dataInput.Type
	}

	dataUpdate["updated_at"] = time.Now()

	noteUpdate, err := s.repository.Update(ID, dataUpdate)

	if err != nil {
		return noteUpdate, err
	}

	// formatNote := FormatNote(noteUpdate)

	return noteUpdate, nil
}

func (s *service) DeleteNote(ID string) (interface{}, error) {
	var dataDelete = map[string]interface{}{}

	note, err := s.repository.FIndByID(ID)

	if err != nil {
		return nil, err
	}

	if note.ID == "" || len(note.ID) <= 1 {
		errResponse := fmt.Sprintf("error note id %s not found", ID)
		return nil, errors.New(errResponse)
	}

	dataDelete["deleted"] = true

	deleteNote, err := s.repository.Update(ID, dataDelete)

	if err != nil {
		return deleteNote, err
	}

	return fmt.Sprintf("note id %s success deleted", ID), nil
}

func (s *service) UnDeleteNote(ID string) (interface{}, error) {
	var dataDelete = map[string]interface{}{}

	note, err := s.repository.FIndByID(ID)

	if err != nil {
		return nil, err
	}

	if note.ID == "" || len(note.ID) <= 1 {
		errResponse := fmt.Sprintf("error note id %s not found", ID)
		return nil, errors.New(errResponse)
	}

	dataDelete["deleted"] = false

	deleteNote, err := s.repository.Update(ID, dataDelete)

	if err != nil {
		return deleteNote, err
	}

	return fmt.Sprintf("note id %s success undeleted", ID), nil
}
