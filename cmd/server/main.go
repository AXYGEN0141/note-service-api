package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/AXYGEN0141/note-service-api/internal/app/api/note_v1"
	desc "github.com/AXYGEN0141/note-service-api/pkg/note_v1"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	fmt.Println("Server is running on port:", port)
	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
