package gRPC

import (
	"context"
	pb "github.com/TekCatZ/imgour-authen-service/generated/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRpcServer struct {
	pb.UnimplementedAuthServiceServer
}

func (gRpcServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (gRpcServer) GetUserProfile(context.Context, *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
