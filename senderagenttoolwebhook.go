// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
)

// SenderAgentToolWebhookService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentToolWebhookService] method instead.
type SenderAgentToolWebhookService struct {
	Options []option.RequestOption
}

// NewSenderAgentToolWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSenderAgentToolWebhookService(opts ...option.RequestOption) (r SenderAgentToolWebhookService) {
	r = SenderAgentToolWebhookService{}
	r.Options = opts
	return
}

// Generate a new signing secret for this tool. The previous one stops working on
// the next call, with no overlap, so update your endpoint first. The tool keeps
// its id, so flows that reference it by name are unaffected.
func (r *SenderAgentToolWebhookService) RotateSecret(ctx context.Context, toolID string, body SenderAgentToolWebhookRotateSecretParams, opts ...option.RequestOption) (res *WebhookSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s/webhook/secret", url.PathEscape(body.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type SenderAgentToolWebhookRotateSecretParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}
