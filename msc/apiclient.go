package msc

import (
	http "github.com/imroc/req"
)

const (
	TRACKING_PATH = "/tracking/v1/"
)

type ApiClient struct {
	Url   string
	Token string
}

func (c *ApiClient) Events(bookingReference string) ([]Event, error) {
	url := c.Url + TRACKING_PATH + "Events"
	r, err := http.Get(url, http.Header{"Authorization": "Bearer " + c.Token}, http.QueryParam{"bookingReference": bookingReference})
	if err != nil {
		return nil, err
	}

	var events []Event
	err = r.ToJSON(&events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
