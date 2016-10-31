package modules

type Doc struct {
	ServiceName     string   `json:"service_name"`
	ServiceUrl      string   `json:"service_url"`
	Version         string   `json:"version"`
	MS              string   `json:"MS"`
	BaseMs          string   `json:base_ms`
	Author          string   `json:"author"`
	AvailableStatus int8     `json:"available_status"`
	Description     string   `json:"description"`
	ModTime         string   `json:"mod_time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
}

type Request struct {
	RequestType       string         `json:"request_type"`
	RequestParamsType string         `json:"request_params_type"`
	RequestParams     []RequestParam `json:"request_params"`
	RequestExample    interface{}    `json:"request_example"`
}

type RequestParam struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Required    bool     `json:"required"`
	Example     string   `json:"example"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
	Remark      string   `json:"remark"`
	Child       []Parmam `json:"child"`
}

type Parmam struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Required    bool        `json:"required"`
	Example     interface{} `json:"example"`
	Default     interface{} `json:"default"`
	Description string      `json:"description"`
	Remark      string      `json:"remark"`
	Child       []Parmam    `json:"child"`
}

type RequestExample struct {
	Meta map[string]interface{} `json:"meta"`
	Data map[string]interface{} `json:"data"`
}

type Response struct {
	ResponseParamsType string          `json:"response_params_type"`
	ResponseParams     []ResponseParam `json:"response_params"`
	ResponseError      []ResponseError `json:"response_error"`
	ResponseExample    interface{}     `json:"response_example"`
}

type ResponseParam struct {
	Description string   `json:"description"`
	Params      []Parmam `json:"params"`
}

type ResponseError struct {
	Status interface{} `json:"status"`
	Data   string      `json:"data"`
}

type ResponseExample struct {
	Meta   string      `json:"meta"`
	Status int16       `json:"status"`
	Data   interface{} `json:"data"`
}

type Base struct {
}
