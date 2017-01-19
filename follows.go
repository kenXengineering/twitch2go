package twitch2go

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var limit = 25

// Follower Follower data for twtich channel
type Follow struct {
	CreatedAt     time.Time         `json:"created_at"`
	Links         map[string]string `json:"_links"`
	Notifications bool              `json:"notifications"`
	User          User              `json:"user"`
	Channel       Channel           `json:"channel"`
}

type followResponse struct {
	Total   int               `json:"_total"`
	Links   map[string]string `json:"_links"`
	Cursor  string            `json:"_cursor"`
	Follows []Follow          `json:"follows"`
}

// GetChannelFollows Get a list for Follows for the given twitch channel
func (c *Client) GetChannelFollows(channelID int64) (*[]Follow, error) {
	url := "/channels/" + strconv.FormatInt(channelID, 10) + "/follows"
	ops := &doOptions{
		params: map[string]string{"limit": "100"},
	}
	// Do the initial request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fr := &followResponse{}
	err = json.NewDecoder(resp.Body).Decode(&fr)
	if err != nil {
		fmt.Println(err)
	}
	cursor := fr.Cursor
	var follows []Follow
	follows = append(follows, fr.Follows...)
	// Paginated request
	for {
		if cursor == "" {
			break
		}
		resp, err = c.do("GET", url, &doOptions{
			params: map[string]string{"cursor": cursor, "limit": "100"},
		})
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fr = &followResponse{}
		err = json.NewDecoder(resp.Body).Decode(&fr)
		if err != nil {
			return nil, err
		}
		follows = append(follows, fr.Follows...)
		cursor = fr.Cursor
	}

	return &follows, nil
}

func (c *Client) GetChannelFollowsStream(channelID int64, dataChan chan<- []Follow, doneChan chan<- bool, errChan chan<- error) {

	url := "/channels/" + strconv.FormatInt(channelID, 10) + "/follows"
	ops := &doOptions{
		params: map[string]string{"limit": "100"},
	}
	// Do the initial request
	resp, err := c.do("GET", url, ops)
	if err != nil {
		errChan <- err
		doneChan <- true
		return
	}
	defer resp.Body.Close()
	fr := &followResponse{}
	err = json.NewDecoder(resp.Body).Decode(&fr)
	if err != nil {
		errChan <- err
		doneChan <- true
		return
	}
	cursor := fr.Cursor
	var follows []Follow
	follows = append(follows, fr.Follows...)
	dataChan <- follows
	// Paginated request
	for {
		if cursor == "" {
			break
		}
		follows = []Follow{}
		resp, err = c.do("GET", url, &doOptions{
			params: map[string]string{"cursor": cursor, "limit": "100"},
		})
		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}
		defer resp.Body.Close()
		fr = &followResponse{}
		err = json.NewDecoder(resp.Body).Decode(&fr)
		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}
		follows = append(follows, fr.Follows...)
		dataChan <- follows
		cursor = fr.Cursor
	}
	doneChan <- true
}

func (c *Client) GetUserChannelsFollowingStreaming(channelID int64, dataChan chan<- []Follow, doneChan chan<- bool, errChan chan<- error) {
	user := strconv.FormatInt(channelID, 10)
	url := "users/" + user + "/follows/channels"
	doOptions := &doOptions{
		params: map[string]string{
			"direction": "DESC",
			"limit":     strconv.Itoa(limit),
			"offset":    "0",
			"sortby":    "created_at",
		},
	}

	// Get the initial response
	resp, err := c.do("GET", url, doOptions)
	if err != nil {
		errChan <- err
		doneChan <- true
		return
	}
	defer resp.Body.Close()
	fr := &followResponse{}
	err = json.NewDecoder(resp.Body).Decode(&fr)
	if err != nil {
		errChan <- err
		doneChan <- true
		return
	}
	var follows []Follow
	follows = append(follows, fr.Follows...)
	dataChan <- follows
	offset := limit
	for {
		if offset > fr.Total {
			break
		}
		doOptions.params["offset"] = strconv.Itoa(offset)
		resp, err = c.do("GET", url, doOptions)
		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}
		defer resp.Body.Close()
		fr = &followResponse{}
		err = json.NewDecoder(resp.Body).Decode(&fr)
		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}
		follows = []Follow{}
		follows = append(follows, fr.Follows...)
		dataChan <- follows
		offset += limit
	}
	doneChan <- true
}

func (c *Client) GetUserChannelsFollowing(id int64) (*[]Follow, error) {
	user := strconv.FormatInt(id, 10)
	url := "users/" + user + "/follows/channels"
	doOptions := &doOptions{
		params: map[string]string{
			"direction": "DESC",
			"limit":     strconv.Itoa(limit),
			"offset":    "0",
			"sortby":    "created_at",
		},
	}

	// Get the initial response
	resp, err := c.do("GET", url, doOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fr := &followResponse{}
	err = json.NewDecoder(resp.Body).Decode(&fr)
	if err != nil {
		return nil, err
	}
	var follows []Follow
	follows = append(follows, fr.Follows...)
	offset := limit
	for {
		if offset > fr.Total {
			break
		}
		doOptions.params["offset"] = strconv.Itoa(offset)
		resp, err = c.do("GET", url, doOptions)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fr = &followResponse{}
		err = json.NewDecoder(resp.Body).Decode(&fr)
		if err != nil {
			return nil, err
		}
		follows = append(follows, fr.Follows...)
		offset += limit
	}
	return &follows, nil
}
