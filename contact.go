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

// ContactService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	Options  []option.RequestOption
	Channels ContactChannelService
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r ContactService) {
	r = ContactService{}
	r.Options = opts
	r.Channels = NewContactChannelService(opts...)
	return
}

// Create a new contact with one or more communication channels.
func (r *ContactService) New(ctx context.Context, body ContactNewParams, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/contacts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get contact
func (r *ContactService) Get(ctx context.Context, contactID string, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update contact
func (r *ContactService) Update(ctx context.Context, contactID string, body ContactUpdateParams, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List contacts with their communication channels.
func (r *ContactService) List(ctx context.Context, query ContactListParams, opts ...option.RequestOption) (res *pagination.Cursor[Contact], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contacts"
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

// List contacts with their communication channels.
func (r *ContactService) ListAutoPaging(ctx context.Context, query ContactListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Contact] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete a contact and its communication channels. Implements
// right-to-erasure obligations under GDPR Art. 17, Ley 19.628 (Chile) Art. 12,
// CCPA § 1798.105, and LGPD Art. 18.VI. The contact, its channels, and any
// associated agent flow sessions and conversation threads are removed. Past
// message records and broadcast delivery logs are retained for billing/audit but
// no longer reference the deleted contact.
func (r *ContactService) Delete(ctx context.Context, contactID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v1/contacts/%s", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Dismiss the merge suggestion for a contact.
func (r *ContactService) DismissMergeSuggestion(ctx context.Context, contactID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v1/contacts/%s/merge-suggestion", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Merge a source contact into this contact. All channels from the source contact
// will be moved to the target contact, and the source contact will be marked as
// merged.
func (r *ContactService) Merge(ctx context.Context, contactID string, body ContactMergeParams, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.Options, opts)
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/%s/merge", url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get contact by phone number
func (r *ContactService) GetByPhone(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *Contact, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/contacts/phone/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Contact struct {
	ID string `json:"id" api:"required"`
	// List of available messaging channels for this contact.
	AvailableChannels []string          `json:"availableChannels" api:"required"`
	CreatedAt         time.Time         `json:"createdAt" api:"required" format:"date-time"`
	Metadata          map[string]string `json:"metadata" api:"required"`
	// Whether this contact has been verified.
	Verified bool `json:"verified" api:"required"`
	// All communication channels for this contact.
	Channels    []ContactChannel `json:"channels"`
	CountryCode string           `json:"countryCode"`
	// Preferred channel for this contact.
	//
	// Any of "sms", "whatsapp", "telegram", "email", "instagram", "messenger",
	// "voice".
	DefaultChannel ContactDefaultChannel `json:"defaultChannel"`
	// Display name for the contact.
	DisplayName string `json:"displayName"`
	// DEPRECATED: Use primaryPhone instead. Primary phone number in E.164 format.
	PhoneNumber string `json:"phoneNumber"`
	// Primary email address.
	PrimaryEmail string `json:"primaryEmail" format:"email"`
	// Primary phone number in E.164 format.
	PrimaryPhone string `json:"primaryPhone"`
	// Contact's WhatsApp profile name. Only available for WhatsApp contacts.
	ProfileName string `json:"profileName" api:"nullable"`
	// ID of a contact suggested for merging.
	SuggestedMergeWith string    `json:"suggestedMergeWith"`
	UpdatedAt          time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AvailableChannels  respjson.Field
		CreatedAt          respjson.Field
		Metadata           respjson.Field
		Verified           respjson.Field
		Channels           respjson.Field
		CountryCode        respjson.Field
		DefaultChannel     respjson.Field
		DisplayName        respjson.Field
		PhoneNumber        respjson.Field
		PrimaryEmail       respjson.Field
		PrimaryPhone       respjson.Field
		ProfileName        respjson.Field
		SuggestedMergeWith respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contact) RawJSON() string { return r.JSON.raw }
func (r *Contact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preferred channel for this contact.
type ContactDefaultChannel string

const (
	ContactDefaultChannelSMS       ContactDefaultChannel = "sms"
	ContactDefaultChannelWhatsapp  ContactDefaultChannel = "whatsapp"
	ContactDefaultChannelTelegram  ContactDefaultChannel = "telegram"
	ContactDefaultChannelEmail     ContactDefaultChannel = "email"
	ContactDefaultChannelInstagram ContactDefaultChannel = "instagram"
	ContactDefaultChannelMessenger ContactDefaultChannel = "messenger"
	ContactDefaultChannelVoice     ContactDefaultChannel = "voice"
)

// A communication channel for a contact.
type ContactChannel struct {
	ID string `json:"id" api:"required"`
	// Channel type.
	//
	// Any of "sms", "whatsapp", "email", "telegram", "instagram", "messenger",
	// "voice".
	Channel   ContactChannelChannel `json:"channel" api:"required"`
	CreatedAt time.Time             `json:"createdAt" api:"required" format:"date-time"`
	// Channel identifier (phone number or email address).
	Identifier string `json:"identifier" api:"required"`
	// Whether this is the primary channel for its type.
	IsPrimary bool `json:"isPrimary" api:"required"`
	// Whether this channel has been verified.
	Verified bool `json:"verified" api:"required"`
	// ISO country code for phone numbers.
	CountryCode string `json:"countryCode"`
	// Optional label for the channel.
	Label string `json:"label"`
	// Last time a message was received on this channel.
	LastInboundAt time.Time         `json:"lastInboundAt" format:"date-time"`
	Metadata      map[string]string `json:"metadata"`
	// Delivery metrics for this channel.
	Metrics   ContactChannelMetrics `json:"metrics"`
	UpdatedAt time.Time             `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Channel       respjson.Field
		CreatedAt     respjson.Field
		Identifier    respjson.Field
		IsPrimary     respjson.Field
		Verified      respjson.Field
		CountryCode   respjson.Field
		Label         respjson.Field
		LastInboundAt respjson.Field
		Metadata      respjson.Field
		Metrics       respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactChannel) RawJSON() string { return r.JSON.raw }
func (r *ContactChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Channel type.
type ContactChannelChannel string

const (
	ContactChannelChannelSMS       ContactChannelChannel = "sms"
	ContactChannelChannelWhatsapp  ContactChannelChannel = "whatsapp"
	ContactChannelChannelEmail     ContactChannelChannel = "email"
	ContactChannelChannelTelegram  ContactChannelChannel = "telegram"
	ContactChannelChannelInstagram ContactChannelChannel = "instagram"
	ContactChannelChannelMessenger ContactChannelChannel = "messenger"
	ContactChannelChannelVoice     ContactChannelChannel = "voice"
)

// Delivery metrics for this channel.
type ContactChannelMetrics struct {
	AvgDeliveryTimeMs float64   `json:"avgDeliveryTimeMs"`
	FailureCount      int64     `json:"failureCount"`
	LastSuccessAt     time.Time `json:"lastSuccessAt" format:"date-time"`
	SuccessCount      int64     `json:"successCount"`
	TotalAttempts     int64     `json:"totalAttempts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgDeliveryTimeMs respjson.Field
		FailureCount      respjson.Field
		LastSuccessAt     respjson.Field
		SuccessCount      respjson.Field
		TotalAttempts     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactChannelMetrics) RawJSON() string { return r.JSON.raw }
func (r *ContactChannelMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactNewParams struct {
	// Communication channels for the contact.
	Channels []ContactNewParamsChannel `json:"channels,omitzero" api:"required"`
	// Display name for the contact.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Arbitrary metadata to associate with the contact.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for creating a contact channel.
//
// The properties Channel, Identifier are required.
type ContactNewParamsChannel struct {
	// Channel type.
	//
	// Any of "sms", "whatsapp", "email", "telegram", "instagram", "messenger",
	// "voice".
	Channel string `json:"channel,omitzero" api:"required"`
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

func (r ContactNewParamsChannel) MarshalJSON() (data []byte, err error) {
	type shadow ContactNewParamsChannel
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactNewParamsChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ContactNewParamsChannel](
		"channel", "sms", "whatsapp", "email", "telegram", "instagram", "messenger", "voice",
	)
}

type ContactUpdateParams struct {
	// Preferred channel for this contact. Set to null to clear.
	//
	// Any of "sms", "whatsapp", "telegram", "email", "instagram", "messenger",
	// "voice".
	DefaultChannel ContactUpdateParamsDefaultChannel `json:"defaultChannel,omitzero"`
	Metadata       map[string]string                 `json:"metadata,omitzero"`
	paramObj
}

func (r ContactUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preferred channel for this contact. Set to null to clear.
type ContactUpdateParamsDefaultChannel string

const (
	ContactUpdateParamsDefaultChannelSMS       ContactUpdateParamsDefaultChannel = "sms"
	ContactUpdateParamsDefaultChannelWhatsapp  ContactUpdateParamsDefaultChannel = "whatsapp"
	ContactUpdateParamsDefaultChannelTelegram  ContactUpdateParamsDefaultChannel = "telegram"
	ContactUpdateParamsDefaultChannelEmail     ContactUpdateParamsDefaultChannel = "email"
	ContactUpdateParamsDefaultChannelInstagram ContactUpdateParamsDefaultChannel = "instagram"
	ContactUpdateParamsDefaultChannelMessenger ContactUpdateParamsDefaultChannel = "messenger"
	ContactUpdateParamsDefaultChannelVoice     ContactUpdateParamsDefaultChannel = "voice"
)

type ContactListParams struct {
	Cursor      param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	PhoneNumber param.Opt[string] `query:"phoneNumber,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContactListParams]'s query parameters as `url.Values`.
func (r ContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContactMergeParams struct {
	// ID of the contact to merge into the target contact. The source contact will be
	// marked as merged.
	SourceContactID string `json:"sourceContactId" api:"required"`
	paramObj
}

func (r ContactMergeParams) MarshalJSON() (data []byte, err error) {
	type shadow ContactMergeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContactMergeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
