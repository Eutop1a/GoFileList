package service

import (
	"OJReconstitution/model"
	"OJReconstitution/serializer"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "Username already exist",
		}
	}
	user.UserName = service.UserName
	// 密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "mysql operation failed",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "register success",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "do not have this user",
			}
		}
		// 如果不是用户不存在，而是其他不可抗拒的因素导致的错误
		return serializer.Response{
			Status: 500,
			Msg:    "mysql error ",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "mysql error ",
		}
	}
	token, err := utils.
}
