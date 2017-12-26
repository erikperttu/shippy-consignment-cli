package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"context"
	pb "github.com/erikperttu/shippy-consignment-service/proto/consignment"
	microClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	cmd.Init()

	// Create a new client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microClient.DefaultClient)

	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("not enough arguments, expecting a file and a token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Failed to create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Failed to list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}

}
