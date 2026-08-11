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
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
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
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the AI agent configuration for a sender.
func (r *SenderAgentService) Get(ctx context.Context, senderID string, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an AI agent's configuration.
func (r *SenderAgentService) Update(ctx context.Context, senderID string, body SenderAgentUpdateParams, opts ...option.RequestOption) (res *AgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete an AI agent.
func (r *SenderAgentService) Delete(ctx context.Context, senderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return err
	}
	path := fmt.Sprintf("v1/senders/%s/agent", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get statistics for an AI agent including invocations, tokens, and costs.
//
// Covers the messaging channels only. Voice calls are not counted here: a call is
// a multi-turn conversation rather than one inbound message and one reply, so it
// is recorded as a call, not an execution. An agent that only answers phone calls
// reports zeros on every field. Use `GET /v1/calls` for voice activity, duration,
// and cost.
func (r *SenderAgentService) Stats(ctx context.Context, senderID string, opts ...option.RequestOption) (res *AgentStats, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/agent/stats", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	MaxTokens int64 `json:"maxTokens" api:"nullable"`
	// Senders this agent answers on. An agent can serve several; `senderId` remains
	// the primary one, for compatibility.
	SenderIDs []string   `json:"senderIds"`
	Stats     AgentStats `json:"stats"`
	// LLM temperature (0-2).
	Temperature float64 `json:"temperature" api:"nullable"`
	// Channels that trigger the agent.
	TriggerOnChannels []string `json:"triggerOnChannels"`
	// Message types that trigger the agent.
	TriggerOnMessageTypes []string `json:"triggerOnMessageTypes"`
	// Voice Agent configuration. When present and enabled, the agent can answer
	// inbound phone calls and place outbound calls with Zavu's managed voice pipeline.
	// Requires the Voice Agents feature to be enabled for your team.
	Voice AgentVoice `json:"voice"`
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
		SenderIDs              respjson.Field
		Stats                  respjson.Field
		Temperature            respjson.Field
		TriggerOnChannels      respjson.Field
		TriggerOnMessageTypes  respjson.Field
		Voice                  respjson.Field
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

// Voice Agent configuration. When present and enabled, the agent can answer
// inbound phone calls and place outbound calls with Zavu's managed voice pipeline.
// Requires the Voice Agents feature to be enabled for your team.
type AgentVoice struct {
	// Whether the agent handles voice calls. When false, the sender's number is not
	// answered by the voice agent and outbound calls are rejected.
	Enabled bool `json:"enabled" api:"required"`
	// Opening line the agent speaks when the call connects. If omitted, the agent
	// waits for the caller to speak first.
	Greeting string `json:"greeting"`
	// Greeting per language, keyed by language code. Used when the caller's language
	// differs from the one `greeting` is written in.
	Greetings map[string]string `json:"greetings"`
	// Whether the caller can interrupt the agent while it is speaking (barge-in). When
	// true, the agent stops talking as soon as the caller starts.
	Interruptible bool `json:"interruptible"`
	// BCP-47 language code used for both speech recognition and speech synthesis (e.g.
	// `en`, `es`, `pt-BR`). Auto-detected from the recipient when omitted.
	Language string `json:"language"`
	// Hard limit on call length in minutes. The call ends automatically when reached.
	MaxCallDurationMinutes int64 `json:"maxCallDurationMinutes"`
	// How long the agent waits during silence before ending the call.
	MaxIdleSeconds int64 `json:"maxIdleSeconds"`
	// Model that runs the conversation, co-located in the voice network for lowest
	// latency. Independent of the model used for text messaging. Derived from the
	// agent's text model when omitted.
	Model string `json:"model"`
	// Whether the call audio is recorded.
	RecordCalls bool `json:"recordCalls"`
	// Speech-recognition model. Uses the default when omitted.
	SttModel string `json:"sttModel"`
	// Speech-recognition provider. Uses the default when omitted.
	SttProvider string `json:"sttProvider"`
	// E.164 phone number the agent can transfer the call to. When set, the agent is
	// given a transfer tool it can use to hand the call to a human.
	TransferPhoneNumber string `json:"transferPhoneNumber"`
	// Speech-synthesis provider. Uses the default when omitted.
	TtsProvider string `json:"ttsProvider"`
	// Identifier of the synthesized voice that speaks. Choose from the voices
	// available in the dashboard. Uses a neutral default when omitted.
	TtsVoiceID string `json:"ttsVoiceId"`
	// What the agent does when an answering machine or voicemail is detected on an
	// outbound call.
	//
	// Any of "hangup", "leave_message".
	VoicemailAction string `json:"voicemailAction"`
	// Message spoken when `voicemailAction` is `leave_message`. Falls back to
	// `greeting` when omitted.
	VoicemailMessage string `json:"voicemailMessage"`
	// Speech rate. 1.0 is natural. Only honoured by voices that support rate control;
	// ignored by the others.
	VoiceSpeed float64 `json:"voiceSpeed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                respjson.Field
		Greeting               respjson.Field
		Greetings              respjson.Field
		Interruptible          respjson.Field
		Language               respjson.Field
		MaxCallDurationMinutes respjson.Field
		MaxIdleSeconds         respjson.Field
		Model                  respjson.Field
		RecordCalls            respjson.Field
		SttModel               respjson.Field
		SttProvider            respjson.Field
		TransferPhoneNumber    respjson.Field
		TtsProvider            respjson.Field
		TtsVoiceID             respjson.Field
		VoicemailAction        respjson.Field
		VoicemailMessage       respjson.Field
		VoiceSpeed             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVoice) RawJSON() string { return r.JSON.raw }
func (r *AgentVoice) UnmarshalJSON(data []byte) error {
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
	Status           AgentExecutionStatus `json:"status" api:"required"`
	ErrorMessage     string               `json:"errorMessage" api:"nullable"`
	InboundMessageID string               `json:"inboundMessageId"`
	// Knowledge-base chunks retrieved for this answer. Zero on an agent that has
	// documents attached means the reply was not grounded in them, which is otherwise
	// indistinguishable from a correct answer in this record. Absent on executions
	// recorded before this field existed, which is not the same as zero.
	KnowledgeChunksUsed int64  `json:"knowledgeChunksUsed" api:"nullable"`
	ResponseMessageID   string `json:"responseMessageId" api:"nullable"`
	ResponseText        string `json:"responseText" api:"nullable"`
	// Tools the agent called while producing this reply. Zero on an agent that has
	// tools configured means it answered without calling any — the case where a reply
	// says it will look something up and nothing ever reaches your endpoint. Absent on
	// executions recorded before this field existed, which is not the same as zero.
	ToolCalls int64 `json:"toolCalls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AgentID             respjson.Field
		Cost                respjson.Field
		CreatedAt           respjson.Field
		InputTokens         respjson.Field
		LatencyMs           respjson.Field
		OutputTokens        respjson.Field
		Status              respjson.Field
		ErrorMessage        respjson.Field
		InboundMessageID    respjson.Field
		KnowledgeChunksUsed respjson.Field
		ResponseMessageID   respjson.Field
		ResponseText        respjson.Field
		ToolCalls           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Voice Agent configuration. Enable this to let the agent answer and place phone
	// calls with Zavu's managed voice pipeline. Requires the Voice Agents feature to
	// be enabled for your team.
	Voice SenderAgentNewParamsVoice `json:"voice,omitzero"`
	paramObj
}

func (r SenderAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice Agent configuration. Enable this to let the agent answer and place phone
// calls with Zavu's managed voice pipeline. Requires the Voice Agents feature to
// be enabled for your team.
//
// The property Enabled is required.
type SenderAgentNewParamsVoice struct {
	// Whether the agent handles voice calls. When false, the sender's number is not
	// answered by the voice agent and outbound calls are rejected.
	Enabled bool `json:"enabled" api:"required"`
	// Opening line the agent speaks when the call connects. If omitted, the agent
	// waits for the caller to speak first.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// Whether the caller can interrupt the agent while it is speaking (barge-in). When
	// true, the agent stops talking as soon as the caller starts.
	Interruptible param.Opt[bool] `json:"interruptible,omitzero"`
	// BCP-47 language code used for both speech recognition and speech synthesis (e.g.
	// `en`, `es`, `pt-BR`). Auto-detected from the recipient when omitted.
	Language param.Opt[string] `json:"language,omitzero"`
	// Hard limit on call length in minutes. The call ends automatically when reached.
	MaxCallDurationMinutes param.Opt[int64] `json:"maxCallDurationMinutes,omitzero"`
	// How long the agent waits during silence before ending the call.
	MaxIdleSeconds param.Opt[int64] `json:"maxIdleSeconds,omitzero"`
	// Model that runs the conversation, co-located in the voice network for lowest
	// latency. Independent of the model used for text messaging. Derived from the
	// agent's text model when omitted.
	Model param.Opt[string] `json:"model,omitzero"`
	// Whether the call audio is recorded.
	RecordCalls param.Opt[bool] `json:"recordCalls,omitzero"`
	// Speech-recognition model. Uses the default when omitted.
	SttModel param.Opt[string] `json:"sttModel,omitzero"`
	// Speech-recognition provider. Uses the default when omitted.
	SttProvider param.Opt[string] `json:"sttProvider,omitzero"`
	// E.164 phone number the agent can transfer the call to. When set, the agent is
	// given a transfer tool it can use to hand the call to a human.
	TransferPhoneNumber param.Opt[string] `json:"transferPhoneNumber,omitzero"`
	// Speech-synthesis provider. Uses the default when omitted.
	TtsProvider param.Opt[string] `json:"ttsProvider,omitzero"`
	// Identifier of the synthesized voice that speaks. Choose from the voices
	// available in the dashboard. Uses a neutral default when omitted.
	TtsVoiceID param.Opt[string] `json:"ttsVoiceId,omitzero"`
	// Message spoken when `voicemailAction` is `leave_message`. Falls back to
	// `greeting` when omitted.
	VoicemailMessage param.Opt[string] `json:"voicemailMessage,omitzero"`
	// Speech rate. 1.0 is natural. Only honoured by voices that support rate control;
	// ignored by the others.
	VoiceSpeed param.Opt[float64] `json:"voiceSpeed,omitzero"`
	// Greeting per language, keyed by language code. Used when the caller's language
	// differs from the one `greeting` is written in.
	Greetings map[string]string `json:"greetings,omitzero"`
	// What the agent does when an answering machine or voicemail is detected on an
	// outbound call.
	//
	// Any of "hangup", "leave_message".
	VoicemailAction string `json:"voicemailAction,omitzero"`
	paramObj
}

func (r SenderAgentNewParamsVoice) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentNewParamsVoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentNewParamsVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentNewParamsVoice](
		"voicemailAction", "hangup", "leave_message",
	)
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
	// Voice Agent configuration. Patch this object to enable voice, change the
	// greeting, or adjust call limits. Requires the Voice Agents feature to be enabled
	// for your team.
	Voice SenderAgentUpdateParamsVoice `json:"voice,omitzero"`
	paramObj
}

func (r SenderAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice Agent configuration. Patch this object to enable voice, change the
// greeting, or adjust call limits. Requires the Voice Agents feature to be enabled
// for your team.
//
// The property Enabled is required.
type SenderAgentUpdateParamsVoice struct {
	// Whether the agent handles voice calls. When false, the sender's number is not
	// answered by the voice agent and outbound calls are rejected.
	Enabled bool `json:"enabled" api:"required"`
	// Opening line the agent speaks when the call connects. If omitted, the agent
	// waits for the caller to speak first.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// Whether the caller can interrupt the agent while it is speaking (barge-in). When
	// true, the agent stops talking as soon as the caller starts.
	Interruptible param.Opt[bool] `json:"interruptible,omitzero"`
	// BCP-47 language code used for both speech recognition and speech synthesis (e.g.
	// `en`, `es`, `pt-BR`). Auto-detected from the recipient when omitted.
	Language param.Opt[string] `json:"language,omitzero"`
	// Hard limit on call length in minutes. The call ends automatically when reached.
	MaxCallDurationMinutes param.Opt[int64] `json:"maxCallDurationMinutes,omitzero"`
	// How long the agent waits during silence before ending the call.
	MaxIdleSeconds param.Opt[int64] `json:"maxIdleSeconds,omitzero"`
	// Model that runs the conversation, co-located in the voice network for lowest
	// latency. Independent of the model used for text messaging. Derived from the
	// agent's text model when omitted.
	Model param.Opt[string] `json:"model,omitzero"`
	// Whether the call audio is recorded.
	RecordCalls param.Opt[bool] `json:"recordCalls,omitzero"`
	// Speech-recognition model. Uses the default when omitted.
	SttModel param.Opt[string] `json:"sttModel,omitzero"`
	// Speech-recognition provider. Uses the default when omitted.
	SttProvider param.Opt[string] `json:"sttProvider,omitzero"`
	// E.164 phone number the agent can transfer the call to. When set, the agent is
	// given a transfer tool it can use to hand the call to a human.
	TransferPhoneNumber param.Opt[string] `json:"transferPhoneNumber,omitzero"`
	// Speech-synthesis provider. Uses the default when omitted.
	TtsProvider param.Opt[string] `json:"ttsProvider,omitzero"`
	// Identifier of the synthesized voice that speaks. Choose from the voices
	// available in the dashboard. Uses a neutral default when omitted.
	TtsVoiceID param.Opt[string] `json:"ttsVoiceId,omitzero"`
	// Message spoken when `voicemailAction` is `leave_message`. Falls back to
	// `greeting` when omitted.
	VoicemailMessage param.Opt[string] `json:"voicemailMessage,omitzero"`
	// Speech rate. 1.0 is natural. Only honoured by voices that support rate control;
	// ignored by the others.
	VoiceSpeed param.Opt[float64] `json:"voiceSpeed,omitzero"`
	// Greeting per language, keyed by language code. Used when the caller's language
	// differs from the one `greeting` is written in.
	Greetings map[string]string `json:"greetings,omitzero"`
	// What the agent does when an answering machine or voicemail is detected on an
	// outbound call.
	//
	// Any of "hangup", "leave_message".
	VoicemailAction string `json:"voicemailAction,omitzero"`
	paramObj
}

func (r SenderAgentUpdateParamsVoice) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentUpdateParamsVoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentUpdateParamsVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SenderAgentUpdateParamsVoice](
		"voicemailAction", "hangup", "leave_message",
	)
}
