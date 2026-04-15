package main

import (
	"log"
	"net"

	
	"github.com/T-AKHMAD/marketplace-backend/internal/db"
	"github.com/T-AKHMAD/marketplace-backend/internal/grpc/pb"
	"github.com/T-AKHMAD/marketplace-backend/internal/repository"
	"github.com/T-AKHMAD/marketplace-backend/internal/service"
	
	"google.golang.org/grpc/reflection"
	grpcpkg "google.golang.org/grpc"
	mygrpc "github.com/T-AKHMAD/marketplace-backend/internal/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pg := db.NewPostgres()
	repo := repository.NewListingPostgresRepository(pg)
	svc := service.NewListingService(repo)

	grpcServer := mygrpc.NewServer(svc)
	s := grpcpkg.NewServer()
	
	pb.RegisterListingServiceServer(s,grpcServer)
	reflection.Register(s)
	
	log.Println("gRPC server running on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
