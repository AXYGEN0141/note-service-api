package note_v1

import (
	"context"
	"fmt"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("Updating Note")
	fmt.Println("Note ID: ", req.GetId())
	fmt.Println("New Title: ", req.GetTitle())
	fmt.Println("New Text: ", req.GetText())
	fmt.Println("New Author: ", req.GetAuthor())

	return &desc.UpdateNoteResponse{
		Id:     req.GetId(),
		Result: 0,
	}, nil
}
