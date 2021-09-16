package client

import (
	"github.com/buyco/go-msc-sdk/model"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/url"
)

var _ = Describe("API", func() {
	var (
		mockCtrl  *gomock.Controller
		client    *MockHttpClient
		response  *MockHttpReponse
		apiClient *Api
		apiUrl    = "http://foo.bar"
		token     = "1234"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		client = NewMockHttpClient(mockCtrl)
		response = NewMockHttpReponse(mockCtrl)
		apiClient = NewApi(client, apiUrl, token)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("creates a new api client", func() {
		Expect(apiClient).To(BeAssignableToTypeOf(&Api{}))
	})

	It("sets bearer header", func() {
		var headers = http.Header{}
		apiClient.setBearerHeader(headers)
		Expect(headers.Get("Authorization")).To(Equal("Bearer " + token))
	})

	It("sets bearer header only one time", func() {
		var headers = http.Header{}
		apiClient.setBearerHeader(headers)
		apiClient.setBearerHeader(headers)
		Expect(headers.Values("Authorization")).To(HaveLen(1))
	})

	Describe("Tracking", func() {
		var (
			urlValues  url.Values
			headers    http.Header
			bookingRef = "ABC123"
		)

		BeforeEach(func() {
			urlValues = url.Values{}
			headers = http.Header{}
			urlValues.Add("bookingReference", bookingRef)
			headers.Add("Authorization", "Bearer "+token)
		})

		It("calls MSC", func() {
			var structParsed = model.TrackingResponse{}
			response.EXPECT().ToJSON(&structParsed).Times(1).Return(nil)
			client.EXPECT().Get(apiUrl+trackingEventsPath, urlValues, headers).Times(1).Return(response, nil)
			res, err := apiClient.EventsByBookingReference(bookingRef)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
			Expect(res).To(BeAssignableToTypeOf(&model.TrackingResponse{}))
		})
	})
})
