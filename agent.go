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

// AgentService contains methods and other services that help with interacting with
// the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	Senders AgentSenderService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	r.Senders = NewAgentSenderService(opts...)
	return
}

// Create an agent without a sender. It is created disabled; connect a sender and
// enable it when you are ready for it to answer.
//
// **Sub-resources.** An agent's tools, flows and knowledge bases are reachable at
// `/v1/agents/{agentId}/tools`, `/v1/agents/{agentId}/flows` and
// `/v1/agents/{agentId}/knowledge-bases`, mirroring the sender-scoped routes
// documented under `/v1/senders/{senderId}/agent/...` exactly. Use the
// agent-scoped form while the agent has no sender: the sender-scoped one cannot
// address it.
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *AgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an agent
func (r *AgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s", url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an agent
func (r *AgentService) Update(ctx context.Context, agentID string, body AgentUpdateParams, opts ...option.RequestOption) (res *AgentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s", url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Every agent in the project, newest first — including agents that are not
// connected to any sender yet, which the sender-scoped routes cannot reach. Each
// item carries `senderIds`, the senders the agent answers on.
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *pagination.Cursor[Agent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/agents"
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

// Every agent in the project, newest first — including agents that are not
// connected to any sender yet, which the sender-scoped routes cannot reach. Each
// item carries `senderIds`, the senders the agent answers on.
func (r *AgentService) ListAutoPaging(ctx context.Context, query AgentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Agent] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete an agent
func (r *AgentService) Delete(ctx context.Context, agentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	path := fmt.Sprintf("v1/agents/%s", url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The voices an agent can speak with, for `voice.ttsVoiceId`. Filter by `language`
// to get the ones that speak it; a voice can still be used with `language: auto`,
// where the agent follows the caller and keeps the chosen voice.
func (r *AgentService) ListVoices(ctx context.Context, query AgentListVoicesParams, opts ...option.RequestOption) (res *AgentListVoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Run the agent's prompt, model and knowledge base against a message and return
// the reply instead of delivering it. Writes nothing and charges nothing, so it is
// safe to call repeatedly while iterating on a prompt.
//
// Note that a dry run never **executes** tools — running them would cause real
// side effects. Live conversations on every channel do call them. When the agent
// has enabled tools, that gap is reported in `warnings` rather than silently
// producing an answer that looks like a tool call happened.
func (r *AgentService) Test(ctx context.Context, agentID string, body AgentTestParams, opts ...option.RequestOption) (res *AgentTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/test", url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AgentNewResponse struct {
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
func (r AgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponse struct {
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
func (r AgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUpdateResponse struct {
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
func (r AgentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListVoicesResponse struct {
	Items []AgentListVoicesResponseItem `json:"items" api:"required"`
	// Languages an agent can be pinned to. `auto` follows the caller.
	Languages []string `json:"languages" api:"required"`
	// Voices in the catalog, before filtering.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Languages   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListVoicesResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListVoicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListVoicesResponseItem struct {
	// Value for `voice.ttsVoiceId`.
	ID       string `json:"id" api:"required"`
	Language string `json:"language" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListVoicesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AgentListVoicesResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTestResponse struct {
	Error       string `json:"error" api:"required"`
	InputTokens int64  `json:"inputTokens" api:"required"`
	// Knowledge-base chunks retrieved for this message. Zero means the answer was not
	// grounded in your documents.
	KnowledgeChunksUsed int64 `json:"knowledgeChunksUsed" api:"required"`
	LatencyMs           int64 `json:"latencyMs" api:"required"`
	OutputTokens        int64 `json:"outputTokens" api:"required"`
	Success             bool  `json:"success" api:"required"`
	// What the agent would reply.
	Text string `json:"text" api:"required"`
	// Things that are true of this agent but that a dry run cannot prove. Surfaced so
	// a passing dry run is never mistaken for proof that the agent works live.
	//
	//   - The agent being disabled.
	//   - Enabled tools that were **not offered to the model** here — the model never
	//     saw them, so a reply that looks like a lookup was invented. Live conversations
	//     on every channel do offer them; running them here would cause real side
	//     effects.
	//   - An agent whose sender has none of the channels it triggers on, which answers
	//     every dry run and no real message.
	//   - Contact metadata that exists on a real conversation but not here.
	Warnings []string `json:"warnings" api:"required"`
	// Tools that actually ran, in order, when the request set `executeTools`. Empty on
	// a normal dry run, where nothing is executed. An entry with `ok: false` means the
	// agent saw an error and answered around it, which is what a customer would have
	// received.
	ExecutedToolCalls []AgentTestResponseExecutedToolCall `json:"executedToolCalls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error               respjson.Field
		InputTokens         respjson.Field
		KnowledgeChunksUsed respjson.Field
		LatencyMs           respjson.Field
		OutputTokens        respjson.Field
		Success             respjson.Field
		Text                respjson.Field
		Warnings            respjson.Field
		ExecutedToolCalls   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTestResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTestResponseExecutedToolCall struct {
	Name  string `json:"name" api:"required"`
	Ok    bool   `json:"ok" api:"required"`
	Error string `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Ok          respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTestResponseExecutedToolCall) RawJSON() string { return r.JSON.raw }
func (r *AgentTestResponseExecutedToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	Model string `json:"model" api:"required"`
	Name  string `json:"name" api:"required"`
	// LLM provider for the AI agent.
	//
	// Any of "openai", "anthropic", "google", "mistral", "zavu".
	Provider               AgentProvider      `json:"provider,omitzero" api:"required"`
	SystemPrompt           string             `json:"systemPrompt" api:"required"`
	ContextWindowMessages  param.Opt[int64]   `json:"contextWindowMessages,omitzero"`
	IncludeContactMetadata param.Opt[bool]    `json:"includeContactMetadata,omitzero"`
	MaxTokens              param.Opt[int64]   `json:"maxTokens,omitzero"`
	Temperature            param.Opt[float64] `json:"temperature,omitzero"`
	TriggerOnChannels      []string           `json:"triggerOnChannels,omitzero"`
	TriggerOnMessageTypes  []string           `json:"triggerOnMessageTypes,omitzero"`
	// Voice Agent configuration on a sender's AI agent. Controls how the agent behaves
	// on inbound and outbound phone calls through Zavu's managed voice pipeline
	// (speech recognition, the agent's LLM, and speech synthesis, with real-time
	// interruption handling). Requires the Voice Agents feature to be enabled for your
	// team.
	Voice AgentNewParamsVoice `json:"voice,omitzero"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice Agent configuration on a sender's AI agent. Controls how the agent behaves
// on inbound and outbound phone calls through Zavu's managed voice pipeline
// (speech recognition, the agent's LLM, and speech synthesis, with real-time
// interruption handling). Requires the Voice Agents feature to be enabled for your
// team.
//
// The property Enabled is required.
type AgentNewParamsVoice struct {
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

func (r AgentNewParamsVoice) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParamsVoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParamsVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentNewParamsVoice](
		"voicemailAction", "hangup", "leave_message",
	)
}

type AgentUpdateParams struct {
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
	Voice AgentUpdateParamsVoice `json:"voice,omitzero"`
	paramObj
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Voice Agent configuration. Patch this object to enable voice, change the
// greeting, or adjust call limits. Requires the Voice Agents feature to be enabled
// for your team.
//
// The property Enabled is required.
type AgentUpdateParamsVoice struct {
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

func (r AgentUpdateParamsVoice) MarshalJSON() (data []byte, err error) {
	type shadow AgentUpdateParamsVoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUpdateParamsVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentUpdateParamsVoice](
		"voicemailAction", "hangup", "leave_message",
	)
}

type AgentListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentListVoicesParams struct {
	// BCP-47 tag (`en`, `es`, `pt-BR`). Omit, or pass `auto`, for every voice.
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListVoicesParams]'s query parameters as `url.Values`.
func (r AgentListVoicesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentTestParams struct {
	// What to say to the agent.
	Message string `json:"message" api:"required"`
	// Run the tools the agent calls instead of reporting the choice and stopping.
	//
	// Off by default because a tool handler talks to the outside world: a rehearsal
	// that charges a card is not a rehearsal. Turn it on to exercise the loop that
	// actually matters — the model picks a tool, the handler answers, the model
	// replies with the result — without sending a message to anyone. What ran comes
	// back in `executedToolCalls`.
	ExecuteTools param.Opt[bool] `json:"executeTools,omitzero"`
	// Set false to skip retrieval and isolate prompt behaviour from the knowledge
	// base.
	UseKnowledgeBase param.Opt[bool] `json:"useKnowledgeBase,omitzero"`
	// Prior turns, oldest first, to exercise multi-turn behaviour without persisting a
	// thread. Trimmed to the agent's context window.
	History []AgentTestParamsHistory `json:"history,omitzero"`
	paramObj
}

func (r AgentTestParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type AgentTestParamsHistory struct {
	Content string `json:"content" api:"required"`
	// Any of "user", "assistant".
	Role string `json:"role,omitzero" api:"required"`
	paramObj
}

func (r AgentTestParamsHistory) MarshalJSON() (data []byte, err error) {
	type shadow AgentTestParamsHistory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTestParamsHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentTestParamsHistory](
		"role", "user", "assistant",
	)
}
