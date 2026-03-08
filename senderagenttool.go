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

// SenderAgentToolService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentToolService] method instead.
type SenderAgentToolService struct {
	Options []option.RequestOption
}

// NewSenderAgentToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSenderAgentToolService(opts ...option.RequestOption) (r SenderAgentToolService) {
	r = SenderAgentToolService{}
	r.Options = opts
	return
}

// Create a new tool for an agent. Tools allow the agent to call external webhooks.
func (r *SenderAgentToolService) New(ctx context.Context, senderID string, body SenderAgentToolNewParams, opts ...option.RequestOption) (res *SenderAgentToolNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific tool.
func (r *SenderAgentToolService) Get(ctx context.Context, toolID string, query SenderAgentToolGetParams, opts ...option.RequestOption) (res *SenderAgentToolGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(query.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a tool.
func (r *SenderAgentToolService) Update(ctx context.Context, toolID string, params SenderAgentToolUpdateParams, opts ...option.RequestOption) (res *SenderAgentToolUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(params.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List tools for an agent.
func (r *SenderAgentToolService) List(ctx context.Context, senderID string, query SenderAgentToolListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentTool], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools", url.PathEscape(senderID))
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

// List tools for an agent.
func (r *SenderAgentToolService) ListAutoPaging(ctx context.Context, senderID string, query SenderAgentToolListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentTool] {
	return pagination.NewCursorAutoPager(r.List(ctx, senderID, query, opts...))
}

// Delete a tool.
func (r *SenderAgentToolService) Delete(ctx context.Context, toolID string, body SenderAgentToolDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(body.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Test a tool by triggering its webhook with test parameters.
func (r *SenderAgentToolService) Test(ctx context.Context, toolID string, params SenderAgentToolTestParams, opts ...option.RequestOption) (res *SenderAgentToolTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s/test", url.PathEscape(params.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AgentTool struct {
	ID        string    `json:"id" api:"required"`
	AgentID   string    `json:"agentId" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description for the LLM to understand when to use this tool.
	Description string              `json:"description" api:"required"`
	Enabled     bool                `json:"enabled" api:"required"`
	Name        string              `json:"name" api:"required"`
	Parameters  AgentToolParameters `json:"parameters" api:"required"`
	UpdatedAt   time.Time           `json:"updatedAt" api:"required" format:"date-time"`
	// HTTPS URL to call when the tool is executed.
	WebhookURL string `json:"webhookUrl" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AgentID     respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Enabled     respjson.Field
		Name        respjson.Field
		Parameters  respjson.Field
		UpdatedAt   respjson.Field
		WebhookURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTool) RawJSON() string { return r.JSON.raw }
func (r *AgentTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentToolParameters struct {
	Properties map[string]AgentToolParametersProperty `json:"properties" api:"required"`
	Required   []string                               `json:"required" api:"required"`
	// Any of "object".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentToolParameters) RawJSON() string { return r.JSON.raw }
func (r *AgentToolParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentToolParametersProperty struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentToolParametersProperty) RawJSON() string { return r.JSON.raw }
func (r *AgentToolParametersProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolNewResponse struct {
	Tool AgentTool `json:"tool" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tool        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolGetResponse struct {
	Tool AgentTool `json:"tool" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tool        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolUpdateResponse struct {
	Tool AgentTool `json:"tool" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tool        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolTestResponse struct {
	Scheduled bool `json:"scheduled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scheduled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolTestResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolNewParams struct {
	Description string                             `json:"description" api:"required"`
	Name        string                             `json:"name" api:"required"`
	Parameters  SenderAgentToolNewParamsParameters `json:"parameters,omitzero" api:"required"`
	// Must be HTTPS.
	WebhookURL string          `json:"webhookUrl" api:"required" format:"uri"`
	Enabled    param.Opt[bool] `json:"enabled,omitzero"`
	// Optional secret for webhook signature verification.
	WebhookSecret param.Opt[string] `json:"webhookSecret,omitzero"`
	paramObj
}

func (r SenderAgentToolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Properties, Required, Type are required.
type SenderAgentToolNewParamsParameters struct {
	Properties map[string]SenderAgentToolNewParamsParametersProperty `json:"properties,omitzero" api:"required"`
	Required   []string                                              `json:"required,omitzero" api:"required"`
	// Any of "object".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r SenderAgentToolNewParamsParameters) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolNewParamsParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolNewParamsParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentToolNewParamsParameters](
		"type", "object",
	)
}

type SenderAgentToolNewParamsParametersProperty struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SenderAgentToolNewParamsParametersProperty) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolNewParamsParametersProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolNewParamsParametersProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolGetParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentToolUpdateParams struct {
	SenderID      string                                `path:"senderId" api:"required" json:"-"`
	WebhookSecret param.Opt[string]                     `json:"webhookSecret,omitzero"`
	Description   param.Opt[string]                     `json:"description,omitzero"`
	Enabled       param.Opt[bool]                       `json:"enabled,omitzero"`
	Name          param.Opt[string]                     `json:"name,omitzero"`
	WebhookURL    param.Opt[string]                     `json:"webhookUrl,omitzero" format:"uri"`
	Parameters    SenderAgentToolUpdateParamsParameters `json:"parameters,omitzero"`
	paramObj
}

func (r SenderAgentToolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Properties, Required, Type are required.
type SenderAgentToolUpdateParamsParameters struct {
	Properties map[string]SenderAgentToolUpdateParamsParametersProperty `json:"properties,omitzero" api:"required"`
	Required   []string                                                 `json:"required,omitzero" api:"required"`
	// Any of "object".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r SenderAgentToolUpdateParamsParameters) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolUpdateParamsParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolUpdateParamsParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentToolUpdateParamsParameters](
		"type", "object",
	)
}

type SenderAgentToolUpdateParamsParametersProperty struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r SenderAgentToolUpdateParamsParametersProperty) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolUpdateParamsParametersProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolUpdateParamsParametersProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolListParams struct {
	Cursor  param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Enabled param.Opt[bool]   `query:"enabled,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentToolListParams]'s query parameters as
// `url.Values`.
func (r SenderAgentToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderAgentToolDeleteParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentToolTestParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	// Parameters to pass to the tool for testing.
	TestParams map[string]any `json:"testParams,omitzero" api:"required"`
	paramObj
}

func (r SenderAgentToolTestParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
