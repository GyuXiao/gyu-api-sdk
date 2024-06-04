package main

import (
	"encoding/json"
	"github.com/GyuXiao/gyu-api-sdk/client"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/GyuXiao/gyu-api-sdk/sdk/request"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/GyuXiao/gyu-api-sdk/service/user"
)

// 用户需要做的，只是配置好一个 config，然后向目标 url 发起请求即可

func main() {
	// 使用 SDK 分 3 步走
	// 下面以向 user 服务发起请求作为例子

	// 1,使用 config 里的配置创建 client
	config := sdk.NewConfig(user.AccessKey, user.SecretKey)
	clt, err := client.NewClient(config)
	if err != nil {
		clt.Logger.Errorf("客户端创建错误: %v", err)
		panic(err)
	}

	// 具体情况具体分析，比如这里的 request body 是一个 user 对象
	userTest := user.NewUser(user.TestUsername)
	userJson, _ := json.Marshal(userTest)

	// 2,构建服务请求
	req := &request.BaseRequest{
		URL:    user.ServiceUrl,
		Method: user.MethodPost,
		Header: nil,
		Body:   string(userJson),
	}
	resp := &response.BaseResponse{}

	// 3,通过 client 向 user 服务发起一个 POST 请求
	err = clt.Send(req, resp)
	if err != nil {
		clt.Logger.Errorf("客户端请求错误: %v", err)
	}
}
