package healthcheck

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/checkhealth/ping", PingHandler)
	e.POST("/startcontainer", StartContainerHandler)
}
