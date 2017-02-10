package twitch2go

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/juju/errors"
)

// GetChannelByOAuth will return a Channel object for the given oauth.  Will return annotated errors.
func (c *Client) GetChannelByOAuth(oauth string) (*Channel, error) {
	if oauth == "" {
		return nil, errors.New("OAuth token required")
	}
	url := "/channel"
	ops := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelByOAuth")
	}
	defer resp.Body.Close()
	ch := &Channel{}
	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return ch, nil
}

// GetChannelByID will return a Channel object for the given channelID.  Will return annotated errors.
func (c *Client) GetChannelByID(channelID string) (*Channel, error) {
	url := fmt.Sprintf("/channels/%s", channelID)
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelByID")
	}
	defer resp.Body.Close()
	ch := &Channel{}
	err = json.NewDecoder(resp.Body).Decode(&ch)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return ch, nil
}

func (c *Client) GetChannelEditors(channelID string, oauth string) (*[]User, error) {
	if oauth == "" {
		return nil, errors.New("OAuth token required")
	}
	url := fmt.Sprintf("/channels/%s/editors", channelID)
	ops := &doOptions{
		oauth: oauth,
	}
	// Do the requst
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelEditors")
	}
	defer resp.Body.Close()
	editors := &Editors{}
	err = json.NewDecoder(resp.Body).Decode(&editors)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return &editors.Users, nil
}

// GetChannelFollows returns a list of users that follow the channel.  The first time you call the function, pass in an empty cursor.  On consecutive calls, pass in the curosr returned from the previous call.  When the cursor object is empty, you have reached the end of the list.  Limit is the size of the list returned.  Default is 25, maximum is 100.
func (c *Client) GetChannelFollows(channelID string, cursor string, limit int) (*Followers, error) {
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
		return nil, errors.Annotate(err, "GetChannelFollows")
	}
	defer resp.Body.Close()
	follows := &Followers{}
	err = json.NewDecoder(resp.Body).Decode(&follows)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return follows, nil
}

func (c *Client) GetChannelSubscribers(channelID string, oauth string, cursor string, limit int) (*Subscribers, error) {
	if limit <= 0 {
		return nil, errors.New("Limit must be a positive number less than 100")
	} else if limit > 100 {
		return nil, errors.New("Limit must be less than or equal to 100")
	}
	if oauth == "" {
		return nil, errors.New("OAuth token required")
	}
	url := fmt.Sprintf("/channels/%s/subscriptions", channelID)
	ops := &doOptions{
		params: map[string]string{"cursor": cursor, "limit": strconv.Itoa(limit)},
		oauth:  oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelSubscribers")
	}
	defer resp.Body.Close()
	subscribers := &Subscribers{}
	err = json.NewDecoder(resp.Body).Decode(&subscribers)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return subscribers, nil
}

func (c *Client) GetChannelSubscriberByUser(channelID string, userID string, oauth string) (*Subscription, error) {
	if oauth == "" {
		return nil, errors.New("OAuth token required")
	}
	url := fmt.Sprintf("/channles/%s/subscriptions/%s", channelID, userID)
	ops := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelSubscriberByUser")
	}
	defer resp.Body.Close()
	subscription := &Subscription{}
	err = json.NewDecoder(resp.Body).Decode(&subscription)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return subscription, nil
}

func (c *Client) GetChannelVideos(channelID string) (*Videos, error) {
	url := fmt.Sprintf("/channels/%s/videos", channelID)
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "GetChannelVideos")
	}
	defer resp.Body.Close()
	videos := &Videos{}
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return videos, nil
}
