package resp

type MovieHotListResp struct {
	Data []struct {
		Id           int    `json:"id"`
		CityId       string `json:"cityId"`
		Director     string `json:"director"`
		PublishDate  string `json:"publishDate"`
		VersionTypes string `json:"versionTypes"`
		Language     string `json:"language"`
		ShowStatus   string `json:"showStatus"`
		Pic          string `json:"pic"`
		FilmTypes    string `json:"filmTypes"`
		LikeNum      string `json:"likeNum"`
		Duration     string `json:"duration"`
		Cast         string `json:"cast"`
		FilmId       int    `json:"filmId"`
		Grade        string `json:"grade"`
		Intro        string `json:"intro"`
		Name         string `json:"name"`
		CreatedAt    string `json:"createdAt"`
	} `json:"data"`
	Message string `json:"message"`
	Code    string `json:"code"`
}
