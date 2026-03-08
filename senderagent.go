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

	"github.com/stainless-sdks/zavudev-go/internal/apijson"
	"github.com/stainless-sdks/zavudev-go/internal/requestconfig"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/stainless-sdks/zavudev-go/packages/param"
	"github.com/stainless-sdks/zavudev-go/packages/respjson"
)

// SenderAgentService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentService] method instead.
type SenderAgentService struct {
	Options        []option.RequestOption
	Executions     SenderAgentExecutionService
	Flows          SenderAgentFlowService
	Tools          SenderAgentToolService
	KnowledgeBases SenderAgentKnowledgeBaseService
}

// NewSenderAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSenderAgentService(opts ...option.RequestOption) (r SenderAgentService) {
	r = SenderAgentService{}
	r.Options = opts
	r.Executions = NewSenderAgentExecutionService(opts...)
	r.Flows = NewSenderAgentFlowService(opts...)
	r.Tools = NewSenderAgentToolService(opts...)
	r.KnowledgeBases = NewSenderAgentKnowledgeBaseService(opts...)
	return
}

// Create an AI agent for a sender. Each sender can have at most one agent.
func (r *SenderAgentService) New(ctx context.Context, senderID string, body SenderAgentNewParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the AI agent configuration for a sender.
func (r *SenderAgentService) Get(ctx context.Context, senderID string, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an AI agent's configuration.
func (r *SenderAgentService) Update(ctx context.Context, senderID string, body SenderAgentUpdateParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete an AI agent.
func (r *SenderAgentService) Delete(ctx context.Context, senderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get statistics for an AI agent including invocations, tokens, and costs.
func (r *SenderAgentService) Stats(ctx context.Context, senderID string, opts ...option.RequestOption) (res *AgentStats, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/stats", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// AI Agent configuration for a sender.
type Agent struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the agent is active.
	Enabled bool `json:"enabled" api:"required"`
	// Model ID (e.g., gpt-4o-mini, claude-3-5-sonnet).
	Model string `json:"model" api:"required"`
	Name  string `json:"name" api:"required"`
	// LLM provider for the AI agent.
	//
	// Any of "openai", "anthropic", "google", "mistral", "zavu".
	Provider AgentProvider `json:"provider" api:"required"`
	SenderID string        `json:"senderId" api:"required"`
	// System prompt for the agent.
	SystemPrompt string    `json:"systemPrompt" api:"required"`
	UpdatedAt    time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Number of previous messages to include as context.
	ContextWindowMessages int64 `json:"contextWindowMessages"`
	// Whether to include contact metadata in context.
	IncludeContactMetadata bool `json:"includeContactMetadata"`
	// Maximum tokens for LLM response.
	MaxTokens int64      `json:"maxTokens" api:"nullable"`
	Stats     AgentStats `json:"stats"`
	// LLM temperature (0-2).
	Temperature float64 `json:"temperature" api:"nullable"`
	// Channels that trigger the agent.
	TriggerOnChannels []string `json:"triggerOnChannels"`
	// Message types that trigger the agent.
	TriggerOnMessageTypes []string `json:"triggerOnMessageTypes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		Enabled                respjson.Field
		Model                  respjson.Field
		Name                   respjson.Field
		Provider               respjson.Field
		SenderID               respjson.Field
		SystemPrompt           respjson.Field
		UpdatedAt              respjson.Field
		ContextWindowMessages  respjson.Field
		IncludeContactMetadata respjson.Field
		MaxTokens              respjson.Field
		Stats                  respjson.Field
		Temperature            respjson.Field
		TriggerOnChannels      respjson.Field
		TriggerOnMessageTypes  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Agent) RawJSON() string { return r.JSON.raw }
func (r *Agent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentStats struct {
	// Total cost in USD.
	TotalCost        float64 `json:"totalCost"`
	TotalInvocations int64   `json:"totalInvocations"`
	TotalTokensUsed  int64   `json:"totalTokensUsed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalCost        respjson.Field
		TotalInvocations respjson.Field
		TotalTokensUsed  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentStats) RawJSON() string { return r.JSON.raw }
func (r *AgentStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentExecution struct {
	ID      string `json:"id" api:"required"`
	AgentID string `json:"agentId" api:"required"`
	// Cost in USD.
	Cost         float64   `json:"cost" api:"required"`
	CreatedAt    time.Time `json:"createdAt" api:"required" format:"date-time"`
	InputTokens  int64     `json:"inputTokens" api:"required"`
	LatencyMs    int64     `json:"latencyMs" api:"required"`
	OutputTokens int64     `json:"outputTokens" api:"required"`
	// Status of an agent execution.
	//
	// Any of "success", "error", "filtered", "rate_limited", "balance_insufficient".
	Status            AgentExecutionStatus `json:"status" api:"required"`
	ErrorMessage      string               `json:"errorMessage" api:"nullable"`
	InboundMessageID  string               `json:"inboundMessageId"`
	ResponseMessageID string               `json:"responseMessageId" api:"nullable"`
	ResponseText      string               `json:"responseText" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AgentID           respjson.Field
		Cost              respjson.Field
		CreatedAt         respjson.Field
		InputTokens       respjson.Field
		LatencyMs         respjson.Field
		OutputTokens      respjson.Field
		Status            respjson.Field
		ErrorMessage      respjson.Field
		InboundMessageID  respjson.Field
		ResponseMessageID respjson.Field
		ResponseText      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentExecution) RawJSON() string { return r.JSON.raw }
func (r *AgentExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of an agent execution.
type AgentExecutionStatus string

const (
	AgentExecutionStatusSuccess             AgentExecutionStatus = "success"
	AgentExecutionStatusError               AgentExecutionStatus = "error"
	AgentExecutionStatusFiltered            AgentExecutionStatus = "filtered"
	AgentExecutionStatusRateLimited         AgentExecutionStatus = "rate_limited"
	AgentExecutionStatusBalanceInsufficient AgentExecutionStatus = "balance_insufficient"
)

// LLM provider for the AI agent.
type AgentProvider string

const (
	AgentProviderOpenAI    AgentProvider = "openai"
	AgentProviderAnthropic AgentProvider = "anthropic"
	AgentProviderGoogle    AgentProvider = "google"
	AgentProviderMistral   AgentProvider = "mistral"
	AgentProviderZavu      AgentProvider = "zavu"
)

type AgentResponse struct {
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
func (r AgentResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentNewParams struct {
	Model string `json:"model" api:"required"`
	Name  string `json:"name" api:"required"`
	// LLM provider for the AI agent.
	//
	// Any of "openai", "anthropic", "google", "mistral", "zavu".
	Provider     AgentProvider `json:"provider,omitzero" api:"required"`
	SystemPrompt string        `json:"systemPrompt" api:"required"`
	// API key for the LLM provider. Required unless provider is 'zavu'.
	APIKey                 param.Opt[string]  `json:"apiKey,omitzero"`
	ContextWindowMessages  param.Opt[int64]   `json:"contextWindowMessages,omitzero"`
	IncludeContactMetadata param.Opt[bool]    `json:"includeContactMetadata,omitzero"`
	MaxTokens              param.Opt[int64]   `json:"maxTokens,omitzero"`
	Temperature            param.Opt[float64] `json:"temperature,omitzero"`
	TriggerOnChannels      []string           `json:"triggerOnChannels,omitzero"`
	TriggerOnMessageTypes  []string           `json:"triggerOnMessageTypes,omitzero"`
	paramObj
}

func (r SenderAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentUpdateParams struct {
	MaxTokens              param.Opt[int64]   `json:"maxTokens,omitzero"`
	Temperature            param.Opt[float64] `json:"temperature,omitzero"`
	APIKey                 param.Opt[string]  `json:"apiKey,omitzero"`
	ContextWindowMessages  param.Opt[int64]   `json:"contextWindowMessages,omitzero"`
	Enabled                param.Opt[bool]    `json:"enabled,omitzero"`
	IncludeContactMetadata param.Opt[bool]    `json:"includeContactMetadata,omitzero"`
	Model                  param.Opt[string]  `json:"model,omitzero"`
	Name                   param.Opt[string]  `json:"name,omitzero"`
	SystemPrompt           param.Opt[string]  `json:"systemPrompt,omitzero"`
	// LLM provider for the AI agent.
	//
	// Any of "openai", "anthropic", "google", "mistral", "zavu".
	Provider              AgentProvider `json:"provider,omitzero"`
	TriggerOnChannels     []string      `json:"triggerOnChannels,omitzero"`
	TriggerOnMessageTypes []string      `json:"triggerOnMessageTypes,omitzero"`
	paramObj
}

func (r SenderAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
