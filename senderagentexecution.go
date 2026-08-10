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
	"github.com/zavudev/sdk-go/internal/apiquery"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/pagination"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// SenderAgentExecutionService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentExecutionService] method instead.
type SenderAgentExecutionService struct {
	Options []option.RequestOption
}

// NewSenderAgentExecutionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSenderAgentExecutionService(opts ...option.RequestOption) (r SenderAgentExecutionService) {
	r = SenderAgentExecutionService{}
	r.Options = opts
	return
}

// Fetch full details for one execution — including `errorMessage`, `errorCode`,
// and `responseText`. Use this to debug failures surfaced by the list endpoint.
func (r *SenderAgentExecutionService) Get(ctx context.Context, executionID string, query SenderAgentExecutionGetParams, opts ...option.RequestOption) (res *SenderAgentExecutionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/executions/%s", url.PathEscape(query.SenderID), url.PathEscape(executionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List recent agent executions with pagination.
//
// An execution is one inbound message answered by the agent, so this covers the
// messaging channels only. Voice calls are never listed here regardless of how
// many the agent handled. Use `GET /v1/calls` (and `GET /v1/calls/{callId}` for
// the transcript) for voice.
func (r *SenderAgentExecutionService) List(ctx context.Context, senderID string, query SenderAgentExecutionListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/executions", url.PathEscape(senderID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List recent agent executions with pagination.
//
// An execution is one inbound message answered by the agent, so this covers the
// messaging channels only. Voice calls are never listed here regardless of how
// many the agent handled. Use `GET /v1/calls` (and `GET /v1/calls/{callId}` for
// the transcript) for voice.
func (r *SenderAgentExecutionService) ListAutoPaging(ctx context.Context, senderID string, query SenderAgentExecutionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentExecution] {
	return pagination.NewCursorAutoPager(r.List(ctx, senderID, query, opts...))
}

type SenderAgentExecutionGetResponse struct {
	Execution AgentExecution `json:"execution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Execution   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentExecutionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentExecutionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentExecutionGetParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentExecutionListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Status of an agent execution.
	//
	// Any of "success", "error", "filtered", "rate_limited", "balance_insufficient".
	Status AgentExecutionStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentExecutionListParams]'s query parameters as
// `url.Values`.
func (r SenderAgentExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
