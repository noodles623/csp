package objects

import (
	"encoding/json"
	"net/http"
)

const MaxListLimit = 200

type GetRequest struct {
	Id string `json:"id"`
}

type ListRequest struct {
	Limit int `json:"limit"`
}

type CreateRequest struct {
	Asset *Asset `json:"asset"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type AssetResponseWrapper struct {
	Asset  *Asset   `json:"asset,omitempty""`
	Assets []*Asset `json:"events,omitempty"`
	Code   int      `json:"-"`
}

func (a *AssetResponseWrapper) Json() []byte {
	if a == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(a)
	return res
}

func (a *AssetResponseWrapper) StatusCode() int {
	if a == nil || a.Code == 0 {
		return http.StatusOK
	}
	return a.Code
}
