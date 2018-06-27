package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite/models"
)

func (ah *AdminHandler) RemoveService(c *gin.Context) {
	removeService(c, ah.Logger)
}

func (ah *AdminHandler) ListServices(c *gin.Context) {
	var (
		userID, nodeID int64
		err            error
	)
	if userIDStr := c.Query("userID"); userIDStr != "" {
		if userID, err = strconv.ParseInt(userIDStr, 10, 64); err != nil {
			c.JSON(http.StatusBadRequest, models.NewErrorResp(err))
			return
		}
	}
	if nodeIDStr := c.Query("nodeID"); nodeIDStr != "" {
		if nodeID, err = strconv.ParseInt(nodeIDStr, 10, 64); err != nil {
			c.JSON(http.StatusBadRequest, models.NewErrorResp(err))
			return
		}
	}

	listServices(c, userID, nodeID, ah.Logger)
}
