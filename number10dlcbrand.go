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

// Number10dlcBrandService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcBrandService] method instead.
type Number10dlcBrandService struct {
	Options []option.RequestOption
}

// NewNumber10dlcBrandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumber10dlcBrandService(opts ...option.RequestOption) (r Number10dlcBrandService) {
	r = Number10dlcBrandService{}
	r.Options = opts
	return
}

// Create a 10DLC brand registration. The brand starts in draft status. Submit it
// for review using the submit endpoint.
func (r *Number10dlcBrandService) New(ctx context.Context, body Number10dlcBrandNewParams, opts ...option.RequestOption) (res *Number10dlcBrandNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/10dlc/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get 10DLC brand
func (r *Number10dlcBrandService) Get(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Number10dlcBrandGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/brands/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a 10DLC brand in draft status. Cannot update after submission.
func (r *Number10dlcBrandService) Update(ctx context.Context, brandID string, body Number10dlcBrandUpdateParams, opts ...option.RequestOption) (res *Number10dlcBrandUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/brands/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List 10DLC brand registrations for this project.
func (r *Number10dlcBrandService) List(ctx context.Context, query Number10dlcBrandListParams, opts ...option.RequestOption) (res *pagination.Cursor[TenDlcBrand], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/10dlc/brands"
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

// List 10DLC brand registrations for this project.
func (r *Number10dlcBrandService) ListAutoPaging(ctx context.Context, query Number10dlcBrandListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TenDlcBrand] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete 10DLC brand
func (r *Number10dlcBrandService) Delete(ctx context.Context, brandID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return err
	}
	path := fmt.Sprintf("v1/10dlc/brands/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// List available use cases for 10DLC campaign registration.
func (r *Number10dlcBrandService) ListUseCases(ctx context.Context, opts ...option.RequestOption) (res *Number10dlcBrandListUseCasesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/10dlc/brands/use-cases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit a draft brand to The Campaign Registry (TCR) for vetting. The brand must
// be in draft status, and the team must have an approved Business Verification
// (KYB) — carriers register a brand against a vetted legal entity, so submitting
// without one returns `403` with code `kyb_required`.
//
// TCR's one-time $4 brand registration fee is charged from your balance at
// submission (passed through at cost) and refunded if the carrier rejects the
// registration. The fee is per BRAND: a team registering a second legal entity
// pays it again. A brand already paid for through the compliance flow is not
// charged twice. Campaign registration is billed separately when a campaign is
// submitted.
func (r *Number10dlcBrandService) Submit(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Number10dlcBrandSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/brands/%s/submit", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Sync the brand status with the registration provider. Use this to check for
// approval updates after submission.
func (r *Number10dlcBrandService) SyncStatus(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Number10dlcBrandSyncStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/brands/%s/sync", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TenDlcBrand struct {
	ID   string `json:"id" api:"required"`
	City string `json:"city" api:"required"`
	// Two-letter ISO country code.
	Country   string    `json:"country" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Display name of the brand.
	DisplayName string `json:"displayName" api:"required"`
	Email       string `json:"email" api:"required" format:"email"`
	// Business entity type for 10DLC brand registration.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT",
	// "SOLE_PROPRIETOR".
	EntityType TenDlcBrandEntityType `json:"entityType" api:"required"`
	// Contact phone number in E.164 format.
	Phone      string `json:"phone" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	State      string `json:"state" api:"required"`
	// Status of a 10DLC brand registration.
	//
	// Any of "draft", "pending", "verified", "rejected".
	Status    TenDlcBrandStatus `json:"status" api:"required"`
	Street    string            `json:"street" api:"required"`
	UpdatedAt time.Time         `json:"updatedAt" api:"required" format:"date-time"`
	// Industry vertical.
	Vertical          string `json:"vertical" api:"required"`
	BrandRelationship string `json:"brandRelationship" api:"nullable"`
	// Trust score assigned by TCR after vetting.
	BrandScore int64 `json:"brandScore" api:"nullable"`
	// Legal company name.
	CompanyName string `json:"companyName" api:"nullable"`
	// Employer Identification Number (EIN).
	Ein string `json:"ein" api:"nullable"`
	// Reason for rejection, if applicable.
	FailureReason string    `json:"failureReason" api:"nullable"`
	FirstName     string    `json:"firstName" api:"nullable"`
	LastName      string    `json:"lastName" api:"nullable"`
	StockExchange string    `json:"stockExchange" api:"nullable"`
	StockSymbol   string    `json:"stockSymbol" api:"nullable"`
	SubmittedAt   time.Time `json:"submittedAt" api:"nullable" format:"date-time"`
	VerifiedAt    time.Time `json:"verifiedAt" api:"nullable" format:"date-time"`
	Website       string    `json:"website" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		City              respjson.Field
		Country           respjson.Field
		CreatedAt         respjson.Field
		DisplayName       respjson.Field
		Email             respjson.Field
		EntityType        respjson.Field
		Phone             respjson.Field
		PostalCode        respjson.Field
		State             respjson.Field
		Status            respjson.Field
		Street            respjson.Field
		UpdatedAt         respjson.Field
		Vertical          respjson.Field
		BrandRelationship respjson.Field
		BrandScore        respjson.Field
		CompanyName       respjson.Field
		Ein               respjson.Field
		FailureReason     respjson.Field
		FirstName         respjson.Field
		LastName          respjson.Field
		StockExchange     respjson.Field
		StockSymbol       respjson.Field
		SubmittedAt       respjson.Field
		VerifiedAt        respjson.Field
		Website           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenDlcBrand) RawJSON() string { return r.JSON.raw }
func (r *TenDlcBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business entity type for 10DLC brand registration.
type TenDlcBrandEntityType string

const (
	TenDlcBrandEntityTypePrivateProfit  TenDlcBrandEntityType = "PRIVATE_PROFIT"
	TenDlcBrandEntityTypePublicProfit   TenDlcBrandEntityType = "PUBLIC_PROFIT"
	TenDlcBrandEntityTypeNonProfit      TenDlcBrandEntityType = "NON_PROFIT"
	TenDlcBrandEntityTypeGovernment     TenDlcBrandEntityType = "GOVERNMENT"
	TenDlcBrandEntityTypeSoleProprietor TenDlcBrandEntityType = "SOLE_PROPRIETOR"
)

// Status of a 10DLC brand registration.
type TenDlcBrandStatus string

const (
	TenDlcBrandStatusDraft    TenDlcBrandStatus = "draft"
	TenDlcBrandStatusPending  TenDlcBrandStatus = "pending"
	TenDlcBrandStatusVerified TenDlcBrandStatus = "verified"
	TenDlcBrandStatusRejected TenDlcBrandStatus = "rejected"
)

type Number10dlcBrandNewResponse struct {
	Brand TenDlcBrand `json:"brand" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandNewResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandGetResponse struct {
	Brand TenDlcBrand `json:"brand" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandUpdateResponse struct {
	Brand TenDlcBrand `json:"brand" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandListUseCasesResponse struct {
	UseCases []Number10dlcBrandListUseCasesResponseUseCase `json:"useCases" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UseCases    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandListUseCasesResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandListUseCasesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandListUseCasesResponseUseCase struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandListUseCasesResponseUseCase) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandListUseCasesResponseUseCase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandSubmitResponse struct {
	Brand TenDlcBrand `json:"brand" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandSyncStatusResponse struct {
	Brand TenDlcBrand `json:"brand" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandSyncStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandSyncStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandNewParams struct {
	City string `json:"city" api:"required"`
	// Two-letter ISO country code.
	Country string `json:"country" api:"required"`
	// Display name of the brand.
	DisplayName string `json:"displayName" api:"required"`
	Email       string `json:"email" api:"required" format:"email"`
	// Business entity type for 10DLC brand registration.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT",
	// "SOLE_PROPRIETOR".
	EntityType Number10dlcBrandNewParamsEntityType `json:"entityType,omitzero" api:"required"`
	// Contact phone in E.164 format.
	Phone      string `json:"phone" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	State      string `json:"state" api:"required"`
	Street     string `json:"street" api:"required"`
	// Industry vertical.
	Vertical string `json:"vertical" api:"required"`
	// Legal company name.
	CompanyName param.Opt[string] `json:"companyName,omitzero"`
	// Employer Identification Number (format: XX-XXXXXXX).
	Ein           param.Opt[string] `json:"ein,omitzero"`
	FirstName     param.Opt[string] `json:"firstName,omitzero"`
	LastName      param.Opt[string] `json:"lastName,omitzero"`
	StockExchange param.Opt[string] `json:"stockExchange,omitzero"`
	StockSymbol   param.Opt[string] `json:"stockSymbol,omitzero"`
	Website       param.Opt[string] `json:"website,omitzero" format:"uri"`
	paramObj
}

func (r Number10dlcBrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business entity type for 10DLC brand registration.
type Number10dlcBrandNewParamsEntityType string

const (
	Number10dlcBrandNewParamsEntityTypePrivateProfit  Number10dlcBrandNewParamsEntityType = "PRIVATE_PROFIT"
	Number10dlcBrandNewParamsEntityTypePublicProfit   Number10dlcBrandNewParamsEntityType = "PUBLIC_PROFIT"
	Number10dlcBrandNewParamsEntityTypeNonProfit      Number10dlcBrandNewParamsEntityType = "NON_PROFIT"
	Number10dlcBrandNewParamsEntityTypeGovernment     Number10dlcBrandNewParamsEntityType = "GOVERNMENT"
	Number10dlcBrandNewParamsEntityTypeSoleProprietor Number10dlcBrandNewParamsEntityType = "SOLE_PROPRIETOR"
)

type Number10dlcBrandUpdateParams struct {
	City          param.Opt[string] `json:"city,omitzero"`
	CompanyName   param.Opt[string] `json:"companyName,omitzero"`
	Country       param.Opt[string] `json:"country,omitzero"`
	DisplayName   param.Opt[string] `json:"displayName,omitzero"`
	Ein           param.Opt[string] `json:"ein,omitzero"`
	Email         param.Opt[string] `json:"email,omitzero" format:"email"`
	FirstName     param.Opt[string] `json:"firstName,omitzero"`
	LastName      param.Opt[string] `json:"lastName,omitzero"`
	Phone         param.Opt[string] `json:"phone,omitzero"`
	PostalCode    param.Opt[string] `json:"postalCode,omitzero"`
	State         param.Opt[string] `json:"state,omitzero"`
	StockExchange param.Opt[string] `json:"stockExchange,omitzero"`
	StockSymbol   param.Opt[string] `json:"stockSymbol,omitzero"`
	Street        param.Opt[string] `json:"street,omitzero"`
	Vertical      param.Opt[string] `json:"vertical,omitzero"`
	Website       param.Opt[string] `json:"website,omitzero" format:"uri"`
	// Business entity type for 10DLC brand registration.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT",
	// "SOLE_PROPRIETOR".
	EntityType Number10dlcBrandUpdateParamsEntityType `json:"entityType,omitzero"`
	paramObj
}

func (r Number10dlcBrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Business entity type for 10DLC brand registration.
type Number10dlcBrandUpdateParamsEntityType string

const (
	Number10dlcBrandUpdateParamsEntityTypePrivateProfit  Number10dlcBrandUpdateParamsEntityType = "PRIVATE_PROFIT"
	Number10dlcBrandUpdateParamsEntityTypePublicProfit   Number10dlcBrandUpdateParamsEntityType = "PUBLIC_PROFIT"
	Number10dlcBrandUpdateParamsEntityTypeNonProfit      Number10dlcBrandUpdateParamsEntityType = "NON_PROFIT"
	Number10dlcBrandUpdateParamsEntityTypeGovernment     Number10dlcBrandUpdateParamsEntityType = "GOVERNMENT"
	Number10dlcBrandUpdateParamsEntityTypeSoleProprietor Number10dlcBrandUpdateParamsEntityType = "SOLE_PROPRIETOR"
)

type Number10dlcBrandListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcBrandListParams]'s query parameters as
// `url.Values`.
func (r Number10dlcBrandListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
