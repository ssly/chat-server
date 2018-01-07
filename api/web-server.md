## 从基本的http包开始构建

Go语言比较简洁，容易上手，初学go，搭建web服务器，可以不需要使用框架，从基本的入手。下面笔者从最基本的语法入手一个web聊天室。

### 思路
1. 登录阶段
    - 由最简单用户、密码登录聊天室
    - 走POST请求，{"name": "xx", "password" "xxx"}
2. 聊天阶段
    - 走websocket协议
    - 消息分两种：登录、聊天
    - 登录成功，客户端连接ws，发送登录请求

### 第一步：实现一个简单的登录接口
下面实现了一个简单的GO服务器，包含接口`/login`
```go
package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlerLogin)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		w.Write([]byte("Request Error"))
	}

    // 取body体，取出来是[]byte
	bodyStr, _ := ioutil.ReadAll(r.Body)
	w.WriteHeader(200)
	w.Write(bodyStr)
}
```
### 第二步：登录成功之后，客户端注册登录

