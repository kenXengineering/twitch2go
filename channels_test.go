package twitch2go

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetChannelByID(t *testing.T) {
	channelID := "6391593"
	jsonChannel := `{
  "mature": false,
  "status": null,
  "broadcaster_language": null,
  "display_name": "Chosenken",
  "game": "Grand Theft Auto V",
  "language": "en",
  "name": "chosenken",
  "created_at": "2009-05-19T00:46:39Z",
  "updated_at": "2017-02-08T21:09:45Z",
  "_id": "6391593",
  "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/chosenken-profile_image-6c79ed9956371448-300x300.jpeg",
  "video_banner": null,
  "profile_banner": null,
  "profile_banner_background_color": null,
  "partner": false,
  "url": "https://www.twitch.tv/chosenken",
  "views": 1,
  "followers": 0
}`
	var expected Channel
	err := json.Unmarshal([]byte(jsonChannel), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonChannel, status: http.StatusOK}
	client := newTestClient(fakeRT)
	channel, err := client.GetChannelByID(channelID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*channel, expected) {
		t.Errorf("GetChannelByID(%q): Expected %#v.  Got %#v.", channelID, expected, channel)
	}
}

func TestGetChannelByOAuth(t *testing.T) {
	oauth := "testoauth"
	jsonChannel := `{
  "mature": false,
  "status": null,
  "broadcaster_language": null,
  "display_name": "Chosenken",
  "game": "Grand Theft Auto V",
  "language": "en",
  "name": "chosenken",
  "created_at": "2009-05-19T00:46:39Z",
  "updated_at": "2017-02-08T21:09:45Z",
  "_id": "6391593",
  "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/chosenken-profile_image-6c79ed9956371448-300x300.jpeg",
  "video_banner": null,
  "profile_banner": null,
  "profile_banner_background_color": null,
  "partner": false,
  "url": "https://www.twitch.tv/chosenken",
  "views": 1,
  "followers": 0,
  "email": "chosenken@gmail.com",
  "stream_key": "live_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}`
	var expected Channel
	err := json.Unmarshal([]byte(jsonChannel), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonChannel, status: http.StatusOK}
	client := newTestClient(fakeRT)
	channel, err := client.GetChannelByOAuth(oauth)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*channel, expected) {
		t.Errorf("GetChannelByOAuth(%q): Expected %#v.  Got %#v.", oauth, expected, channel)
	}
}

func TestGetChannelEditors(t *testing.T) {
	id := "123456"
	authToken := "fakeauthtoken1"
	jsonResponse := `{
  "users": [
    {
      "display_name": "usbanksy",
      "type": "user",
      "bio": "",
      "created_at": "2012-06-15T05:45:17Z",
      "updated_at": "2017-01-26T17:00:23Z",
      "name": "usbanksy",
      "_id": "31337919",
      "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/usbanksy-profile_image-2cd6c687b6e45093-300x300.jpeg"
    }
  ]
}`
	var editors Editors
	err := json.Unmarshal([]byte(jsonResponse), &editors)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	users, err := client.GetChannelEditors(id, authToken)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(users, &editors.Users) {
		t.Errorf("GetChannelEditors(%q, %q): Exptected %#v.  Got %#v.", id, authToken, editors.Users, users)
	}
}

func TestGetChannelFollows(t *testing.T) {
	cursor := ""
	limit := 10
	id := "69222531"
	direction := ASC
	jsonResponse := `{
  "_total": 1194,
  "_cursor": "1486145103566302000",
  "follows": [
    {
      "created_at": "2017-02-05T16:04:10.599099Z",
      "notifications": false,
      "user": {
        "display_name": "Jananton",
        "_id": "93876011",
        "name": "jananton",
        "type": "user",
        "bio": null,
        "created_at": "2015-06-19T00:17:43.665655Z",
        "updated_at": "2016-10-28T23:02:03.025273Z",
        "logo": null
      }
    },
    {
      "created_at": "2017-02-04T20:59:37.248363Z",
      "notifications": false,
      "user": {
        "display_name": "n1ckA_",
        "_id": "146934573",
        "name": "n1cka_",
        "type": "user",
        "bio": null,
        "created_at": "2017-02-04T20:54:58.931904Z",
        "updated_at": "2017-02-09T05:53:49.67037Z",
        "logo": null
      }
    },
    {
      "created_at": "2017-02-04T13:56:25.963255Z",
      "notifications": false,
      "user": {
        "display_name": "KnightScorpius",
        "_id": "43513646",
        "name": "knightscorpius",
        "type": "user",
        "bio": "Soy un Padajuan de la fuerza. O algo. Suelo stremear cuanto puedo y de todo un poco, abierto a peticiones :P ",
        "created_at": "2013-05-13T17:07:36.848871Z",
        "updated_at": "2017-02-10T01:00:49.08248Z",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/knightscorpius-profile_image-f66992a00d332851-300x300.jpeg"
      }
    },
    {
      "created_at": "2017-02-04T07:21:17.024085Z",
      "notifications": true,
      "user": {
        "display_name": "can070124",
        "_id": "82207149",
        "name": "can070124",
        "type": "user",
        "bio": null,
        "created_at": "2015-02-09T02:07:11.39462Z",
        "updated_at": "2017-01-19T23:34:40.786643Z",
        "logo": "https://static-cdn.jtvnw.net/jtv_user_pictures/can070124-profile_image-4e5c6a93e432ee60-300x300.png"
      }
    },
    {
      "created_at": "2017-02-03T18:05:03.566302Z",
      "notifications": false,
      "user": {
        "display_name": "photoguy1876",
        "_id": "142221416",
        "name": "photoguy1876",
        "type": "user",
        "bio": null,
        "created_at": "2016-12-18T00:54:45.39435Z",
        "updated_at": "2016-12-18T00:54:47.008188Z",
        "logo": null
      }
    }
  ]
}`
	var expected Followers
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	follows, err := client.GetChannelFollows(id, cursor, limit, direction)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*follows, expected) {
		t.Errorf("GetChannelFollows(%q, %q, %q, %q): Exptected %#v.  Got %#v.", id, cursor, limit, direction, expected, follows)
	}
}

