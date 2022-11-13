package note_v1

import (
	"context"
	"fmt"
	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
	"strconv"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {

	idsStr := ""
	for i := 0; i < len(req.GetIds()); i++ {
		idsStr += strconv.FormatInt(req.GetIds()[i], 10)
		if i != len(req.GetIds())-1 {
			idsStr += ", "
		}
	}

	fmt.Println("Getting List of Notes")
	fmt.Println("Ids: ", idsStr)

	note1 := &desc.GetListNoteResponse_Result{
		Title:  "GetListNote1",
		Text:   "Text of GetListNote1",
		Author: "Danila",
	}

	note2 := &desc.GetListNoteResponse_Result{
		Title:  "GetListNote2",
		Text:   "Text of GetListNote2",
		Author: "Dan",
	}

	return &desc.GetListNoteResponse{
		Notes: []*desc.GetListNoteResponse_Result{note1, note2},
	}, nil
}
