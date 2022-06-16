package healthcheck

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
)

func PingHandler(c *gin.Context) {
	//返回Pong
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

type ServiceInfo struct {
	Ip     string `json:"ip"`
	Image  string `json:"image"`
	Cmd    string `json:"cmd"`
	Weight int    `json:"weight"`
}

func StartContainerHandler(c *gin.Context) {
	var serviceInfo ServiceInfo
	c.BindJSON(&serviceInfo)
	fmt.Println("serviceInfo:")
	fmt.Println(serviceInfo)
	command := "config/run.sh" + " " + serviceInfo.Image + " " + serviceInfo.Cmd + " " + serviceInfo.Ip + " " + strconv.Itoa(serviceInfo.Weight)
	cmd := exec.Command("/bin/bash", "config/run.sh", serviceInfo.Image, serviceInfo.Cmd, serviceInfo.Ip, strconv.Itoa(serviceInfo.Weight))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s\n", command, err.Error())
		c.JSON(http.StatusAccepted, gin.H{
			"message": "exec " + command + "failed",
		})
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s\n", command, string(output))
	c.JSON(http.StatusOK, gin.H{
		"message": "exec " + command + "success",
	})
}
