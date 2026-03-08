// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/zavudev/sdk-go"
	"github.com/zavudev/sdk-go/internal/testutil"
	"github.com/zavudev/sdk-go/option"
)

func TestSenderAgentToolNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Senders.Agent.Tools.New(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentToolNewParams{
			Description: "Get the status of a customer order",
			Name:        "get_order_status",
			Parameters: zavudev.SenderAgentToolNewParamsParameters{
				Properties: map[string]zavudev.SenderAgentToolNewParamsParametersProperty{
					"order_id": {
						Description: zavudev.String("The order ID to look up"),
						Type:        zavudev.String("string"),
					},
				},
				Required: []string{"order_id"},
				Type:     "object",
			},
			WebhookURL:    "https://api.example.com/webhooks/order-status",
			Enabled:       zavudev.Bool(true),
			WebhookSecret: zavudev.String("whsec_..."),
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentToolGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Senders.Agent.Tools.Get(
		context.TODO(),
		"toolId",
		zavudev.SenderAgentToolGetParams{
			SenderID: "senderId",
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentToolUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Senders.Agent.Tools.Update(
		context.TODO(),
		"toolId",
		zavudev.SenderAgentToolUpdateParams{
			SenderID:    "senderId",
			Description: zavudev.String("description"),
			Enabled:     zavudev.Bool(true),
			Name:        zavudev.String("name"),
			Parameters: zavudev.SenderAgentToolUpdateParamsParameters{
				Properties: map[string]zavudev.SenderAgentToolUpdateParamsParametersProperty{
					"foo": {
						Description: zavudev.String("description"),
						Type:        zavudev.String("type"),
					},
				},
				Required: []string{"string"},
				Type:     "object",
			},
			WebhookSecret: zavudev.String("webhookSecret"),
			WebhookURL:    zavudev.String("https://example.com"),
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentToolListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Senders.Agent.Tools.List(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentToolListParams{
			Cursor:  zavudev.String("cursor"),
			Enabled: zavudev.Bool(true),
			Limit:   zavudev.Int(100),
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentToolDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	err := client.Senders.Agent.Tools.Delete(
		context.TODO(),
		"toolId",
		zavudev.SenderAgentToolDeleteParams{
			SenderID: "senderId",
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentToolTest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Senders.Agent.Tools.Test(
		context.TODO(),
		"toolId",
		zavudev.SenderAgentToolTestParams{
			SenderID: "senderId",
			TestParams: map[string]any{
				"order_id": "bar",
			},
		},
	)
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
