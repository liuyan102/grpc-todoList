package handler

import (
	"context"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/e"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// UserLogin 用户登录 token
func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (*service.UserDetailResponse, error) {
	var user repository.User
	resp := new(service.UserDetailResponse)
	resp.Code = e.Success
	err := user.GetUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}

// UserRegister 用户注册
func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest) (*service.UserDetailResponse, error) {
	var user repository.User
	resp := new(service.UserDetailResponse)
	resp.Code = e.Success
	err := user.CreateUser(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = repository.BuildUser(user)
	return resp, nil
}
