package twitch2go

import (
	"encoding/json"
	"time"
)

// User Twitch User Data
type User struct {
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Links       map[string]string `json:"_links"`
	Logo        string            `json:"logo"`
	ID          int64             `json:"_id"`
	DisplayName string            `json:"display_name"`
	Bio         string            `json:"bio"`
}

func (c *Client) GetUser(user string) (*User, error) {

	resp, err := c.do("GET", "/users/"+user, doOptions{})
	if err != nil {
		return nil, err
	}

	u := &User{}
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
