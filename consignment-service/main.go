package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	micro "github.com/micro/go-micro"
	pb "github.com/vbrown608/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/vbrown608/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
