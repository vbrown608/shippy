package main

import (
	"context"
	"log"

	pb "github.com/vbrown608/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRespository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (respository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := respository.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Failed to Find on consignments collection")
		return nil, err
	}
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		// var consignment *pb.Consignment
		consignment := &pb.Consignment{}
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
