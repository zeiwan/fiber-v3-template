package whiteCatApi

type CommonResp struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    string `json:"code"`
}
