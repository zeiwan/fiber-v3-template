package wechatApi

import (
	"fiber/pkg"
	"github.com/imroc/req/v3"
)

var (
	host = "http://10.0.10.24:8800"
)

// WeChatClient 客户端结构体
type WeChatClient struct {
	invoker pkg.Invoker
}

// NewWeChatClient 创建客户端
func NewWeChatClient() *WeChatClient {
	client := req.C()
	client.SetBaseURL(host)
	c := &WeChatClient{
		invoker: pkg.Invoker{Client: client},
	}
	return c
}

func (c WeChatClient) WeChatCode() (resp CommonResp, err error) {
	body := map[string]any{
		"type": 11136,
		"data": map[string]string{
			"appid": "wx1c21fe0869e8f11e",
		},
	}
	err = c.invoker.Post("/ClientId", nil, nil, body, &resp)
	return
}
