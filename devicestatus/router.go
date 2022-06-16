package devicestatus

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/devicestatus", DeviceStatusHandler)
	e.GET("/devicedspstatus", DeviceDSPStatusHandler)
	e.GET("/devicefpgastatus", DeviceFPGAStatusHandler)
}
