// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/zavudev/sdk-go/internal/apijson"
	"github.com/zavudev/sdk-go/internal/apiquery"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/pagination"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// SenderAgentFlowService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentFlowService] method instead.
type SenderAgentFlowService struct {
	Options []option.RequestOption
}

// NewSenderAgentFlowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSenderAgentFlowService(opts ...option.RequestOption) (r SenderAgentFlowService) {
	r = SenderAgentFlowService{}
	r.Options = opts
	return
}

// Create a new flow for an agent.
func (r *SenderAgentFlowService) New(ctx context.Context, senderID string, body SenderAgentFlowNewParams, opts ...option.RequestOption) (res *SenderAgentFlowNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific flow.
func (r *SenderAgentFlowService) Get(ctx context.Context, flowID string, query SenderAgentFlowGetParams, opts ...option.RequestOption) (res *SenderAgentFlowGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows/%s", url.PathEscape(query.SenderID), url.PathEscape(flowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a flow.
func (r *SenderAgentFlowService) Update(ctx context.Context, flowID string, params SenderAgentFlowUpdateParams, opts ...option.RequestOption) (res *SenderAgentFlowUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows/%s", url.PathEscape(params.SenderID), url.PathEscape(flowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List flows for an agent.
func (r *SenderAgentFlowService) List(ctx context.Context, senderID string, query SenderAgentFlowListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentFlow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows", url.PathEscape(senderID))
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

// List flows for an agent.
func (r *SenderAgentFlowService) ListAutoPaging(ctx context.Context, senderID string, query SenderAgentFlowListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentFlow] {
	return pagination.NewCursorAutoPager(r.List(ctx, senderID, query, opts...))
}

// Delete a flow. Cannot delete flows with active sessions.
func (r *SenderAgentFlowService) Delete(ctx context.Context, flowID string, body SenderAgentFlowDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows/%s", url.PathEscape(body.SenderID), url.PathEscape(flowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a copy of an existing flow with a new name.
func (r *SenderAgentFlowService) Duplicate(ctx context.Context, flowID string, params SenderAgentFlowDuplicateParams, opts ...option.RequestOption) (res *SenderAgentFlowDuplicateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if flowID == "" {
		err = errors.New("missing required flowId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/flows/%s/duplicate", url.PathEscape(params.SenderID), url.PathEscape(flowID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AgentFlow struct {
	ID        string    `json:"id" api:"required"`
	AgentID   string    `json:"agentId" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Enabled   bool      `json:"enabled" api:"required"`
	Name      string    `json:"name" api:"required"`
	// Priority when multiple flows match (higher = more priority).
	Priority    int64            `json:"priority" api:"required"`
	Steps       []AgentFlowStep  `json:"steps" api:"required"`
	Trigger     AgentFlowTrigger `json:"trigger" api:"required"`
	UpdatedAt   time.Time        `json:"updatedAt" api:"required" format:"date-time"`
	Description string           `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AgentID     respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Priority    respjson.Field
		Steps       respjson.Field
		Trigger     respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFlow) RawJSON() string { return r.JSON.raw }
func (r *AgentFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFlowStep struct {
	// Unique step identifier.
	ID string `json:"id" api:"required"`
	// Step configuration (varies by type).
	Config map[string]any `json:"config" api:"required"`
	// Type of flow step.
	//
	// Any of "message", "collect", "condition", "tool", "llm", "transfer".
	Type string `json:"type" api:"required"`
	// ID of the next step to execute.
	NextStepID string `json:"nextStepId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		Type        respjson.Field
		NextStepID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFlowStep) RawJSON() string { return r.JSON.raw }
func (r *AgentFlowStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentFlowTrigger struct {
	// Type of trigger for a flow.
	//
	// Any of "keyword", "intent", "always", "manual".
	Type string `json:"type" api:"required"`
	// Intent that triggers the flow (for intent type).
	Intent string `json:"intent"`
	// Keywords that trigger the flow (for keyword type).
	Keywords []string `json:"keywords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Intent      respjson.Field
		Keywords    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFlowTrigger) RawJSON() string { return r.JSON.raw }
func (r *AgentFlowTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentFlowNewResponse struct {
	Flow AgentFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentFlowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentFlowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentFlowGetResponse struct {
	Flow AgentFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentFlowGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentFlowGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentFlowUpdateResponse struct {
	Flow AgentFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentFlowUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentFlowUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentFlowDuplicateResponse struct {
	Flow AgentFlow `json:"flow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flow        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentFlowDuplicateResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentFlowDuplicateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentFlowNewParams struct {
	Name        string                          `json:"name" api:"required"`
	Steps       []SenderAgentFlowNewParamsStep  `json:"steps,omitzero" api:"required"`
	Trigger     SenderAgentFlowNewParamsTrigger `json:"trigger,omitzero" api:"required"`
	Description param.Opt[string]               `json:"description,omitzero"`
	Enabled     param.Opt[bool]                 `json:"enabled,omitzero"`
	Priority    param.Opt[int64]                `json:"priority,omitzero"`
	paramObj
}

func (r SenderAgentFlowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Config, Type are required.
type SenderAgentFlowNewParamsStep struct {
	// Unique step identifier.
	ID string `json:"id" api:"required"`
	// Step configuration (varies by type).
	Config map[string]any `json:"config,omitzero" api:"required"`
	// Type of flow step.
	//
	// Any of "message", "collect", "condition", "tool", "llm", "transfer".
	Type string `json:"type,omitzero" api:"required"`
	// ID of the next step to execute.
	NextStepID param.Opt[string] `json:"nextStepId,omitzero"`
	paramObj
}

func (r SenderAgentFlowNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentFlowNewParamsStep](
		"type", "message", "collect", "condition", "tool", "llm", "transfer",
	)
}

// The property Type is required.
type SenderAgentFlowNewParamsTrigger struct {
	// Type of trigger for a flow.
	//
	// Any of "keyword", "intent", "always", "manual".
	Type string `json:"type,omitzero" api:"required"`
	// Intent that triggers the flow (for intent type).
	Intent param.Opt[string] `json:"intent,omitzero"`
	// Keywords that trigger the flow (for keyword type).
	Keywords []string `json:"keywords,omitzero"`
	paramObj
}

func (r SenderAgentFlowNewParamsTrigger) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowNewParamsTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowNewParamsTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentFlowNewParamsTrigger](
		"type", "keyword", "intent", "always", "manual",
	)
}

type SenderAgentFlowGetParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentFlowUpdateParams struct {
	SenderID    string                             `path:"senderId" api:"required" json:"-"`
	Description param.Opt[string]                  `json:"description,omitzero"`
	Enabled     param.Opt[bool]                    `json:"enabled,omitzero"`
	Name        param.Opt[string]                  `json:"name,omitzero"`
	Priority    param.Opt[int64]                   `json:"priority,omitzero"`
	Steps       []SenderAgentFlowUpdateParamsStep  `json:"steps,omitzero"`
	Trigger     SenderAgentFlowUpdateParamsTrigger `json:"trigger,omitzero"`
	paramObj
}

func (r SenderAgentFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Config, Type are required.
type SenderAgentFlowUpdateParamsStep struct {
	// Unique step identifier.
	ID string `json:"id" api:"required"`
	// Step configuration (varies by type).
	Config map[string]any `json:"config,omitzero" api:"required"`
	// Type of flow step.
	//
	// Any of "message", "collect", "condition", "tool", "llm", "transfer".
	Type string `json:"type,omitzero" api:"required"`
	// ID of the next step to execute.
	NextStepID param.Opt[string] `json:"nextStepId,omitzero"`
	paramObj
}

func (r SenderAgentFlowUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentFlowUpdateParamsStep](
		"type", "message", "collect", "condition", "tool", "llm", "transfer",
	)
}

// The property Type is required.
type SenderAgentFlowUpdateParamsTrigger struct {
	// Type of trigger for a flow.
	//
	// Any of "keyword", "intent", "always", "manual".
	Type string `json:"type,omitzero" api:"required"`
	// Intent that triggers the flow (for intent type).
	Intent param.Opt[string] `json:"intent,omitzero"`
	// Keywords that trigger the flow (for keyword type).
	Keywords []string `json:"keywords,omitzero"`
	paramObj
}

func (r SenderAgentFlowUpdateParamsTrigger) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowUpdateParamsTrigger
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowUpdateParamsTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentFlowUpdateParamsTrigger](
		"type", "keyword", "intent", "always", "manual",
	)
}

type SenderAgentFlowListParams struct {
	Cursor  param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Enabled param.Opt[bool]   `query:"enabled,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentFlowListParams]'s query parameters as
// `url.Values`.
func (r SenderAgentFlowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderAgentFlowDeleteParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentFlowDuplicateParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	NewName  string `json:"newName" api:"required"`
	paramObj
}

func (r SenderAgentFlowDuplicateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentFlowDuplicateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentFlowDuplicateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
