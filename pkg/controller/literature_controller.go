package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"webdict-server/pkg/entity"
	"webdict-server/pkg/service"
)

type LiteratureController struct {
	service *service.LiteratureService
}

func NewLiteratureController() *LiteratureController {
	return &LiteratureController{service: service.NewLiteratureService()}
}

func (controller LiteratureController) FindAll(c *gin.Context) {
	literatures := controller.service.FindAll(c)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(literatures), "data": literatures})
}

func (controller LiteratureController) FuzzyFind(c *gin.Context) {
	key := c.Param("key")
	literatures := controller.service.FuzzyFind(c, key)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": len(literatures), "data": literatures})
}

func (controller LiteratureController) Insert(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	literature := entity.Literature{}
	json.Unmarshal(bytes, &literature)
	if err != nil {
		panic(err)
	}
	isOk := controller.service.Insert(c, literature)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "insert", "count": 0, "data": isOk})
}

func (controller LiteratureController) Update(c *gin.Context) {
	body := c.Request.Body
	jsonBytes, err := ioutil.ReadAll(body)
	literature := entity.Literature{}
	json.Unmarshal(jsonBytes, &literature)
	if err != nil {
		panic(err)
	}
	isOk := controller.service.Update(c, literature)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "update", "count": 0, "data": isOk})
}
