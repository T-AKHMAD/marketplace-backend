package grpc

import (
	"context"

	"github.com/T-AKHMAD/marketplace-backend/internal/grpc/pb"
	"github.com/T-AKHMAD/marketplace-backend/internal/service"
)

type Server struct {
	pb.UnimplementedListingServiceServer
	svc *service.ListingService
}

func NewServer(svc *service.ListingService) *Server {
	return &Server{svc: svc}
}

func (s *Server) GetListing(ctx context.Context, req *pb.GetListingRequest) (*pb.GetListingResponse, error) {
	l, err := s.svc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetListingResponse{
		Id: l.ID,
		Title: l.Title,
		Description: l.Description,
		Price: l.Price,
		CreatedAt: l.CreatedAt.String(),
	}, nil
}