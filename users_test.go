package twitch2go

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetUserByOAuth(t *testing.T) {
	oauth := "fakeoauth"
	jsonResponse := `{
  "display_name": "Chosenken",
  "_id": "6391593",
  "name": "chosenken",
  "type": "user",
  "bio": "I play video games....some times...maybe.",
  "created_at": "2009-05-19T00:46:39.693935Z",
  "updated_at": "2017-02-08T21:09:45.334344Z",
  "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/chosenken-profile_image-6c79ed9956371448-300x300.jpeg",
  "email": "chosenken@gmail.com",
  "email_verified": true,
  "partnered": false,
  "twitter_connected": false,
  "notifications": {
    "push": true,
    "email": true
  }
}`
	var expected User
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	user, err := client.GetUserByOAuth(oauth)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*user, expected) {
		t.Errorf("GetUserByOAuth(%q): Expected %#v.  Got %#v.", oauth, expected, user)
	}
}

func TestGetUserFollows(t *testing.T) {
	userID := "123456"
	limit := 100
	offset := 0
	direction := ASC
	sortBy := CreatedAt
	jsonResponse := `
{
    "_total": 1,
    "follows": [
    {
      "created_at": "2015-02-27T05:07:13Z",
      "notifications": true,
      "channel": {
        "background": null,
        "banner": null,
        "broadcaster_language": "en",
        "display_name": "Colonelwill",
        "game": "Factorio",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/colonelwill-profile_image-eb352dbcdc14deac-300x300.jpeg",
        "mature": false,
        "status": "Guest Stream With Mojo - 1 RPM Base!",
        "partner": false,
        "url": "https://www.twitch.tv/colonelwill",
        "video_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/colonelwill-channel_offline_image-3ff63f95c54c33c9-640x360.jpeg",
        "_id": "54499069",
        "name": "colonelwill",
        "created_at": "2014-01-05T22:12:56Z",
        "updated_at": "2017-02-14T21:30:39Z",
        "delay": null,
        "followers": 5250,
        "profile_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/colonelwill-profile_banner-2034bc4b07be0bad-480.jpeg",
        "profile_banner_background_color": "#0723f6",
        "views": 152372,
        "language": "en"
      }}]
}`
	var expected Followers
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	follows, err := client.GetUserFollows(userID, limit, offset, direction, sortBy)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*follows, expected) {
		t.Errorf("GetUserFollows(%q, %q, %q, %q, %q): Expected %#v.  Got %#v.", userID, limit, offset, direction, sortBy, expected, follows)
	}
}
