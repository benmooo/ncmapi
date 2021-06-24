package apitypes

import "encoding/json"

type H map[string]interface{}

func (h H) Set(k string, v interface{}) H {
	h[k] = v
	return h
}

func (h H) Json() []byte {
	b, _ := json.Marshal(h)
	return b
}

func (h H) Merge(hs ...H) H {
	for _, _h := range hs {
		for k, v := range _h {
			h[k] = v
		}
	}
	return h
}
