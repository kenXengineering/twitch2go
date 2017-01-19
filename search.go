package twitch2go

import (
	"encoding/json"
)

type Result struct {
	Totals   int64     `json:"_total"`
	Channels []Channel `json:"channels"`
}

var searchURL = "search/channels"

// SearchChannels Searches for the given channel and returns the results
func (c *Client) SearchChannels(channel string) (*[]Channel, error) {
	doOptions := &doOptions{
		params: map[string]string{
			"query": channel,
		},
	}

	resp, err := c.do("GET", searchURL, doOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := &Result{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result.Channels, nil
}

// SearchExactChannel Searches for the given channel and returns
// the channel if the channel name is an exact match
func (c *Client) SearchExactChannel(channel string) (*Channel, error) {
	channels, err := c.SearchChannels(channel)
	if err != nil {
		return nil, err
	}
	for _, c := range *channels {
		if c.Name == channel {
			return &c, nil
		}
	}
	return nil, nil
}
