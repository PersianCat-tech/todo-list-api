package service

import (
	"net/http"
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做， 1是已做
}

type ShowTaskService struct {
}

// 新增一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := http.StatusOK
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0, //默认未完成
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = http.StatusInternalServerError
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// 展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := http.StatusOK
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = http.StatusInternalServerError
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}
