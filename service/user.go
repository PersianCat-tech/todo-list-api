package service

import (
	"net/http"
	"todo_list/model"
	"todo_list/pkg/utils"
	"todo_list/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"pass_word" json:"pass_word" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)

	if count == 1 {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "have registerd",
		}
	}

	user.UserName = service.UserName
	//加密
	if err := user.SetPassWord(service.PassWord); err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "create user failed",
		}
	}

	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "registe success",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//先查询数据库判断数据库中是否有这个人
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) { //判断错误类型是否是"记录未找到"
			return serializer.Response{
				Status: http.StatusBadRequest,
				Msg:    "user not exist, please register first",
			}
		}

		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}

	if user.CheckPassword(service.PassWord) == false {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "error password",
		}
	}

	//发一个token，为了其他功能需要身份验证所给前端存储的
	//创建备忘录这个功能需要token，不然都不知道是谁创建的
	token, err := utils.GenerateToken(user.ID, service.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "Token 签发错误",
		}
	}

	return serializer.Response{
		Status: http.StatusOK,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "login success",
	}
}
