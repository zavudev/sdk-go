// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"net/http"
	"slices"

	"github.com/zavudev/sdk-go/internal/apijson"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// IntrospectService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntrospectService] method instead.
type IntrospectService struct {
	Options []option.RequestOption
}

// NewIntrospectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntrospectService(opts ...option.RequestOption) (r IntrospectService) {
	r = IntrospectService{}
	r.Options = opts
	return
}

// Heuristic email validation to run before sending: catches invalid syntax, dead
// domains (no MX/A records), disposable inboxes, role-based addresses (info@,
// contacto@, sales@), and addresses already on your project's suppression list.
// Use it to clean a list before a broadcast and keep your bounce rate low.
//
// No mailbox-level (SMTP) probe is performed, so a `deliverable` verdict is not a
// delivery guarantee — it means no negative signal was found. Treat `risky`
// addresses with care and drop `undeliverable` ones.
//
// Accepts a single `email` or an `emails` batch (max 100 per request).
func (r *IntrospectService) ValidateEmail(ctx context.Context, body IntrospectValidateEmailParams, opts ...option.RequestOption) (res *IntrospectValidateEmailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/introspect/email"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Validate a phone number and check if a WhatsApp conversation window is open.
func (r *IntrospectService) ValidatePhone(ctx context.Context, body IntrospectValidatePhoneParams, opts ...option.RequestOption) (res *IntrospectValidatePhoneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/introspect/phone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Type of phone line.
type LineType string

const (
	LineTypeMobile   LineType = "mobile"
	LineTypeLandline LineType = "landline"
	LineTypeVoip     LineType = "voip"
	LineTypeTollFree LineType = "toll_free"
	LineTypeUnknown  LineType = "unknown"
)

type IntrospectValidateEmailResponse struct {
	// One result per submitted address, in the same order.
	Results []IntrospectValidateEmailResponseResult `json:"results" api:"required"`
	Summary IntrospectValidateEmailResponseSummary  `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntrospectValidateEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *IntrospectValidateEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntrospectValidateEmailResponseResult struct {
	// Domain part of the address. Null when the syntax is invalid.
	Domain string `json:"domain" api:"required"`
	// The address exactly as submitted.
	Email string `json:"email" api:"required"`
	// Lowercased, trimmed form of the address. Null when the syntax is invalid.
	Normalized string `json:"normalized" api:"required"`
	// Signals behind the verdict. Empty for a clean `deliverable` address.
	//
	// Any of "invalid_syntax", "domain_not_found", "domain_no_mx",
	// "disposable_domain", "role_address", "suppressed_hard_bounce",
	// "suppressed_soft_bounce", "suppressed_complaint", "suppressed_manual",
	// "suppressed_unsubscribe".
	Reasons []string `json:"reasons" api:"required"`
	// Validation verdict.
	//
	//   - `deliverable`: nothing suggests the address will bounce.
	//   - `risky`: sendable, but a signal predicts elevated bounce/complaint odds (role
	//     address, disposable domain, MX-less domain, prior soft bounce).
	//   - `undeliverable`: will bounce or is blocked (invalid syntax, dead domain, or
	//     the address is on your suppression list after a hard bounce/complaint).
	//
	// Any of "deliverable", "risky", "undeliverable".
	Verdict string `json:"verdict" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Email       respjson.Field
		Normalized  respjson.Field
		Reasons     respjson.Field
		Verdict     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntrospectValidateEmailResponseResult) RawJSON() string { return r.JSON.raw }
func (r *IntrospectValidateEmailResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntrospectValidateEmailResponseSummary struct {
	Deliverable   int64 `json:"deliverable" api:"required"`
	Risky         int64 `json:"risky" api:"required"`
	Total         int64 `json:"total" api:"required"`
	Undeliverable int64 `json:"undeliverable" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deliverable   respjson.Field
		Risky         respjson.Field
		Total         respjson.Field
		Undeliverable respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntrospectValidateEmailResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *IntrospectValidateEmailResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntrospectValidatePhoneResponse struct {
	CountryCode string `json:"countryCode" api:"required"`
	PhoneNumber string `json:"phoneNumber" api:"required"`
	ValidNumber bool   `json:"validNumber" api:"required"`
	// List of available messaging channels for this phone number.
	AvailableChannels []string `json:"availableChannels"`
	// Carrier information for the phone number.
	Carrier IntrospectValidatePhoneResponseCarrier `json:"carrier"`
	// Type of phone line.
	//
	// Any of "mobile", "landline", "voip", "toll_free", "unknown".
	LineType LineType `json:"lineType"`
	// Phone number in national format.
	NationalFormat string `json:"nationalFormat"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode       respjson.Field
		PhoneNumber       respjson.Field
		ValidNumber       respjson.Field
		AvailableChannels respjson.Field
		Carrier           respjson.Field
		LineType          respjson.Field
		NationalFormat    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntrospectValidatePhoneResponse) RawJSON() string { return r.JSON.raw }
func (r *IntrospectValidatePhoneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carrier information for the phone number.
type IntrospectValidatePhoneResponseCarrier struct {
	// Carrier name.
	Name string `json:"name" api:"nullable"`
	// Type of phone line.
	//
	// Any of "mobile", "landline", "voip", "toll_free", "unknown".
	Type LineType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntrospectValidatePhoneResponseCarrier) RawJSON() string { return r.JSON.raw }
func (r *IntrospectValidatePhoneResponseCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntrospectValidateEmailParams struct {
	// Single email address to validate.
	Email param.Opt[string] `json:"email,omitzero"`
	// Batch of email addresses to validate (max 100).
	Emails []string `json:"emails,omitzero"`
	paramObj
}

func (r IntrospectValidateEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow IntrospectValidateEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntrospectValidateEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntrospectValidatePhoneParams struct {
	PhoneNumber string `json:"phoneNumber" api:"required"`
	paramObj
}

func (r IntrospectValidatePhoneParams) MarshalJSON() (data []byte, err error) {
	type shadow IntrospectValidatePhoneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntrospectValidatePhoneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
