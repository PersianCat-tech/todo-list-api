package api

import (
	"log"
	"net/http"
	"strings"
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	claim, err := utils.ParseToken(token)
	if err != nil {
		log.Fatal("CreateTask: ", err)
	}
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := createTask.Create(claim.Id)
	c.JSON(http.StatusOK, res)
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService

	if err := c.ShouldBind(&showTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := showTask.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// 列表返回用户所有备忘录

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claim, err := utils.ParseToken(token)

	if err != nil {
		log.Fatal("ListTask", err)
	}

	if err := c.ShouldBind(&listTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := listTask.List(claim.Id)
	c.JSON(http.StatusOK, res)
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService

	if err := c.ShouldBind(&updateTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := updateTask.Update(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	claim, _ := utils.ParseToken(token)

	if err := c.ShouldBind(&searchTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := searchTask.Search(claim.Id)
	c.JSON(http.StatusOK, res)
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService

	if err := c.ShouldBind(&deleteTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := deleteTask.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
