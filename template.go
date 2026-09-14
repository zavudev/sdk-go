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

// TemplateService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Create a WhatsApp message template. Note: Templates must be approved by Meta
// before use.
func (r *TemplateService) New(ctx context.Context, body TemplateNewParams, opts ...option.RequestOption) (res *Template, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get template
func (r *TemplateService) Get(ctx context.Context, templateID string, opts ...option.RequestOption) (res *Template, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/templates/%s", url.PathEscape(templateID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List WhatsApp message templates for this project.
func (r *TemplateService) List(ctx context.Context, query TemplateListParams, opts ...option.RequestOption) (res *pagination.Cursor[Template], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/templates"
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

// List WhatsApp message templates for this project.
func (r *TemplateService) ListAutoPaging(ctx context.Context, query TemplateListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Template] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete template
func (r *TemplateService) Delete(ctx context.Context, templateID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return err
	}
	path := fmt.Sprintf("v1/templates/%s", url.PathEscape(templateID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Submit a WhatsApp template to Meta for approval. The template must be in draft
// status and associated with a sender that has a WhatsApp Business Account
// configured.
func (r *TemplateService) Submit(ctx context.Context, templateID string, body TemplateSubmitParams, opts ...option.RequestOption) (res *Template, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateID == "" {
		err = errors.New("missing required templateId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/templates/%s/submit", url.PathEscape(templateID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reconcile this project's templates against WhatsApp. Two things happen per
// connected WhatsApp Business Account: templates that exist on Meta but not in
// Zavu are imported (or linked to an existing template with the same name), and
// the approval status of the templates Zavu already knows about is refreshed from
// Meta.
//
// This is what to call when a template was created outside Zavu — in Meta Business
// Manager, or by another tool — or when a `template.status_changed` webhook was
// missed and a template is stuck in `pending`. Status changes normally arrive by
// webhook; this endpoint is the recovery path and the only path for a template
// Zavu never created.
//
// Templates that Meta reports as rejected or disabled are not imported; they are
// counted in `skipped`. Existing local templates are matched first by Meta
// template ID, then by name.
//
// By default every sender in the project with a WhatsApp Business Account is
// synced. Pass `senderId` to sync only that sender's account. The call is
// synchronous — it waits for Meta and returns what changed — so it can take a few
// seconds per account. A failure on one account does not fail the request: it is
// reported in `errors` and the remaining accounts are still synced.
func (r *TemplateService) Sync(ctx context.Context, body TemplateSyncParams, opts ...option.RequestOption) (res *TemplateSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/templates/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Template struct {
	ID string `json:"id" api:"required"`
	// Default template body with variables: positional ({{1}}, {{2}}) or named
	// ({{customer_name}}, {{contact.first_name}}). Templates created in Zavu are
	// submitted to Meta as positional; templates imported from a WhatsApp Business
	// Account keep their original format (named or positional). Used when no
	// channel-specific body is set.
	Body string `json:"body" api:"required"`
	// WhatsApp template category.
	//
	// Any of "UTILITY", "MARKETING", "AUTHENTICATION".
	Category WhatsappCategory `json:"category" api:"required"`
	// Language code.
	Language string `json:"language" api:"required"`
	// Template name. For WhatsApp, must match the approved template name in Meta.
	Name string `json:"name" api:"required"`
	// Add 'Do not share this code' disclaimer. Only for AUTHENTICATION templates.
	AddSecurityRecommendation bool `json:"addSecurityRecommendation"`
	// Template buttons.
	Buttons []TemplateButton `json:"buttons"`
	// Code expiration time in minutes. Only for AUTHENTICATION templates.
	CodeExpirationMinutes int64     `json:"codeExpirationMinutes"`
	CreatedAt             time.Time `json:"createdAt" format:"date-time"`
	// Footer text for the template.
	Footer string `json:"footer"`
	// Header content (text or media URL).
	HeaderContent string `json:"headerContent"`
	// Type of header (text, image, video, document).
	HeaderType string `json:"headerType"`
	// Channel-specific body for Instagram messages. Falls back to `body` if not set.
	InstagramBody string `json:"instagramBody"`
	// Channel-specific body for SMS messages. Falls back to `body` if not set.
	SMSBody string `json:"smsBody"`
	// Any of "draft", "pending", "approved", "rejected".
	Status TemplateStatus `json:"status"`
	// Channel-specific body for Telegram messages. Falls back to `body` if not set.
	TelegramBody string    `json:"telegramBody"`
	UpdatedAt    time.Time `json:"updatedAt" format:"date-time"`
	// List of variable names for documentation.
	Variables []string `json:"variables"`
	// WhatsApp-specific template information.
	Whatsapp TemplateWhatsapp `json:"whatsapp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Body                      respjson.Field
		Category                  respjson.Field
		Language                  respjson.Field
		Name                      respjson.Field
		AddSecurityRecommendation respjson.Field
		Buttons                   respjson.Field
		CodeExpirationMinutes     respjson.Field
		CreatedAt                 respjson.Field
		Footer                    respjson.Field
		HeaderContent             respjson.Field
		HeaderType                respjson.Field
		InstagramBody             respjson.Field
		SMSBody                   respjson.Field
		Status                    respjson.Field
		TelegramBody              respjson.Field
		UpdatedAt                 respjson.Field
		Variables                 respjson.Field
		Whatsapp                  respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Template) RawJSON() string { return r.JSON.raw }
func (r *Template) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateButton struct {
	// Sample value used to substitute `{{1}}` in the URL when submitting the template
	// to Meta for review. Only present for dynamic URL buttons.
	Example string `json:"example"`
	// OTP button type. Required when type is 'otp'.
	//
	// Any of "COPY_CODE", "ONE_TAP".
	OtpType string `json:"otpType"`
	// Android package name. Required for ONE_TAP buttons.
	PackageName string `json:"packageName"`
	PhoneNumber string `json:"phoneNumber"`
	// Android app signature hash. Required for ONE_TAP buttons.
	SignatureHash string `json:"signatureHash"`
	Text          string `json:"text"`
	// Any of "quick_reply", "url", "phone", "otp", "request_contact_info".
	Type string `json:"type"`
	URL  string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Example       respjson.Field
		OtpType       respjson.Field
		PackageName   respjson.Field
		PhoneNumber   respjson.Field
		SignatureHash respjson.Field
		Text          respjson.Field
		Type          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateButton) RawJSON() string { return r.JSON.raw }
func (r *TemplateButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateStatus string

const (
	TemplateStatusDraft    TemplateStatus = "draft"
	TemplateStatusPending  TemplateStatus = "pending"
	TemplateStatusApproved TemplateStatus = "approved"
	TemplateStatusRejected TemplateStatus = "rejected"
)

// WhatsApp-specific template information.
type TemplateWhatsapp struct {
	// WhatsApp Business Account namespace.
	Namespace string `json:"namespace"`
	// WhatsApp approval status.
	Status string `json:"status"`
	// WhatsApp template name.
	TemplateName string `json:"templateName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Namespace    respjson.Field
		Status       respjson.Field
		TemplateName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *TemplateWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp template category.
type WhatsappCategory string

const (
	WhatsappCategoryUtility        WhatsappCategory = "UTILITY"
	WhatsappCategoryMarketing      WhatsappCategory = "MARKETING"
	WhatsappCategoryAuthentication WhatsappCategory = "AUTHENTICATION"
)

type TemplateSyncResponse struct {
	// WhatsApp Business Accounts reconciled in this call.
	AccountsSynced int64 `json:"accountsSynced" api:"required"`
	// Problems hit while syncing. Non-empty with a 200 means part of the sync did not
	// complete — the rest still did.
	Errors []string `json:"errors" api:"required"`
	// Templates that existed on Meta and were created in Zavu by this call.
	Imported int64 `json:"imported" api:"required"`
	// Existing Zavu templates that were matched to a Meta template by name and bound
	// to its Meta ID.
	Linked int64 `json:"linked" api:"required"`
	// Meta templates left alone: already linked to a Zavu template, or
	// rejected/disabled on Meta.
	Skipped int64 `json:"skipped" api:"required"`
	// Templates whose approval status changed to match Meta.
	Updated int64 `json:"updated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountsSynced respjson.Field
		Errors         respjson.Field
		Imported       respjson.Field
		Linked         respjson.Field
		Skipped        respjson.Field
		Updated        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplateSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateNewParams struct {
	// Default template body. Used when no channel-specific body is set.
	Body     string `json:"body" api:"required"`
	Language string `json:"language" api:"required"`
	Name     string `json:"name" api:"required"`
	// Add 'Do not share this code' disclaimer. Only for AUTHENTICATION templates.
	AddSecurityRecommendation param.Opt[bool] `json:"addSecurityRecommendation,omitzero"`
	// Code expiration time in minutes. Only for AUTHENTICATION templates.
	CodeExpirationMinutes param.Opt[int64] `json:"codeExpirationMinutes,omitzero"`
	// Footer text for the template.
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Header content (text string or media URL).
	HeaderContent param.Opt[string] `json:"headerContent,omitzero"`
	// Channel-specific body for Instagram. Falls back to `body` if not set.
	InstagramBody param.Opt[string] `json:"instagramBody,omitzero"`
	// Channel-specific body for SMS. Falls back to `body` if not set.
	SMSBody param.Opt[string] `json:"smsBody,omitzero"`
	// Channel-specific body for Telegram. Falls back to `body` if not set.
	TelegramBody param.Opt[string] `json:"telegramBody,omitzero"`
	// Template buttons (max 3).
	Buttons []TemplateNewParamsButton `json:"buttons,omitzero"`
	// Type of header for the template.
	//
	// Any of "text", "image", "video", "document".
	HeaderType TemplateNewParamsHeaderType `json:"headerType,omitzero"`
	Variables  []string                    `json:"variables,omitzero"`
	// WhatsApp template category.
	//
	// Any of "UTILITY", "MARKETING", "AUTHENTICATION".
	WhatsappCategory WhatsappCategory `json:"whatsappCategory,omitzero"`
	paramObj
}

func (r TemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type TemplateNewParamsButton struct {
	// `request_contact_info` renders a fixed **Share Contact Info** button that asks
	// the recipient to share their phone number — useful when a contact adopted a
	// WhatsApp username and you only know their BSUID. It takes no other fields.
	//
	// Any of "quick_reply", "url", "phone", "otp", "request_contact_info".
	Type string `json:"type,omitzero" api:"required"`
	// Sample value Meta uses to review templates with a dynamic URL button.
	// Substituted into `{{1}}` of the URL when the template is submitted to Meta. Only
	// meaningful when `url` contains `{{1}}`; ignored for static URLs.
	Example param.Opt[string] `json:"example,omitzero"`
	// Android package name. Required for ONE_TAP buttons.
	PackageName param.Opt[string] `json:"packageName,omitzero"`
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Android app signature hash. Required for ONE_TAP buttons.
	SignatureHash param.Opt[string] `json:"signatureHash,omitzero"`
	// Button label. Required for every type except `request_contact_info`, whose label
	// is fixed by WhatsApp.
	Text param.Opt[string] `json:"text,omitzero"`
	// Button destination. Use `{{1}}` exactly once for a dynamic URL (e.g.
	// `https://example.com/orders/{{1}}`); WhatsApp only accepts the strict `{{1}}`
	// form. Static URLs must not contain any `{{...}}` placeholder.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Required when type is 'otp'. COPY_CODE shows copy button, ONE_TAP enables
	// Android autofill.
	//
	// Any of "COPY_CODE", "ONE_TAP".
	OtpType string `json:"otpType,omitzero"`
	paramObj
}

func (r TemplateNewParamsButton) MarshalJSON() (data []byte, err error) {
	type shadow TemplateNewParamsButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateNewParamsButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TemplateNewParamsButton](
		"type", "quick_reply", "url", "phone", "otp", "request_contact_info",
	)
	apijson.RegisterFieldValidator[TemplateNewParamsButton](
		"otpType", "COPY_CODE", "ONE_TAP",
	)
}

// Type of header for the template.
type TemplateNewParamsHeaderType string

const (
	TemplateNewParamsHeaderTypeText     TemplateNewParamsHeaderType = "text"
	TemplateNewParamsHeaderTypeImage    TemplateNewParamsHeaderType = "image"
	TemplateNewParamsHeaderTypeVideo    TemplateNewParamsHeaderType = "video"
	TemplateNewParamsHeaderTypeDocument TemplateNewParamsHeaderType = "document"
)

type TemplateListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TemplateListParams]'s query parameters as `url.Values`.
func (r TemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TemplateSubmitParams struct {
	// The sender ID with the WhatsApp Business Account to submit the template to.
	SenderID string `json:"senderId" api:"required"`
	// Template category. If not provided, uses the category set on the template.
	//
	// Any of "UTILITY", "MARKETING", "AUTHENTICATION".
	Category WhatsappCategory `json:"category,omitzero"`
	paramObj
}

func (r TemplateSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplateSyncParams struct {
	// Sync only the WhatsApp Business Account attached to this sender. If omitted,
	// every WhatsApp sender in the project is synced.
	SenderID param.Opt[string] `json:"senderId,omitzero"`
	paramObj
}

func (r TemplateSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplateSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplateSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
