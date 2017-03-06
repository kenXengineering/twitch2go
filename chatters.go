package twitch2go

import "encoding/json"

// GetChatters returns the chatters for the given channel
func (c *Client) GetChatters(channel string) (*ChatterResponse, error) {
	resp, err := c.doChatters("GET", channel)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := &ChatterResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
