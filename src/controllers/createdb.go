package controllers

import (
	"io/ioutil"
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDatabase(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	DBname := string(data)
	status := models.CreateDB(DBname)
	if status {
		c.String(http.StatusOK, "ok")
	} else {
		c.String(http.StatusOK, "err")
	}
}
