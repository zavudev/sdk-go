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

func TestMessageGet(t *testing.T) {
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
	_, err := client.Messages.Get(context.TODO(), "messageId")
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.List(context.TODO(), zavudev.MessageListParams{
		Channel: zavudev.ChannelAuto,
		Cursor:  zavudev.String("cursor"),
		Limit:   zavudev.Int(100),
		Status:  zavudev.MessageStatusQueued,
		To:      zavudev.String("to"),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageReactWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.React(
		context.TODO(),
		"messageId",
		zavudev.MessageReactParams{
			Emoji:      "👍",
			ZavuSender: zavudev.String("sender_12345"),
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

func TestMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Send(context.TODO(), zavudev.MessageSendParams{
		To:      "+56912345678",
		Channel: zavudev.ChannelAuto,
		Content: zavudev.MessageContentParam{
			Buttons: []zavudev.MessageContentButtonParam{{
				ID:    "id",
				Title: "title",
			}},
			Contacts: []zavudev.MessageContentContactParam{{
				Name:   zavudev.String("name"),
				Phones: []string{"string"},
			}},
			Emoji:            zavudev.String("emoji"),
			Filename:         zavudev.String("invoice.pdf"),
			Latitude:         zavudev.Float(0),
			ListButton:       zavudev.String("listButton"),
			LocationAddress:  zavudev.String("locationAddress"),
			LocationName:     zavudev.String("locationName"),
			Longitude:        zavudev.Float(0),
			MediaID:          zavudev.String("mediaId"),
			MediaURL:         zavudev.String("https://example.com/image.jpg"),
			MimeType:         zavudev.String("image/jpeg"),
			ReactToMessageID: zavudev.String("reactToMessageId"),
			Sections: []zavudev.MessageContentSectionParam{{
				Rows: []zavudev.MessageContentSectionRowParam{{
					ID:          "id",
					Title:       "title",
					Description: zavudev.String("description"),
				}},
				Title: "title",
			}},
			TemplateID: zavudev.String("templateId"),
			TemplateVariables: map[string]string{
				"1": "John",
				"2": "ORD-12345",
			},
		},
		FallbackEnabled: zavudev.Bool(true),
		HTMLBody:        zavudev.String("htmlBody"),
		IdempotencyKey:  zavudev.String("msg_01HZY4ZP7VQY2J3BRW7Z6G0QGE"),
		MessageType:     zavudev.MessageTypeText,
		Metadata: map[string]string{
			"foo": "string",
		},
		ReplyTo:       zavudev.String("support@example.com"),
		Subject:       zavudev.String("Your order confirmation"),
		Text:          zavudev.String("Your verification code is 123456"),
		VoiceLanguage: zavudev.String("es-ES"),
		ZavuSender:    zavudev.String("sender_12345"),
	})
	if err != nil {
		var apierr *zavudev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
