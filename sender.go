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
	"github.com/stainless-sdks/zavudev-go/internal/apiquery"
	"github.com/stainless-sdks/zavudev-go/internal/requestconfig"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/stainless-sdks/zavudev-go/packages/pagination"
	"github.com/stainless-sdks/zavudev-go/packages/param"
	"github.com/stainless-sdks/zavudev-go/packages/respjson"
)

// SenderService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderService] method instead.
type SenderService struct {
	Options []option.RequestOption
	Agent   SenderAgentService
}

// NewSenderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSenderService(opts ...option.RequestOption) (r SenderService) {
	r = SenderService{}
	r.Options = opts
	r.Agent = NewSenderAgentService(opts...)
	return
}

// Create sender
func (r *SenderService) New(ctx context.Context, body SenderNewParams, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/senders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get sender
func (r *SenderService) Get(ctx context.Context, senderID string, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update sender
func (r *SenderService) Update(ctx context.Context, senderID string, body SenderUpdateParams, opts ...option.RequestOption) (res *Sender, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("v1/senders/%s", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get the WhatsApp Business profile for a sender. The sender must have a WhatsApp
// Business Account connected.
func (r *SenderService) GetProfile(ctx context.Context, senderID string, opts ...option.RequestOption) (res *WhatsappBusinessProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/profile", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Regenerate the webhook secret for a sender. The old secret will be invalidated
// immediately.
func (r *SenderService) RegenerateWebhookSecret(ctx context.Context, senderID string, opts ...option.RequestOption) (res *WebhookSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/webhook/secret", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Update the WhatsApp Business profile for a sender. The sender must have a
// WhatsApp Business Account connected.
func (r *SenderService) UpdateProfile(ctx context.Context, senderID string, body SenderUpdateProfileParams, opts ...option.RequestOption) (res *SenderUpdateProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/profile", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Upload a new profile picture for the WhatsApp Business profile. The image will
// be uploaded to Meta and set as the profile picture.
func (r *SenderService) UploadProfilePicture(ctx context.Context, senderID string, body SenderUploadProfilePictureParams, opts ...option.RequestOption) (res *SenderUploadProfilePictureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/profile/picture", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Sender struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Phone number in E.164 format.
	PhoneNumber string    `json:"phoneNumber" api:"required"`
	CreatedAt   time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt             respjson.Field
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
	// HTTPS URL that will receive webhook events.
	URL string `json:"url" api:"required" format:"uri"`
	// Webhook secret for signature verification. Only returned on create or
	// regenerate.
	Secret string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Events      respjson.Field
		URL         respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWebhook) RawJSON() string { return r.JSON.raw }
func (r *SenderWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of event that triggers the webhook.
//
// **Message lifecycle events:**
//
//   - `message.queued`: Message created and queued for sending. `data.status` =
//     `queued`
//   - `message.sent`: Message accepted by the provider. `data.status` = `sent`
//   - `message.delivered`: Message delivered to recipient. `data.status` =
//     `delivered`
//   - `message.failed`: Message failed to send. `data.status` = `failed`
//
// **Inbound events:**
//
//   - `message.inbound`: New message received from a contact. Reactions are
//     delivered as `message.inbound` with `messageType='reaction'`
//   - `message.unsupported`: Received a message type that is not supported
//
// **Other events:**
//
// - `conversation.new`: New conversation started with a contact
// - `template.status_changed`: WhatsApp template approval status changed
type WebhookEvent string

const (
	WebhookEventMessageQueued         WebhookEvent = "message.queued"
	WebhookEventMessageSent           WebhookEvent = "message.sent"
	WebhookEventMessageDelivered      WebhookEvent = "message.delivered"
	WebhookEventMessageFailed         WebhookEvent = "message.failed"
	WebhookEventMessageInbound        WebhookEvent = "message.inbound"
	WebhookEventMessageUnsupported    WebhookEvent = "message.unsupported"
	WebhookEventConversationNew       WebhookEvent = "conversation.new"
	WebhookEventTemplateStatusChanged WebhookEvent = "template.status_changed"
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
	Name         string          `json:"name" api:"required"`
	PhoneNumber  string          `json:"phoneNumber" api:"required"`
	SetAsDefault param.Opt[bool] `json:"setAsDefault,omitzero"`
	// HTTPS URL for webhook events.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// Events to subscribe to.
	WebhookEvents []WebhookEvent `json:"webhookEvents,omitzero"`
	paramObj
}

func (r SenderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderUpdateParams struct {
	// HTTPS URL for webhook events. Set to null to remove webhook.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// Enable or disable inbound email receiving for this sender.
	EmailReceivingEnabled param.Opt[bool]   `json:"emailReceivingEnabled,omitzero"`
	Name                  param.Opt[string] `json:"name,omitzero"`
	SetAsDefault          param.Opt[bool]   `json:"setAsDefault,omitzero"`
	// Whether the webhook is active.
	WebhookActive param.Opt[bool] `json:"webhookActive,omitzero"`
	// Events to subscribe to.
	WebhookEvents []WebhookEvent `json:"webhookEvents,omitzero"`
	paramObj
}

func (r SenderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
