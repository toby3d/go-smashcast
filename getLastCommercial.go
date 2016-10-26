package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// LastCommercial is a response body about last commercial.
type LastCommercial struct {
	SecondsAgo string `json:"seconds_ago"`
	AdCount    string `json:"ad_count"`
	TimeOut    int64  `json:"timeout"`
}

// GetLastCommercial returns last commercial break object.
func GetLastCommercial(channel string) (LastCommercial, error) {
	requestURL := fmt.Sprintf("%s/ws/combreak/%s", API, channel)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return LastCommercial{}, err
	}
	var obj LastCommercial
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return LastCommercial{}, err
	}
	return obj, nil
}
