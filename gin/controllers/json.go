package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJSON(c *gin.Context) {
	data := new(interface{})

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
