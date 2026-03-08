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

func TestSenderAgentExecutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Senders.Agent.Executions.List(
		context.TODO(),
		"senderId",
		zavudev.SenderAgentExecutionListParams{
			Cursor: zavudev.String("cursor"),
			Limit:  zavudev.Int(100),
			Status: zavudev.AgentExecutionStatusSuccess,
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
