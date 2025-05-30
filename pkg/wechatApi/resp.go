package wechatApi

type CommonResp struct {
	Data struct {
		AppIconUrl   string `json:"appIconUrl"`
		AppName      string `json:"appName"`
		BaseResponse struct {
			ErrMsg struct {
				String string `json:"string"`
			} `json:"errMsg"`
			Ret int `json:"ret"`
		} `json:"baseResponse"`
		Code              string `json:"code"`
		JsApiBaseResponse struct {
			Errcode int    `json:"errcode"`
			Errmsg  string `json:"errmsg"`
		} `json:"jsApiBaseResponse"`
		LiftSpan      int           `json:"liftSpan"`
		OpenId        string        `json:"openId"`
		ScopeList     []interface{} `json:"scopeList"`
		SessionKey    string        `json:"sessionKey"`
		SessionTicket string        `json:"sessionTicket"`
		Signature     string        `json:"signature"`
		State         string        `json:"state"`
	} `json:"data"`
	Type int `json:"type"`
}
