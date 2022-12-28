package commonData

type CommonData struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	DataCount int         `json:"data_count"`
	TotalData int         `json:"total_data"`
}

func (c *CommonData) DefaultResponse() {
	c.Message = "Ok"
}