func TestGetChannelSubscribers(t *testing.T) {
	limit := 10
	id := "69222531"
	oauth := "fakeoauth"
	offset := 0
	direction := ASC
	jsonResponse := `{
   "_total": 1,
   "subscriptions": [{
      "_id": "67123294ed8305ce3a8ef09649d2237c5a300590",
      "created_at": "2014-05-19T23:38:53Z",
      "user": {
            "_id": "44322889",
            "bio": null,
            "created_at": "2014-01-28T00:50:38Z",
            "display_name": "dallas",
            "logo": null,
            "name": "dallas",
            "type": "staff",
            "updated_at": "2016-05-05T20:47:07Z"
      }
   }]
}`
	var expected Subscribers
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	subscribers, err := client.GetChannelSubscribers(id, oauth, limit, offset, direction)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*subscribers, expected) {
		t.Errorf("GetChannelSubscribers(%q, %q, %q, %q, %q): Exptected %#v.  Got %#v.", id, oauth, limit, offset, direction, expected, subscribers)
	}
}

func TestGetChannelSubscriberByUser(t *testing.T) {
	channelID := "123456"
	userID := "654321"
	oauth := "fakeoauth"
	jsonResponse := `{
   "_id": "180c61c3304a0460f771ebc53ed29007a9745abc",
   "created_at": "2016-11-30T00:37:00Z",
   "user": {
      "_id": "44322889",
      "bio": "",
      "created_at": "2015-04-20T06:18:47Z",
      "display_name": "dallas",
      "logo": null,
      "name": "dallas",
      "type": "user",
      "updated_at": "2016-12-10T09:30:47Z"
   }
}`
	var expected Subscription
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	subscriber, err := client.GetChannelSubscriberByUser(channelID, userID, oauth)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*subscriber, expected) {
		t.Errorf("GetChannelSubscriberByUser(%q, %q, %q):  Expected %#v.  Got %#v.", channelID, userID, oauth, expected, subscriber)
	}
}

func TestGetChannelVideos(t *testing.T) {
	channelID := "123456"
	limit := 10
	offset := 0
	broadcastType := ""
	language := ""
	sort := Views
	jsonResponse := `{
      "_total": 15,
  "videos": [
    {
      "title": "What happens when you don't obey the signals",
      "description": "Yeah, I cut a train in half. Brought to you by Champdoom, thanks for finding the time code for me.",
      "description_html": "Yeah, I cut a train in half. Brought to you by Champdoom, thanks for finding the time code for me.<br>",
      "broadcast_id": 16367731856,
      "broadcast_type": "highlight",
      "status": "recorded",
      "tag_list": "",
      "views": 110,
      "url": "https://www.twitch.tv/videos/17215738",
      "language": "en",
      "viewable": "public",
      "viewable_at": null,
      "animated_preview_url": "https://vod-storyboards.twitch.tv/f684a86731_rootnegative_16367731856_305882873/storyboards/17215738-strip-0.jpg",
      "_id": "v17215738",
      "recorded_at": "2015-09-14T09:50:41Z",
      "game": "Factorio",
      "length": 109,
      "preview": {
        "small": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-80x45.jpg",
        "medium": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-320x180.jpg",
        "large": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-640x360.jpg",
        "template": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-{width}x{height}.jpg"
      },
      "thumbnails": {
        "small": [
          {
            "url": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-80x45.jpg",
            "type": "generated"
          }
        ],
        "medium": [
          {
            "url": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-320x180.jpg",
            "type": "generated"
          }
        ],
        "large": [
          {
            "url": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-640x360.jpg",
            "type": "generated"
          }
        ],
        "template": [
          {
            "url": "https://static-cdn.jtvnw.net/s3_vods/f684a86731_rootnegative_16367731856_305882873//thumb/thumb17215738-{width}x{height}.jpg",
            "type": "generated"
          }
        ]
      },
      "fps": {
        "audio_only": 0,
        "chunked": 29.0552212841774,
        "high": 29.0481005452066,
        "low": 19.4368714842515,
        "medium": 29.0481005452066
      },
      "resolutions": {
        "chunked": "852x480",
        "high": "852x480",
        "low": "400x226",
        "medium": "638x360"
      },
      "created_at": "2015-09-14T04:06:48Z",
      "published_at": "2015-09-14T04:06:48Z",
      "channel": {
        "name": "rootnegative",
        "display_name": "rootnegative",
        "_id": "69222531"
      }
    }
    ]}`
	var expected Videos
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	videos, err := client.GetChannelVideos(channelID, limit, offset, broadcastType, language, sort)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*videos, expected) {
		t.Errorf("GetChannelVideos(%q, %q, %q, %q, %q, %q):  Expected %#v.  Got %#v.", channelID, limit, offset, broadcastType, language, sort, expected, videos)
	}
}
