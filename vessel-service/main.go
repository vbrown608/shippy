package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vbrown608/shippy/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// TODO: write to respository
	// vessels := []*pb.Vessel{
	// 	&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// }

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
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

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	h := &handler{repository}

	pb.RegisterVesselServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
