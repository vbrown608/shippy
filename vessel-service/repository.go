package main

import (
	"context"

	pb "github.com/vbrown608/shippy/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *MongoRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection.FindOne(context.Background(), bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).Decode(vessel)
	return vessel, err
}

func (repo *MongoRepository) Create(vessel *pb.Vessel) error {
	_, err := repo.collection.InsertOne(context.Background(), vessel)
	return err
}
