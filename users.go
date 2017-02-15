package twitch2go

import (
	"encoding/json"
	"strconv"

	"github.com/juju/errors"
)

/*
GetUserByOauth returns a user from the given oauth token.
*/
func (c *Client) GetUserByOAuth(oauth string) (*User, error) {
	url := "/user"
	opts := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, opts)
	if err != nil {
		return nil, errors.Annotate(err, "GetUserByOAuth")
	}
	defer resp.Body.Close()
	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return user, nil
}

// GetUserByID will return ther user for the given ID.
func (c *Client) GetUserByID(userID string) (*User, error) {
	url := "/users/" + userID
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "GetUserByID")
	}
	defer resp.Body.Close()
	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, errors.Annotate(err, "Error deocding JSON")
	}
	return user, nil
}

/*
GetUserFollows returns a list of channles the given user follows.

The function takes in five parameters:
	userID:
		The User ID

	limit:
		The number of channels to return.  Max is 100, default is 25.

	offset:
		The offset in the list to return.  If a list has more entries then the limit, the returned list will be a subset of the whole.  The offset lets you continue where you left off.

	direction:
		The direction of the returned list.  Values are `ASC` and `DESC` (newest first).  Default is `DESC`.

	sortBy:
		The sort of the returned list.  Values are `CreatedAt`, `LastBroadcast`, and `Login`  Default is `CreatedAt`.
*/
func (c *Client) GetUserFollows(userID string, limit int, offset int, direction Direction, sortBy SortBy) (*Followers, error) {
	url := "/users/" + userID + "/follows/channels"
	opts := &doOptions{
		params: map[string]string{
			"limit":     strconv.Itoa(limit),
			"offset":    strconv.Itoa(offset),
			"direction": string(direction),
			"sortby":    string(sortBy),
		},
	}
	// Do the request
	resp, err := c.do("GET", url, opts)
	if err != nil {
		return nil, errors.Annotate(err, "GetUserFollows")
	}
	defer resp.Body.Close()
	follows := &Followers{}
	err = json.NewDecoder(resp.Body).Decode(follows)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return follows, nil
}

// CheckUserSubscriptionByChannel will return a user object if the given user is subscribed to the given channel.  If no user and an error is returned, then either the user is not subscribed to the channel, or the channel does not have a subscription program.
func (c *Client) CheckUserSubscriptionByChannel(userID string, channelID string, oauth string) (*User, error) {
	url := "/users/" + userID + "/subscriptions/" + channelID
	opts := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, opts)
	if err != nil {
		return nil, errors.Annotate(err, "CheckUserSubscriptionByChannel")
	}
	defer resp.Body.Close()
	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return user, nil
}

// CheckUserFollowsChannel will return the user object if the user is following the given channel.  If no users and an error is returned, then the user is not following the channel.
func (c *Client) CheckUserFollowsChannel(userID string, channelID string) (*User, error) {
	url := "/users/" + userID + "/follows/channels/" + channelID
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "CheckUserFollowsChannel")
	}
	defer resp.Body.Close()
	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return user, nil
}
