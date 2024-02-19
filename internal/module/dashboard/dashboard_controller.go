package dashboard

import (
	"net/http"
	"web/internal/helpers/jwthelpers"
	"web/internal/helpers/responses"

	"github.com/labstack/echo/v4"
)

type ProtectedResponse struct {
	Message  string `json:"message"`
	UserId   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type DashboardController struct {
	AuthMiddleware *echo.MiddlewareFunc
}

func (c *DashboardController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/dashboard", (*c.AuthMiddleware)(c.protectedHandler))
}

func (controller *DashboardController) protectedHandler(c echo.Context) error {
	user := jwthelpers.CurrentUser(c)

	return c.JSON(http.StatusOK, responses.NewOkResponse(ProtectedResponse{
		Message:  "welcome to the protected route",
		UserId:   user.Id,
		UserRole: user.Role,
	}))
}
