package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AccessLevels is about permissions of user on channel.
type AccessLevels struct {
	AccessUserID  string `json:"access_user_id"`
	Account       string `json:"account"`
	Admin         string `json:"admin"`
	Broadcast     string `json:"broadcast"`
	Chat          string `json:"chat"`
	Following     string `json:"following"`
	Inbox         string `json:"inbox"`
	IsFollower    bool   `json:"isFollower"`
	IsSubscriber  bool   `json:"isSubscriber"`
	Livestreams   string `json:"livestreams"`
	Partner       string `json:"partner"`
	Payments      string `json:"payments"`
	Recordings    string `json:"recordings"`
	Revenues      string `json:"revenues"`
	Settings      string `json:"settings"`
	Statistics    string `json:"statistics"`
	Subscriptions string `json:"subscriptions"`
	Superadmin    string `json:"superadmin"`
	Teams         string `json:"teams"`
	UserID        string `json:"user_id"`
	Videos        string `json:"videos"`
}

// GetUserAccessLevels return access levels that auth has in channel.
//
// If you have never been granted Moderator or Editor in channel, this API will only return IsSubscriber and IsFollower.
func (account *Account) GetUserAccessLevels(channel string) (*AccessLevels, error) {
	if err := checkGetUserAccessLevels(account, channel); err != nil {
		return nil, err
	}

	url := fmt.Sprint(API, "/user/access/", channel, "/", account.AuthToken)
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj AccessLevels
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func checkGetUserAccessLevels(account *Account, channel string) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case channel == "":
		return errors.New("channel can not be empty")
	}
	return nil
}
