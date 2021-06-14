package helper

import "github.com/afistapratama12/projectBackend/note"

type NoteFormat struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Type   string `json:"type"`
	Secret string `json:"secret"`
}

func FormatNote(note note.Note) NoteFormat {
	return NoteFormat{
		ID:     note.ID,
		Title:  note.Title,
		Body:   note.Body,
		Type:   note.Type,
		Secret: note.Secret,
	}
}

func FormatNotes(notes []note.Note) []NoteFormat {
	var formatNotes []NoteFormat

	for _, note := range notes {
		formatNotes = append(formatNotes, FormatNote(note))
	}

	return formatNotes
}
