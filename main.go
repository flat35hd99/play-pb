package main

import (
	pb "github.com/flat35hd99/play-pb/tutorialpb"
	"google.golang.org/protobuf/proto"

	"encoding/json"
	"fmt"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	// Serialize/Marshal to protocol buffer
	bytes, err := proto.Marshal(&p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", bytes)

	// Serialize/Marchal to json
	bytes, err = json.Marshal(&p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)
}
