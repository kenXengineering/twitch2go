package twitch2go

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/juju/errors"
)

// ChannelByOAuth will return a Channel object for the given oauth.  Will return annotated errors.
func (c *Client) ChannelByOAuth(oauth string) (*Channel, error) {
	url := "/channel"
	ops := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "ChannelByOAuth")
	}
	defer resp.Body.Close()
	ch := &Channel{}
	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return ch, nil
}

// ChannelByID will return a Channel object for the given channelID.  Will return annotated errors.
func (c *Client) ChannelByID(channelID string) (*Channel, error) {
	url := fmt.Sprintf("/channels/%s", channelID)
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "ChannelByID")
	}
	defer resp.Body.Close()
	ch := &Channel{}
	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return ch, nil
}

func (c *Client) ChannelEditors(channelID string, oauth string) (*[]User, error) {
	url := fmt.Sprintf("/channels/%s/editors", channelID)
	ops := &doOptions{
		oauth: oauth,
	}
	// Do the requst
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "ChannelEditors")
	}
	defer resp.Body.Close()
	editors := &Editors{}
	err = json.NewDecoder(resp.Body).Decode(&editors)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return editors.Users, nil
}

// ChannelFollows returns a list of users that follow the channel.  The first time you call the function, pass in an empty cursor.  On consecutive calls, pass in the curosr returned from the previous call.  When the cursor object is empty, you have reached the end of the list.  Limit is the size of the list returned.  Default is 25, maximum is 100.
func (c *Client) ChannelFollows(channelID string, cursor string, limit int) (*Follows, error) {
	if limit <= 0 {
		return nil, errors.New("Limit must be a positive number less than 100")
	} else if limit > 100 {
		return nil, errors.New("Limit must be less than or equal to 100")
	}
	url := fmt.Sprintf("/channels/%s/follows", channelID)
	ops := &doOptions{
		params: map[string]string{"cursor": cursor, "limit": strconv.Itoa(limit)},
	}
	// Do the request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	follows := &Follows{}
	err = json.NewDecoder(resp.Body).Decode(&follows)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return follows, nil
}
