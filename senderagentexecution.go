// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/zavudev/sdk-go/internal/apiquery"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/pagination"
	"github.com/zavudev/sdk-go/packages/param"
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

// List recent agent executions with pagination.
func (r *SenderAgentExecutionService) List(ctx context.Context, senderID string, query SenderAgentExecutionListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
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
func (r *SenderAgentExecutionService) ListAutoPaging(ctx context.Context, senderID string, query SenderAgentExecutionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentExecution] {
	return pagination.NewCursorAutoPager(r.List(ctx, senderID, query, opts...))
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
