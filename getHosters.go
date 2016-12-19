package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// HostersList contains information about hosted channels.
type HostersList struct {
	Hosters []struct {
		UserLogo string `json:"user_logo"`
		UserName string `json:"user_name"`
	} `json:"hosters"`
}

// GetHosters returns a list of channels hosting channel.
//
// Editors can read this API.
func (account *Account) GetHosters(channel string) (*HostersList, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case channel == "":
		return nil, errors.New("channel can not be empty")
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/hosters/", channel)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj HostersList
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
