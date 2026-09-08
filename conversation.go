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

// ConversationService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	Options []option.RequestOption
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.Options = opts
	return
}

// Get conversation
func (r *ConversationService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *ConversationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s", url.PathEscape(conversationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List inbox threads, most recently active first. A conversation groups every
// message with one contact across channels, which is what you need to build an
// inbox: `GET /v1/messages` returns a flat log with no thread to hang it on.
//
// Use `senderId` to scope the list to a single number, and `channel` to keep only
// threads that have carried that channel.
func (r *ConversationService) List(ctx context.Context, query ConversationListParams, opts ...option.RequestOption) (res *pagination.Cursor[ConversationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/conversations"
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

// List inbox threads, most recently active first. A conversation groups every
// message with one contact across channels, which is what you need to build an
// inbox: `GET /v1/messages` returns a flat log with no thread to hang it on.
//
// Use `senderId` to scope the list to a single number, and `channel` to keep only
// threads that have carried that channel.
func (r *ConversationService) ListAutoPaging(ctx context.Context, query ConversationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ConversationListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Messages in this thread, newest first, across every channel it has carried.
// Reply with `POST /v1/messages`, passing the conversation's `senderId` as the
// `Zavu-Sender` header so the answer leaves from the number the contact already
// knows.
func (r *ConversationService) ListMessages(ctx context.Context, conversationID string, query ConversationListMessagesParams, opts ...option.RequestOption) (res *pagination.Cursor[Message], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/messages", url.PathEscape(conversationID))
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

// Messages in this thread, newest first, across every channel it has carried.
// Reply with `POST /v1/messages`, passing the conversation's `senderId` as the
// `Zavu-Sender` header so the answer leaves from the number the contact already
// knows.
func (r *ConversationService) ListMessagesAutoPaging(ctx context.Context, conversationID string, query ConversationListMessagesParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Message] {
	return pagination.NewCursorAutoPager(r.ListMessages(ctx, conversationID, query, opts...))
}

// Reset the thread's `unreadCount` to zero. Marks the thread read in your own
// inbox only: it does not send a read receipt to the contact.
func (r *ConversationService) MarkAsRead(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *ConversationMarkAsReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/read", url.PathEscape(conversationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ConversationGetResponse struct {
	// An inbox thread with one contact. A conversation groups every message exchanged
	// with that contact across channels, so a contact who writes on WhatsApp and later
	// by email stays in one thread.
	Conversation ConversationGetResponseConversation `json:"conversation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversation respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An inbox thread with one contact. A conversation groups every message exchanged
// with that contact across channels, so a contact who writes on WhatsApp and later
// by email stays in one thread.
type ConversationGetResponseConversation struct {
	ID string `json:"id" api:"required"`
	// Every channel this thread has carried messages on.
	Channels []string `json:"channels" api:"required"`
	// The key this thread is filed under: a phone number in E.164, a WhatsApp
	// business-scoped user ID (BSUID), a numeric chat ID
	// (Telegram/Instagram/Messenger), or a group JID. It is not always a phone number,
	// so do not parse it as one.
	ContactIdentifier string    `json:"contactIdentifier" api:"required"`
	CreatedAt         time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Denormalized preview of the most recent message, so a thread list needs no extra
	// fetch.
	LastMessage  ConversationGetResponseConversationLastMessage `json:"lastMessage" api:"required"`
	MessageCount int64                                          `json:"messageCount" api:"required"`
	// Inbound messages not yet marked read. Reset with POST
	// /v1/conversations/{conversationId}/read.
	UnreadCount int64     `json:"unreadCount" api:"required"`
	UpdatedAt   time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the contact this thread belongs to. Absent on group threads and on threads
	// whose contact has not been resolved yet.
	ContactID string `json:"contactId"`
	// Email address of the thread, when the contact was reached by email.
	Email string `json:"email"`
	// Present when the thread is a group chat rather than a one-to-one conversation.
	Group ConversationGetResponseConversationGroup `json:"group"`
	// Sender that last handled this thread. Use it as the `Zavu-Sender` header when
	// replying so the answer leaves from the same number the contact knows.
	SenderID string `json:"senderId"`
	// WhatsApp identity, present when the contact adopted a username.
	Whatsapp ConversationGetResponseConversationWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Channels          respjson.Field
		ContactIdentifier respjson.Field
		CreatedAt         respjson.Field
		LastMessage       respjson.Field
		MessageCount      respjson.Field
		UnreadCount       respjson.Field
		UpdatedAt         respjson.Field
		ContactID         respjson.Field
		Email             respjson.Field
		Group             respjson.Field
		SenderID          respjson.Field
		Whatsapp          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationGetResponseConversation) RawJSON() string { return r.JSON.raw }
func (r *ConversationGetResponseConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Denormalized preview of the most recent message, so a thread list needs no extra
// fetch.
type ConversationGetResponseConversationLastMessage struct {
	ID string    `json:"id" api:"required"`
	At time.Time `json:"at" api:"required" format:"date-time"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "messenger", "voice".
	Channel Channel `json:"channel" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Text or caption. Empty when the last message carried no text (e.g. media).
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		At          respjson.Field
		Channel     respjson.Field
		Direction   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationGetResponseConversationLastMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationGetResponseConversationLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present when the thread is a group chat rather than a one-to-one conversation.
type ConversationGetResponseConversationGroup struct {
	ID               string `json:"id" api:"required"`
	ParticipantCount int64  `json:"participantCount"`
	Subject          string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ParticipantCount respjson.Field
		Subject          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationGetResponseConversationGroup) RawJSON() string { return r.JSON.raw }
func (r *ConversationGetResponseConversationGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp identity, present when the contact adopted a username.
type ConversationGetResponseConversationWhatsapp struct {
	// Business-scoped user ID. Can be used as `to` when sending.
	Bsuid    string `json:"bsuid"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bsuid       respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationGetResponseConversationWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *ConversationGetResponseConversationWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An inbox thread with one contact. A conversation groups every message exchanged
// with that contact across channels, so a contact who writes on WhatsApp and later
// by email stays in one thread.
type ConversationListResponse struct {
	ID string `json:"id" api:"required"`
	// Every channel this thread has carried messages on.
	Channels []string `json:"channels" api:"required"`
	// The key this thread is filed under: a phone number in E.164, a WhatsApp
	// business-scoped user ID (BSUID), a numeric chat ID
	// (Telegram/Instagram/Messenger), or a group JID. It is not always a phone number,
	// so do not parse it as one.
	ContactIdentifier string    `json:"contactIdentifier" api:"required"`
	CreatedAt         time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Denormalized preview of the most recent message, so a thread list needs no extra
	// fetch.
	LastMessage  ConversationListResponseLastMessage `json:"lastMessage" api:"required"`
	MessageCount int64                               `json:"messageCount" api:"required"`
	// Inbound messages not yet marked read. Reset with POST
	// /v1/conversations/{conversationId}/read.
	UnreadCount int64     `json:"unreadCount" api:"required"`
	UpdatedAt   time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the contact this thread belongs to. Absent on group threads and on threads
	// whose contact has not been resolved yet.
	ContactID string `json:"contactId"`
	// Email address of the thread, when the contact was reached by email.
	Email string `json:"email"`
	// Present when the thread is a group chat rather than a one-to-one conversation.
	Group ConversationListResponseGroup `json:"group"`
	// Sender that last handled this thread. Use it as the `Zavu-Sender` header when
	// replying so the answer leaves from the same number the contact knows.
	SenderID string `json:"senderId"`
	// WhatsApp identity, present when the contact adopted a username.
	Whatsapp ConversationListResponseWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Channels          respjson.Field
		ContactIdentifier respjson.Field
		CreatedAt         respjson.Field
		LastMessage       respjson.Field
		MessageCount      respjson.Field
		UnreadCount       respjson.Field
		UpdatedAt         respjson.Field
		ContactID         respjson.Field
		Email             respjson.Field
		Group             respjson.Field
		SenderID          respjson.Field
		Whatsapp          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Denormalized preview of the most recent message, so a thread list needs no extra
// fetch.
type ConversationListResponseLastMessage struct {
	ID string    `json:"id" api:"required"`
	At time.Time `json:"at" api:"required" format:"date-time"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "messenger", "voice".
	Channel Channel `json:"channel" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Text or caption. Empty when the last message carried no text (e.g. media).
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		At          respjson.Field
		Channel     respjson.Field
		Direction   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseLastMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present when the thread is a group chat rather than a one-to-one conversation.
type ConversationListResponseGroup struct {
	ID               string `json:"id" api:"required"`
	ParticipantCount int64  `json:"participantCount"`
	Subject          string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ParticipantCount respjson.Field
		Subject          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseGroup) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp identity, present when the contact adopted a username.
type ConversationListResponseWhatsapp struct {
	// Business-scoped user ID. Can be used as `to` when sending.
	Bsuid    string `json:"bsuid"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bsuid       respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationListResponseWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *ConversationListResponseWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationMarkAsReadResponse struct {
	// An inbox thread with one contact. A conversation groups every message exchanged
	// with that contact across channels, so a contact who writes on WhatsApp and later
	// by email stays in one thread.
	Conversation ConversationMarkAsReadResponseConversation `json:"conversation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversation respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMarkAsReadResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationMarkAsReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An inbox thread with one contact. A conversation groups every message exchanged
// with that contact across channels, so a contact who writes on WhatsApp and later
// by email stays in one thread.
type ConversationMarkAsReadResponseConversation struct {
	ID string `json:"id" api:"required"`
	// Every channel this thread has carried messages on.
	Channels []string `json:"channels" api:"required"`
	// The key this thread is filed under: a phone number in E.164, a WhatsApp
	// business-scoped user ID (BSUID), a numeric chat ID
	// (Telegram/Instagram/Messenger), or a group JID. It is not always a phone number,
	// so do not parse it as one.
	ContactIdentifier string    `json:"contactIdentifier" api:"required"`
	CreatedAt         time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Denormalized preview of the most recent message, so a thread list needs no extra
	// fetch.
	LastMessage  ConversationMarkAsReadResponseConversationLastMessage `json:"lastMessage" api:"required"`
	MessageCount int64                                                 `json:"messageCount" api:"required"`
	// Inbound messages not yet marked read. Reset with POST
	// /v1/conversations/{conversationId}/read.
	UnreadCount int64     `json:"unreadCount" api:"required"`
	UpdatedAt   time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the contact this thread belongs to. Absent on group threads and on threads
	// whose contact has not been resolved yet.
	ContactID string `json:"contactId"`
	// Email address of the thread, when the contact was reached by email.
	Email string `json:"email"`
	// Present when the thread is a group chat rather than a one-to-one conversation.
	Group ConversationMarkAsReadResponseConversationGroup `json:"group"`
	// Sender that last handled this thread. Use it as the `Zavu-Sender` header when
	// replying so the answer leaves from the same number the contact knows.
	SenderID string `json:"senderId"`
	// WhatsApp identity, present when the contact adopted a username.
	Whatsapp ConversationMarkAsReadResponseConversationWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Channels          respjson.Field
		ContactIdentifier respjson.Field
		CreatedAt         respjson.Field
		LastMessage       respjson.Field
		MessageCount      respjson.Field
		UnreadCount       respjson.Field
		UpdatedAt         respjson.Field
		ContactID         respjson.Field
		Email             respjson.Field
		Group             respjson.Field
		SenderID          respjson.Field
		Whatsapp          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMarkAsReadResponseConversation) RawJSON() string { return r.JSON.raw }
func (r *ConversationMarkAsReadResponseConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Denormalized preview of the most recent message, so a thread list needs no extra
// fetch.
type ConversationMarkAsReadResponseConversationLastMessage struct {
	ID string    `json:"id" api:"required"`
	At time.Time `json:"at" api:"required" format:"date-time"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "messenger", "voice".
	Channel Channel `json:"channel" api:"required"`
	// Any of "inbound", "outbound".
	Direction string `json:"direction" api:"required"`
	// Text or caption. Empty when the last message carried no text (e.g. media).
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		At          respjson.Field
		Channel     respjson.Field
		Direction   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMarkAsReadResponseConversationLastMessage) RawJSON() string { return r.JSON.raw }
func (r *ConversationMarkAsReadResponseConversationLastMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Present when the thread is a group chat rather than a one-to-one conversation.
type ConversationMarkAsReadResponseConversationGroup struct {
	ID               string `json:"id" api:"required"`
	ParticipantCount int64  `json:"participantCount"`
	Subject          string `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ParticipantCount respjson.Field
		Subject          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMarkAsReadResponseConversationGroup) RawJSON() string { return r.JSON.raw }
func (r *ConversationMarkAsReadResponseConversationGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp identity, present when the contact adopted a username.
type ConversationMarkAsReadResponseConversationWhatsapp struct {
	// Business-scoped user ID. Can be used as `to` when sending.
	Bsuid    string `json:"bsuid"`
	Username string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bsuid       respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationMarkAsReadResponseConversationWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *ConversationMarkAsReadResponseConversationWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationListParams struct {
	// Opaque cursor from a previous response's `nextCursor`. Do not construct it.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Search threads by identity: phone number (any format — `+1 (555) 123-4567` and
	// `15551234567` both match), email address (full or local part), WhatsApp group
	// subject, WhatsApp username, or BSUID. Matching is by whole word, with prefix
	// matching on the last term, so `mar` finds `maria@example.com` and `+1555` finds
	// `+15551234567`; a fragment from the middle or end of a number (`4567`) does not
	// match.
	//
	// It does **not** search message bodies — only who the thread is with.
	//
	// Results come back ranked by relevance rather than by recency, so the usual "most
	// recently active first" ordering does not apply while `q` is set. `senderId` and
	// `channel` still narrow the results, and `cursor` paginates them as usual. An
	// empty or whitespace-only `q` returns no items rather than the full list.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Keep only threads last handled by this sender.
	SenderID param.Opt[string] `query:"senderId,omitzero" json:"-"`
	// Keep only threads that have carried this channel.
	//
	// Any of "sms", "sms_oneway", "whatsapp", "email", "telegram", "instagram",
	// "messenger", "voice".
	Channel ConversationListParamsChannel `query:"channel,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationListParams]'s query parameters as `url.Values`.
func (r ConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Keep only threads that have carried this channel.
type ConversationListParamsChannel string

const (
	ConversationListParamsChannelSMS       ConversationListParamsChannel = "sms"
	ConversationListParamsChannelSMSOneway ConversationListParamsChannel = "sms_oneway"
	ConversationListParamsChannelWhatsapp  ConversationListParamsChannel = "whatsapp"
	ConversationListParamsChannelEmail     ConversationListParamsChannel = "email"
	ConversationListParamsChannelTelegram  ConversationListParamsChannel = "telegram"
	ConversationListParamsChannelInstagram ConversationListParamsChannel = "instagram"
	ConversationListParamsChannelMessenger ConversationListParamsChannel = "messenger"
	ConversationListParamsChannelVoice     ConversationListParamsChannel = "voice"
)

type ConversationListMessagesParams struct {
	// Opaque cursor from a previous response's `nextCursor`.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationListMessagesParams]'s query parameters as
// `url.Values`.
func (r ConversationListMessagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
