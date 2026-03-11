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

// Purchase an available phone number. The first US phone number is free for each
// team.
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

// Get regulatory requirements for purchasing phone numbers in a specific country.
// Some countries require additional documentation (addresses, identity documents)
// before phone numbers can be activated.
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
	// Any of "active", "suspended", "pending".
	Status PhoneNumberStatus `json:"status" api:"required"`
	// Optional custom name for the phone number.
	Name            string    `json:"name"`
	NextRenewalDate time.Time `json:"nextRenewalDate" format:"date-time"`
	// Sender ID if the phone number is assigned to a sender.
	SenderID  string    `json:"senderId"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Capabilities    respjson.Field
		CreatedAt       respjson.Field
		PhoneNumber     respjson.Field
		Pricing         respjson.Field
		Status          respjson.Field
		Name            respjson.Field
		NextRenewalDate respjson.Field
		SenderID        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnedPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *OwnedPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Whether this number qualifies for the free first US number offer.
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

type PhoneNumberStatus string

const (
	PhoneNumberStatusActive    PhoneNumberStatus = "active"
	PhoneNumberStatusSuspended PhoneNumberStatus = "suspended"
	PhoneNumberStatusPending   PhoneNumberStatus = "pending"
)

type PhoneNumberType string

const (
	PhoneNumberTypeLocal    PhoneNumberType = "local"
	PhoneNumberTypeMobile   PhoneNumberType = "mobile"
	PhoneNumberTypeTollFree PhoneNumberType = "tollFree"
)

// A group of requirements for a specific country/phone type combination.
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
	// Sender ID to assign the phone number to. Set to null to unassign.
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
	// Any of "active", "suspended", "pending".
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
	paramObj
}

func (r PhoneNumberPurchaseParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberPurchaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberPurchaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberRequirementsParams struct {
	// Two-letter ISO country code.
	CountryCode string `query:"countryCode" api:"required" json:"-"`
	// Type of phone number (local, mobile, tollFree).
	//
	// Any of "local", "mobile", "tollFree".
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
	// Search for numbers containing this string.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Type of phone number to search for.
	//
	// Any of "local", "mobile", "tollFree".
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
