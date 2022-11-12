package note_v1

import desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"

type Note struct {
	desc.UnimplementedNoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}
