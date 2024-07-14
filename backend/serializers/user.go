package serializers

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpResponseStatus int

const (
	HttpResponseFail HttpResponseStatus = iota
	HttpResponseSuccess
)

type HttpResponse struct {
	Status HttpResponseStatus `json:"status"` // 0 : error , 1 :success
	Error  string             `json:"error"`
	Data   interface{}        `json:"data"`
}

type DocumentUpload struct {
	Data  string `json:"data"`
	Owner string `json:"owner"`
}
