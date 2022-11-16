package note_v1

import (
	"context"
	"fmt"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Println("Deleting Note")
	fmt.Println("Note ID: ", req.GetId())

	return &desc.Empty{}, nil
}
