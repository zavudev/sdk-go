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
		return nil, err
	}
	path := fmt.Sprintf("v1/messages/%s", url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
		return nil, err
	}
	path := fmt.Sprintf("v1/messages/%s/reactions", url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
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
// **Plan allowances and email billing:**
//
//   - WhatsApp, Telegram, Instagram and Messenger share an allowance of 2,000
//     messages per month on Free. Over it, sends return 429 with code
//     `a2p_limit_exceeded` and upgrade details; the counter resets on the 1st of
//     each month. Paid plans have no message caps
//   - Email is billed from your prepaid balance in 1,000-message blocks: $0.40 per
//     1,000 transactional emails, $0.80 per 1,000 marketing (broadcast) emails. A
//     block is charged when your monthly count crosses each 1,000 boundary, and at
//     zero balance email sends return 402 with code `insufficient_balance`. Free
//     teams start with $2 of credit and additionally cap at 3,000 emails/month and
//     100/day. Teams on earlier plans keep their original email quotas instead
//   - SMS and voice are billed per message from your balance on every plan
//
// **Email recipient pre-flight:** Email messages are validated automatically
// before dispatch. Sends that would be a guaranteed hard bounce are failed instead
// of sent, protecting your bounce rate: the message transitions to `failed`
// (visible via `GET /v1/messages/{messageId}` and the `message.failed` webhook)
// with `errorCode` set to `EMAIL_INVALID_RECIPIENT` (malformed address),
// `EMAIL_DOMAIN_NOT_FOUND` (recipient domain has no MX or A records), or
// `EMAIL_RECIPIENT_SUPPRESSED` (address is on your suppression list after a
// previous bounce or complaint). Advisory signals (role addresses, disposable
// domains) do not block sends — check them beforehand with
// `POST /v1/introspect/email`.
func (r *MessageService) Send(ctx context.Context, params MessageSendParams, opts ...option.RequestOption) (res *MessageResponse, err error) {
	if !param.IsOmitted(params.ZavuSender) {
		opts = append(opts, option.WithHeader("Zavu-Sender", fmt.Sprintf("%v", params.ZavuSender.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Mark an inbound WhatsApp message as read and display a typing indicator to the
// user while you prepare a response. The indicator is automatically dismissed when
// you send a reply, or after 25 seconds — whichever comes first. Only valid for
// inbound WhatsApp messages. Use this when a reply will take more than a couple of
// seconds (LLM agent, tool call, lookup) to improve the recipient's experience.
func (r *MessageService) ShowTyping(ctx context.Context, messageID string, body MessageShowTypingParams, opts ...option.RequestOption) (res *MessageShowTypingResponse, err error) {
	if !param.IsOmitted(body.ZavuSender) {
		opts = append(opts, option.WithHeader("Zavu-Sender", fmt.Sprintf("%v", body.ZavuSender.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messages/%s/typing", url.PathEscape(messageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
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
	ChannelMessenger Channel = "messenger"
	ChannelVoice     Channel = "voice"
)

type Message struct {
	ID string `json:"id" api:"required"`
	// Delivery channel. Use 'auto' for intelligent routing.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "messenger", "voice".
	Channel   Channel   `json:"channel" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Type of message. Non-text types are supported by WhatsApp and Telegram (varies
	// by type).
	//
	// `location_request` asks the recipient to share their location and is
	// WhatsApp-only. It takes no `content` object — the prompt goes in `text` (max
	// 1024 characters) and the button label is fixed by WhatsApp. The recipient's
	// answer arrives as an inbound `location` message whose `content.replyToMessageId`
	// is the ID of the request.
	//
	// `request_contact_info` asks the recipient to share their phone number and is
	// WhatsApp-only. Like `location_request` it takes no `content` object — the prompt
	// goes in `text` (max 1024 characters) and WhatsApp renders a fixed **Share
	// Contact Info** button. The answer arrives as an inbound `contact` message. Use
	// it to recover the phone number of a contact who adopted a WhatsApp username and
	// is only known by their business-scoped user ID (BSUID); when they share it, Zavu
	// automatically links the phone number to that contact.
	//
	// Any of "text", "image", "video", "audio", "document", "sticker", "location",
	// "contact", "buttons", "list", "cta_url", "request_contact_info",
	// "location_request", "reaction", "template".
	MessageType MessageType `json:"messageType" api:"required"`
	// Any of "queued", "sending", "sent", "delivered", "read", "failed", "received",
	// "pending_url_verification".
	Status MessageStatus `json:"status" api:"required"`
	To     string        `json:"to" api:"required"`
	// Content for non-text message types (WhatsApp and Telegram).
	Content MessageContent `json:"content"`
	// ID of the conversation (inbox thread) this message belongs to. Use it to build a
	// direct dashboard link:
	// `https://dashboard.zavu.dev/{locale}/inbox?conv={conversationId}`. Omitted only
	// on legacy messages created before conversation threading.
	ConversationID string `json:"conversationId"`
	// Zavu platform charge in USD for this message. Messaging is billed against your
	// plan's monthly limits plus usage-based overage.
	Cost float64 `json:"cost" api:"nullable"`
	// Carrier and delivery cost in USD.
	CostProvider float64 `json:"costProvider" api:"nullable"`
	// Total cost in USD (platform charge + delivery cost).
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
		ConversationID    respjson.Field
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
	// Button label for cta_url messages.
	CtaDisplayText string `json:"ctaDisplayText"`
	// Public HTTPS URL of the header media when ctaHeaderType is 'image', 'video', or
	// 'document'. WhatsApp fetches this URL — it must be publicly reachable and return
	// the declared content type.
	CtaHeaderMediaURL string `json:"ctaHeaderMediaUrl" format:"uri"`
	// Header text when ctaHeaderType is 'text'.
	CtaHeaderText string `json:"ctaHeaderText"`
	// Optional header type for cta_url messages.
	//
	// Any of "text", "image", "video", "document".
	CtaHeaderType MessageContentCtaHeaderType `json:"ctaHeaderType"`
	// Destination URL opened in the device's default browser when the button is
	// tapped. Used with messageType=cta_url. WhatsApp requires HTTPS in production.
	CtaURL string `json:"ctaUrl" format:"uri"`
	// Emoji for reaction messages.
	Emoji string `json:"emoji"`
	// Filename for documents.
	Filename string `json:"filename"`
	// Optional footer text for cta_url messages.
	FooterText string `json:"footerText"`
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
	// Sender of the quoted message (phone number in E.164 format).
	ReplyToFrom string `json:"replyToFrom"`
	// Zavu message ID of the quoted message this message replies to. Present on
	// inbound messages that quote an earlier message. Omitted when the quoted message
	// is not found in Zavu (e.g. an old or unknown message) — use
	// replyToProviderMessageId in that case.
	ReplyToMessageID string `json:"replyToMessageId"`
	// Type of the quoted message (text, image, video, etc.).
	ReplyToMessageType string `json:"replyToMessageType"`
	// Provider message ID (WhatsApp WAMID) of the quoted message. Present whenever an
	// inbound message is a reply, even if the quoted message is not stored in Zavu.
	ReplyToProviderMessageID string `json:"replyToProviderMessageId"`
	// Truncated snippet of the quoted message's text, for display. Empty when the
	// quoted message has no text (e.g. media).
	ReplyToText string `json:"replyToText"`
	// Sections for list messages.
	Sections []MessageContentSection `json:"sections"`
	// Variables for dynamic button placeholders (URL buttons and OTP buttons). Keys
	// are the button index (0, 1, 2) in the template's `buttons` array — not the
	// placeholder name. Values substitute the `{{1}}` placeholder inside that button's
	// URL.
	//
	// **WhatsApp constraints:**
	//
	//   - URL buttons only accept `{{1}}` — positional, numeric, no whitespace, no name.
	//     Named placeholders like `{{token}}` are stored as literal URL text by Meta and
	//     cannot be substituted.
	//   - At most one placeholder per URL button.
	//   - A template may have at most three buttons.
	//   - Static URL buttons (no placeholder) and `quick_reply` buttons are not included
	//     here.
	TemplateButtonVariables map[string]string `json:"templateButtonVariables"`
	// Value for a text-header variable, keyed by `1` (WhatsApp text headers allow at
	// most one variable). Optional override. If omitted, Zavu resolves the header from
	// `templateVariables` using the header placeholder's name (e.g. `novios`). Static
	// text headers need no value.
	TemplateHeaderVariables map[string]string `json:"templateHeaderVariables"`
	// Template ID for template messages.
	TemplateID string `json:"templateId"`
	// Variables for body placeholders. Key them to match the template body: by
	// position (`1`, `2`, ...) for positional templates, or by name (e.g.
	// `customer_name`) for named templates. Zavu detects the template's format and
	// sends the correct payload to Meta. Named keys also resolve a named text-header
	// variable. Do not mix positional and named keys in the same request.
	TemplateVariables map[string]string `json:"templateVariables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons                  respjson.Field
		Contacts                 respjson.Field
		CtaDisplayText           respjson.Field
		CtaHeaderMediaURL        respjson.Field
		CtaHeaderText            respjson.Field
		CtaHeaderType            respjson.Field
		CtaURL                   respjson.Field
		Emoji                    respjson.Field
		Filename                 respjson.Field
		FooterText               respjson.Field
		Latitude                 respjson.Field
		ListButton               respjson.Field
		LocationAddress          respjson.Field
		LocationName             respjson.Field
		Longitude                respjson.Field
		MediaID                  respjson.Field
		MediaURL                 respjson.Field
		MimeType                 respjson.Field
		ReactToMessageID         respjson.Field
		ReplyToFrom              respjson.Field
		ReplyToMessageID         respjson.Field
		ReplyToMessageType       respjson.Field
		ReplyToProviderMessageID respjson.Field
		ReplyToText              respjson.Field
		Sections                 respjson.Field
		TemplateButtonVariables  respjson.Field
		TemplateHeaderVariables  respjson.Field
		TemplateID               respjson.Field
		TemplateVariables        respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
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

// Optional header type for cta_url messages.
type MessageContentCtaHeaderType string

const (
	MessageContentCtaHeaderTypeText     MessageContentCtaHeaderType = "text"
	MessageContentCtaHeaderTypeImage    MessageContentCtaHeaderType = "image"
	MessageContentCtaHeaderTypeVideo    MessageContentCtaHeaderType = "video"
	MessageContentCtaHeaderTypeDocument MessageContentCtaHeaderType = "document"
)

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
	// Button label for cta_url messages.
	CtaDisplayText param.Opt[string] `json:"ctaDisplayText,omitzero"`
	// Public HTTPS URL of the header media when ctaHeaderType is 'image', 'video', or
	// 'document'. WhatsApp fetches this URL — it must be publicly reachable and return
	// the declared content type.
	CtaHeaderMediaURL param.Opt[string] `json:"ctaHeaderMediaUrl,omitzero" format:"uri"`
	// Header text when ctaHeaderType is 'text'.
	CtaHeaderText param.Opt[string] `json:"ctaHeaderText,omitzero"`
	// Destination URL opened in the device's default browser when the button is
	// tapped. Used with messageType=cta_url. WhatsApp requires HTTPS in production.
	CtaURL param.Opt[string] `json:"ctaUrl,omitzero" format:"uri"`
	// Emoji for reaction messages.
	Emoji param.Opt[string] `json:"emoji,omitzero"`
	// Filename for documents.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Optional footer text for cta_url messages.
	FooterText param.Opt[string] `json:"footerText,omitzero"`
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
	// Sender of the quoted message (phone number in E.164 format).
	ReplyToFrom param.Opt[string] `json:"replyToFrom,omitzero"`
	// Zavu message ID of the quoted message this message replies to. Present on
	// inbound messages that quote an earlier message. Omitted when the quoted message
	// is not found in Zavu (e.g. an old or unknown message) — use
	// replyToProviderMessageId in that case.
	ReplyToMessageID param.Opt[string] `json:"replyToMessageId,omitzero"`
	// Type of the quoted message (text, image, video, etc.).
	ReplyToMessageType param.Opt[string] `json:"replyToMessageType,omitzero"`
	// Provider message ID (WhatsApp WAMID) of the quoted message. Present whenever an
	// inbound message is a reply, even if the quoted message is not stored in Zavu.
	ReplyToProviderMessageID param.Opt[string] `json:"replyToProviderMessageId,omitzero"`
	// Truncated snippet of the quoted message's text, for display. Empty when the
	// quoted message has no text (e.g. media).
	ReplyToText param.Opt[string] `json:"replyToText,omitzero"`
	// Template ID for template messages.
	TemplateID param.Opt[string] `json:"templateId,omitzero"`
	// Interactive buttons (max 3).
	Buttons []MessageContentButtonParam `json:"buttons,omitzero"`
	// Contact cards for contact messages.
	Contacts []MessageContentContactParam `json:"contacts,omitzero"`
	// Optional header type for cta_url messages.
	//
	// Any of "text", "image", "video", "document".
	CtaHeaderType MessageContentCtaHeaderType `json:"ctaHeaderType,omitzero"`
	// Sections for list messages.
	Sections []MessageContentSectionParam `json:"sections,omitzero"`
	// Variables for dynamic button placeholders (URL buttons and OTP buttons). Keys
	// are the button index (0, 1, 2) in the template's `buttons` array — not the
	// placeholder name. Values substitute the `{{1}}` placeholder inside that button's
	// URL.
	//
	// **WhatsApp constraints:**
	//
	//   - URL buttons only accept `{{1}}` — positional, numeric, no whitespace, no name.
	//     Named placeholders like `{{token}}` are stored as literal URL text by Meta and
	//     cannot be substituted.
	//   - At most one placeholder per URL button.
	//   - A template may have at most three buttons.
	//   - Static URL buttons (no placeholder) and `quick_reply` buttons are not included
	//     here.
	TemplateButtonVariables map[string]string `json:"templateButtonVariables,omitzero"`
	// Value for a text-header variable, keyed by `1` (WhatsApp text headers allow at
	// most one variable). Optional override. If omitted, Zavu resolves the header from
	// `templateVariables` using the header placeholder's name (e.g. `novios`). Static
	// text headers need no value.
	TemplateHeaderVariables map[string]string `json:"templateHeaderVariables,omitzero"`
	// Variables for body placeholders. Key them to match the template body: by
	// position (`1`, `2`, ...) for positional templates, or by name (e.g.
	// `customer_name`) for named templates. Zavu detects the template's format and
	// sends the correct payload to Meta. Named keys also resolve a named text-header
	// variable. Do not mix positional and named keys in the same request.
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
	MessageStatusRead                   MessageStatus = "read"
	MessageStatusFailed                 MessageStatus = "failed"
	MessageStatusReceived               MessageStatus = "received"
	MessageStatusPendingURLVerification MessageStatus = "pending_url_verification"
)

// Type of message. Non-text types are supported by WhatsApp and Telegram (varies
// by type).
//
// `location_request` asks the recipient to share their location and is
// WhatsApp-only. It takes no `content` object — the prompt goes in `text` (max
// 1024 characters) and the button label is fixed by WhatsApp. The recipient's
// answer arrives as an inbound `location` message whose `content.replyToMessageId`
// is the ID of the request.
//
// `request_contact_info` asks the recipient to share their phone number and is
// WhatsApp-only. Like `location_request` it takes no `content` object — the prompt
// goes in `text` (max 1024 characters) and WhatsApp renders a fixed **Share
// Contact Info** button. The answer arrives as an inbound `contact` message. Use
// it to recover the phone number of a contact who adopted a WhatsApp username and
// is only known by their business-scoped user ID (BSUID); when they share it, Zavu
// automatically links the phone number to that contact.
type MessageType string

const (
	MessageTypeText               MessageType = "text"
	MessageTypeImage              MessageType = "image"
	MessageTypeVideo              MessageType = "video"
	MessageTypeAudio              MessageType = "audio"
	MessageTypeDocument           MessageType = "document"
	MessageTypeSticker            MessageType = "sticker"
	MessageTypeLocation           MessageType = "location"
	MessageTypeContact            MessageType = "contact"
	MessageTypeButtons            MessageType = "buttons"
	MessageTypeList               MessageType = "list"
	MessageTypeCtaURL             MessageType = "cta_url"
	MessageTypeRequestContactInfo MessageType = "request_contact_info"
	MessageTypeLocationRequest    MessageType = "location_request"
	MessageTypeReaction           MessageType = "reaction"
	MessageTypeTemplate           MessageType = "template"
)

type MessageShowTypingResponse struct {
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageShowTypingResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageShowTypingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	To     param.Opt[string] `query:"to,omitzero" json:"-"`
	// Filter by delivery channel.
	//
	// Any of "sms", "sms_oneway", "whatsapp", "email", "telegram", "instagram",
	// "messenger", "voice".
	Channel MessageListParamsChannel `query:"channel,omitzero" json:"-"`
	// Filter by status. Not all stored statuses are filterable.
	//
	// Any of "queued", "sending", "sent", "delivered", "failed", "received".
	Status MessageListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageListParams]'s query parameters as `url.Values`.
func (r MessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by delivery channel.
type MessageListParamsChannel string

const (
	MessageListParamsChannelSMS       MessageListParamsChannel = "sms"
	MessageListParamsChannelSMSOneway MessageListParamsChannel = "sms_oneway"
	MessageListParamsChannelWhatsapp  MessageListParamsChannel = "whatsapp"
	MessageListParamsChannelEmail     MessageListParamsChannel = "email"
	MessageListParamsChannelTelegram  MessageListParamsChannel = "telegram"
	MessageListParamsChannelInstagram MessageListParamsChannel = "instagram"
	MessageListParamsChannelMessenger MessageListParamsChannel = "messenger"
	MessageListParamsChannelVoice     MessageListParamsChannel = "voice"
)

// Filter by status. Not all stored statuses are filterable.
type MessageListParamsStatus string

const (
	MessageListParamsStatusQueued    MessageListParamsStatus = "queued"
	MessageListParamsStatusSending   MessageListParamsStatus = "sending"
	MessageListParamsStatusSent      MessageListParamsStatus = "sent"
	MessageListParamsStatusDelivered MessageListParamsStatus = "delivered"
	MessageListParamsStatusFailed    MessageListParamsStatus = "failed"
	MessageListParamsStatusReceived  MessageListParamsStatus = "received"
)

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
	// Recipient phone number in E.164 format, email address, WhatsApp business-scoped
	// user ID (BSUID, e.g. `US.13491208655302741918`), or numeric chat ID (for
	// Telegram/Instagram/Messenger). A BSUID is routed to WhatsApp and sent via the
	// `recipient` field; use it to message a contact who adopted a username and whose
	// phone number is hidden.
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
	// Email attachments. Only supported when channel is 'email'. Maximum 40MB total
	// size.
	Attachments []MessageSendParamsAttachment `json:"attachments,omitzero"`
	// Delivery channel. Use 'auto' for intelligent routing. If omitted, channel is
	// auto-selected based on sender capabilities and recipient type. For email
	// recipients, defaults to 'email'.
	//
	// Any of "auto", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "messenger", "voice".
	Channel Channel `json:"channel,omitzero"`
	// Additional content for non-text message types.
	Content MessageContentParam `json:"content,omitzero"`
	// Type of message. Defaults to 'text'.
	//
	// Any of "text", "image", "video", "audio", "document", "sticker", "location",
	// "contact", "buttons", "list", "cta_url", "request_contact_info",
	// "location_request", "reaction", "template".
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

// Email attachment. Provide either `content` (base64) or `path` (URL), not both.
//
// The property Filename is required.
type MessageSendParamsAttachment struct {
	// Name of the attached file.
	Filename string `json:"filename" api:"required"`
	// Content of the attached file as a Base64-encoded string.
	Content param.Opt[string] `json:"content,omitzero"`
	// Content ID for inline images. Reference in HTML as
	// `<img src="cid:your_content_id">`.
	ContentID param.Opt[string] `json:"content_id,omitzero"`
	// MIME type of the attachment. If not set, will be derived from the filename.
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// URL where the attachment file is hosted. The server will fetch the file.
	Path param.Opt[string] `json:"path,omitzero" format:"uri"`
	paramObj
}

func (r MessageSendParamsAttachment) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParamsAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParamsAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageShowTypingParams struct {
	ZavuSender param.Opt[string] `header:"Zavu-Sender,omitzero" json:"-"`
	paramObj
}
