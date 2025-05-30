package wechatApi

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	client := NewWeChatClient()
	err, resp := client.WeChatCode()
	fmt.Println(resp, err)
}
