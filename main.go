package main

import (
	"rc/devicestatus"
	"rc/dockerope"
	"rc/healthcheck"
)

func main() {
	// 绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	Include(healthcheck.Routers, dockerope.Routers, devicestatus.Routers)
	r := Init()
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8001")

}
