package twitch2go

import "encoding/json"

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
