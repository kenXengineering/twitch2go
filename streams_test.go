package twitch2go

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetStreamByChannel(t *testing.T) {
	channelID := "123456"
	jsonResponse := `{
  "stream": {
    "_id": 24487762416,
    "game": "For Honor",
    "community_id": "",
    "viewers": 5350,
    "video_height": 1080,
    "average_fps": 60.7488222167,
    "delay": 0,
    "created_at": "2017-02-10T19:00:30Z",
    "is_playlist": false,
    "preview": {
      "small": "https://static-cdn.jtvnw.net/previews-ttv/live_user_cohhcarnage-80x45.jpg",
      "medium": "https://static-cdn.jtvnw.net/previews-ttv/live_user_cohhcarnage-320x180.jpg",
      "large": "https://static-cdn.jtvnw.net/previews-ttv/live_user_cohhcarnage-640x360.jpg",
      "template": "https://static-cdn.jtvnw.net/previews-ttv/live_user_cohhcarnage-{width}x{height}.jpg"
    },
    "channel": {
      "mature": false,
      "status": "For Honor! \\o/ - Nobushi main! - Powered by Twitch Prime! - @CohhCarnage - !Prime - !Achievements - !4Year",
      "broadcaster_language": "en",
      "display_name": "CohhCarnage",
      "game": "For Honor",
      "language": "en",
      "_id": 26610234,
      "name": "cohhcarnage",
      "created_at": "2011-12-06T18:20:34Z",
      "updated_at": "2017-02-10T21:06:30Z",
      "partner": true,
      "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/cohhcarnage-profile_image-92dc409e41560047-300x300.png",
      "video_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/cohhcarnage-channel_offline_image-527afff8d1773bb0-1920x1080.png",
      "profile_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/cohhcarnage-profile_banner-bcb1b1b8e6194799-480.png",
      "profile_banner_background_color": null,
      "url": "https://www.twitch.tv/cohhcarnage",
      "views": 51587401,
      "followers": 730841
    }
  }
}`
	var expected StreamResponse
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	stream, err := client.GetStreamByChannel(channelID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*stream, expected) {
		t.Errorf("GetStreamByChannel(%q): Expected %#v.  Got %#v.", channelID, expected, stream)
	}

}

func TestGetFollowedStream(t *testing.T) {
	oauth := "fakeoauth"
	jsonResponse := `{
   "_total": 1,
   "streams": [
      {
         "_id": 23937446096,
         "average_fps": 60,
         "channel": {
               "_id": 121059319,
               "broadcaster_language": "en",
               "created_at": "2016-04-06T04:12:40Z",
               "display_name": "MOONMOON_OW",
               "followers": 251103,
               "game": "Overwatch",
               "language": "en",
               "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/moonmoon_ow-profile_image-0fe586039bb28259-300x300.png",
               "mature": true,
               "name": "moonmoon_ow",
               "partner": true,
               "profile_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/moonmoon_ow-profile_banner-13fbfa1ba07bcd8a-480.png",
               "profile_banner_background_color": null,
               "status": "KKona where my Darryl subs at KKona",
               "updated_at": "2016-12-15T19:34:46Z",
               "url": "https://www.twitch.tv/moonmoon_ow",
               "video_banner": "https://static-cdn.jtvnw.net/jtv_user_pictures/moonmoon_ow-channel_offline_image-2b3302e20384eee8-1920x1080.png",
               "views": 9865358
         },
         "created_at": "2016-12-15T14:55:49Z",
         "delay": 0,
         "game": "Overwatch",
         "is_playlist": false,
         "preview": {
               "large": "https://static-cdn.jtvnw.net/previews-ttv/live_user_moonmoon_ow-640x360.jpg",
               "medium": "https://static-cdn.jtvnw.net/previews-ttv/live_user_moonmoon_ow-320x180.jpg",
               "small": "https://static-cdn.jtvnw.net/previews-ttv/live_user_moonmoon_ow-80x45.jpg",
               "template": "https://static-cdn.jtvnw.net/previews-ttv/live_user_moonmoon_ow-{width}x{height}.jpg"
         },
         "video_height": 720,
         "viewers": 11211
      }
   ]
}
`
	var expected FollowedStream
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	stream, err := client.GetFollowedStreams(oauth)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*stream, expected) {
		t.Errorf("GetFollowedStream(%q): Expected %#v.  Got %#v.", oauth, expected, stream)
	}
}
