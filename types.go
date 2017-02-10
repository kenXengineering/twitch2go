package twitch2go

import "time"

// Channel Twitch Channel Data
type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLanguage          string    `json:"broadcaster_language"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Language                     string    `json:"language"`
	ID                           string    `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Logo                         string    `json:"logo"`
	VideoBanner                  string    `json:"video_banner"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        int64     `json:"views"`
	Followers                    int64     `json:"followers"`
}

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json":created_at"`
	Deleted   bool      `json:"deleted"`
	Emotes    []string  `json:"emotes"`
	Body      string    `json:"body"`
	User      User      `json:"user"`
}

// Follower data for twtich channel
type Follow struct {
	CreatedAt     time.Time         `json:"created_at"`
	Links         map[string]string `json:"_links"`
	Notifications bool              `json:"notifications"`
	User          User              `json:"user"`
	Channel       Channel           `json:"channel"`
}

type Followers struct {
	Total   int64    `json:"_total"`
	Cursor  string   `json:"_cursor"`
	Follows []Follow `json:"follows"`
}

type Subscription struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type Subscribers struct {
	Total         int64          `json:"_total"`
	Cursor        string         `json:"_cursor"`
	Subscriptions []Subscription `json:"subscriptions"`
}

type UserSearchResult struct {
	Total int64  `json:"_total"`
	Users []User `json:"users"`
}

// User Twitch User Data
type User struct {
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        string    `json:"logo"`
	ID          string    `json:"_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
}

type Editors struct {
	Users []User `json:"Users"`
}
