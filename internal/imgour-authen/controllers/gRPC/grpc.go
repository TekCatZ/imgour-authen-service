package gRPC

import (
	"context"
	"fmt"
	pb "github.com/TekCatZ/imgour-authen-service/generated/auth"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func (authServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (authServer) GetUserProfile(context.Context, *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}

func Setup(config configs.GrpcConfig) (*grpc.Server, *net.Listener) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatalf("Setup gRPC, failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &authServer{})

	return s, &lis
}
