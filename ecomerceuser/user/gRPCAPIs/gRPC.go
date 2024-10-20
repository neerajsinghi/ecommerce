package grpcapis

import (
	"context"
	"ecommerceuser/model"
	grpc_api "ecommerceuser/proto"
)

type UserMessengerServer struct {
	grpc_api.UnimplementedUserMessengerServer
}

func NewGrpcServer() *UserMessengerServer {
	return &UserMessengerServer{}
}

func (s *UserMessengerServer) GetUserList(ctx context.Context, in *grpc_api.GetUserListRequest) (*grpc_api.GetUserListResponse, error) {
	db := model.GetDB()
	var users []model.User
	input := in.UserIds
	db.Where("id IN ?", input).Find(&users)
	var grpcUsers []*grpc_api.User
	for _, user := range users {
		grpcUser := &grpc_api.User{
			Id:      int32(user.ID),
			Name:    user.Name,
			Email:   user.Email,
			PhoneNo: user.PhoneNo,
		}
		grpcUsers = append(grpcUsers, grpcUser)
	}
	return &grpc_api.GetUserListResponse{
		Users: grpcUsers,
	}, nil
}
