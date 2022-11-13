package main

import (
	"context"
	"google.golang.org/grpc"
	"log"

	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)

	// Creating Note
	resCreate, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm surprised!",
		Author: "Danila",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Note with ID %d has been created\n\n", resCreate.Id)

	// Getting Note
	resGet, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Note is Found")
	log.Printf("Title: %v", resGet.Title)
	log.Printf("Text: %v", resGet.Text)
	log.Printf("Author: %v \n\n", resGet.Author)

	// Getting List of Notes
	resGetList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		Ids: []int64{1, 2},
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Following Notes were received")
	for i := 0; i < len(resGetList.GetNotes()); i++ {
		log.Printf("Title: %v", resGetList.GetNotes()[i].Title)
		log.Printf("Text: %v", resGetList.GetNotes()[i].Text)
		log.Printf("Author: %v \n\n", resGetList.GetNotes()[i].Author)
	}

	// Updating Note
	resUpd, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:     1,
		Title:  "Updated Title",
		Text:   "Updated Text",
		Author: "Updated Author",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Note with ID %d was Updated", resUpd.Id)
	log.Printf("Status Code: %d \n\n", resUpd.Result)

	// Deleting Note
	resDel, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Note with ID %d was Deleted", resDel.Id)
	log.Printf("Status Code: %d \n\n", resDel.Result)
}
