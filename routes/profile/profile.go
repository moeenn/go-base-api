package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileResponse struct {
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

func ProfileHandler(c *gin.Context) {
	userId, ok := c.MustGet("userId").(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	res := ProfileResponse{
		Message: "Hello world",
		UserId:  userId,
	}

	c.JSON(http.StatusOK, res)
}
