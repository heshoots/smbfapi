package facebook

import (
	"encoding/json"
	"github.com/heshoots/smbfapi/gen/models"
	"io/ioutil"
	"net/http"
)

type FacebookEvents []*models.Event

type GetEventsResponse struct {
	Data FacebookEvents `json:"data"`
}

func GetEvents(accessToken string) ([]*models.Event, error) {
	resp, err := http.Get("https://graph.facebook.com/v3.1/153549318022206/events?access_token=" + accessToken + "&debug=all&fields=name%2Cid%2Cattending_count%2Cstart_time%2Cend_time&format=json&method=get&pretty=0&suppress_http_code=1")
	if err != nil {
		return nil, err
	}
	var data GetEventsResponse
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(out, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}

func GetEvent(id string, accessToken string) (e *models.Event, err error) {
	resp, err := http.Get("https://graph.facebook.com/v3.1/" + id + "?access_token=" + accessToken + "&debug=all&fields=name%2Cid%2Cattending_count%2Cstart_time%2Cend_time&format=json&method=get&pretty=0&suppress_http_code=1")
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(out, &e)
	if err != nil {
		return nil, err
	}
	return e, nil

}
