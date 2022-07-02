package gRPC

import (
	"context"
	"fmt"
	pb "github.com/TekCatZ/imgour-authen-service/generated/auth"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/configs"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/models"
	"github.com/TekCatZ/imgour-authen-service/internal/imgour-authen/entities/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func (authServer) GetRoles(_ context.Context, req *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	out := new(pb.GetRolesResponse)

	uid := req.GetUid()
	if uid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "UID is required")
	}

	user, err := repositories.GetUser(uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting user: %v", err)
	}

	resultRoles := parseRole(user)

	out.Roles = resultRoles
	return out, nil
}
func (authServer) GetUserProfile(_ context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	out := new(pb.GetUserProfileResponse)

	uid := req.GetUid()
	if uid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "UID is required")
	}

	user, err := repositories.GetUser(uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting user: %v", err)
	}

	resultRoles := parseRole(user)

	out.Profile = &pb.UserProfile{
		Uid:   user.Uid,
		Email: user.Email,
		Name:  user.Name,
		Phone: user.PhoneNumber,
		Roles: resultRoles,
	}

	return out, nil
}

func parseRole(user *models.UserProfile) []pb.Role {
	var resultRoles []pb.Role
	for _, role := range user.Roles {
		resultRoles = append(resultRoles, pb.Role(role))
	}
	return resultRoles
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
