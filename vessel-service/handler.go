package main

import (
	"context"

	pb "github.com/vbrown608/shippy/vessel-service/proto/vessel"
)

// Our grpc service handler
type handler struct {
	repo Repository
}

func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}
