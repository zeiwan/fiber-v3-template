package whiteCatApi

import (
	"fiber/pkg"
	"fiber/pkg/wechatApi"
	"github.com/imroc/req/v3"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/url"
)

var (
	host = "https://mapi.imax.kim/api/app/v1"
)

// CatClient 客户端结构体
type CatClient struct {
	invoker pkg.Invoker
}

// NewCatClient 创建客户端
func NewCatClient() *CatClient {
	client := req.C().DevMode()
	client.SetBaseURL(host)

	c := &CatClient{
		invoker: pkg.Invoker{Client: client},
	}
	// 设置重试次数
	client.SetCommonRetryCount(3).
		AddCommonRetryCondition(c.shouldRetry)
	return c
}

// CityList 城市列表
func (c *CatClient) CityList() (resp CommonResp, err error) {
	err = c.invoker.Get("/dyp/getCityList", nil, &resp)
	return
}

// GetHotList 热门电影
func (c *CatClient) GetHotList(cityId string) (resp CommonResp, err error) {
	params := url.Values{
		"cityId": []string{cityId},
	}
	err = c.invoker.Get("/dyp/getHotList", params, &resp)
	return
}

// CinemaBrand 影院品牌
func (c *CatClient) CinemaBrand() (resp CommonResp, err error) {
	body := map[string]interface{}{
		"dicCode": "dyp_cinema_list",
	}
	err = c.invoker.Post("/common/systemDict", nil, nil, body, &resp)
	return
}

// Geocoder 地理编码
func (c *CatClient) Geocoder(lat, lng string) (resp CommonResp, err error) {
	params := url.Values{
		"lat": []string{lat},
		"lng": []string{lng},
	}
	err = c.invoker.Get("/common/geocoder", params, &resp)
	return
}

// GetCityArea 城市区域
func (c *CatClient) GetCityArea(cityId string) (resp CommonResp, err error) {
	params := url.Values{
		"cityId": []string{cityId},
	}
	err = c.invoker.Get("/dyp/getCityArea", params, &resp)
	return
}

// GetDateList 日期列表
func (c *CatClient) GetDateList(filmId, cityId string) (resp CommonResp, err error) {
	params := url.Values{
		"filmId": []string{filmId},
		"cityId": []string{cityId},
	}
	err = c.invoker.Get("/dyp/getShowDate", params, &resp)
	return
}

// GetCinemas 影院列表
func (c *CatClient) GetCinemas(filmId, cityId, date, latitude, longitude string) (resp CommonResp, err error) {

	body := map[string]string{
		"page":      "1",
		"limit":     "100",
		"filmId":    filmId,
		"cityId":    cityId,
		"area":      "",
		"date":      date,
		"latitude":  latitude,
		"longitude": longitude,
	}

	err = c.invoker.Post("/dyp/getShowList", nil, body, nil, &resp)
	return
}

// GetCinemaSeat 影院座次
func (c *CatClient) GetCinemaSeat(showId string) (resp CommonResp, err error) {
	params := url.Values{
		"showId": []string{showId},
	}
	err = c.invoker.Get("/dyp/getSeat", params, &resp)
	return
}

// GetTabsCinemaList 影院列表Tabs 筛选条件
func (c *CatClient) GetTabsCinemaList(cityId, lng, lat, regionName, keyword string) (resp CommonResp, err error) {
	params := url.Values{
		"cityId":     []string{cityId},
		"lng":        []string{lng},
		"lat":        []string{lat},
		"regionName": []string{},
		"keyword":    []string{},
	}
	err = c.invoker.Get("/dyp/getCinemaList", params, &resp)
	return
}

// shouldRetry 判断是否需要重试
func (c *CatClient) shouldRetry(resp *req.Response, err error) bool {

	// 3. 如果有响应体，检查业务错误码
	if resp != nil && resp.Response != nil && resp.Response.Body != nil {
		defer resp.Response.Body.Close()

		body, err := io.ReadAll(resp.Response.Body)
		if err != nil {
			return true // 读取失败，可能需要重试
		}

		// 使用 jsoniter 解析业务错误码
		code := jsoniter.Get(body, "code").ToString()
		if code == "A0230" { // Token 失效，需要刷新
			client := wechatApi.NewWeChatClient()
			commonResp, _ := client.WeChatCode()
			// 更新 token
			c.setToken(commonResp.Data.Code)
			return true
		}
	}

	// 其他情况不重试
	return false
}

// setToken 设置新的 token，并更新到 HTTP Client 的公共请求头
func (c *CatClient) setToken(code string) {
	params := url.Values{
		"code":  []string{code},
		"appid": []string{"wx1c21fe0869e8f11e"},
		"idNum": []string{},
	}

	var resp CommonResp
	err := c.invoker.Get("/auth/modelXcxLogin", params, &resp)
	if err != nil {
		return
	}
	//  code 00000 表示成功 并且设置 token
	if resp.Code == "00000" {
		if token, ok := resp.Data.(string); ok {
			c.invoker.Client.SetCommonHeader("token", token)
			return
		}
		return
	}
	return
}
