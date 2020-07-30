package controllers

import (
	"net/http"
	"util"

	"github.com/gin-gonic/gin"
)

func GetDatabases(c *gin.Context) {
	c.String(http.StatusOK, util.GetArbitraryDB())
}
