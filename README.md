# ReliabilityClient

#### 介绍
SDHV可靠性设计客户端 程序

#### 软件架构
各级目录/文件介绍

1.  devicestatus: 设备状态信息监测，包括ARM、DSP和FPGA
    1.  include：是58所提供的状态信息检测的C头文件
    2.  lib：58所提供的状态信息检测静态库
    3.  router.go: 状态检测路由，`/devicestatus，/devicedspstatus，/devicefpgastatus`分别是ARM、DSP和FPGA的状态检测路由
    4.  handler.go: 路由对应handler，使用CGO从GO调用C程序，[CGO使用链接]([第2章 CGO编程 · Go语言高级编程 (chai2010.cn)](https://chai2010.cn/advanced-go-programming-book/ch2-cgo/readme.html))
2.  dockerope: docker操作
    1.  controller.go: 各个针对docker操作的路由请求先由对应controller处理
    2.  router.go: 路由
    3.  handler.go: controller处理后交由handler处理
3.  healthcheck: 健康检查（该目录暂时用不到）
4.  param: 数据解析参数，包括请求接受时的入参和返回时的参数
5.  system: linux设置reliabilityclient开机自启需要的文件
    1.  rc.sh: 启动可执行程序的脚本文件
    2.  reliabilityclient.service：把该文件放到/lib/systemed/system目录下，执行`sudo systemctl enable reliablilityclient`和`sudo systemctl start reliabilityclient`
6.  main.go：主程序，通过自定义的Include函数注册了每个文件下的所有路由，这是一种设计模式
7.  routers.go：路由注册和初始化

#### 使用说明

​	编译过程：先执行`go mod tidy`下载依赖库，然后执行`go build`，生成rc程序，直接放到ARM板上`/home/ubuntu/dang/code/reliabilityclient`目录下，然后`sudo chmod a+x rc`

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

6.  https://gitee.com/gitee-stars/)
