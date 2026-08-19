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

// CallService contains methods and other services that help with interacting with
// the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r CallService) {
	r = CallService{}
	r.Options = opts
	return
}

// Place an outbound voice call answered by the voice agent configured on the
// sender. Zavu dials the recipient and runs the conversation through its managed
// voice pipeline (speech recognition, the agent's LLM, and speech synthesis, with
// real-time interruption handling).
//
// **Requirements:**
//
// - The Voice Agents feature must be enabled for your team (otherwise `403`).
// - The sender's agent must have `voice.enabled` set to `true`.
// - Not available with test-mode API keys.
//
// **Billing:** Voice calls are billed per minute of connected time plus telephony,
// deducted from your prepaid balance. A short-duration estimate is reserved when
// the call is placed; you are charged for the actual duration when the call ends.
func (r *CallService) New(ctx context.Context, body CallNewParams, opts ...option.RequestOption) (res *CallNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single voice call, including its full transcript once the
// conversation has produced turns.
func (r *CallService) Get(ctx context.Context, callID string, opts ...option.RequestOption) (res *CallGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callID == "" {
		err = errors.New("missing required callId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/calls/%s", url.PathEscape(callID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List voice calls for this project, most recent first. Transcripts are omitted
// from the list; fetch a single call to get its transcript.
func (r *CallService) List(ctx context.Context, query CallListParams, opts ...option.RequestOption) (res *pagination.Cursor[CallListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/calls"
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

// List voice calls for this project, most recent first. Transcripts are omitted
// from the list; fetch a single call to get its transcript.
func (r *CallService) ListAutoPaging(ctx context.Context, query CallListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CallListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// End an active voice call. The call must still be ringing or in progress. Not
// available with test-mode API keys.
func (r *CallService) Hangup(ctx context.Context, callID string, opts ...option.RequestOption) (res *CallHangupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callID == "" {
		err = errors.New("missing required callId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/calls/%s/hangup", url.PathEscape(callID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CallNewResponse struct {
	Call CallNewResponseCall `json:"call" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Call        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CallNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewResponseCall struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the call was placed by Zavu (outbound) or received from a caller
	// (inbound).
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Caller phone number in E.164 format. Your sender's number for outbound calls;
	// the caller's number for inbound calls.
	From string `json:"from" api:"required"`
	// Lifecycle status of a voice call.
	//
	// - `queued`: outbound call created, not yet dialing.
	// - `ringing`: dialing (outbound) or received and ringing (inbound).
	// - `in_progress`: answered, the agent is connected.
	// - `completed`: ended after a conversation.
	// - `failed`: could not be completed.
	// - `busy`: the line was busy.
	// - `no_answer`: rang but was not answered.
	// - `canceled`: canceled before it was answered.
	//
	// Any of "queued", "ringing", "in_progress", "completed", "failed", "busy",
	// "no_answer", "canceled".
	Status string `json:"status" api:"required"`
	// Callee phone number in E.164 format.
	To string `json:"to" api:"required"`
	// When the call was answered.
	AnsweredAt time.Time `json:"answeredAt" api:"nullable" format:"date-time"`
	// Total cost of the call in USD, combining the managed voice pipeline per-minute
	// charge and telephony. Available once the call has ended.
	Cost float64 `json:"cost" api:"nullable"`
	// Billable talk time in seconds, measured from answer to hangup.
	DurationSeconds int64 `json:"durationSeconds" api:"nullable"`
	// When the call ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// Why the call ended (e.g. `agent_ended`, `max_duration`, `transfer`, `hangup`).
	// Present once the call is no longer active.
	EndReason string `json:"endReason" api:"nullable"`
	// Arbitrary metadata you attached when creating the call.
	Metadata map[string]string `json:"metadata"`
	// Ordered transcript of the call. Included when retrieving a single call; omitted
	// from list responses.
	Transcript []CallNewResponseCallTranscript `json:"transcript"`
	// Number of conversation turns exchanged during the call.
	TurnCount int64     `json:"turnCount" api:"nullable"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Direction       respjson.Field
		From            respjson.Field
		Status          respjson.Field
		To              respjson.Field
		AnsweredAt      respjson.Field
		Cost            respjson.Field
		DurationSeconds respjson.Field
		EndedAt         respjson.Field
		EndReason       respjson.Field
		Metadata        respjson.Field
		Transcript      respjson.Field
		TurnCount       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallNewResponseCall) RawJSON() string { return r.JSON.raw }
func (r *CallNewResponseCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single turn in a voice call transcript.
type CallNewResponseCallTranscript struct {
	// Who produced the turn. `tool` records a tool call the agent made during the
	// conversation.
	//
	// Any of "user", "assistant", "tool".
	Role string `json:"role" api:"required"`
	// Ordinal position of the turn within the call, starting at 0.
	Seq int64 `json:"seq" api:"required"`
	// Transcribed speech for `user` and `assistant` turns, or a JSON summary of the
	// tool call for `tool` turns.
	Text string `json:"text" api:"required"`
	// When the turn ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// When the turn started.
	StartedAt time.Time `json:"startedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Seq         respjson.Field
		Text        respjson.Field
		EndedAt     respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallNewResponseCallTranscript) RawJSON() string { return r.JSON.raw }
func (r *CallNewResponseCallTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetResponse struct {
	Call CallGetResponseCall `json:"call" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Call        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CallGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetResponseCall struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the call was placed by Zavu (outbound) or received from a caller
	// (inbound).
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Caller phone number in E.164 format. Your sender's number for outbound calls;
	// the caller's number for inbound calls.
	From string `json:"from" api:"required"`
	// Lifecycle status of a voice call.
	//
	// - `queued`: outbound call created, not yet dialing.
	// - `ringing`: dialing (outbound) or received and ringing (inbound).
	// - `in_progress`: answered, the agent is connected.
	// - `completed`: ended after a conversation.
	// - `failed`: could not be completed.
	// - `busy`: the line was busy.
	// - `no_answer`: rang but was not answered.
	// - `canceled`: canceled before it was answered.
	//
	// Any of "queued", "ringing", "in_progress", "completed", "failed", "busy",
	// "no_answer", "canceled".
	Status string `json:"status" api:"required"`
	// Callee phone number in E.164 format.
	To string `json:"to" api:"required"`
	// When the call was answered.
	AnsweredAt time.Time `json:"answeredAt" api:"nullable" format:"date-time"`
	// Total cost of the call in USD, combining the managed voice pipeline per-minute
	// charge and telephony. Available once the call has ended.
	Cost float64 `json:"cost" api:"nullable"`
	// Billable talk time in seconds, measured from answer to hangup.
	DurationSeconds int64 `json:"durationSeconds" api:"nullable"`
	// When the call ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// Why the call ended (e.g. `agent_ended`, `max_duration`, `transfer`, `hangup`).
	// Present once the call is no longer active.
	EndReason string `json:"endReason" api:"nullable"`
	// Arbitrary metadata you attached when creating the call.
	Metadata map[string]string `json:"metadata"`
	// Ordered transcript of the call. Included when retrieving a single call; omitted
	// from list responses.
	Transcript []CallGetResponseCallTranscript `json:"transcript"`
	// Number of conversation turns exchanged during the call.
	TurnCount int64     `json:"turnCount" api:"nullable"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Direction       respjson.Field
		From            respjson.Field
		Status          respjson.Field
		To              respjson.Field
		AnsweredAt      respjson.Field
		Cost            respjson.Field
		DurationSeconds respjson.Field
		EndedAt         respjson.Field
		EndReason       respjson.Field
		Metadata        respjson.Field
		Transcript      respjson.Field
		TurnCount       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetResponseCall) RawJSON() string { return r.JSON.raw }
func (r *CallGetResponseCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single turn in a voice call transcript.
type CallGetResponseCallTranscript struct {
	// Who produced the turn. `tool` records a tool call the agent made during the
	// conversation.
	//
	// Any of "user", "assistant", "tool".
	Role string `json:"role" api:"required"`
	// Ordinal position of the turn within the call, starting at 0.
	Seq int64 `json:"seq" api:"required"`
	// Transcribed speech for `user` and `assistant` turns, or a JSON summary of the
	// tool call for `tool` turns.
	Text string `json:"text" api:"required"`
	// When the turn ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// When the turn started.
	StartedAt time.Time `json:"startedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Seq         respjson.Field
		Text        respjson.Field
		EndedAt     respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetResponseCallTranscript) RawJSON() string { return r.JSON.raw }
func (r *CallGetResponseCallTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallListResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the call was placed by Zavu (outbound) or received from a caller
	// (inbound).
	//
	// Any of "inbound", "outbound".
	Direction CallListResponseDirection `json:"direction" api:"required"`
	// Caller phone number in E.164 format. Your sender's number for outbound calls;
	// the caller's number for inbound calls.
	From string `json:"from" api:"required"`
	// Lifecycle status of a voice call.
	//
	// - `queued`: outbound call created, not yet dialing.
	// - `ringing`: dialing (outbound) or received and ringing (inbound).
	// - `in_progress`: answered, the agent is connected.
	// - `completed`: ended after a conversation.
	// - `failed`: could not be completed.
	// - `busy`: the line was busy.
	// - `no_answer`: rang but was not answered.
	// - `canceled`: canceled before it was answered.
	//
	// Any of "queued", "ringing", "in_progress", "completed", "failed", "busy",
	// "no_answer", "canceled".
	Status CallListResponseStatus `json:"status" api:"required"`
	// Callee phone number in E.164 format.
	To string `json:"to" api:"required"`
	// When the call was answered.
	AnsweredAt time.Time `json:"answeredAt" api:"nullable" format:"date-time"`
	// Total cost of the call in USD, combining the managed voice pipeline per-minute
	// charge and telephony. Available once the call has ended.
	Cost float64 `json:"cost" api:"nullable"`
	// Billable talk time in seconds, measured from answer to hangup.
	DurationSeconds int64 `json:"durationSeconds" api:"nullable"`
	// When the call ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// Why the call ended (e.g. `agent_ended`, `max_duration`, `transfer`, `hangup`).
	// Present once the call is no longer active.
	EndReason string `json:"endReason" api:"nullable"`
	// Arbitrary metadata you attached when creating the call.
	Metadata map[string]string `json:"metadata"`
	// Ordered transcript of the call. Included when retrieving a single call; omitted
	// from list responses.
	Transcript []CallListResponseTranscript `json:"transcript"`
	// Number of conversation turns exchanged during the call.
	TurnCount int64     `json:"turnCount" api:"nullable"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Direction       respjson.Field
		From            respjson.Field
		Status          respjson.Field
		To              respjson.Field
		AnsweredAt      respjson.Field
		Cost            respjson.Field
		DurationSeconds respjson.Field
		EndedAt         respjson.Field
		EndReason       respjson.Field
		Metadata        respjson.Field
		Transcript      respjson.Field
		TurnCount       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallListResponse) RawJSON() string { return r.JSON.raw }
func (r *CallListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the call was placed by Zavu (outbound) or received from a caller
// (inbound).
type CallListResponseDirection string

const (
	CallListResponseDirectionInbound  CallListResponseDirection = "inbound"
	CallListResponseDirectionOutbound CallListResponseDirection = "outbound"
)

// Lifecycle status of a voice call.
//
// - `queued`: outbound call created, not yet dialing.
// - `ringing`: dialing (outbound) or received and ringing (inbound).
// - `in_progress`: answered, the agent is connected.
// - `completed`: ended after a conversation.
// - `failed`: could not be completed.
// - `busy`: the line was busy.
// - `no_answer`: rang but was not answered.
// - `canceled`: canceled before it was answered.
type CallListResponseStatus string

const (
	CallListResponseStatusQueued     CallListResponseStatus = "queued"
	CallListResponseStatusRinging    CallListResponseStatus = "ringing"
	CallListResponseStatusInProgress CallListResponseStatus = "in_progress"
	CallListResponseStatusCompleted  CallListResponseStatus = "completed"
	CallListResponseStatusFailed     CallListResponseStatus = "failed"
	CallListResponseStatusBusy       CallListResponseStatus = "busy"
	CallListResponseStatusNoAnswer   CallListResponseStatus = "no_answer"
	CallListResponseStatusCanceled   CallListResponseStatus = "canceled"
)

// A single turn in a voice call transcript.
type CallListResponseTranscript struct {
	// Who produced the turn. `tool` records a tool call the agent made during the
	// conversation.
	//
	// Any of "user", "assistant", "tool".
	Role string `json:"role" api:"required"`
	// Ordinal position of the turn within the call, starting at 0.
	Seq int64 `json:"seq" api:"required"`
	// Transcribed speech for `user` and `assistant` turns, or a JSON summary of the
	// tool call for `tool` turns.
	Text string `json:"text" api:"required"`
	// When the turn ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// When the turn started.
	StartedAt time.Time `json:"startedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Seq         respjson.Field
		Text        respjson.Field
		EndedAt     respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallListResponseTranscript) RawJSON() string { return r.JSON.raw }
func (r *CallListResponseTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallHangupResponse struct {
	Call CallHangupResponseCall `json:"call" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Call        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupResponse) RawJSON() string { return r.JSON.raw }
func (r *CallHangupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallHangupResponseCall struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the call was placed by Zavu (outbound) or received from a caller
	// (inbound).
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Caller phone number in E.164 format. Your sender's number for outbound calls;
	// the caller's number for inbound calls.
	From string `json:"from" api:"required"`
	// Lifecycle status of a voice call.
	//
	// - `queued`: outbound call created, not yet dialing.
	// - `ringing`: dialing (outbound) or received and ringing (inbound).
	// - `in_progress`: answered, the agent is connected.
	// - `completed`: ended after a conversation.
	// - `failed`: could not be completed.
	// - `busy`: the line was busy.
	// - `no_answer`: rang but was not answered.
	// - `canceled`: canceled before it was answered.
	//
	// Any of "queued", "ringing", "in_progress", "completed", "failed", "busy",
	// "no_answer", "canceled".
	Status string `json:"status" api:"required"`
	// Callee phone number in E.164 format.
	To string `json:"to" api:"required"`
	// When the call was answered.
	AnsweredAt time.Time `json:"answeredAt" api:"nullable" format:"date-time"`
	// Total cost of the call in USD, combining the managed voice pipeline per-minute
	// charge and telephony. Available once the call has ended.
	Cost float64 `json:"cost" api:"nullable"`
	// Billable talk time in seconds, measured from answer to hangup.
	DurationSeconds int64 `json:"durationSeconds" api:"nullable"`
	// When the call ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// Why the call ended (e.g. `agent_ended`, `max_duration`, `transfer`, `hangup`).
	// Present once the call is no longer active.
	EndReason string `json:"endReason" api:"nullable"`
	// Arbitrary metadata you attached when creating the call.
	Metadata map[string]string `json:"metadata"`
	// Ordered transcript of the call. Included when retrieving a single call; omitted
	// from list responses.
	Transcript []CallHangupResponseCallTranscript `json:"transcript"`
	// Number of conversation turns exchanged during the call.
	TurnCount int64     `json:"turnCount" api:"nullable"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Direction       respjson.Field
		From            respjson.Field
		Status          respjson.Field
		To              respjson.Field
		AnsweredAt      respjson.Field
		Cost            respjson.Field
		DurationSeconds respjson.Field
		EndedAt         respjson.Field
		EndReason       respjson.Field
		Metadata        respjson.Field
		Transcript      respjson.Field
		TurnCount       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupResponseCall) RawJSON() string { return r.JSON.raw }
func (r *CallHangupResponseCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single turn in a voice call transcript.
type CallHangupResponseCallTranscript struct {
	// Who produced the turn. `tool` records a tool call the agent made during the
	// conversation.
	//
	// Any of "user", "assistant", "tool".
	Role string `json:"role" api:"required"`
	// Ordinal position of the turn within the call, starting at 0.
	Seq int64 `json:"seq" api:"required"`
	// Transcribed speech for `user` and `assistant` turns, or a JSON summary of the
	// tool call for `tool` turns.
	Text string `json:"text" api:"required"`
	// When the turn ended.
	EndedAt time.Time `json:"endedAt" api:"nullable" format:"date-time"`
	// When the turn started.
	StartedAt time.Time `json:"startedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Seq         respjson.Field
		Text        respjson.Field
		EndedAt     respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupResponseCallTranscript) RawJSON() string { return r.JSON.raw }
func (r *CallHangupResponseCallTranscript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewParams struct {
	// Recipient phone number in E.164 format.
	To string `json:"to" api:"required"`
	// Overrides the agent's configured greeting for this call only.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// Language the agent speaks on this call only, as a BCP-47 tag (`en`, `es`,
	// `es-ES`, `pt-BR`), or `auto` to detect the caller's language and follow it.
	// Overrides the agent's configured language for speech recognition, the agent's
	// replies, and the synthesized voice. If the agent uses a custom voice you
	// supplied, that voice is kept and only the language changes. When omitted, the
	// agent's configured language is used.
	Language param.Opt[string] `json:"language,omitzero"`
	// Overrides the agent's maximum call duration for this call only.
	MaxDurationMinutes param.Opt[int64] `json:"maxDurationMinutes,omitzero"`
	// Sender profile that places the call. Uses the project's default sender if
	// omitted. The sender's agent must have voice enabled.
	SenderID param.Opt[string] `json:"senderId,omitzero"`
	// Arbitrary metadata to associate with the call. Returned on the call object and
	// included in voice webhooks.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r CallNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CallNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Whether the call was placed by Zavu (outbound) or received from a caller
	// (inbound).
	//
	// Any of "inbound", "outbound".
	Direction CallListParamsDirection `query:"direction,omitzero" json:"-"`
	// Lifecycle status of a voice call.
	//
	// - `queued`: outbound call created, not yet dialing.
	// - `ringing`: dialing (outbound) or received and ringing (inbound).
	// - `in_progress`: answered, the agent is connected.
	// - `completed`: ended after a conversation.
	// - `failed`: could not be completed.
	// - `busy`: the line was busy.
	// - `no_answer`: rang but was not answered.
	// - `canceled`: canceled before it was answered.
	//
	// Any of "queued", "ringing", "in_progress", "completed", "failed", "busy",
	// "no_answer", "canceled".
	Status CallListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallListParams]'s query parameters as `url.Values`.
func (r CallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether the call was placed by Zavu (outbound) or received from a caller
// (inbound).
type CallListParamsDirection string

const (
	CallListParamsDirectionInbound  CallListParamsDirection = "inbound"
	CallListParamsDirectionOutbound CallListParamsDirection = "outbound"
)

// Lifecycle status of a voice call.
//
// - `queued`: outbound call created, not yet dialing.
// - `ringing`: dialing (outbound) or received and ringing (inbound).
// - `in_progress`: answered, the agent is connected.
// - `completed`: ended after a conversation.
// - `failed`: could not be completed.
// - `busy`: the line was busy.
// - `no_answer`: rang but was not answered.
// - `canceled`: canceled before it was answered.
type CallListParamsStatus string

const (
	CallListParamsStatusQueued     CallListParamsStatus = "queued"
	CallListParamsStatusRinging    CallListParamsStatus = "ringing"
	CallListParamsStatusInProgress CallListParamsStatus = "in_progress"
	CallListParamsStatusCompleted  CallListParamsStatus = "completed"
	CallListParamsStatusFailed     CallListParamsStatus = "failed"
	CallListParamsStatusBusy       CallListParamsStatus = "busy"
	CallListParamsStatusNoAnswer   CallListParamsStatus = "no_answer"
	CallListParamsStatusCanceled   CallListParamsStatus = "canceled"
)
