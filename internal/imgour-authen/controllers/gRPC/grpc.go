package gRPC

import (
	"context"
	pb "github.com/TekCatZ/imgour-authen-service/generated/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const CERTIFICATE = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhki...\n-----END PUBLIC KEY-----"

type gRpcServer struct {
	pb.UnimplementedAuthServiceServer
}

func (gRpcServer) GetRoles(ctx context.Context, req *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	out := new(pb.GetRolesResponse)

	uid := req.GetUid()
	if uid == "" {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}

	return out, nil
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (gRpcServer) GetUserProfile(context.Context, *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
