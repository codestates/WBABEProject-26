package dto


type ResponseBody struct {
	Result bool `json:"result" binding:"required"`
	Data interface{} `json:"data"`
	Msg interface{} `json:"msg" binding:"omitempty"`
}
