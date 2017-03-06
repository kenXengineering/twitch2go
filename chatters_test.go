package twitch2go

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetChatters(t *testing.T) {
	channel := "testChannel"
	jsonResponse := `{
  "_links": {},
  "chatter_count": 28,
  "chatters": {
    "moderators": [
      "colonelwill",
      "kaydeethree",
      "machtigemik3",
      "nightbot"
    ],
    "staff": [],
    "admins": [],
    "global_mods": [],
    "viewers": [
      "add7799",
      "biiggzgaming",
      "cars101507",
      "croccydile",
      "failcamper",
      "hsquishy",
      "hugorm77",
      "johnywishbone",
      "kilbil2",
      "kriell",
      "lexxreal",
      "lorddeson",
      "master_satate",
      "nardonius",
      "pascaldulieu",
      "proteanprojects",
      "psychoi3oy",
      "selterde",
      "sonyc88",
      "spamsac",
      "ssllqq",
      "talekouale",
      "tweeetythebird",
      "uxikoll"
    ]
  }
}`
	var expected ChatterResponse
	err := json.Unmarshal([]byte(jsonResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}
	fakeRT := &FakeRoundTripper{message: jsonResponse, status: http.StatusOK}
	client := newTestClient(fakeRT)
	chatters, err := client.GetChatters(channel)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(*chatters, expected) {
		t.Errorf("GetChatters(%q): Expected %#v.  Got %#v.", channel, expected, chatters)
	}
}
