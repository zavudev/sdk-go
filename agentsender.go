// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/zavudev/sdk-go/internal/apijson"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// AgentSenderService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentSenderService] method instead.
type AgentSenderService struct {
	Options []option.RequestOption
}

// NewAgentSenderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentSenderService(opts ...option.RequestOption) (r AgentSenderService) {
	r = AgentSenderService{}
	r.Options = opts
	return
}

// Make the agent answer on this sender. An agent can serve several senders; a
// sender answers with at most one agent, so connecting one that is already in use
// returns `400` naming the agent that holds it.
func (r *AgentSenderService) Connect(ctx context.Context, agentID string, body AgentSenderConnectParams, opts ...option.RequestOption) (res *AgentSenderConnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/senders", url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop the agent answering on this sender. The agent's primary sender is part of
// the agent itself and cannot be disconnected here.
func (r *AgentSenderService) Disconnect(ctx context.Context, senderID string, body AgentSenderDisconnectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AgentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return err
	}
	path := fmt.Sprintf("v1/agents/%s/senders/%s", url.PathEscape(body.AgentID), url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AgentSenderConnectResponse struct {
	// AI Agent configuration for a sender.
	Agent Agent `json:"agent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSenderConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentSenderConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSenderConnectParams struct {
	// Sender to connect.
	SenderID string `json:"senderId" api:"required"`
	paramObj
}

func (r AgentSenderConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentSenderConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSenderConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSenderDisconnectParams struct {
	AgentID string `path:"agentId" api:"required" json:"-"`
	paramObj
}
