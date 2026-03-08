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

// AddressService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r AddressService) {
	r = AddressService{}
	r.Options = opts
	return
}

// Create a regulatory address for phone number purchases. Some countries require a
// verified address before phone numbers can be activated.
func (r *AddressService) New(ctx context.Context, body AddressNewParams, opts ...option.RequestOption) (res *AddressNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific regulatory address.
func (r *AddressService) Get(ctx context.Context, addressID string, opts ...option.RequestOption) (res *AddressGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if addressID == "" {
		err = errors.New("missing required addressId parameter")
		return
	}
	path := fmt.Sprintf("v1/addresses/%s", url.PathEscape(addressID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List regulatory addresses for this project.
func (r *AddressService) List(ctx context.Context, query AddressListParams, opts ...option.RequestOption) (res *pagination.Cursor[Address], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/addresses"
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

// List regulatory addresses for this project.
func (r *AddressService) ListAutoPaging(ctx context.Context, query AddressListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Address] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a regulatory address. Cannot delete addresses that are in use.
func (r *AddressService) Delete(ctx context.Context, addressID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if addressID == "" {
		err = errors.New("missing required addressId parameter")
		return
	}
	path := fmt.Sprintf("v1/addresses/%s", url.PathEscape(addressID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A regulatory address for phone number requirements.
type Address struct {
	ID          string    `json:"id" api:"required"`
	CountryCode string    `json:"countryCode" api:"required"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	Locality    string    `json:"locality" api:"required"`
	PostalCode  string    `json:"postalCode" api:"required"`
	// Any of "pending", "verified", "rejected".
	Status             AddressStatus `json:"status" api:"required"`
	StreetAddress      string        `json:"streetAddress" api:"required"`
	AdministrativeArea string        `json:"administrativeArea" api:"nullable"`
	BusinessName       string        `json:"businessName" api:"nullable"`
	ExtendedAddress    string        `json:"extendedAddress" api:"nullable"`
	FirstName          string        `json:"firstName" api:"nullable"`
	LastName           string        `json:"lastName" api:"nullable"`
	UpdatedAt          time.Time     `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CountryCode        respjson.Field
		CreatedAt          respjson.Field
		Locality           respjson.Field
		PostalCode         respjson.Field
		Status             respjson.Field
		StreetAddress      respjson.Field
		AdministrativeArea respjson.Field
		BusinessName       respjson.Field
		ExtendedAddress    respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressStatus string

const (
	AddressStatusPending  AddressStatus = "pending"
	AddressStatusVerified AddressStatus = "verified"
	AddressStatusRejected AddressStatus = "rejected"
)

type AddressNewResponse struct {
	// A regulatory address for phone number requirements.
	Address Address `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressGetResponse struct {
	// A regulatory address for phone number requirements.
	Address Address `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressNewParams struct {
	CountryCode        string            `json:"countryCode" api:"required"`
	Locality           string            `json:"locality" api:"required"`
	PostalCode         string            `json:"postalCode" api:"required"`
	StreetAddress      string            `json:"streetAddress" api:"required"`
	AdministrativeArea param.Opt[string] `json:"administrativeArea,omitzero"`
	BusinessName       param.Opt[string] `json:"businessName,omitzero"`
	ExtendedAddress    param.Opt[string] `json:"extendedAddress,omitzero"`
	FirstName          param.Opt[string] `json:"firstName,omitzero"`
	LastName           param.Opt[string] `json:"lastName,omitzero"`
	paramObj
}

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParams]'s query parameters as `url.Values`.
func (r AddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
