package dockerope

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"time"
)

var ctx = context.Background()
var cli, err1 = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func GetImagesHandler() ([]types.ImageSummary, error) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		fmt.Println("get local docker images failed: %v", err.Error())
		return nil, err
	}
	return images, nil
}

func GetAllContainersHandler() ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println("get all contaniers failed: %v", err.Error())
		return nil, err
	}
	return containers, nil
}

func CreateAndRunContainerHandler(imageName string, cmd []string, network string, tdy bool, openstdin bool, containername string) error {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:     imageName, //镜像名称
		Tty:       tdy,       //docker run命令中的-t选项
		OpenStdin: openstdin, //docker run命令中的-i选项
		Cmd:       cmd,       //docker 容器中执行的命令
	}, &container.HostConfig{
		NetworkMode: "host", //host模式
	}, nil, nil, containername)
	if err != nil {
		fmt.Printf("create container from image %s failed: %v\n", imageName, err.Error())
		return err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("start container %.12s faild: %v\n", resp.ID, err.Error())
		return err
	}
	return nil
}

func RunContainerHandler(containerid string) error {
	if err := cli.ContainerStart(ctx, containerid, types.ContainerStartOptions{}); err != nil {
		fmt.Printf("start container %.12s faild: %v\n", containerid, err.Error())
		return err
	}
	return nil
}

func StopContainerHandler(containerid string) error {
	timeout := time.Second * 1
	if err := cli.ContainerStop(ctx, containerid, &timeout); err != nil {
		fmt.Printf("stop container %.12s faild: %v\n", containerid, err.Error())
		return err
	}
	return nil
}

func RemoveContainerHandler(containerid string) error {
	if err := cli.ContainerRemove(ctx, containerid, types.ContainerRemoveOptions{RemoveVolumes: true, Force: true}); err != nil {
		fmt.Printf("remove container %.12s failed: %v\n", containerid, err.Error())
		return err
	}
	return nil
}

func RemoveImageHandler(imageid string) error {
	if _, err := cli.ImageRemove(ctx, imageid, types.ImageRemoveOptions{Force: true}); err != nil {
		fmt.Printf("remove image %.12s failed: %v\n", imageid, err.Error())
		return err
	}
	return nil
}

func ReceiveAndLoadImageHandler(input io.Reader) error {
	response, err := cli.ImageLoad(ctx, input, true)
	if err != nil {
		fmt.Println("image load err. %v", err.Error())
		return err
	}
	defer response.Body.Close()
	return nil
}
