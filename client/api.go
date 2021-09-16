package client

import (
	"github.com/buyco/go-msc-sdk/model"
	"net/http"
	"net/url"
)

const (
	trackingPath       = "/trackandtrace/v1/"
	trackingEventsPath = trackingPath + "events"
)

type Api struct {
	httpClient HttpClient
	url        string
	token      string
}

func NewApi(httpClient HttpClient, url string, token string) *Api {
	return &Api{
		httpClient: httpClient,
		url:        url,
		token:      token,
	}
}

func (a Api) EventsByBookingReference(bookingReference string) (*model.TrackingResponse, error) {
	fullUrl := a.url + trackingEventsPath

	// We use stdlib structs to ensure most clients are compatible
	var urlValues = url.Values{}
	urlValues.Set("bookingReference", bookingReference)
	var headers = http.Header{}
	a.setBearerHeader(headers)

	response, err := a.httpClient.Get(
		fullUrl,
		urlValues,
		headers,
	)
	if err != nil {
		return nil, err
	}

	var trackingResponse model.TrackingResponse
	err = response.ToJSON(&trackingResponse)
	if err != nil {
		return nil, err
	}

	return &trackingResponse, nil
}

func (a Api) setBearerHeader(headers http.Header) {
	headers.Set("Authorization", "Bearer "+a.token)
}
