package whiteCatApi

import (
	"fiber/pkg"
	"github.com/imroc/req/v3"
	"net/url"
)

var (
	host         = "https://mapi.imax.kim"
	defaultToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NDg1NjI0MDMsInJvbGUiOiJkeXAiLCJleHAiOjE3ODAwOTg0MDMsInN1YiI6NDk2MjgzfQ.UqVhDxEiapIyJNl3Y9fAN4SXLEpupM-Qw-Mn4UZrVnM"
)

type CatClient struct {
	invoker pkg.Invoker
}

func NewCatClient() *CatClient {
	client := req.C().DevMode()
	client.SetBaseURL(host)
	client.SetCommonHeader("token", defaultToken)
	return &CatClient{
		invoker: pkg.Invoker{Client: client},
	}
}
func (c CatClient) SetToken(token string) {
	c.invoker.Client.SetCommonHeader("token", token)
}

func (c CatClient) GetCityList() {
	c.invoker.Get("/api/app/v1/dyp/getCityList", nil, nil)
}
func (c CatClient) GetHotList(cityId string) (err error, resp CommonResp) {
	values := url.Values{}
	values.Set("cityId", cityId)
	err = c.invoker.Get("api/app/v1/dyp/getHotList", values, &resp)
	return err, resp
}
