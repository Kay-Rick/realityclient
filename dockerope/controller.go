package dockerope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"rc/param"
	"strings"
	"time"
)

func GetTimeStr(timestamp int64) (timestr string) {
	timeago := time.Unix(timestamp, 0)
	subTime := time.Now().Sub(timeago)
	if hours := int64(subTime.Hours()); hours != 0 {
		if hours >= 24 {
			if hours >= 24*30 {
				timestr = fmt.Sprintf("%d months ago", hours/(24*30))
			} else {
				timestr = fmt.Sprintf("%d days ago", hours/24)
			}
		} else {
			timestr = fmt.Sprintf("%d hours ago", hours)
		}
	} else if minutes := int64(subTime.Minutes()); minutes != 0 {
		timestr = fmt.Sprintf("%d minutes ago", minutes)
	} else {
		timestr = fmt.Sprintf("%d seconds ago", int64(subTime.Seconds()))
	}
	return
}

func GetImagesController(c *gin.Context) {
	imageList, err := GetImagesHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	images := make([]param.Image, len(imageList), len(imageList))
	for i, imagel := range imageList {
		images[i].ImageId = strings.Split(imagel.ID, ":")[1][:12] //获取imageid前12个字符
		subrepotags := strings.Split(imagel.RepoTags[0], ":")
		images[i].Repository = subrepotags[0]
		images[i].Tag = subrepotags[1]
		sizeDecimal, _ := decimal.NewFromFloat(float64(imagel.Size) / 1000000).Round(1).Float64()
		images[i].Size = fmt.Sprintf("%.1fMB", sizeDecimal) //转成1位字符串
		images[i].Created = GetTimeStr(imagel.Created)
	}
	c.JSON(http.StatusOK, images)
}

func RunContainerController(c *gin.Context) {
	//通过cli客户端对象去执行ContainerList(其实docker ps 不就是一个docker正在运行容器的一个list嘛)
	//containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	containerid := c.Query("containerid")
	if err := RunContainerHandler(containerid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "run container success",
	})
}

func CreateAndRunContainerController(c *gin.Context) {
	var p param.CARInParam
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println("parse param failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := CreateAndRunContainerHandler(p.ImageName, p.Cmd, p.Network, p.Tdy, p.Openstdin, p.ContainerName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "create and run container success",
	})
}

func GetAllContainersController(c *gin.Context) {
	containersl, err := GetAllContainersHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var containers []param.Container = make([]param.Container, len(containersl), len(containersl))
	for i, con := range containersl {
		containers[i].ContainerId = con.ID[:12]
		containers[i].Cmd = con.Command
		containers[i].Status = con.Status
		containers[i].Names = make([]string, len(con.Names), len(con.Names))
		for i2, name := range con.Names {
			containers[i].Names[i2] = name[1:]
		}
		containers[i].ImageName = con.Image
		containers[i].Created = GetTimeStr(con.Created)
		fmt.Println("created: %s", containers[i].Created)
		containers[i].Running = con.State == "running"
	}
	c.JSON(http.StatusOK, containers)
}

func StopContainerController(c *gin.Context) {
	containerid := c.Query("containerid")
	if err := StopContainerHandler(containerid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "stop container success",
	})
}

func RemoveContainerController(c *gin.Context) {
	containerid := c.Query("containerid")
	if err := RemoveContainerHandler(containerid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "remove container success",
	})
}

func RemoveImageController(c *gin.Context) {
	imageid := c.Query("imageid")
	if err := RemoveImageHandler(imageid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "remove image success",
	})
}

func ReceiveAndLoadImageController(c *gin.Context) {
	ifile, err := c.FormFile("imagefile")
	if err != nil {
		fmt.Printf("get file failed:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fh, err := ifile.Open()
	if err != nil {
		fmt.Println("imagefile open faild:%v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ReceiveAndLoadImageHandler(fh); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": "receive and load image success",
	})
	return
}
