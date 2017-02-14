package twitch2go

import (
	"encoding/json"
	"strconv"

	"github.com/juju/errors"
)

// GetChannelByOAuth will return a Channel object for the given oauth.  Will return annotated errors.
func (c *Client) GetChannelByOAuth(oauth string) (*Channel, error) {
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
	url := "/channels" + channelID
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

// GetChannelEditors returns a list of Users that are editors for the given channel.  Requires users oauth token.
func (c *Client) GetChannelEditors(channelID string, oauth string) (*[]User, error) {
	url := "/channels/" + channelID + "/editors"
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

/*
GetChannelFollows returns a list of followers for the given channel.  Since a channel can have thousands of followers, you must make multiple requests of smaller sizes.

The function takes four parameters:

	channelID:
		The channel ID

	cursor:
		Tells the server where to start fetching the next set of results, in a multi-page response.

	limit:
		Number of records to return.  Maximum is 100, default is 25.

	direction:
		Direction of sorting.  Valid values are `ASC`, and `DESC` (newest first).  Default is `DESC`.
*/
func (c *Client) GetChannelFollows(channelID string, cursor string, limit int, direction Direction) (*Followers, error) {
	if limit <= 0 {
		limit = 25
	} else if limit > 100 {
		limit = 100
	}
	url := "/channels/" + channelID + "/follows"
	ops := &doOptions{
		params: map[string]string{
			"cursor":    cursor,
			"limit":     strconv.Itoa(limit),
			"direction": string(direction),
		},
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

/*
GetChannelSubscribers returns a list of users that are subscribe to the channel.

Function takes five parameters:

	channelID:
		ID of the Channel

	oauth:
		User oauth token

	limit:
		Maximum number of objects top return.  Default 25, Maximum 100.

	direction:
		Sorting direction.  Valid values are `ASC` and `DESC`.  Default `ASC` (oldest first)
*/
func (c *Client) GetChannelSubscribers(channelID string, oauth string, limit int, offset int, direction Direction) (*Subscribers, error) {
	if limit <= 0 {
		limit = 25
	} else if limit > 100 {
		limit = 100
	}
	url := "/channels/" + channelID + "/subscriptions"
	ops := &doOptions{
		params: map[string]string{
			"limit":     strconv.Itoa(limit),
			"offset":    strconv.Itoa(offset),
			"direction": string(direction),
		},
		oauth: oauth,
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

// GetChannelSubscriberByUser will return the subscriber information if the user is subscribed to the channel.
func (c *Client) GetChannelSubscriberByUser(channelID string, userID string, oauth string) (*Subscription, error) {
	url := "/channels/" + channelID + "/subscribtions/" + userID
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

/*
GetChannelVideos returns a list of all videos on the given channel.

Function takes size parameters:

	channelID:
		Channel ID

	limit:
		Maximum number of objects to return.  Default 10, Maximum 100.

	offset:
		The offset in the list to return.  If a list has more entries then the limit, the returned list will be a subset of the whole.  The offset lets you continue where you left off.

	broadcastType:
		Comma separated list with any combination of `archive`, `highlight`, and `upload`.  Default is `highlight`.

	language:
		Constrains the language of the videos returned.  For example `en,es`.  Default is all languages.

	sort:
		Sorting order of returned videos.  Valid values `views` and `time`.  Default is `time` (most recent first).
*/
func (c *Client) GetChannelVideos(channelID string, limit int, offset int, broadcastType string, language string, sort VideoSort) (*Videos, error) {
	url := "/channels/" + channelID + "/videos"
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	opts := &doOptions{
		params: map[string]string{
			"limit":          strconv.Itoa(limit),
			"offset":         strconv.Itoa(offset),
			"boradcast_type": broadcastType,
			"language":       language,
			"sort":           string(sort),
		},
	}

	// Do the request
	resp, err := c.do("GET", url, opts)
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
