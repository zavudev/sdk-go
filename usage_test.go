// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev_test

import (
	"context"
	"os"
	"testing"

	"github.com/zavudev/sdk-go"
	"github.com/zavudev/sdk-go/internal/testutil"
	"github.com/zavudev/sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := zavudev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Mock server tests are disabled")
	messageResponse, err := client.Messages.Send(context.TODO(), zavudev.MessageSendParams{
		To:   "+14155551234",
		Text: zavudev.String("Hello from Zavu!"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", messageResponse.Message)
}
