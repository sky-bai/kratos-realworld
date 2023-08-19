package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "kratos-realworld/api/realworld/v1"
)

// dto --> do 数据传输对象 转成领域对象 传到 biz 层

func (r *RealWorldService) Login(context.Context, *v1.LoginRequest) (*v1.UserReply, error) {

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "blj",
		},
	}, nil
}

func (r *RealWorldService) Register(context.Context, *v1.RegisterRequest) (*v1.UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
