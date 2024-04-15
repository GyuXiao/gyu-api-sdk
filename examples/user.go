package main

import (
	"encoding/json"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/GyuXiao/gyu-api-sdk/sdk/request"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/GyuXiao/gyu-api-sdk/service/user"
)

var accessKey string = "gyu"
var secretKey string = "abcdefg"

// 用户需要做的，只是配置好一个 config，然后向目标 url 发起请求即可

func main() {
	// 使用 config 里的配置创建 client
	config := sdk.NewConfig(accessKey, secretKey)
	client, err := user.NewClient(config)
	if err != nil {
		client.Logger.Errorf("客户端创建错误: %v", err)
		panic(err)
	}
	// 创建一个 user 对象
	userTest := user.NewUser("user_test1")
	userJson, _ := json.Marshal(userTest)
	// 创建请求
	req := &request.BaseRequest{
		URL:    "http://127.0.0.1:8123/api/user",
		Method: "POST",
		Header: nil,
		Body:   string(userJson),
	}
	// 通过 client 发起请求
	err = client.Send(req, &response.BaseResponse{})
	if err != nil {
		client.Logger.Errorf("客户端请求错误: %v", err)
	}
}
