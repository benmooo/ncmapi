package apitypes

// basic json response
type ImplicitResult struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Time    int         `json:"time"`
	Result  interface{} `json:"result"`
	Data    interface{} `json:"data"`
}
