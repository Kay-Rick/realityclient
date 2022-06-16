package main

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) {
	options = append(options, opts...)
}

//初始化
func Init() *gin.Engine {
	r := gin.Default()
	//每个模块下的路由进行注册
	for _, opt := range options {
		opt(r)
	}
	return r
}
