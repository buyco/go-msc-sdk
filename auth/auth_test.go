package auth

import (
	mockhttp "github.com/buyco/go-msc-sdk/v2/auth/http/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		mockCtrl   *gomock.Controller
		client     *mockhttp.MockHTTPClient
		authClient *Client
		apiUrl     = "http://foo.bar"
		tenantID   = "1234"
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		client = mockhttp.NewMockHTTPClient(mockCtrl)
		authClient = NewClient(
			client,
			apiUrl,
			"azerty",
			"-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCT3O/qC06UUu54\nSGM3hZy6c0ipXIxCxjqS3RGF5q7+UFMwjjCzXHrTfZ6K7vxiEYZo56/OjXpd19K3\n7SwsFrlgHUpXBRez2F/dsgy8pvt0uazL5eUqGYE79uqN6RhX0JiMpoY6qsflNPee\nZBJOx7s++YnyQvS/mYZMHtsfGvY8liK3lT3fNtzZtGsF0uY3JlBgkxqQ7vdl6ohL\n3HKTbf7niRV8dqQZGyDxOfkBEdif66MkPcyLfQ3F2s8CtKrvGdwPDitjW9cxlH8k\ntv9gxACDFtAWnry8izCpgQNNISHDI/SZ8Rf80G0ql3BsPzF66cX5BSVpwmNP0kSE\n+4Dggfw7AgMBAAECggEAKH+U/oeGSD3Grw80jZp86Nx2hFyi1g8xL9R43jHmsCUU\nA/KOCDJGOfLoH6mBWuLt64G5t1ssrtNUFahSNukqcNbU66yrZ0jWSQRhVLJvoPLS\nDy6ya6t8qA3jBGdZkYPCpJNfpGXuRisRv0IteYJfGMqEK+SG4IuOKv8wiP57fvA9\nkcesFszv94ALLB/TjEkr4wRE0AiU65W+XFvgsgeI8br7FanNCF4V1610r3hfRbjk\na3X5JsUszXQWTnxYv5tG8SLxjTAjF7dsDE96wDXpZ41e6ZZb1k661tC+83ykuxlz\nD9waP1yu7xeYSFD693bkKfl9CXVRnNn5UT9BS4h2oQKBgQDDuCtgKdeeaIpXyoxl\nTBZDSbb+3Y4G3cUkVjFvhcwHvtVmmdegDLsTqEWapnE0WDRGBVz4bfFdi1Z5XTEO\nckYvsI6EaBW3Bo1koDfNXNxT2hHhm2fIVJOoEXvnXLmx7a2IbP4xyd3CHh0Cjx9Z\nAMOCqmTuqQPR1VB0EmbPqEsf8wKBgQDBZ3NAkJcjrjvADdn5xXycGrVRhUO3vSIy\nU7UJh4oiYgHMQWSISB3J1HICm9FK4rYp6mOZ1LhVg6rfm34z5FbCo+X8Q/TqLMKm\nbhWAAN88p8gcykOgn2FeW1PWewjvT1ArhbKpkLhydgIgDWztQ7py302Jr66jBKc0\n1WDxTRGMmQKBgAfCN0X6oqeO8V0FlIc3evJz66My2TyAch48pH0NSsdL013b32Zi\n2s+urgOxcW9nx7q237ahdR4GNgldnmI6OXoOf7fUAHhe9B/3Ef88HSfdzzOoW3bf\nk3LoLoc/b8UT7PsphvImVHorg27kiZOXqih15MZpQNOCp0vSpuy4eTHtAoGAN/YK\nCC2OPfnFOi4H21jEVJr5ygvIa1rjkTJdWNOKKaa4JHTrdO+BBwxcrNqPNZ7h3MEA\nbtt5Nu0xPSBN5Q/19r3b5yF2tWecLvH9cJtP/MoDgikYZlqXnujIGnBhRnVpmh5G\ncv/4Ds6MkN+xm/mT8ncghW17F5paE1SGh2uoX0kCgYEAuf5fYLyTw48unXO1hyFu\nfBB/+QZpfqIqZxmweJxD0zX35dGO0F2zs1H0Ob7fRP0z4dqrXzyROCzNpPq5VJG8\nw6LOfaPU795axQLDVfa0Hxp+p1aqn04HbatSePbDsWv2tP21ZJZRLcxnYMlaJPSu\nKiElKOVAmMLlRxie4xWNR10=\n-----END PRIVATE KEY-----\n",
			"666",
			"1",
			tenantID,
		)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("creates a valid custom claim", func() {
		Expect(customClaims{}.Valid()).ToNot(HaveOccurred())
	})

	It("creates X.509 certificate thumbprint", func() {
		Expect(authClient.x5t("foo:bar")).ToNot(HaveLen(0))
	})

	It("generates a random hexadecimal number", func() {
		Expect(authClient.randomHex(5)).ToNot(HaveLen(0))
	})

	It("generates headers", func() {
		headers := authClient.buildHeaders()
		Expect(headers["Accept"]).To(Equal("application/json"))
		Expect(headers["Content-Type"]).To(Equal("application/x-www-form-urlencoded"))
	})

	Context("When private key is valid", func() {
		It("generates request body", func() {
			body, err := authClient.buildParams()
			Expect(err).ToNot(HaveOccurred())
			Expect(body["tenant"]).ToNot(BeEmpty())
			Expect(body["client_id"]).ToNot(BeEmpty())
			Expect(body["scope"]).ToNot(BeEmpty())
			Expect(body["client_assertion_type"]).ToNot(BeEmpty())
			Expect(body["client_assertion"]).ToNot(BeEmpty())
			Expect(body["grant_type"]).ToNot(BeEmpty())
		})

		It("generates a signed client assertion", func() {
			signed, err := authClient.buildClientAssertion()
			Expect(err).ToNot(HaveOccurred())
			Expect(signed).ToNot(HaveLen(0))
		})
	})

	Context("When private key is invalid", func() {
		BeforeEach(func() {
			authClient = NewClient(
				client,
				"http://foo.bar",
				"azerty",
				"",
				"666",
				"1",
				"2",
			)
		})

		It("fails to generate request body", func() {
			_, err := authClient.buildParams()
			Expect(err).To(HaveOccurred())
		})

		It("fails to generates a signed client assertion", func() {
			_, err := authClient.buildClientAssertion()
			Expect(err).To(HaveOccurred())
		})
	})
})
