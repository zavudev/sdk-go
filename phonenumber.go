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

// PhoneNumberService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberService] method instead.
type PhoneNumberService struct {
	Options []option.RequestOption
}

// NewPhoneNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhoneNumberService(opts ...option.RequestOption) (r PhoneNumberService) {
	r = PhoneNumberService{}
	r.Options = opts
	return
}

// Get details of a specific phone number.
func (r *PhoneNumberService) Get(ctx context.Context, phoneNumberID string, opts ...option.RequestOption) (res *PhoneNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phoneNumberId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/phone-numbers/%s", url.PathEscape(phoneNumberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a phone number's name or sender assignment.
func (r *PhoneNumberService) Update(ctx context.Context, phoneNumberID string, body PhoneNumberUpdateParams, opts ...option.RequestOption) (res *PhoneNumberUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phoneNumberId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/phone-numbers/%s", url.PathEscape(phoneNumberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all phone numbers owned by this project.
func (r *PhoneNumberService) List(ctx context.Context, query PhoneNumberListParams, opts ...option.RequestOption) (res *pagination.Cursor[OwnedPhoneNumber], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/phone-numbers"
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

// List all phone numbers owned by this project.
func (r *PhoneNumberService) ListAutoPaging(ctx context.Context, query PhoneNumberListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[OwnedPhoneNumber] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Purchase an available phone number. Requires a paid plan: the Free plan cannot
// purchase phone numbers and receives `402` with code `paid_plan_required`.
//
// **The included number.** A paid plan includes one number at no charge, once per
// account: it must be a US or Canadian number (a +1 number) costing $20 a month or
// less. `isFreeEligible` in `GET /v1/phone-numbers/available` marks the numbers
// that qualify. Claiming it spends the benefit for good, across every team the
// account owner owns, so releasing that number does not make another one free.
//
// **Numbers with regulatory requirements.** Which numbers need regulatory
// information is decided per number, not by a fixed country list. The purchase
// looks the requirements up for the exact number before charging anything:
//
//  1. `GET /v1/phone-numbers/requirements?phoneNumber=...`. If `items` is empty,
//     buy normally.
//  2. Create what it asks for: addresses with `POST /v1/addresses`, documents with
//     `POST /v1/documents`.
//  3. Purchase with `type` and `regulatoryRequirements`. The number is bought and
//     billed at once with `regulatoryStatus: pending_review`.
//  4. Poll `GET /v1/phone-numbers/{phoneNumberId}` until `regulatoryStatus` is
//     `approved`. Assign it to a sender before or after approval; it starts
//     carrying messages once approved.
//
// **Reuse.** Information you submitted is kept for your project, per country and
// `type`, and a later purchase there may omit `regulatoryRequirements`. Reuse only
// happens when what is kept still covers every requirement of the new number and
// every address and document in it belongs to the project. Otherwise, or when
// nothing is kept, the purchase returns `400 regulatory_compliance_required` with
// the missing requirements in `details`.
//
// Invalid values (a missing, unknown or repeated requirement id, an address or
// document from another project, or one rejected in review) return
// `400 invalid_request`. If an address or document cannot be registered for
// review, the purchase returns `400 invalid_request` naming the requirement. If
// the requirements cannot be looked up, the purchase returns
// `502 requirements_unavailable`, except for US and Canadian numbers, which are
// sold as numbers without requirements. None of these errors charge anything.
func (r *PhoneNumberService) Purchase(ctx context.Context, body PhoneNumberPurchaseParams, opts ...option.RequestOption) (res *PhoneNumberPurchaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/phone-numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Release a phone number. The phone number must not be assigned to a sender.
func (r *PhoneNumberService) Release(ctx context.Context, phoneNumberID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumberID == "" {
		err = errors.New("missing required phoneNumberId parameter")
		return err
	}
	path := fmt.Sprintf("v1/phone-numbers/%s", url.PathEscape(phoneNumberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get the regulatory information needed to buy a phone number, for one specific
// number or for a country and number type. Prefer `phoneNumber`: the response is
// then exactly the list the purchase of that number validates against. Pass each
// `requirementTypes[].id` back as `requirementType` in `regulatoryRequirements` on
// `POST /v1/phone-numbers`.
//
// For `phoneNumber`, the requirements of that exact number are returned. When they
// cannot be resolved for the number itself, the list for its country and `type` is
// returned instead, and the purchase uses the same list. An empty `items` array
// means the number needs no regulatory information. If the requirements cannot be
// retrieved at all, the response is `502 requirements_unavailable`, never an empty
// list.
//
// URL-encode the `+` of `phoneNumber` as `%2B`. An unencoded `+` is also accepted.
func (r *PhoneNumberService) Requirements(ctx context.Context, query PhoneNumberRequirementsParams, opts ...option.RequestOption) (res *PhoneNumberRequirementsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/phone-numbers/requirements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search for available phone numbers to purchase by country and type.
func (r *PhoneNumberService) SearchAvailable(ctx context.Context, query PhoneNumberSearchAvailableParams, opts ...option.RequestOption) (res *PhoneNumberSearchAvailableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/phone-numbers/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AvailablePhoneNumber struct {
	Capabilities PhoneNumberCapabilities `json:"capabilities" api:"required"`
	PhoneNumber  string                  `json:"phoneNumber" api:"required"`
	Pricing      PhoneNumberPricing      `json:"pricing" api:"required"`
	FriendlyName string                  `json:"friendlyName"`
	Locality     string                  `json:"locality"`
	Region       string                  `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities respjson.Field
		PhoneNumber  respjson.Field
		Pricing      respjson.Field
		FriendlyName respjson.Field
		Locality     respjson.Field
		Region       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OwnedPhoneNumber struct {
	ID           string                  `json:"id" api:"required"`
	Capabilities []string                `json:"capabilities" api:"required"`
	CreatedAt    time.Time               `json:"createdAt" api:"required" format:"date-time"`
	PhoneNumber  string                  `json:"phoneNumber" api:"required"`
	Pricing      OwnedPhoneNumberPricing `json:"pricing" api:"required"`
	// Regulatory review state. Numbers that need no review are `approved` immediately.
	// A number bought with regulatory information is owned and billed from purchase
	// and starts `pending_review`; it cannot send messages or place calls until this
	// is `approved`. The state is re-checked every 6 hours: poll
	// `GET /v1/phone-numbers/{phoneNumberId}` to follow it.
	//
	// Assign it to a sender with `PATCH /v1/phone-numbers/{phoneNumberId}`
	// (`senderId`) before or after approval. A number assigned while under review is
	// recorded and connected to that sender when it is approved; the connection is
	// retried until it succeeds. A sender created over the API is set up for SMS as
	// part of the assignment. `rejected` means review refused the information: the
	// number cannot be assigned to a sender. A number that stays `pending_review` may
	// be waiting on information the API cannot supply; contact support.
	//
	// Any of "approved", "pending_review", "rejected".
	RegulatoryStatus OwnedPhoneNumberRegulatoryStatus `json:"regulatoryStatus" api:"required"`
	// Billing state of an owned number, separate from `regulatoryStatus`. `pending` is
	// legacy and is not written to numbers today. The SDKs carry `active`, `suspended`
	// and `pending` only; `releasing` and `released` are returned by the REST API
	// until their next release.
	//
	// Any of "active", "suspended", "pending", "releasing", "released".
	Status PhoneNumberStatus `json:"status" api:"required"`
	// Optional custom name for the phone number.
	Name            string    `json:"name"`
	NextRenewalDate time.Time `json:"nextRenewalDate" format:"date-time"`
	// Sender ID if the phone number is assigned to a sender.
	SenderID  string    `json:"senderId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Capabilities     respjson.Field
		CreatedAt        respjson.Field
		PhoneNumber      respjson.Field
		Pricing          respjson.Field
		RegulatoryStatus respjson.Field
		Status           respjson.Field
		Name             respjson.Field
		NextRenewalDate  respjson.Field
		SenderID         respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnedPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *OwnedPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Regulatory review state. Numbers that need no review are `approved` immediately.
// A number bought with regulatory information is owned and billed from purchase
// and starts `pending_review`; it cannot send messages or place calls until this
// is `approved`. The state is re-checked every 6 hours: poll
// `GET /v1/phone-numbers/{phoneNumberId}` to follow it.
//
// Assign it to a sender with `PATCH /v1/phone-numbers/{phoneNumberId}`
// (`senderId`) before or after approval. A number assigned while under review is
// recorded and connected to that sender when it is approved; the connection is
// retried until it succeeds. A sender created over the API is set up for SMS as
// part of the assignment. `rejected` means review refused the information: the
// number cannot be assigned to a sender. A number that stays `pending_review` may
// be waiting on information the API cannot supply; contact support.
type OwnedPhoneNumberRegulatoryStatus string

const (
	OwnedPhoneNumberRegulatoryStatusApproved      OwnedPhoneNumberRegulatoryStatus = "approved"
	OwnedPhoneNumberRegulatoryStatusPendingReview OwnedPhoneNumberRegulatoryStatus = "pending_review"
	OwnedPhoneNumberRegulatoryStatusRejected      OwnedPhoneNumberRegulatoryStatus = "rejected"
)

type OwnedPhoneNumberPricing struct {
	// Whether this is a free number.
	IsFreeNumber bool `json:"isFreeNumber"`
	// Monthly cost in cents.
	MonthlyCost float64 `json:"monthlyCost"`
	// Monthly price in USD.
	MonthlyPrice float64 `json:"monthlyPrice"`
	// One-time purchase cost in cents.
	UpfrontCost float64 `json:"upfrontCost"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFreeNumber respjson.Field
		MonthlyCost  respjson.Field
		MonthlyPrice respjson.Field
		UpfrontCost  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnedPhoneNumberPricing) RawJSON() string { return r.JSON.raw }
func (r *OwnedPhoneNumberPricing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCapabilities struct {
	Mms   bool `json:"mms"`
	SMS   bool `json:"sms"`
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mms         respjson.Field
		SMS         respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCapabilities) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberPricing struct {
	// Whether this number qualifies as the plan-included number: a US or Canadian
	// number (a +1 number) costing $20 a month or less. The benefit is one per
	// account: it is never offered again once claimed, not even after the number is
	// released.
	IsFreeEligible bool `json:"isFreeEligible"`
	// Monthly price in USD.
	MonthlyPrice float64 `json:"monthlyPrice"`
	// One-time purchase price in USD.
	UpfrontPrice float64 `json:"upfrontPrice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFreeEligible respjson.Field
		MonthlyPrice   respjson.Field
		UpfrontPrice   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberPricing) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberPricing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing state of an owned number, separate from `regulatoryStatus`. `pending` is
// legacy and is not written to numbers today. The SDKs carry `active`, `suspended`
// and `pending` only; `releasing` and `released` are returned by the REST API
// until their next release.
type PhoneNumberStatus string

const (
	PhoneNumberStatusActive    PhoneNumberStatus = "active"
	PhoneNumberStatusSuspended PhoneNumberStatus = "suspended"
	PhoneNumberStatusPending   PhoneNumberStatus = "pending"
	PhoneNumberStatusReleasing PhoneNumberStatus = "releasing"
	PhoneNumberStatusReleased  PhoneNumberStatus = "released"
)

// Type of phone number. `mobile` is stocked in countries where no geographic
// (`local`) or non-geographic (`national`) inventory exists, and in several
// markets it is the only type that can receive SMS.
type PhoneNumberType string

const (
	PhoneNumberTypeLocal    PhoneNumberType = "local"
	PhoneNumberTypeNational PhoneNumberType = "national"
	PhoneNumberTypeTollFree PhoneNumberType = "tollFree"
	PhoneNumberTypeMobile   PhoneNumberType = "mobile"
)

// The requirements for ordering a number: for a country and number type, or for
// one specific number when requested with `phoneNumber` (then `id` is that phone
// number and `countryCode` is taken from it).
type Requirement struct {
	ID               string            `json:"id" api:"required"`
	Action           string            `json:"action" api:"required"`
	CountryCode      string            `json:"countryCode" api:"required"`
	PhoneNumberType  string            `json:"phoneNumberType" api:"required"`
	RequirementTypes []RequirementType `json:"requirementTypes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Action           respjson.Field
		CountryCode      respjson.Field
		PhoneNumberType  respjson.Field
		RequirementTypes respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Requirement) RawJSON() string { return r.JSON.raw }
func (r *Requirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Acceptance criteria for a requirement.
type RequirementAcceptanceCriteria struct {
	AllowedValues []string `json:"allowedValues" api:"nullable"`
	MaxLength     int64    `json:"maxLength" api:"nullable"`
	MinLength     int64    `json:"minLength" api:"nullable"`
	RegexPattern  string   `json:"regexPattern" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedValues respjson.Field
		MaxLength     respjson.Field
		MinLength     respjson.Field
		RegexPattern  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementAcceptanceCriteria) RawJSON() string { return r.JSON.raw }
func (r *RequirementAcceptanceCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of requirement field.
type RequirementFieldType string

const (
	RequirementFieldTypeTextual  RequirementFieldType = "textual"
	RequirementFieldTypeAddress  RequirementFieldType = "address"
	RequirementFieldTypeDocument RequirementFieldType = "document"
	RequirementFieldTypeAction   RequirementFieldType = "action"
)

// A specific requirement type within a requirement group.
type RequirementType struct {
	// Send this as `requirementType` in `regulatoryRequirements` when purchasing.
	ID          string `json:"id" api:"required"`
	Description string `json:"description" api:"required"`
	Name        string `json:"name" api:"required"`
	// Type of requirement field.
	//
	// Any of "textual", "address", "document", "action".
	Type RequirementFieldType `json:"type" api:"required"`
	// Acceptance criteria for a requirement.
	AcceptanceCriteria RequirementAcceptanceCriteria `json:"acceptanceCriteria"`
	Example            string                        `json:"example" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Description        respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		AcceptanceCriteria respjson.Field
		Example            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementType) RawJSON() string { return r.JSON.raw }
func (r *RequirementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberGetResponse struct {
	PhoneNumber OwnedPhoneNumber `json:"phoneNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberUpdateResponse struct {
	PhoneNumber OwnedPhoneNumber `json:"phoneNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberPurchaseResponse struct {
	PhoneNumber OwnedPhoneNumber `json:"phoneNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberPurchaseResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberPurchaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberRequirementsResponse struct {
	Items []Requirement `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberRequirementsResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberRequirementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberSearchAvailableResponse struct {
	Items []AvailablePhoneNumber `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberSearchAvailableResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberSearchAvailableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberUpdateParams struct {
	// Custom name for the phone number. Set to null to clear.
	Name param.Opt[string] `json:"name,omitzero"`
	// Sender ID to assign the phone number to. Set to null to unassign. A number under
	// regulatory review is recorded now and connected to the sender when approved; a
	// rejected number is refused.
	SenderID param.Opt[string] `json:"senderId,omitzero"`
	paramObj
}

func (r PhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListParams struct {
	// Pagination cursor.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Filter by phone number status.
	//
	// Any of "active", "suspended", "pending", "releasing", "released".
	Status PhoneNumberStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParams]'s query parameters as `url.Values`.
func (r PhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PhoneNumberPurchaseParams struct {
	// Phone number in E.164 format.
	PhoneNumber string `json:"phoneNumber" api:"required"`
	// Optional custom name for the phone number.
	Name param.Opt[string] `json:"name,omitzero"`
	// Regulatory information, for numbers whose requirements list is not empty. Get
	// the list with `GET /v1/phone-numbers/requirements?phoneNumber=...` and send one
	// entry per requirement id, except `action` requirements, which take no value.
	// Every required id must be present, once, and no unknown id may be sent;
	// otherwise the purchase is refused with `400 invalid_request` before anything is
	// charged.
	//
	// The information is kept for your project under the number's country and `type`.
	// A later purchase there may omit this field if what is kept still covers that
	// number's requirements. Omit it for numbers without requirements.
	RegulatoryRequirements []PhoneNumberPurchaseParamsRegulatoryRequirement `json:"regulatoryRequirements,omitzero"`
	// Type of phone number. `mobile` is stocked in countries where no geographic
	// (`local`) or non-geographic (`national`) inventory exists, and in several
	// markets it is the only type that can receive SMS.
	//
	// Any of "local", "national", "tollFree", "mobile".
	Type PhoneNumberType `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberPurchaseParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberPurchaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberPurchaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FieldValue, RequirementType are required.
type PhoneNumberPurchaseParamsRegulatoryRequirement struct {
	// Depends on the requirement's `type`: the text itself for `textual`; for
	// `address`, the `id` of an address created in this project with
	// `POST /v1/addresses`; for `document`, the `id` of a document created with
	// `POST /v1/documents`. An address or document from another project, or one
	// rejected in review, is refused.
	FieldValue string `json:"fieldValue" api:"required"`
	// A `requirementTypes[].id` from `GET /v1/phone-numbers/requirements`. Each id may
	// appear only once.
	RequirementType string `json:"requirementType" api:"required"`
	paramObj
}

func (r PhoneNumberPurchaseParamsRegulatoryRequirement) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberPurchaseParamsRegulatoryRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberPurchaseParamsRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberRequirementsParams struct {
	// Two-letter ISO country code. Required unless `phoneNumber` is given.
	CountryCode param.Opt[string] `query:"countryCode,omitzero" json:"-"`
	// E.164 number from `GET /v1/phone-numbers/available`, with `+` encoded as `%2B`.
	// Returns the requirements the purchase of that number checks. Takes precedence
	// over `countryCode`.
	PhoneNumber param.Opt[string] `query:"phoneNumber,omitzero" json:"-"`
	// Type of phone number (local, national, mobile, tollFree). Defaults to `local`.
	// With `phoneNumber`, used only when the number's own requirements cannot be
	// resolved and the country list is returned.
	//
	// Any of "local", "national", "tollFree", "mobile".
	Type PhoneNumberType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberRequirementsParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberRequirementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PhoneNumberSearchAvailableParams struct {
	// Two-letter ISO country code.
	CountryCode string `query:"countryCode" api:"required" json:"-"`
	// Comma-separated capabilities the number must have: `sms`, `voice`, `mms`.
	// Numbers missing any of them are dropped.
	Capabilities param.Opt[string] `query:"capabilities,omitzero" json:"-"`
	// Search for numbers containing this string.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Type of phone number to search for.
	//
	// Any of "local", "national", "tollFree", "mobile".
	Type PhoneNumberType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSearchAvailableParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberSearchAvailableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
