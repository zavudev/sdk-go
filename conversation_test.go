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

func TestConversationGet(t *testing.T) {
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
	_, err := client.Conversations.Get(context.TODO(), "conversationId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.List(context.TODO(), zavudev.ConversationListParams{
		Channel:  zavudev.ConversationListParamsChannelSMS,
		Cursor:   zavudev.String("cursor"),
		Limit:    zavudev.Int(100),
		Search:   zavudev.String("+56912345678"),
		SenderID: zavudev.String("senderId"),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversationListMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.ListMessages(
		context.TODO(),
		"conversationId",
		zavudev.ConversationListMessagesParams{
			Cursor: zavudev.String("cursor"),
			Limit:  zavudev.Int(100),
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

func TestConversationMarkAsRead(t *testing.T) {
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
	_, err := client.Conversations.MarkAsRead(context.TODO(), "conversationId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
