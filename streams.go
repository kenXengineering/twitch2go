package twitch2go

import (
	"encoding/json"

	"github.com/juju/errors"
)

// GetStreamByChannel returns a StreamResponse for the given channel
func (c *Client) GetStreamByChannel(channelID string) (*StreamResponse, error) {
	url := "/streams/" + channelID
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

// GetFollowedStreams returns a list of streams the user follows, based on user auth token.
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
