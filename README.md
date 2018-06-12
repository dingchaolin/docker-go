# docker-go

## 1. 安装

#### 1.1 centos内核版本要求高于3.10
- uname -r  查看内核版本

#### 1.2 更新yum
- sudo yum update

#### 1.3 卸载旧版本（如果安装过）
- sudo yum remove docker  docker-common docker-selinux docker-engine


#### 1.4 安装需要的软件包， yum-util 提供yum-config-manager功能，另外两个是devicemapper驱动依赖的
- sudo yum install -y yum-utils device-mapper-persistent-data lvm2

#### 1.5 安装docker
- yum -y install docker-io

#### 1.6 启动docker后台服务
- service docker start


## 2. 一个go+docker的实践
- server.go
```
package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go + docker !")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server start on 8888!")
	http.ListenAndServe(":8888", nil)
}
```

- Dockerfile
```
# 得到最新的 golang docker 镜像
FROM  daocloud.io/library/golang:1.9-rc-alpine
# 在容器内部创建一个目录来存储我们的 web 应用，接着使它成为工作目录。
RUN mkdir -p /go/src/web-app
WORKDIR /go/src/web-app
# 复制 web-app 目录到容器中
COPY . /go/src/web-app
# 下载并安装第三方依赖到容器中
RUN go-wrapper download
RUN go-wrapper install
# 设置 PORT 环境变量
ENV PORT 8888
# 给主机暴露 8888 端口，这样外部网络可以访问你的应用
EXPOSE 8888
# 告诉 Docker 启动容器运行的命令
CMD ["go-wrapper", "run"]
```

## 3. 生成镜像
- docker build --rm -t gotestimg .

## 4. 启动
- docker run -p 9527:8888 --name="testgo" gotestimg

## 5. Dockerfile FROM
- https://hub.daocloud.io/repos/514a7477-98ca-4262-83d4-1baaf4fe8ce4
- 这里可以找到不用翻墙的
- 如果生成镜像的过程中出现连接超时的操作 可以考虑换一个镜像
- 1.9-rc-alpine 可用

## 6. go-wrapper
- https://github.com/docker-library/golang/blob/master/go-wrapper


