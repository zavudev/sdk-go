// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/zavudev-go"
	"github.com/stainless-sdks/zavudev-go/internal/testutil"
	"github.com/stainless-sdks/zavudev-go/option"
)

func TestSenderAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.New(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentNewParams{
			Model:                  "gpt-4o-mini",
			Name:                   "Customer Support",
			Provider:               zavudev.AgentProviderOpenAI,
			SystemPrompt:           "You are a helpful customer support agent. Be friendly and concise.",
			APIKey:                 zavudev.String("sk-..."),
			ContextWindowMessages:  zavudev.Int(1),
			IncludeContactMetadata: zavudev.Bool(true),
			MaxTokens:              zavudev.Int(1),
			Temperature:            zavudev.Float(0),
			TriggerOnChannels:      []string{"string"},
			TriggerOnMessageTypes:  []string{"string"},
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

func TestSenderAgentGet(t *testing.T) {
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
	_, err := client.Senders.Agent.Get(context.TODO(), "senderId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.Update(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentUpdateParams{
			APIKey:                 zavudev.String("apiKey"),
			ContextWindowMessages:  zavudev.Int(1),
			Enabled:                zavudev.Bool(true),
			IncludeContactMetadata: zavudev.Bool(true),
			MaxTokens:              zavudev.Int(1),
			Model:                  zavudev.String("model"),
			Name:                   zavudev.String("name"),
			Provider:               zavudev.AgentProviderOpenAI,
			SystemPrompt:           zavudev.String("systemPrompt"),
			Temperature:            zavudev.Float(0),
			TriggerOnChannels:      []string{"string"},
			TriggerOnMessageTypes:  []string{"string"},
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

func TestSenderAgentDelete(t *testing.T) {
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
	err := client.Senders.Agent.Delete(context.TODO(), "senderId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSenderAgentStats(t *testing.T) {
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
	_, err := client.Senders.Agent.Stats(context.TODO(), "senderId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
