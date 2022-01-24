package main

import (
	"context"
	"sync"

	// import the generated protobuf code
	pb "github.com/bulidiriba/shippy-service-consignment/proto/consignment"
)

const (
	port = "50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository -- Dummy repository, this simulates the use of a datastore of some kind.
// We'll replace this with a real implementation later on.
type Repository struct {
	mu          sync.RWMutex
	consignment []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// Service should implement all of the methods to satisfy the service
// we defined our protobuf definition. You can check the interface in the generated code itsefl
// for the exact method signatures etc to give you a better idea.
type service struct {
	repo repository
}

// createConsignment - we created just one method on our service,
// which is a created method, which takes a context and a requested as an argument,
// these are handled by the gRPC server
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// save our consignment
	consignmnet, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
}
func main() {

}
