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

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm surprised!",
		Author: "Danila",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(res.Id)

	res2, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(res2.Title)
	log.Println(res2.Text)
	log.Println(res2.Author)
}
