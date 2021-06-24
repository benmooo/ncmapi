package apitypes

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	// kind of useless, reserve for future using
	Header *http.Header

	// bytes of http.Response.Body
	Data []byte
}

func NewAPIResponse(header *http.Header, data []byte) *APIResponse {
	return &APIResponse{
		Header: header,
		Data:   data,
	}
}

func (r *APIResponse) String() string {
	return string(r.Data)
}

func (r *APIResponse) Deserialize(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}

// 4 testing, otherwise had better return []byte ?
// remind me never ever do this aggin
func (r *APIResponse) DeserializeToImplicitResult() (*ImplicitResult, error) {
	var res ImplicitResult
	if err := json.Unmarshal(r.Data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
