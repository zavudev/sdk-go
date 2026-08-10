// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/zavudev/sdk-go/internal/apijson"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// ContactChannelService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactChannelService] method instead.
type ContactChannelService struct {
	Options []option.RequestOption
}

// NewContactChannelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContactChannelService(opts ...option.RequestOption) (r ContactChannelService) {
	r = ContactChannelService{}
	r.Options = opts
	return
}

// Update a contact's channel properties.
func (r *ContactChannelService) Update(ctx context.Context, channelID string, params ContactChannelUpdateParams, opts ...option.RequestOption) (res *ContactChannelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ContactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s/channels/%s", url.PathEscape(params.ContactID), url.PathEscape(channelID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Add a new communication channel to an existing contact.
func (r *ContactChannelService) Add(ctx context.Context, contactID string, body ContactChannelAddParams, opts ...option.RequestOption) (res *ContactChannelAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s/channels", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a communication channel from a contact. Cannot remove the last channel.
func (r *ContactChannelService) Remove(ctx context.Context, channelID string, body ContactChannelRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ContactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return err
	}
	path := fmt.Sprintf("v1/contacts/%s/channels/%s", url.PathEscape(body.ContactID), url.PathEscape(channelID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Set a channel as the primary channel for its type.
func (r *ContactChannelService) SetPrimary(ctx context.Context, channelID string, body ContactChannelSetPrimaryParams, opts ...option.RequestOption) (res *ContactChannelSetPrimaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ContactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	if channelID == "" {
		err = errors.New("missing required channelId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s/channels/%s/primary", url.PathEscape(body.ContactID), url.PathEscape(channelID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ContactChannelUpdateResponse struct {
	// A communication channel for a contact.
	Channel ContactChannel `json:"channel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactChannelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactChannelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactChannelAddResponse struct {
	// A communication channel for a contact.
	Channel ContactChannel `json:"channel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactChannelAddResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactChannelAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactChannelSetPrimaryResponse struct {
	// A communication channel for a contact.
	Channel ContactChannel `json:"channel" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactChannelSetPrimaryResponse) RawJSON() string { return r.JSON.raw }
func (r *ContactChannelSetPrimaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactChannelUpdateParams struct {
	ContactID string `path:"contactId" api:"required" json:"-"`
	// Optional label for the channel. Set to null to clear.
	Label param.Opt[string] `json:"label,omitzero"`
	// Whether the channel is verified.
	Verified param.Opt[bool]   `json:"verified,omitzero"`
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ContactChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactChannelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactChannelAddParams struct {
	// Channel type.
	//
	// Any of "sms", "whatsapp", "email", "telegram", "instagram", "messenger",
	// "voice".
	Channel ContactChannelAddParamsChannel `json:"channel,omitzero" api:"required"`
	// Channel identifier (phone number in E.164 format or email address).
	Identifier string `json:"identifier" api:"required"`
	// ISO country code for phone numbers.
	CountryCode param.Opt[string] `json:"countryCode,omitzero"`
	// Whether this should be the primary channel for its type.
	IsPrimary param.Opt[bool] `json:"isPrimary,omitzero"`
	// Optional label for the channel.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r ContactChannelAddParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactChannelAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactChannelAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Channel type.
type ContactChannelAddParamsChannel string

const (
	ContactChannelAddParamsChannelSMS       ContactChannelAddParamsChannel = "sms"
	ContactChannelAddParamsChannelWhatsapp  ContactChannelAddParamsChannel = "whatsapp"
	ContactChannelAddParamsChannelEmail     ContactChannelAddParamsChannel = "email"
	ContactChannelAddParamsChannelTelegram  ContactChannelAddParamsChannel = "telegram"
	ContactChannelAddParamsChannelInstagram ContactChannelAddParamsChannel = "instagram"
	ContactChannelAddParamsChannelMessenger ContactChannelAddParamsChannel = "messenger"
	ContactChannelAddParamsChannelVoice     ContactChannelAddParamsChannel = "voice"
)

type ContactChannelRemoveParams struct {
	ContactID string `path:"contactId" api:"required" json:"-"`
	paramObj
}

type ContactChannelSetPrimaryParams struct {
	ContactID string `path:"contactId" api:"required" json:"-"`
	paramObj
}
