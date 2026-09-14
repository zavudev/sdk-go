// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"encoding/json"
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
	Webhook SenderAgentToolWebhookService
}

// NewSenderAgentToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSenderAgentToolService(opts ...option.RequestOption) (r SenderAgentToolService) {
	r = SenderAgentToolService{}
	r.Options = opts
	r.Webhook = NewSenderAgentToolWebhookService(opts...)
	return
}

// Create a new tool for an agent. Tools allow the agent to call external webhooks.
func (r *SenderAgentToolService) New(ctx context.Context, senderID string, body SenderAgentToolNewParams, opts ...option.RequestOption) (res *SenderAgentToolNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a specific tool.
func (r *SenderAgentToolService) Get(ctx context.Context, toolID string, query SenderAgentToolGetParams, opts ...option.RequestOption) (res *SenderAgentToolGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(query.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a tool.
func (r *SenderAgentToolService) Update(ctx context.Context, toolID string, params SenderAgentToolUpdateParams, opts ...option.RequestOption) (res *SenderAgentToolUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(params.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List tools for an agent.
func (r *SenderAgentToolService) List(ctx context.Context, senderID string, query SenderAgentToolListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentTool], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
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
		return err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s", url.PathEscape(body.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Recent runs of this tool triggered from the test endpoint, newest first. Covers
// manual tests only: a tool called by an agent during a real conversation is not
// recorded here.
func (r *SenderAgentToolService) ListTestRuns(ctx context.Context, toolID string, params SenderAgentToolListTestRunsParams, opts ...option.RequestOption) (res *SenderAgentToolListTestRunsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s/test-runs", url.PathEscape(params.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Run a tool with the parameters you supply and return what it answered.
//
// The call is synchronous: the response carries the tool's status, body, and
// duration, so a green result is evidence the tool ran rather than evidence it was
// accepted. Each run is also recorded and readable afterwards via
// `GET /v1/senders/{senderId}/agent/tools/{toolId}/test-runs`.
//
// A tool that answers with an error is reported as a run with `success: false` —
// the endpoint itself still returns 200. This fires the tool's real webhook, so a
// test has whatever side effects the tool has.
func (r *SenderAgentToolService) Test(ctx context.Context, toolID string, params SenderAgentToolTestParams, opts ...option.RequestOption) (res *SenderAgentToolTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	if toolID == "" {
		err = errors.New("missing required toolId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/tools/%s/test", url.PathEscape(params.SenderID), url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type AgentTool struct {
	ID        string    `json:"id" api:"required"`
	AgentID   string    `json:"agentId" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description for the LLM to understand when to use this tool.
	Description string             `json:"description" api:"required"`
	Enabled     bool               `json:"enabled" api:"required"`
	Name        string             `json:"name" api:"required"`
	Parameters  ToolParametersResp `json:"parameters" api:"required"`
	UpdatedAt   time.Time          `json:"updatedAt" api:"required" format:"date-time"`
	// HTTPS URL to call when the tool is executed.
	WebhookURL string `json:"webhookUrl" api:"required" format:"uri"`
	// Signing secret for this tool's webhook. **Returned only when the tool is
	// created**, never on a later read.
	//
	// Zavu generates one if you do not supply it, and signs every call to this tool
	// with it: `X-Zavu-Signature: <hex>`, the HMAC-SHA256 of the request body. Verify
	// it before trusting the call. Lost it? Rotate with
	// `POST /v1/senders/{senderId}/agent/tools/{toolId}/webhook/secret`.
	WebhookSecret string `json:"webhookSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AgentID       respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Enabled       respjson.Field
		Name          respjson.Field
		Parameters    respjson.Field
		UpdatedAt     respjson.Field
		WebhookURL    respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTool) RawJSON() string { return r.JSON.raw }
func (r *AgentTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolParametersResp struct {
	Properties map[string]ToolParametersPropertyResp `json:"properties" api:"required"`
	Required   []string                              `json:"required" api:"required"`
	// Any of "object".
	Type ToolParametersType `json:"type" api:"required"`
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
func (r ToolParametersResp) RawJSON() string { return r.JSON.raw }
func (r *ToolParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolParametersResp to a ToolParameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolParameters.Overrides()
func (r ToolParametersResp) ToParam() ToolParameters {
	return param.Override[ToolParameters](json.RawMessage(r.RawJSON()))
}

type ToolParametersPropertyResp struct {
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
func (r ToolParametersPropertyResp) RawJSON() string { return r.JSON.raw }
func (r *ToolParametersPropertyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolParametersType string

const (
	ToolParametersTypeObject ToolParametersType = "object"
)

// The properties Properties, Required, Type are required.
type ToolParameters struct {
	Properties map[string]ToolParametersProperty `json:"properties,omitzero" api:"required"`
	Required   []string                          `json:"required,omitzero" api:"required"`
	// Any of "object".
	Type ToolParametersType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolParameters) MarshalJSON() (data []byte, err error) {
	type shadow ToolParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolParametersProperty struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ToolParametersProperty) MarshalJSON() (data []byte, err error) {
	type shadow ToolParametersProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolParametersProperty) UnmarshalJSON(data []byte) error {
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

type SenderAgentToolListTestRunsResponse struct {
	Items []SenderAgentToolListTestRunsResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolListTestRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolListTestRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One run of a tool triggered from the test endpoint. Recorded so a test is
// verifiable after the fact rather than only visible in the response.
type SenderAgentToolListTestRunsResponseItem struct {
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	DurationMs int64     `json:"durationMs" api:"required"`
	// Whether the tool returned without error. A tool that answered with a non-2xx
	// status is a failed run, not an error of this endpoint.
	Success bool   `json:"success" api:"required"`
	ToolID  string `json:"toolId" api:"required"`
	// Why the run failed, when it did.
	Error string `json:"error" api:"nullable"`
	// The parameters the tool was called with.
	Params map[string]any `json:"params"`
	// The tool's response body, truncated.
	Response string `json:"response" api:"nullable"`
	// HTTP status the tool's webhook returned. Absent for tools that do not go over
	// HTTP.
	StatusCode int64 `json:"statusCode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DurationMs  respjson.Field
		Success     respjson.Field
		ToolID      respjson.Field
		Error       respjson.Field
		Params      respjson.Field
		Response    respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolListTestRunsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolListTestRunsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolTestResponse struct {
	// One run of a tool triggered from the test endpoint. Recorded so a test is
	// verifiable after the fact rather than only visible in the response.
	Run SenderAgentToolTestResponseRun `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolTestResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One run of a tool triggered from the test endpoint. Recorded so a test is
// verifiable after the fact rather than only visible in the response.
type SenderAgentToolTestResponseRun struct {
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	DurationMs int64     `json:"durationMs" api:"required"`
	// Whether the tool returned without error. A tool that answered with a non-2xx
	// status is a failed run, not an error of this endpoint.
	Success bool   `json:"success" api:"required"`
	ToolID  string `json:"toolId" api:"required"`
	// Why the run failed, when it did.
	Error string `json:"error" api:"nullable"`
	// The parameters the tool was called with.
	Params map[string]any `json:"params"`
	// The tool's response body, truncated.
	Response string `json:"response" api:"nullable"`
	// HTTP status the tool's webhook returned. Absent for tools that do not go over
	// HTTP.
	StatusCode int64 `json:"statusCode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DurationMs  respjson.Field
		Success     respjson.Field
		ToolID      respjson.Field
		Error       respjson.Field
		Params      respjson.Field
		Response    respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentToolTestResponseRun) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentToolTestResponseRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentToolNewParams struct {
	Description string         `json:"description" api:"required"`
	Name        string         `json:"name" api:"required"`
	Parameters  ToolParameters `json:"parameters,omitzero" api:"required"`
	// Must be HTTPS.
	WebhookURL string          `json:"webhookUrl" api:"required" format:"uri"`
	Enabled    param.Opt[bool] `json:"enabled,omitzero"`
	// Signing secret for the webhook. Optional: Zavu generates one when omitted and
	// returns it on this response only. Supply your own if you already have a secret
	// you want reused.
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

type SenderAgentToolGetParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentToolUpdateParams struct {
	SenderID      string            `path:"senderId" api:"required" json:"-"`
	WebhookSecret param.Opt[string] `json:"webhookSecret,omitzero"`
	Description   param.Opt[string] `json:"description,omitzero"`
	Enabled       param.Opt[bool]   `json:"enabled,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	WebhookURL    param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	Parameters    ToolParameters    `json:"parameters,omitzero"`
	paramObj
}

func (r SenderAgentToolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentToolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentToolUpdateParams) UnmarshalJSON(data []byte) error {
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

type SenderAgentToolListTestRunsParams struct {
	SenderID string           `path:"senderId" api:"required" json:"-"`
	Limit    param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentToolListTestRunsParams]'s query parameters as
// `url.Values`.
func (r SenderAgentToolListTestRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
