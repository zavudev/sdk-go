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
	"github.com/zavudev/sdk-go/packages/respjson"
)

// AgentTemplateService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentTemplateService] method instead.
type AgentTemplateService struct {
	Options []option.RequestOption
}

// NewAgentTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentTemplateService(opts ...option.RequestOption) (r AgentTemplateService) {
	r = AgentTemplateService{}
	r.Options = opts
	return
}

// Fetch a single factory agent fully rendered: the function files to scaffold (an
// `index.ts` that declares the agent with `defineAgent` and its skills with
// `defineTool`) plus the secrets it needs. This is what
// `npx zavudev agents pull <id>` writes to disk before `npx zavudev deploy`.
func (r *AgentTemplateService) Get(ctx context.Context, templateID string, opts ...option.RequestOption) (res *AgentTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agent-templates/%s", url.PathEscape(templateID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the factory agents available to scaffold with `npx zavudev agents pull`.
// Each entry is a ready-made voice or text agent (system prompt, skills, and — for
// voice agents — a co-located voice config).
func (r *AgentTemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *AgentTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agent-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentTemplateGetResponse struct {
	// A fully rendered factory agent: the function files to scaffold plus the secrets
	// it needs. Returned by GET /v1/agent-templates/{templateId} and consumed by
	// `npx zavudev agents pull`.
	Template AgentTemplateGetResponseTemplate `json:"template" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fully rendered factory agent: the function files to scaffold plus the secrets
// it needs. Returned by GET /v1/agent-templates/{templateId} and consumed by
// `npx zavudev agents pull`.
type AgentTemplateGetResponseTemplate struct {
	ID string `json:"id" api:"required"`
	// Any of "sales", "support", "frontDesk", "ops".
	Category    string `json:"category" api:"required"`
	DefaultSlug string `json:"defaultSlug" api:"required"`
	// npm dependencies for the scaffolded function.
	Dependencies    map[string]string                                `json:"dependencies" api:"required"`
	Files           []AgentTemplateGetResponseTemplateFile           `json:"files" api:"required"`
	Name            string                                           `json:"name" api:"required"`
	RequiredSecrets []AgentTemplateGetResponseTemplateRequiredSecret `json:"requiredSecrets" api:"required"`
	Summary         string                                           `json:"summary" api:"required"`
	Voice           bool                                             `json:"voice" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Category        respjson.Field
		DefaultSlug     respjson.Field
		Dependencies    respjson.Field
		Files           respjson.Field
		Name            respjson.Field
		RequiredSecrets respjson.Field
		Summary         respjson.Field
		Voice           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseTemplate) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateGetResponseTemplateFile struct {
	// File contents to write verbatim.
	Content string `json:"content" api:"required"`
	Path    string `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseTemplateFile) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseTemplateFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateGetResponseTemplateRequiredSecret struct {
	Hint string `json:"hint" api:"required"`
	Key  string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hint        respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseTemplateRequiredSecret) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseTemplateRequiredSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateListResponse struct {
	Items []AgentTemplateListResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compact catalog entry for a factory agent.
type AgentTemplateListResponseItem struct {
	ID string `json:"id" api:"required"`
	// Any of "sales", "support", "frontDesk", "ops".
	Category  string `json:"category" api:"required"`
	Name      string `json:"name" api:"required"`
	Summary   string `json:"summary" api:"required"`
	ToolCount int64  `json:"toolCount" api:"required"`
	// Whether this agent answers phone calls.
	Voice bool `json:"voice" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		Name        respjson.Field
		Summary     respjson.Field
		ToolCount   respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
