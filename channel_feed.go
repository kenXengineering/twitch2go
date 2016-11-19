package twitch2go

import "time"

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json":created_at"`
	Deleted   bool      `json:"deleted"`
	Emotes    []string  `json:"emotes"`
	Body      string    `json:"body"`
	User      User      `json:"user"`
}
