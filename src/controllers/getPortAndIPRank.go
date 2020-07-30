package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPortAndIPRank(c *gin.Context) {
	c.String(http.StatusOK, models.RankByMostCon())
}
