package controller

import (
	"net/http"
	"web/internal/helpers/jwthelpers"

	"github.com/labstack/echo/v4"
)

type ProtectedResponse struct {
	Message  string `json:"message"`
	UserId   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

func ProtectedHandler(c echo.Context) error {
	user := jwthelpers.CurrentUser(c)

	return c.JSON(http.StatusOK, ProtectedResponse{
		Message:  "welcome to the protected route",
		UserId:   user.Id,
		UserRole: user.Role,
	})
}
