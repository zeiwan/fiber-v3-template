package whiteCatApi

type CommonResp struct {
	Message any    `json:"message"`
	Code    string `json:"code"`
	Data    any    `json:"data"`
}
