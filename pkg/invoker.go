package pkg

import (
	"github.com/imroc/req/v3"
	"net/url"
)

type Invoker struct {
	Client *req.Client
}

func (i *Invoker) Get(path string, params url.Values, data any) error {
	client := i.Client.DevMode().R()

	client.QueryParams = params
	client.SetHeader("Accept", "application/json, text/plain, */*")
	return i.do(client, "GET", path, &data)
}

func (i *Invoker) post(path string, params url.Values, body map[string]any, data any) error {
	client := i.Client.DevMode().R()

	client.FormData = params

	client.SetBody(body)

	client.SetHeader("Accept", "application/json, text/plain, */*")
	client.SetHeader("Content-Type", "application/json;charset=UTF-8")
	err := i.do(client, "POST", path, &data)
	return err
}

func (i *Invoker) do(client *req.Request, method string, path string, data any) (err error) {
	resp, err := client.Send(method, path)
	if err != nil {
		return
	}
	err = resp.Into(&data)
	if err != nil {
		return
	}
	return
}
func (i *Invoker) put(path string, filePath string, data any) error {
	client := i.Client.R()
	client = client.SetFile("file", filePath)
	return i.do(client, "POST", path, &data)
}
