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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), zavudev.AgentNewParams{
		Model:                  "model",
		Name:                   "name",
		Provider:               zavudev.AgentProviderOpenAI,
		SystemPrompt:           "systemPrompt",
		ContextWindowMessages:  zavudev.Int(1),
		IncludeContactMetadata: zavudev.Bool(true),
		MaxTokens:              zavudev.Int(1),
		Temperature:            zavudev.Float(0),
		TriggerOnChannels:      []string{"string"},
		TriggerOnMessageTypes:  []string{"string"},
		Voice: zavudev.AgentNewParamsVoice{
			Enabled:  true,
			Greeting: zavudev.String("Hi, thanks for calling Acme. How can I help you today?"),
			Greetings: map[string]string{
				"es": "Hola, soy Atlas. Preguntame lo que quieras.",
			},
			Interruptible:          zavudev.Bool(true),
			Language:               zavudev.String("en"),
			MaxCallDurationMinutes: zavudev.Int(1),
			MaxIdleSeconds:         zavudev.Int(5),
			Model:                  zavudev.String("openai/gpt-4o"),
			RecordCalls:            zavudev.Bool(true),
			SttModel:               zavudev.String("sttModel"),
			SttProvider:            zavudev.String("sttProvider"),
			TransferPhoneNumber:    zavudev.String("+14155551234"),
			TtsProvider:            zavudev.String("ttsProvider"),
			TtsVoiceID:             zavudev.String("aria"),
			VoicemailAction:        "hangup",
			VoicemailMessage:       zavudev.String("voicemailMessage"),
			VoiceSpeed:             zavudev.Float(0.5),
		},
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGet(t *testing.T) {
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
	_, err := client.Agents.Get(context.TODO(), "agentId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Update(
		context.TODO(),
		"agentId",
		zavudev.AgentUpdateParams{
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
			Voice: zavudev.AgentUpdateParamsVoice{
				Enabled:  true,
				Greeting: zavudev.String("Hi, thanks for calling Acme. How can I help you today?"),
				Greetings: map[string]string{
					"es": "Hola, soy Atlas. Preguntame lo que quieras.",
				},
				Interruptible:          zavudev.Bool(true),
				Language:               zavudev.String("en"),
				MaxCallDurationMinutes: zavudev.Int(1),
				MaxIdleSeconds:         zavudev.Int(5),
				Model:                  zavudev.String("openai/gpt-4o"),
				RecordCalls:            zavudev.Bool(true),
				SttModel:               zavudev.String("sttModel"),
				SttProvider:            zavudev.String("sttProvider"),
				TransferPhoneNumber:    zavudev.String("+14155551234"),
				TtsProvider:            zavudev.String("ttsProvider"),
				TtsVoiceID:             zavudev.String("aria"),
				VoicemailAction:        "hangup",
				VoicemailMessage:       zavudev.String("voicemailMessage"),
				VoiceSpeed:             zavudev.Float(0.5),
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

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.List(context.TODO(), zavudev.AgentListParams{
		Cursor: zavudev.String("cursor"),
		Limit:  zavudev.Int(100),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDelete(t *testing.T) {
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
	err := client.Agents.Delete(context.TODO(), "agentId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentListVoicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.ListVoices(context.TODO(), zavudev.AgentListVoicesParams{
		Language: zavudev.String("es"),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentTestWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Test(
		context.TODO(),
		"agentId",
		zavudev.AgentTestParams{
			Message:      "Where is order ORD-12345?",
			ExecuteTools: zavudev.Bool(true),
			History: []zavudev.AgentTestParamsHistory{{
				Content: "content",
				Role:    "user",
			}},
			UseKnowledgeBase: zavudev.Bool(true),
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
