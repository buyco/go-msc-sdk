package client

import (
	http "github.com/imroc/req"
)

const (
	trackingPath = "/trackandtrace/v1/"
)

type Api struct {
	Url   string
	Token string
}

func (a *Api) EventsByBookingReference(bookingReference string) (string, error) {
	url := a.Url + trackingPath + "events"
	r, err := http.Get(url, http.Header{"Authorization": "Bearer " + a.Token}, http.QueryParam{"bookingReference": bookingReference})
	if err != nil {
		return "", err
	}

	e, err := r.ToString()
	if err != nil {
		return "", err
	}

	return e, nil
}
