package handlers

import (
	"gin-quick-start/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags        API
// @Summary		Hello World
// @Description	Hello World
// @Success		200	{object}	models.Result[string]		"success"
// @Failure		403	{object}	models.Result[string]		"System error"
// @Failure		500	{object}	models.Result[string]		"System error"
// @Router			/v1/hello-word [GET]
func HelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, models.Ok("Hello World"))
}

// @Tags        API
// @Summary		User Login
// @Description	User Login
// @Success		200	{object}	models.Result[string]		"login success"
// @Failure		403	{object}	models.Result[string]		"System error"
// @Failure		500	{object}	models.Result[string]		"System error"
// @Router			/v1/login	[POST]
func UserLogin(c *gin.Context) {
	var userLogin models.UserLogin
	if err := c.ShouldBindBodyWithJSON(userLogin); err != nil {
		c.JSON(http.StatusForbidden, models.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.Ok(models.UserLoginResponseDTO{
		Token: "token",
	}))
}
