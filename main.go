package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/websocket"
)

type User struct {
	Name string `json:"name"`
	id   *websocket.Conn
}

var userList []User

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", handlerLogin)
	http.Handle("/chat", websocket.Handler(handlerChatroom))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}

// handlerLogin 登录接口
func handlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(400)
		w.Write([]byte("Request Error"))
	}

	bodyStr, _ := ioutil.ReadAll(r.Body)
	w.WriteHeader(200)
	w.Write(bodyStr)
}

// handlerChatroom 聊天websocket
func handlerChatroom(ws *websocket.Conn) {
	defer ws.Close()
	var reciveStr []byte
	var reciveData map[string]string

WSFor:
	for {
		// 接收信息
		websocket.Message.Receive(ws, &reciveStr)
		err := json.Unmarshal(reciveStr, &reciveData)
		if err != nil {
			reciveData = map[string]string{}
		}
		fmt.Println("收到的信息：", reciveData)

		switch reciveData["type"] {
		case "login":
			var user = &User{
				Name: reciveData["username"],
				id:   ws,
			}

			isContain := func() bool {
				for i := range userList {
					if userList[i].Name == user.Name {
						// 如果名称相同，则需要替换id
						userList[i].id = user.id

						return true
					}
				}
				return false
			}()

			if !isContain {
				userList = append(userList, *user)
			}

			fmt.Println("登录的列表", userList)
			// 返回聊天列表给所有的userList
			sendData := &struct {
				Type     string `json:"type"`
				UserList []User `json:"userList"`
			}{
				Type:     "login",
				UserList: userList,
			}

			sendStr, _ := json.Marshal(sendData)

			if err := sendMessage(string(sendStr)); err != nil {
				break WSFor
			}
		case "chat":
			fmt.Println("聊天")
			if err := sendMessage(string(reciveStr)); err != nil {
				break WSFor
			}
		}

	}
}

func sendMessage(msg string) error {
	for i, v := range userList {
		err := websocket.Message.Send(v.id, msg)
		if err != nil {
			fmt.Println("发送出错", err)
			fmt.Println("准备发送的信息", msg)
			// 删除发送失败的用户
			userList = append(userList[:i], userList[i+1:]...)
			fmt.Println("删除后剩余的用户", userList)

			// 返回聊天列表给所有的userList
			sendData := &struct {
				Type     string `json:"type"`
				UserList []User `json:"userList"`
			}{
				Type:     "login",
				UserList: userList,
			}
			sendStr, _ := json.Marshal(sendData)
			sendMessage(string(sendStr))
			return err
		}
	}

	return nil
}
