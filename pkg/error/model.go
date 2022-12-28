package errorTools

type Enum struct {
	Type       string
	Error      error
	StatusCode int
}

func (e *Enum) IsNotEmpty() bool {
	if e.Type == "" && e.Error == nil && e.StatusCode == 0 {
		return false
	}
	return true
}
func (e *Enum) Empty() {
	e.Type = ""
	e.Error = nil
	e.StatusCode = 0
}
func (e *Enum) JSONResponse() (int, ErrorResponse) {
	var err ErrorStruct
	err.ErrorEnum = *e
	return err.Response()
}

type Validation struct {
	Field   string
	Message string
}
type Validations []Validation

type ErrorStruct struct {
	ErrorEnum   Enum
	Message     string
	Validations Validations
}

func (e *ErrorStruct) Response() (int, ErrorResponse) {
	var errorResp ErrorResponse
	if !e.IsNotEmpty() {
		return InternalServerError, ErrorResponse{}
	}
	errorResp.Type = e.ErrorEnum.Type
	if e.Message != "" {
		errorResp.Message += " : " + e.Message
		if errorResp.Message != "" {
			errorResp.Message += " : "
		}
	}
	errorResp.Message += e.ErrorEnum.Error.Error()
	errorResp.StatusCode = e.ErrorEnum.StatusCode
	errorResp.Validations = e.Validations
	return errorResp.StatusCode, errorResp
}
func (e *ErrorStruct) Empty() ErrorStruct {
	return ErrorStruct{}
}
func (e *ErrorStruct) AddFunctionName(function string) {
	e.Message = function + " : " + e.Message
}
func (e *ErrorStruct) AddMessage(msg string) {
	e.Message += " : " + msg
}
func (e *ErrorStruct) IsNotEmpty() bool {
	if e.Message == "" && !e.ErrorEnum.IsNotEmpty() {
		return false
	}
	return true
}

type ErrorResponse struct {
	Type        string      `json:"type"`
	Message     string      `json:"message"`
	StatusCode  int         `json:"status_code"`
	Validations Validations `json:"validations"`
}
