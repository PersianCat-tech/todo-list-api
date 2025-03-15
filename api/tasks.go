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
	log.Println("CreateTask")
	var createTask service.CreateTaskService
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	claims, err := utils.ParseToken(token)
	if err != nil || claims== nil {
		log.Fatal(err, "claims nil")
	}
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}

	res := createTask.Create(claims.Id)
	c.JSON(http.StatusOK, res)

}
