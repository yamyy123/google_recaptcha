package routes

import (
	"recaptcha/service"

	"github.com/gin-gonic/gin"
)

func Newroute(router *gin.Engine) {
	router.Static("/", ".")
	router.POST("/submit", service.SubmitHandler)
}
