package note_v1

import (
	"context"
	"fmt"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("Creating Note")
	fmt.Println("Title: ", req.GetTitle())
	fmt.Println("Text: ", req.GetText())
	fmt.Println("Author: ", req.GetAuthor())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
