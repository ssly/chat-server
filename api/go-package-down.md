---
title: Go包安装，无梯子装golang.org/x
date: 2018-01-06 22:57
categories: Go
tags: Go
---

大家在下载golong包时，有时候会遇到没有梯子下载不了的包，下面就`golang.org/x`内的包下载进行说明。
<!-- more -->

就`golang.org/x`下的`net/websocket`包举例说明。

首先正常情况下，下面的命令肯定是没有用的
```
go get golang.org/x/net/websocket
```

`golang.org/x`的包在Github下都是有备份的，我们可以先下载Github的
```
cd $GOPATH/src/golang.org/x
mkdir net
git clone https://github.com/golang/net.git

go install golang.org/x/net/websocket
```

> windows下，`$GOPATH`对应`%GOPATH%`，前提是`%GOPATH%`必须配置在环境变量内，可参考[GO环境变量配置](https://lius.me/blog/)

<br>
如果想要安装其他包同样执行`go install pkg`，可以在`/net`文件夹内查看包是否存在