package note_v1

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	var sb strings.Builder
	for i := 0; i < len(req.GetIds()); i++ {
		sb.WriteString(strconv.FormatInt(req.GetIds()[i], 10))
		if i != len(req.GetIds())-1 {
			sb.WriteString(", ")
		}
	}

	fmt.Println("Getting List of Notes")
	fmt.Println("Ids: ", sb)

	return &desc.GetListNoteResponse{
		Notes: []*desc.GetListNoteResponse_Result{
			{
				Title:  "GetListNote1",
				Text:   "Text of GetListNote1",
				Author: "Danila",
			},
			{
				Title:  "GetListNote2",
				Text:   "Text of GetListNote2",
				Author: "Dan",
			}},
	}, nil
}
