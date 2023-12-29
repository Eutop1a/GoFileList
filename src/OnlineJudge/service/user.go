package service

import "OnlineJudge/serializer"

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"utop1a"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16" example:"utop1a123"`
}

func (service *UserService) Register() serializer.Response {

}

func (service *UserService) Login() serializer.Response {

}
