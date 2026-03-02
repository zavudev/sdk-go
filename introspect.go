// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/zavudev-go/internal/apijson"
	"github.com/stainless-sdks/zavudev-go/internal/requestconfig"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/stainless-sdks/zavudev-go/packages/param"
	"github.com/stainless-sdks/zavudev-go/packages/respjson"
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

// Validate a phone number and check if a WhatsApp conversation window is open.
func (r *IntrospectService) ValidatePhone(ctx context.Context, body IntrospectValidatePhoneParams, opts ...option.RequestOption) (res *IntrospectValidatePhoneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/introspect/phone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
