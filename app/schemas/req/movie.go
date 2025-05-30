package req

// MovieCityIdReq 城市id请求参数结构体
type MovieCityIdReq struct {
	CityId string `query:"cityId" `
}

// CinemaListReq 影院列表请求参数结构体
type CinemaListReq struct {
	CityId     string `query:"cityId" `
	Lng        string `query:"lng" `
	Lat        string `query:"lat" `
	RegionName string `query:"regionName"`
	Keyword    string `query:"keyword"`
}
