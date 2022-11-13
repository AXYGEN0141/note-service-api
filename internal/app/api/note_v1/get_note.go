package note_v1

import (
	"context"
	"fmt"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("Getting Note")
	fmt.Println("Note ID: ", req.GetId())

	return &desc.GetNoteResponse{
		Title:  "Test Title",
		Text:   "GetNoteHandler",
		Author: "Danila",
	}, nil
}
