package client

import (
	"errors"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
)

type Client struct {
	sdk.Client
}

func NewClient(config *sdk.Config) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		return nil, errors.New("配置不能为空")
	}
	client.Init().WithConfig(config)
	return
}
