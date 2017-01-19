package twitch2go

import (
	"encoding/json"
	"strconv"
	"time"
)

type UserSearchResult struct {
	Total int64        `json:"_total"`
	Users []SearchUser `json:"users"`
}

// User Twitch User Data
type User struct {
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
	ID          int64     `json:"_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
}

type SearchUser struct {
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
	ID          string    `json:"_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
}

var searchUserUrl = "search/users"

func (c *Client) SearchUsers(user string) (*[]User, error) {
	doOptions := &doOptions{
		params: map[string]string{
			"query": user,
		},
	}

	resp, err := c.do("GET", searchUserUrl, doOptions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := &UserSearchResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	// So right now twitch returns the user object with a string ID when searched,
	// and uses an INT everywhere else, gumble grumble
	users := []User{}
	for _, u := range result.Users {
		id, _ := strconv.ParseInt(u.ID, 10, 64)
		uu := User{
			Type:        u.Type,
			Name:        u.Name,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
			Logo:        u.Logo,
			ID:          id,
			DisplayName: u.DisplayName,
			Bio:         u.Bio,
		}
		users = append(users, uu)
	}
	return &users, nil
}

func (c *Client) SearchExactUser(user string) (*User, error) {
	users, err := c.SearchUsers(user)
	if err != nil {
		return nil, err
	}
	for _, u := range *users {
		if u.Name == user {
			return &u, nil
		}
	}
	return nil, nil
}
