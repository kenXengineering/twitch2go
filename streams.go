package twitch2go

import (
	"encoding/json"
	"fmt"

	"github.com/juju/errors"
)

func (c *Client) GetStreamByChannel(channelID string) (*StreamResponse, error) {
	url := fmt.Sprintf("/streams/%s", channelID)
	// Do the request
	resp, err := c.do("GET", url, &doOptions{})
	if err != nil {
		return nil, errors.Annotate(err, "GetStreamByChannel")
	}
	defer resp.Body.Close()
	stream := &StreamResponse{}
	err = json.NewDecoder(resp.Body).Decode(&stream)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding JSON")
	}
	return stream, nil
}

func (c *Client) GetFollowedStreams(oauth string) (*FollowedStream, error) {
	url := "/streams/followed"
	opts := &doOptions{
		oauth: oauth,
	}
	// Do the request
	resp, err := c.do("GET", url, opts)
	if err != nil {
		return nil, errors.Annotate(err, "GetFollowedStreams")
	}
	defer resp.Body.Close()
	fs := &FollowedStream{}
	err = json.NewDecoder(resp.Body).Decode(&fs)
	if err != nil {
		return nil, errors.Annotate(err, "Error decoding json")
	}
	return fs, nil
}
