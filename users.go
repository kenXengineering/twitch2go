package twitch2go

import (
	"encoding/json"
	"fmt"
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
	url := fmt.Sprintf("/users/%s/follows/channels", userID)
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
	return &result.Users, nil
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

// func (c *Client) UserFollowsStreaming(channelID int64, dataChan chan<- []Follow, doneChan chan<- bool, errChan chan<- error) {
// 	user := strconv.FormatInt(channelID, 10)
// 	url := "users/" + user + "/follows/channels"
// 	doOptions := &doOptions{
// 		params: map[string]string{
// 			"direction": "DESC",
// 			"limit":     strconv.FormatInt(limit, 10),
// 			"offset":    "0",
// 			"sortby":    "created_at",
// 		},
// 	}

// 	// Get the initial response
// 	resp, err := c.do("GET", url, doOptions)
// 	if err != nil {
// 		errChan <- err
// 		doneChan <- true
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fr := &followResponse{}
// 	err = json.NewDecoder(resp.Body).Decode(&fr)
// 	if err != nil {
// 		errChan <- err
// 		doneChan <- true
// 		return
// 	}
// 	var follows []Follow
// 	follows = append(follows, fr.Follows...)
// 	dataChan <- follows
// 	offset := limit
// 	for {
// 		if offset > fr.Total {
// 			break
// 		}
// 		doOptions.params["offset"] = strconv.FormatInt(offset, 10)
// 		resp, err = c.do("GET", url, doOptions)
// 		if err != nil {
// 			errChan <- err
// 			doneChan <- true
// 			return
// 		}
// 		defer resp.Body.Close()
// 		fr = &followResponse{}
// 		err = json.NewDecoder(resp.Body).Decode(&fr)
// 		if err != nil {
// 			errChan <- err
// 			doneChan <- true
// 			return
// 		}
// 		follows = []Follow{}
// 		follows = append(follows, fr.Follows...)
// 		dataChan <- follows
// 		offset += limit
// 	}
// 	doneChan <- true
// }

// func (c *Client) UserFollows(id int64) (*[]Follow, error) {
// 	user := strconv.FormatInt(id, 10)
// 	url := "users/" + user + "/follows/channels"
// 	doOptions := &doOptions{
// 		params: map[string]string{
// 			"direction": "DESC",
// 			"limit":     strconv.FormatInt(limit, 10),
// 			"offset":    "0",
// 			"sortby":    "created_at",
// 		},
// 	}

// 	// Get the initial response
// 	resp, err := c.do("GET", url, doOptions)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	fr := &followResponse{}
// 	err = json.NewDecoder(resp.Body).Decode(&fr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var follows []Follow
// 	follows = append(follows, fr.Follows...)
// 	offset := limit
// 	for {
// 		if offset > fr.Total {
// 			break
// 		}
// 		doOptions.params["offset"] = strconv.FormatInt(offset, 10)
// 		resp, err = c.do("GET", url, doOptions)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer resp.Body.Close()
// 		fr = &followResponse{}
// 		err = json.NewDecoder(resp.Body).Decode(&fr)
// 		if err != nil {
// 			return nil, err
// 		}
// 		follows = append(follows, fr.Follows...)
// 		offset += limit
// 	}
// 	return &follows, nil
// }
