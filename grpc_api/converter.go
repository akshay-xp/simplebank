package grpc_api

import (
	db "github.com/akshay-xp/simplebank/db/sqlc"
	"github.com/akshay-xp/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		CreatedAt:         timestamppb.New(user.PasswordChangedAt),
		PasswordChangedAt: timestamppb.New(user.CreatedAt),
	}
}
