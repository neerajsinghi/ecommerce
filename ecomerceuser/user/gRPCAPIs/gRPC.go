package grpcapis

import (
	"context"
	"ecommerceuser/model"
	"ecommerceuser/proto"
)

type UserMessengerServer struct {
	proto.UnimplementedUserMessengerServer
}

func NewGrpcServer() *UserMessengerServer {
	return &UserMessengerServer{}
}

func (s *UserMessengerServer) GetUserList(ctx context.Context, in *proto.GetUserListRequest) (*proto.GetUserListResponse, error) {
	db := model.GetDB()
	var users []model.User
	input := in.UserIds
	db.Where("id IN ?", input).Find(&users)
	var grpcUsers []*proto.User
	for _, user := range users {
		grpcUser := &proto.User{
			Id:      int32(user.ID),
			Name:    user.Name,
			Email:   user.Email,
			PhoneNo: user.PhoneNo,
		}
		grpcUsers = append(grpcUsers, grpcUser)
	}
	return &proto.GetUserListResponse{
		Users: grpcUsers,
	}, nil
}
