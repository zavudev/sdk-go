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

func TestSenderAgentFlowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.Flows.New(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentFlowNewParams{
			Name: "Lead Capture",
			Steps: []zavudev.SenderAgentFlowNewParamsStep{{
				ID: "welcome",
				Config: map[string]any{
					"text": "bar",
				},
				Type:       "message",
				NextStepID: zavudev.String("ask_name"),
			}, {
				ID: "ask_name",
				Config: map[string]any{
					"variable": "bar",
					"prompt":   "bar",
				},
				Type:       "collect",
				NextStepID: zavudev.String("nextStepId"),
			}},
			Trigger: zavudev.SenderAgentFlowNewParamsTrigger{
				Type:     "keyword",
				Intent:   zavudev.String("intent"),
				Keywords: []string{"info", "pricing", "demo"},
			},
			Description: zavudev.String("Capture lead information"),
			Enabled:     zavudev.Bool(true),
			Priority:    zavudev.Int(0),
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

func TestSenderAgentFlowGet(t *testing.T) {
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
	_, err := client.Senders.Agent.Flows.Get(
		context.TODO(),
		"flowId",
		zavudev.SenderAgentFlowGetParams{
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

func TestSenderAgentFlowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.Flows.Update(
		context.TODO(),
		"flowId",
		zavudev.SenderAgentFlowUpdateParams{
			SenderID:    "senderId",
			Description: zavudev.String("description"),
			Enabled:     zavudev.Bool(true),
			Name:        zavudev.String("name"),
			Priority:    zavudev.Int(0),
			Steps: []zavudev.SenderAgentFlowUpdateParamsStep{{
				ID: "id",
				Config: map[string]any{
					"foo": "bar",
				},
				Type:       "message",
				NextStepID: zavudev.String("nextStepId"),
			}},
			Trigger: zavudev.SenderAgentFlowUpdateParamsTrigger{
				Type:     "keyword",
				Intent:   zavudev.String("intent"),
				Keywords: []string{"string"},
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

func TestSenderAgentFlowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.Flows.List(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentFlowListParams{
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

func TestSenderAgentFlowDelete(t *testing.T) {
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
	err := client.Senders.Agent.Flows.Delete(
		context.TODO(),
		"flowId",
		zavudev.SenderAgentFlowDeleteParams{
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

func TestSenderAgentFlowDuplicate(t *testing.T) {
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
	_, err := client.Senders.Agent.Flows.Duplicate(
		context.TODO(),
		"flowId",
		zavudev.SenderAgentFlowDuplicateParams{
			SenderID: "senderId",
			NewName:  "Lead Capture (Copy)",
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
