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

	"github.com/stainless-sdks/zavudev-go/internal/apijson"
	"github.com/stainless-sdks/zavudev-go/internal/apiquery"
	"github.com/stainless-sdks/zavudev-go/internal/requestconfig"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/stainless-sdks/zavudev-go/packages/pagination"
	"github.com/stainless-sdks/zavudev-go/packages/param"
	"github.com/stainless-sdks/zavudev-go/packages/respjson"
)

// MessageService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	return
}

// Get message by ID
func (r *MessageService) Get(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s", url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List messages previously sent by this project.
func (r *MessageService) List(ctx context.Context, query MessageListParams, opts ...option.RequestOption) (res *pagination.Cursor[Message], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/messages"
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

// List messages previously sent by this project.
func (r *MessageService) ListAutoPaging(ctx context.Context, query MessageListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Message] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Send an emoji reaction to an existing WhatsApp message. Reactions are only
// supported for WhatsApp messages.
func (r *MessageService) React(ctx context.Context, messageID string, params MessageReactParams, opts ...option.RequestOption) (res *MessageResponse, err error) {
	if !param.IsOmitted(params.ZavuSender) {
		opts = append(opts, option.WithHeader("Zavu-Sender", fmt.Sprintf("%v", params.ZavuSender.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return
	}
	path := fmt.Sprintf("v1/messages/%s/reactions", url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Send a message to a recipient via SMS or WhatsApp.
//
// **Channel selection:**
//
// - If `channel` is omitted and `messageType` is `text`, defaults to SMS
// - If `messageType` is anything other than `text`, WhatsApp is used automatically
//
// **WhatsApp 24-hour window:**
//
// - Free-form messages (non-template) require an open 24h window
// - Window opens when the user messages you first
// - Use template messages to initiate conversations outside the window
//
// **Email requirements:**
//
//   - Email channel requires KYC verification. Complete identity verification in the
//     dashboard before sending emails.
func (r *MessageService) Send(ctx context.Context, params MessageSendParams, opts ...option.RequestOption) (res *MessageResponse, err error) {
	if !param.IsOmitted(params.ZavuSender) {
		opts = append(opts, option.WithHeader("Zavu-Sender", fmt.Sprintf("%v", params.ZavuSender.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delivery channel. Use 'auto' for intelligent routing.
type Channel string

const (
	ChannelAuto      Channel = "auto"
	ChannelSMS       Channel = "sms"
	ChannelSMSOneway Channel = "sms_oneway"
	ChannelWhatsapp  Channel = "whatsapp"
	ChannelTelegram  Channel = "telegram"
	ChannelEmail     Channel = "email"
	ChannelInstagram Channel = "instagram"
	ChannelVoice     Channel = "voice"
)

type Message struct {
	ID string `json:"id" api:"required"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "voice".
	Channel   Channel   `json:"channel" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Type of message. Non-text types are supported by WhatsApp and Telegram (varies
	// by type).
	//
	// Any of "text", "image", "video", "audio", "document", "sticker", "location",
	// "contact", "buttons", "list", "reaction", "template".
	MessageType MessageType `json:"messageType" api:"required"`
	// Any of "queued", "sending", "sent", "delivered", "failed", "received",
	// "pending_url_verification".
	Status MessageStatus `json:"status" api:"required"`
	To     string        `json:"to" api:"required"`
	// Content for non-text message types (WhatsApp and Telegram).
	Content MessageContent `json:"content"`
	// MAU cost in USD (charged for first contact of the month).
	Cost float64 `json:"cost" api:"nullable"`
	// Provider cost in USD (Telnyx, SES, etc.).
	CostProvider float64 `json:"costProvider" api:"nullable"`
	// Total cost in USD (MAU + provider cost).
	CostTotal    float64           `json:"costTotal" api:"nullable"`
	ErrorCode    string            `json:"errorCode" api:"nullable"`
	ErrorMessage string            `json:"errorMessage" api:"nullable"`
	From         string            `json:"from"`
	Metadata     map[string]string `json:"metadata"`
	// Message ID from the delivery provider.
	ProviderMessageID string `json:"providerMessageId"`
	SenderID          string `json:"senderId"`
	// Text content or caption.
	Text      string    `json:"text"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Channel           respjson.Field
		CreatedAt         respjson.Field
		MessageType       respjson.Field
		Status            respjson.Field
		To                respjson.Field
		Content           respjson.Field
		Cost              respjson.Field
		CostProvider      respjson.Field
		CostTotal         respjson.Field
		ErrorCode         respjson.Field
		ErrorMessage      respjson.Field
		From              respjson.Field
		Metadata          respjson.Field
		ProviderMessageID respjson.Field
		SenderID          respjson.Field
		Text              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content for non-text message types (WhatsApp and Telegram).
type MessageContent struct {
	// Interactive buttons (max 3).
	Buttons []MessageContentButton `json:"buttons"`
	// Contact cards for contact messages.
	Contacts []MessageContentContact `json:"contacts"`
	// Emoji for reaction messages.
	Emoji string `json:"emoji"`
	// Filename for documents.
	Filename string `json:"filename"`
	// Latitude for location messages.
	Latitude float64 `json:"latitude"`
	// Button text for list messages.
	ListButton string `json:"listButton"`
	// Address of the location.
	LocationAddress string `json:"locationAddress"`
	// Name of the location.
	LocationName string `json:"locationName"`
	// Longitude for location messages.
	Longitude float64 `json:"longitude"`
	// WhatsApp media ID if already uploaded.
	MediaID string `json:"mediaId"`
	// URL of the media file (for image, video, audio, document, sticker).
	MediaURL string `json:"mediaUrl"`
	// MIME type of the media.
	MimeType string `json:"mimeType"`
	// Message ID to react to.
	ReactToMessageID string `json:"reactToMessageId"`
	// Sections for list messages.
	Sections []MessageContentSection `json:"sections"`
	// Template ID for template messages.
	TemplateID string `json:"templateId"`
	// Variables for template rendering. Keys are variable positions (1, 2, 3...).
	TemplateVariables map[string]string `json:"templateVariables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons           respjson.Field
		Contacts          respjson.Field
		Emoji             respjson.Field
		Filename          respjson.Field
		Latitude          respjson.Field
		ListButton        respjson.Field
		LocationAddress   respjson.Field
		LocationName      respjson.Field
		Longitude         respjson.Field
		MediaID           respjson.Field
		MediaURL          respjson.Field
		MimeType          respjson.Field
		ReactToMessageID  respjson.Field
		Sections          respjson.Field
		TemplateID        respjson.Field
		TemplateVariables respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContent) RawJSON() string { return r.JSON.raw }
func (r *MessageContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageContent to a MessageContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageContentParam.Overrides()
func (r MessageContent) ToParam() MessageContentParam {
	return param.Override[MessageContentParam](json.RawMessage(r.RawJSON()))
}

type MessageContentButton struct {
	ID    string `json:"id" api:"required"`
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentButton) RawJSON() string { return r.JSON.raw }
func (r *MessageContentButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentContact struct {
	Name   string   `json:"name"`
	Phones []string `json:"phones"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Phones      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentContact) RawJSON() string { return r.JSON.raw }
func (r *MessageContentContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentSection struct {
	Rows  []MessageContentSectionRow `json:"rows" api:"required"`
	Title string                     `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentSection) RawJSON() string { return r.JSON.raw }
func (r *MessageContentSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentSectionRow struct {
	ID          string `json:"id" api:"required"`
	Title       string `json:"title" api:"required"`
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageContentSectionRow) RawJSON() string { return r.JSON.raw }
func (r *MessageContentSectionRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content for non-text message types (WhatsApp and Telegram).
type MessageContentParam struct {
	// Emoji for reaction messages.
	Emoji param.Opt[string] `json:"emoji,omitzero"`
	// Filename for documents.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Latitude for location messages.
	Latitude param.Opt[float64] `json:"latitude,omitzero"`
	// Button text for list messages.
	ListButton param.Opt[string] `json:"listButton,omitzero"`
	// Address of the location.
	LocationAddress param.Opt[string] `json:"locationAddress,omitzero"`
	// Name of the location.
	LocationName param.Opt[string] `json:"locationName,omitzero"`
	// Longitude for location messages.
	Longitude param.Opt[float64] `json:"longitude,omitzero"`
	// WhatsApp media ID if already uploaded.
	MediaID param.Opt[string] `json:"mediaId,omitzero"`
	// URL of the media file (for image, video, audio, document, sticker).
	MediaURL param.Opt[string] `json:"mediaUrl,omitzero"`
	// MIME type of the media.
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	// Message ID to react to.
	ReactToMessageID param.Opt[string] `json:"reactToMessageId,omitzero"`
	// Template ID for template messages.
	TemplateID param.Opt[string] `json:"templateId,omitzero"`
	// Interactive buttons (max 3).
	Buttons []MessageContentButtonParam `json:"buttons,omitzero"`
	// Contact cards for contact messages.
	Contacts []MessageContentContactParam `json:"contacts,omitzero"`
	// Sections for list messages.
	Sections []MessageContentSectionParam `json:"sections,omitzero"`
	// Variables for template rendering. Keys are variable positions (1, 2, 3...).
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r MessageContentParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Title are required.
type MessageContentButtonParam struct {
	ID    string `json:"id" api:"required"`
	Title string `json:"title" api:"required"`
	paramObj
}

func (r MessageContentButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageContentContactParam struct {
	Name   param.Opt[string] `json:"name,omitzero"`
	Phones []string          `json:"phones,omitzero"`
	paramObj
}

func (r MessageContentContactParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Rows, Title are required.
type MessageContentSectionParam struct {
	Rows  []MessageContentSectionRowParam `json:"rows,omitzero" api:"required"`
	Title string                          `json:"title" api:"required"`
	paramObj
}

func (r MessageContentSectionParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentSectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentSectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Title are required.
type MessageContentSectionRowParam struct {
	ID          string            `json:"id" api:"required"`
	Title       string            `json:"title" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r MessageContentSectionRowParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageContentSectionRowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageContentSectionRowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageResponse struct {
	Message Message `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageStatus string

const (
	MessageStatusQueued                 MessageStatus = "queued"
	MessageStatusSending                MessageStatus = "sending"
	MessageStatusSent                   MessageStatus = "sent"
	MessageStatusDelivered              MessageStatus = "delivered"
	MessageStatusFailed                 MessageStatus = "failed"
	MessageStatusReceived               MessageStatus = "received"
	MessageStatusPendingURLVerification MessageStatus = "pending_url_verification"
)

// Type of message. Non-text types are supported by WhatsApp and Telegram (varies
// by type).
type MessageType string

const (
	MessageTypeText     MessageType = "text"
	MessageTypeImage    MessageType = "image"
	MessageTypeVideo    MessageType = "video"
	MessageTypeAudio    MessageType = "audio"
	MessageTypeDocument MessageType = "document"
	MessageTypeSticker  MessageType = "sticker"
	MessageTypeLocation MessageType = "location"
	MessageTypeContact  MessageType = "contact"
	MessageTypeButtons  MessageType = "buttons"
	MessageTypeList     MessageType = "list"
	MessageTypeReaction MessageType = "reaction"
	MessageTypeTemplate MessageType = "template"
)

type MessageListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	To     param.Opt[string] `query:"to,omitzero" json:"-"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "voice".
	Channel Channel `query:"channel,omitzero" json:"-"`
	// Any of "queued", "sending", "sent", "delivered", "failed", "received",
	// "pending_url_verification".
	Status MessageStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageReactParams struct {
	// Single emoji character to react with.
	Emoji      string            `json:"emoji" api:"required"`
	ZavuSender param.Opt[string] `header:"Zavu-Sender,omitzero" json:"-"`
	paramObj
}

func (r MessageReactParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageReactParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageReactParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendParams struct {
	// Recipient phone number in E.164 format or email address.
	To string `json:"to" api:"required"`
	// Whether to enable automatic fallback to SMS if WhatsApp fails. Defaults to true.
	FallbackEnabled param.Opt[bool] `json:"fallbackEnabled,omitzero"`
	// HTML body for email messages. If provided, email will be sent as multipart with
	// both text and HTML.
	HTMLBody param.Opt[string] `json:"htmlBody,omitzero"`
	// Optional idempotency key to avoid duplicate sends.
	IdempotencyKey param.Opt[string] `json:"idempotencyKey,omitzero"`
	// Reply-To email address for email messages.
	ReplyTo param.Opt[string] `json:"replyTo,omitzero" format:"email"`
	// Email subject line. Required when channel is 'email' or recipient is an email
	// address.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Text body for text messages or caption for media messages.
	Text param.Opt[string] `json:"text,omitzero"`
	// Language code for voice text-to-speech (e.g., 'en-US', 'es-ES', 'pt-BR'). If
	// omitted, language is auto-detected from recipient's country code.
	VoiceLanguage param.Opt[string] `json:"voiceLanguage,omitzero"`
	ZavuSender    param.Opt[string] `header:"Zavu-Sender,omitzero" json:"-"`
	// Delivery channel. Use 'auto' for intelligent routing. If omitted with non-text
	// messageType, WhatsApp is used. For email recipients, defaults to 'email'.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "voice".
	Channel Channel `json:"channel,omitzero"`
	// Additional content for non-text message types.
	Content MessageContentParam `json:"content,omitzero"`
	// Type of message. Defaults to 'text'.
	//
	// Any of "text", "image", "video", "audio", "document", "sticker", "location",
	// "contact", "buttons", "list", "reaction", "template".
	MessageType MessageType `json:"messageType,omitzero"`
	// Arbitrary metadata to associate with the message.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
