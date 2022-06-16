package dockerope

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/runcontainer", RunContainerController) //参数containerid
	e.GET("/getimages", GetImagesController)       //参数结构体
	e.POST("/createandruncontainer", CreateAndRunContainerController)
	e.GET("/getallcontainers", GetAllContainersController)
	e.GET("/stopcontainer", StopContainerController)
	e.GET("/removecontainer", RemoveContainerController)
	e.GET("/removeimage", RemoveImageController)
	e.POST("/receiveandloadimage", ReceiveAndLoadImageController)
}
