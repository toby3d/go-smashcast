package hitGox

/*
import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// GetChatServers
func GetChatServers() (*[]string, error) {
	url := API + "/chat/servers"
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	var obj = []struct {
		ServerIP string `json:"server_ip"`
	}{}
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
*/
