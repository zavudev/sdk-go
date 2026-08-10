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

// SenderService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderService] method instead.
type SenderService struct {
	Options      []option.RequestOption
	Agent        SenderAgentService
	WhatsappSync SenderWhatsappSyncService
}

// NewSenderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSenderService(opts ...option.RequestOption) (r SenderService) {
	r = SenderService{}
	r.Options = opts
	r.Agent = NewSenderAgentService(opts...)
	r.WhatsappSync = NewSenderWhatsappSyncService(opts...)
	return
}

// Create sender
func (r *SenderService) New(ctx context.Context, body SenderNewParams, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/senders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get sender
func (r *SenderService) Get(ctx context.Context, senderID string, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update sender
func (r *SenderService) Update(ctx context.Context, senderID string, body SenderUpdateParams, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List senders
func (r *SenderService) List(ctx context.Context, query SenderListParams, opts ...option.RequestOption) (res *pagination.Cursor[Sender], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/senders"
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

// List senders
func (r *SenderService) ListAutoPaging(ctx context.Context, query SenderListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Sender] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete sender
func (r *SenderService) Delete(ctx context.Context, senderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return err
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get the WhatsApp Business profile for a sender. The sender must have a WhatsApp
// Business Account connected.
func (r *SenderService) GetProfile(ctx context.Context, senderID string, opts ...option.RequestOption) (res *WhatsappBusinessProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/profile", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Regenerate the webhook secret for a sender. The old secret will be invalidated
// immediately.
func (r *SenderService) RegenerateWebhookSecret(ctx context.Context, senderID string, opts ...option.RequestOption) (res *WebhookSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/webhook/secret", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Update the WhatsApp Business profile for a sender. The sender must have a
// WhatsApp Business Account connected.
func (r *SenderService) UpdateProfile(ctx context.Context, senderID string, body SenderUpdateProfileParams, opts ...option.RequestOption) (res *SenderUpdateProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/profile", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Upload a new profile picture for the WhatsApp Business profile. The image will
// be uploaded to Meta and set as the profile picture.
func (r *SenderService) UploadProfilePicture(ctx context.Context, senderID string, body SenderUploadProfilePictureParams, opts ...option.RequestOption) (res *SenderUploadProfilePictureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/profile/picture", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Sender struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Phone number in E.164 format.
	PhoneNumber string `json:"phoneNumber" api:"required"`
	// Channels this sender can actually send on right now, computed from its
	// configuration. Empty means the sender cannot send or receive anything yet: a
	// phoneNumber alone does not enable SMS or voice. Check this rather than inferring
	// capability from phoneNumber or emailAddress.
	Channels  []string  `json:"channels"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// From-address for the email channel, if configured.
	EmailAddress string `json:"emailAddress"`
	// Whether catch-all receiving is enabled. When true (and emailReceivingEnabled is
	// true), this sender receives email addressed to any local part at its domain, not
	// just its own address. The original recipient is delivered in the message.inbound
	// webhook's data.to.
	EmailCatchAllEnabled bool `json:"emailCatchAllEnabled"`
	// Whether inbound email receiving is enabled for this sender.
	EmailReceivingEnabled bool `json:"emailReceivingEnabled"`
	// Whether this sender is the project's default.
	IsDefault bool      `json:"isDefault"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Webhook configuration for the sender.
	Webhook SenderWebhook `json:"webhook"`
	// WhatsApp Business Account information. Only present if a WABA is connected.
	Whatsapp SenderWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Name                  respjson.Field
		PhoneNumber           respjson.Field
		Channels              respjson.Field
		CreatedAt             respjson.Field
		EmailAddress          respjson.Field
		EmailCatchAllEnabled  respjson.Field
		EmailReceivingEnabled respjson.Field
		IsDefault             respjson.Field
		UpdatedAt             respjson.Field
		Webhook               respjson.Field
		Whatsapp              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Sender) RawJSON() string { return r.JSON.raw }
func (r *Sender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp Business Account information. Only present if a WABA is connected.
type SenderWhatsapp struct {
	// Display phone number.
	DisplayPhoneNumber string `json:"displayPhoneNumber"`
	// Payment configuration status from Meta.
	PaymentStatus SenderWhatsappPaymentStatus `json:"paymentStatus"`
	// WhatsApp phone number ID from Meta.
	PhoneNumberID string `json:"phoneNumberId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayPhoneNumber respjson.Field
		PaymentStatus      respjson.Field
		PhoneNumberID      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *SenderWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment configuration status from Meta.
type SenderWhatsappPaymentStatus struct {
	// Whether template messages can be sent. Requires setupStatus=COMPLETE and
	// methodStatus=VALID.
	CanSendTemplates bool `json:"canSendTemplates"`
	// Payment method status (VALID, NONE, etc.).
	MethodStatus string `json:"methodStatus"`
	// Payment setup status (COMPLETE, NOT_STARTED, etc.).
	SetupStatus string `json:"setupStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanSendTemplates respjson.Field
		MethodStatus     respjson.Field
		SetupStatus      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWhatsappPaymentStatus) RawJSON() string { return r.JSON.raw }
func (r *SenderWhatsappPaymentStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook configuration for the sender.
type SenderWebhook struct {
	// Whether the webhook is active.
	Active bool `json:"active" api:"required"`
	// List of events the webhook is subscribed to.
	Events []WebhookEvent `json:"events" api:"required"`
	// Which `X-Zavu-Signature` scheme this receiver is sent.
	//
	//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
	//     configurable. Existing webhooks stay on it until you move them.
	//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
	//     default for new senders. It signs the timestamp together with the body.
	//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
	//     reading either one works, so you can deploy and confirm your new verifier
	//     before switching over.
	//
	// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
	// https://docs.zavu.dev/guides/receiving-messages/signature-migration
	//
	// Any of "v1", "v1+v2", "v2".
	SignatureVersion SenderWebhookSignatureVersion `json:"signatureVersion" api:"required"`
	// HTTPS URL that will receive webhook events.
	URL string `json:"url" api:"required" format:"uri"`
	// Webhook secret for signature verification. Only returned on create or
	// regenerate.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active           respjson.Field
		Events           respjson.Field
		SignatureVersion respjson.Field
		URL              respjson.Field
		Secret           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWebhook) RawJSON() string { return r.JSON.raw }
func (r *SenderWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which `X-Zavu-Signature` scheme this receiver is sent.
//
//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
//     configurable. Existing webhooks stay on it until you move them.
//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
//     default for new senders. It signs the timestamp together with the body.
//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
//     reading either one works, so you can deploy and confirm your new verifier
//     before switching over.
//
// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
// https://docs.zavu.dev/guides/receiving-messages/signature-migration
type SenderWebhookSignatureVersion string

const (
	SenderWebhookSignatureVersionV1   SenderWebhookSignatureVersion = "v1"
	SenderWebhookSignatureVersionV1V2 SenderWebhookSignatureVersion = "v1+v2"
	SenderWebhookSignatureVersionV2   SenderWebhookSignatureVersion = "v2"
)

// Type of event that triggers the webhook.
//
// **Message lifecycle events:**
//
//   - `message.queued`: Message created and queued for sending. `data.status` =
//     `queued`
//   - `message.sent`: Message accepted by the provider. `data.status` = `sent`
//   - `message.delivered`: Message delivered to recipient. `data.status` =
//     `delivered`
//   - `message.read`: Message was read by the recipient (WhatsApp only).
//     `data.status` = `read`
//   - `message.failed`: Message failed to send. `data.status` = `failed`
//
// **Inbound events:**
//
//   - `message.inbound`: New message received from a contact. `data.conversationId`
//     is the inbox thread id (deep-link with
//     `https://dashboard.zavu.dev/{locale}/inbox?conv={conversationId}`); it is
//     `null` while the conversation row is still being created (the first message of
//     a brand-new thread, or several near-simultaneous first messages), where
//     `conversation.new` carries the id instead — `GET /v1/messages/{messageId}`
//     always has it. Reactions are delivered as `message.inbound` with
//     `messageType='reaction'`. When the contact replied to (quoted) an earlier
//     message, `data.content` carries the reply context: `replyToMessageId`,
//     `replyToProviderMessageId`, `replyToFrom`, `replyToText`, and
//     `replyToMessageType`. `data.providerTimestamp` is the provider's original
//     receive time in Unix milliseconds (the moment the channel received the message
//     from the contact — WhatsApp, Telegram, Instagram, Messenger; `null` for SMS
//     and email). Compare it against the top-level `timestamp` (when Zavu dispatched
//     the webhook) to detect and ignore delayed deliveries.
//   - `message.unsupported`: Received a message type that is not supported
//
// **Broadcast events:**
//
//   - `broadcast.status_changed`: Broadcast status changed (pending_review,
//     approved, rejected, sending, completed, cancelled)
//
// **Other events:**
//
//   - `conversation.new`: New conversation started with a contact. `data` carries
//     `conversationId` (the inbox thread id — deep-link with
//     `https://dashboard.zavu.dev/{locale}/inbox?conv={conversationId}`), the
//     `phoneNumber` or `email` key, `channel`, `firstMessageId`, `firstMessageText`,
//     and `profileName`.
//   - `template.status_changed`: WhatsApp template approval status changed
//
// **Partner events:**
//
//   - `invitation.status_changed`: A partner invitation status changed (pending,
//     in_progress, completed, cancelled, failed). `data` carries `invitationId`,
//     `clientName`, `clientEmail`, `connectionType` (`whatsapp_waba` or
//     `messenger`), `previousStatus`, and `currentStatus`. On `completed` it also
//     carries `senderId` and `connectedAccount` (`channel`, `id`, `name`) — the
//     WhatsApp number or Facebook Page that was linked. On `failed` it carries
//     `failureReason`; the invitation link stays usable, so a client can retry it.
//
// **Voice Agent events:** For every voice event, `data` carries `callId`,
// `direction`, `from`, `to`, `status`, `durationSeconds`, `endReason`, and
// `transcriptAvailable`. The terminal events (`call.completed`, `call.failed`)
// additionally carry `cost` — what the call was billed, in USD, combining
// telephony and the managed voice pipeline — and `currency`. They are dispatched
// after the call is charged, so `cost` is populated rather than zero; telephony
// can still be settling on an outbound call, in which case
// `GET /v1/calls/{callId}` holds the reconciled figure.
//
//   - `call.initiated`: An outbound call was created and is dialing, or an inbound
//     call was received. `data.status` = `ringing`
//   - `call.answered`: The call was answered and the voice agent is connected.
//     `data.status` = `in_progress`
//   - `call.completed`: The call ended after a conversation. `data.status` =
//     `completed`; `durationSeconds` and `endReason` describe how it ended, and
//     `transcriptAvailable` indicates whether a transcript can be fetched.
//   - `call.failed`: The call could not be completed (busy, no answer, canceled, or
//     an error). `data.status` is the terminal status and `endReason` explains the
//     cause.
//
// **Custom domain events:**
//
//   - `domain.verified`: A custom email domain passed verification (DKIM, and
//     SPF/DMARC/MAIL FROM if enhanced records are enabled)
//   - `domain.failed`: A custom email domain failed verification or is partially
//     verified
type WebhookEvent string

const (
	WebhookEventMessageQueued           WebhookEvent = "message.queued"
	WebhookEventMessageSent             WebhookEvent = "message.sent"
	WebhookEventMessageDelivered        WebhookEvent = "message.delivered"
	WebhookEventMessageRead             WebhookEvent = "message.read"
	WebhookEventMessageStatus           WebhookEvent = "message.status"
	WebhookEventMessageFailed           WebhookEvent = "message.failed"
	WebhookEventMessageInbound          WebhookEvent = "message.inbound"
	WebhookEventMessageUnsupported      WebhookEvent = "message.unsupported"
	WebhookEventBroadcastStatusChanged  WebhookEvent = "broadcast.status_changed"
	WebhookEventConversationNew         WebhookEvent = "conversation.new"
	WebhookEventTemplateStatusChanged   WebhookEvent = "template.status_changed"
	WebhookEventInvitationStatusChanged WebhookEvent = "invitation.status_changed"
	WebhookEventCallInitiated           WebhookEvent = "call.initiated"
	WebhookEventCallAnswered            WebhookEvent = "call.answered"
	WebhookEventCallCompleted           WebhookEvent = "call.completed"
	WebhookEventCallFailed              WebhookEvent = "call.failed"
	WebhookEventDomainVerified          WebhookEvent = "domain.verified"
	WebhookEventDomainFailed            WebhookEvent = "domain.failed"
)

type WebhookSecretResponse struct {
	// The new webhook secret.
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp Business profile information.
type WhatsappBusinessProfile struct {
	// Short description of the business (max 139 characters).
	About string `json:"about"`
	// Physical address of the business (max 256 characters).
	Address string `json:"address"`
	// Extended description of the business (max 512 characters).
	Description string `json:"description"`
	// Business email address.
	Email string `json:"email" format:"email"`
	// URL of the business profile picture.
	ProfilePictureURL string `json:"profilePictureUrl" format:"uri"`
	// Business category for WhatsApp Business profile.
	//
	// Any of "UNDEFINED", "OTHER", "AUTO", "BEAUTY", "APPAREL", "EDU", "ENTERTAIN",
	// "EVENT_PLAN", "FINANCE", "GROCERY", "GOVT", "HOTEL", "HEALTH", "NONPROFIT",
	// "PROF_SERVICES", "RETAIL", "TRAVEL", "RESTAURANT", "NOT_A_BIZ".
	Vertical WhatsappBusinessProfileVertical `json:"vertical"`
	// Business website URLs (maximum 2).
	Websites []string `json:"websites" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		About             respjson.Field
		Address           respjson.Field
		Description       respjson.Field
		Email             respjson.Field
		ProfilePictureURL respjson.Field
		Vertical          respjson.Field
		Websites          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessProfile) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessProfileResponse struct {
	// WhatsApp Business profile information.
	Profile WhatsappBusinessProfile `json:"profile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business category for WhatsApp Business profile.
type WhatsappBusinessProfileVertical string

const (
	WhatsappBusinessProfileVerticalUndefined    WhatsappBusinessProfileVertical = "UNDEFINED"
	WhatsappBusinessProfileVerticalOther        WhatsappBusinessProfileVertical = "OTHER"
	WhatsappBusinessProfileVerticalAuto         WhatsappBusinessProfileVertical = "AUTO"
	WhatsappBusinessProfileVerticalBeauty       WhatsappBusinessProfileVertical = "BEAUTY"
	WhatsappBusinessProfileVerticalApparel      WhatsappBusinessProfileVertical = "APPAREL"
	WhatsappBusinessProfileVerticalEdu          WhatsappBusinessProfileVertical = "EDU"
	WhatsappBusinessProfileVerticalEntertain    WhatsappBusinessProfileVertical = "ENTERTAIN"
	WhatsappBusinessProfileVerticalEventPlan    WhatsappBusinessProfileVertical = "EVENT_PLAN"
	WhatsappBusinessProfileVerticalFinance      WhatsappBusinessProfileVertical = "FINANCE"
	WhatsappBusinessProfileVerticalGrocery      WhatsappBusinessProfileVertical = "GROCERY"
	WhatsappBusinessProfileVerticalGovt         WhatsappBusinessProfileVertical = "GOVT"
	WhatsappBusinessProfileVerticalHotel        WhatsappBusinessProfileVertical = "HOTEL"
	WhatsappBusinessProfileVerticalHealth       WhatsappBusinessProfileVertical = "HEALTH"
	WhatsappBusinessProfileVerticalNonprofit    WhatsappBusinessProfileVertical = "NONPROFIT"
	WhatsappBusinessProfileVerticalProfServices WhatsappBusinessProfileVertical = "PROF_SERVICES"
	WhatsappBusinessProfileVerticalRetail       WhatsappBusinessProfileVertical = "RETAIL"
	WhatsappBusinessProfileVerticalTravel       WhatsappBusinessProfileVertical = "TRAVEL"
	WhatsappBusinessProfileVerticalRestaurant   WhatsappBusinessProfileVertical = "RESTAURANT"
	WhatsappBusinessProfileVerticalNotABiz      WhatsappBusinessProfileVertical = "NOT_A_BIZ"
)

type SenderUpdateProfileResponse struct {
	// WhatsApp Business profile information.
	Profile WhatsappBusinessProfile `json:"profile" api:"required"`
	Success bool                    `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUpdateProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderUpdateProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderUploadProfilePictureResponse struct {
	// WhatsApp Business profile information.
	Profile WhatsappBusinessProfile `json:"profile" api:"required"`
	Success bool                    `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profile     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderUploadProfilePictureResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderUploadProfilePictureResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderNewParams struct {
	Name string `json:"name" api:"required"`
	// From-address for the email channel (e.g. noreply@yourdomain.com). The address's
	// domain must be a verified email domain in your project. Setting this attaches
	// the email channel to the sender.
	EmailAddress param.Opt[string] `json:"emailAddress,omitzero" format:"email"`
	// ID of the verified email domain to attach. Optional — resolved from
	// `emailAddress`'s domain when omitted.
	EmailDomainID param.Opt[string] `json:"emailDomainId,omitzero"`
	// Display name shown in the recipient's inbox for the email channel.
	EmailFromName param.Opt[string] `json:"emailFromName,omitzero"`
	// Enable inbound email receiving on this sender. Requires a verified MX record on
	// the domain; ignored otherwise.
	EmailReceivingEnabled param.Opt[bool] `json:"emailReceivingEnabled,omitzero"`
	// Enable the one-way SMS channel (`sms_oneway`). Needs nothing else — no phone
	// number, no credential — so it is the fastest way to get a sender that can send.
	// Recipients cannot reply. Confirm with `sms_oneway` in the `channels` array on
	// the response.
	EnableSMSOneway param.Opt[bool] `json:"enableSmsOneway,omitzero"`
	// Let this sender place and answer phone calls. Requires `phoneNumber`; enabling
	// it without one returns 400. Check the `channels` array on the response to
	// confirm `voice` is on.
	EnableVoice param.Opt[bool] `json:"enableVoice,omitzero"`
	// Phone number in E.164 format, and it must be a number your project already owns
	// (see `GET /v1/phone-numbers`). The number is routed to the sender as part of
	// this call, which is what turns the SMS channel on. Passing a number the project
	// does not own, or one already attached to another sender, returns 400 rather than
	// creating a sender that cannot send. Omit for an email-only sender.
	PhoneNumber  param.Opt[string] `json:"phoneNumber,omitzero"`
	SetAsDefault param.Opt[bool]   `json:"setAsDefault,omitzero"`
	// HTTPS URL for webhook events.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// Events to subscribe to.
	WebhookEvents []WebhookEvent `json:"webhookEvents,omitzero"`
	// Which `X-Zavu-Signature` scheme this receiver is sent.
	//
	//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
	//     configurable. Existing webhooks stay on it until you move them.
	//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
	//     default for new senders. It signs the timestamp together with the body.
	//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
	//     reading either one works, so you can deploy and confirm your new verifier
	//     before switching over.
	//
	// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
	// https://docs.zavu.dev/guides/receiving-messages/signature-migration
	//
	// Any of "v1", "v1+v2", "v2".
	WebhookSignatureVersion SenderNewParamsWebhookSignatureVersion `json:"webhookSignatureVersion,omitzero"`
	paramObj
}

func (r SenderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which `X-Zavu-Signature` scheme this receiver is sent.
//
//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
//     configurable. Existing webhooks stay on it until you move them.
//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
//     default for new senders. It signs the timestamp together with the body.
//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
//     reading either one works, so you can deploy and confirm your new verifier
//     before switching over.
//
// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
// https://docs.zavu.dev/guides/receiving-messages/signature-migration
type SenderNewParamsWebhookSignatureVersion string

const (
	SenderNewParamsWebhookSignatureVersionV1   SenderNewParamsWebhookSignatureVersion = "v1"
	SenderNewParamsWebhookSignatureVersionV1V2 SenderNewParamsWebhookSignatureVersion = "v1+v2"
	SenderNewParamsWebhookSignatureVersionV2   SenderNewParamsWebhookSignatureVersion = "v2"
)

type SenderUpdateParams struct {
	// HTTPS URL for webhook events. Set to null to remove webhook.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// Attach or change the sender's email from-address (e.g. noreply@yourdomain.com).
	// The domain must be a verified email domain in your project.
	EmailAddress param.Opt[string] `json:"emailAddress,omitzero" format:"email"`
	// Enable or disable domain catch-all. When enabled (with emailReceivingEnabled
	// true), this sender receives email for any address at its domain. Ignored
	// (treated as false) if receiving is not enabled.
	EmailCatchAllEnabled param.Opt[bool] `json:"emailCatchAllEnabled,omitzero"`
	// ID of the verified email domain to attach. Optional — resolved from
	// `emailAddress`'s domain when omitted.
	EmailDomainID param.Opt[string] `json:"emailDomainId,omitzero"`
	// Display name shown in the recipient's inbox for the email channel.
	EmailFromName param.Opt[string] `json:"emailFromName,omitzero"`
	// Enable or disable inbound email receiving for this sender.
	EmailReceivingEnabled param.Opt[bool] `json:"emailReceivingEnabled,omitzero"`
	// Turn the one-way SMS channel on or off. Enabling needs nothing else and takes
	// effect immediately; disabling removes the channel from the sender. Confirm with
	// the `channels` array on the response.
	EnableSMSOneway param.Opt[bool] `json:"enableSmsOneway,omitzero"`
	// Turn the voice channel on or off. The sender must already have a phone number
	// provisioned for calls; enabling it otherwise returns 400 instead of storing a
	// flag that changes nothing. Confirm with the `channels` array on the response.
	EnableVoice  param.Opt[bool]   `json:"enableVoice,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	SetAsDefault param.Opt[bool]   `json:"setAsDefault,omitzero"`
	// Whether the webhook is active.
	WebhookActive param.Opt[bool] `json:"webhookActive,omitzero"`
	// Events to subscribe to.
	WebhookEvents []WebhookEvent `json:"webhookEvents,omitzero"`
	// Which `X-Zavu-Signature` scheme this receiver is sent.
	//
	//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
	//     configurable. Existing webhooks stay on it until you move them.
	//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
	//     default for new senders. It signs the timestamp together with the body.
	//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
	//     reading either one works, so you can deploy and confirm your new verifier
	//     before switching over.
	//
	// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
	// https://docs.zavu.dev/guides/receiving-messages/signature-migration
	//
	// Any of "v1", "v1+v2", "v2".
	WebhookSignatureVersion SenderUpdateParamsWebhookSignatureVersion `json:"webhookSignatureVersion,omitzero"`
	paramObj
}

func (r SenderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which `X-Zavu-Signature` scheme this receiver is sent.
//
//   - `v1`: `v1=HMAC_SHA256(secret, body)`. The scheme used before this was
//     configurable. Existing webhooks stay on it until you move them.
//   - `v2`: `v2=HMAC_SHA256(secret, "{t}.{body}")`. The current scheme, and the
//     default for new senders. It signs the timestamp together with the body.
//   - `v1+v2`: both signatures, sharing one `t`. The migration setting: a receiver
//     reading either one works, so you can deploy and confirm your new verifier
//     before switching over.
//
// Moving from `v1` straight to `v2` returns `400`. Set `v1+v2` first. See
// https://docs.zavu.dev/guides/receiving-messages/signature-migration
type SenderUpdateParamsWebhookSignatureVersion string

const (
	SenderUpdateParamsWebhookSignatureVersionV1   SenderUpdateParamsWebhookSignatureVersion = "v1"
	SenderUpdateParamsWebhookSignatureVersionV1V2 SenderUpdateParamsWebhookSignatureVersion = "v1+v2"
	SenderUpdateParamsWebhookSignatureVersionV2   SenderUpdateParamsWebhookSignatureVersion = "v2"
)

type SenderListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderListParams]'s query parameters as `url.Values`.
func (r SenderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderUpdateProfileParams struct {
	// Short description of the business (max 139 characters).
	About param.Opt[string] `json:"about,omitzero"`
	// Physical address of the business (max 256 characters).
	Address param.Opt[string] `json:"address,omitzero"`
	// Extended description of the business (max 512 characters).
	Description param.Opt[string] `json:"description,omitzero"`
	// Business email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Business category for WhatsApp Business profile.
	//
	// Any of "UNDEFINED", "OTHER", "AUTO", "BEAUTY", "APPAREL", "EDU", "ENTERTAIN",
	// "EVENT_PLAN", "FINANCE", "GROCERY", "GOVT", "HOTEL", "HEALTH", "NONPROFIT",
	// "PROF_SERVICES", "RETAIL", "TRAVEL", "RESTAURANT", "NOT_A_BIZ".
	Vertical WhatsappBusinessProfileVertical `json:"vertical,omitzero"`
	// Business website URLs (maximum 2).
	Websites []string `json:"websites,omitzero" format:"uri"`
	paramObj
}

func (r SenderUpdateProfileParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateProfileParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateProfileParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderUploadProfilePictureParams struct {
	// URL of the image to upload.
	ImageURL string `json:"imageUrl" api:"required" format:"uri"`
	// MIME type of the image.
	//
	// Any of "image/jpeg", "image/png".
	MimeType SenderUploadProfilePictureParamsMimeType `json:"mimeType,omitzero" api:"required"`
	paramObj
}

func (r SenderUploadProfilePictureParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderUploadProfilePictureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUploadProfilePictureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MIME type of the image.
type SenderUploadProfilePictureParamsMimeType string

const (
	SenderUploadProfilePictureParamsMimeTypeImageJpeg SenderUploadProfilePictureParamsMimeType = "image/jpeg"
	SenderUploadProfilePictureParamsMimeTypeImagePng  SenderUploadProfilePictureParamsMimeType = "image/png"
)
