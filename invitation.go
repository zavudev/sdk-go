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

// InvitationService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvitationService] method instead.
type InvitationService struct {
	Options []option.RequestOption
}

// NewInvitationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvitationService(opts ...option.RequestOption) (r InvitationService) {
	r = InvitationService{}
	r.Options = opts
	return
}

// Create a partner invitation link for a client to connect a Meta channel. The
// client opens the returned `url` and authorizes with Meta; the resulting sender
// is created in your project when they finish, and the invitation transitions to
// `completed`.
//
// `connectionType` picks the channel:
//
//   - `whatsapp_waba` (default): Meta's embedded signup links an official WhatsApp
//     Business Account.
//   - `messenger`: the client picks a Facebook Page they administer; its Messenger
//     inbox (including Marketplace chats) is routed to Zavu.
//
// One invitation connects one channel — create one per channel to onboard a client
// on several. `phoneNumberId` and `allowedPhoneCountries` apply to `whatsapp_waba`
// only.
func (r *InvitationService) New(ctx context.Context, body InvitationNewParams, opts ...option.RequestOption) (res *InvitationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/invitations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get invitation
func (r *InvitationService) Get(ctx context.Context, invitationID string, opts ...option.RequestOption) (res *InvitationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invitationID == "" {
		err = errors.New("missing required invitationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/invitations/%s", url.PathEscape(invitationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List partner invitations for this project.
func (r *InvitationService) List(ctx context.Context, query InvitationListParams, opts ...option.RequestOption) (res *pagination.Cursor[Invitation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/invitations"
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

// List partner invitations for this project.
func (r *InvitationService) ListAutoPaging(ctx context.Context, query InvitationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Invitation] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Cancel an active invitation. The client will no longer be able to use the
// invitation link.
func (r *InvitationService) Cancel(ctx context.Context, invitationID string, opts ...option.RequestOption) (res *InvitationCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invitationID == "" {
		err = errors.New("missing required invitationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/invitations/%s/cancel", url.PathEscape(invitationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Invitation struct {
	ID string `json:"id" api:"required"`
	// Unique invitation token.
	Token     string    `json:"token" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Current status of the partner invitation.
	//
	// `failed` means the client started the connection and it did not finish (they
	// cancelled Meta's dialog, denied a permission, or abandoned the tab). A failed
	// invitation is still usable: the same link can be retried, and it moves back to
	// `in_progress` when the client tries again.
	//
	// Any of "pending", "in_progress", "completed", "expired", "cancelled", "failed".
	Status    InvitationStatus `json:"status" api:"required"`
	UpdatedAt time.Time        `json:"updatedAt" api:"required" format:"date-time"`
	// Full URL to share with the client.
	URL         string    `json:"url" api:"required"`
	ClientEmail string    `json:"clientEmail" api:"nullable"`
	ClientName  string    `json:"clientName" api:"nullable"`
	ClientPhone string    `json:"clientPhone" api:"nullable"`
	CompletedAt time.Time `json:"completedAt" api:"nullable" format:"date-time"`
	// The account the client linked, populated once the invitation is `completed`.
	// Null before that. Use it to show the partner what was connected without fetching
	// the sender.
	ConnectedAccount InvitationConnectedAccount `json:"connectedAccount" api:"nullable"`
	// Which Meta channel the client connects: `whatsapp_waba` (official WhatsApp Cloud
	// API via embedded signup) or `messenger` (a Facebook Page's Messenger inbox,
	// including Marketplace chats).
	//
	// Any of "whatsapp_waba", "messenger".
	ConnectionType InvitationConnectionType `json:"connectionType"`
	FailedAt       time.Time                `json:"failedAt" api:"nullable" format:"date-time"`
	// Stable code for why the last attempt failed, present when `status` is `failed`.
	// Values include `fb_cancelled` (client closed Meta's dialog), `fb_not_authorized`
	// (permission denied), `signup_abandoned` (started but never finished),
	// `meta_no_pages` (the client administers no Facebook Page), and `internal_error`.
	// Treat unknown codes as a generic failure.
	FailureReason string `json:"failureReason" api:"nullable"`
	// ID of a pre-assigned Zavu phone number for WhatsApp registration. Always null
	// for `messenger` invitations.
	PhoneNumberID string `json:"phoneNumberId" api:"nullable"`
	// ID of the sender created when invitation is completed.
	SenderID  string    `json:"senderId" api:"nullable"`
	StartedAt time.Time `json:"startedAt" api:"nullable" format:"date-time"`
	ViewedAt  time.Time `json:"viewedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Token            respjson.Field
		CreatedAt        respjson.Field
		ExpiresAt        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		URL              respjson.Field
		ClientEmail      respjson.Field
		ClientName       respjson.Field
		ClientPhone      respjson.Field
		CompletedAt      respjson.Field
		ConnectedAccount respjson.Field
		ConnectionType   respjson.Field
		FailedAt         respjson.Field
		FailureReason    respjson.Field
		PhoneNumberID    respjson.Field
		SenderID         respjson.Field
		StartedAt        respjson.Field
		ViewedAt         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invitation) RawJSON() string { return r.JSON.raw }
func (r *Invitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the partner invitation.
//
// `failed` means the client started the connection and it did not finish (they
// cancelled Meta's dialog, denied a permission, or abandoned the tab). A failed
// invitation is still usable: the same link can be retried, and it moves back to
// `in_progress` when the client tries again.
type InvitationStatus string

const (
	InvitationStatusPending    InvitationStatus = "pending"
	InvitationStatusInProgress InvitationStatus = "in_progress"
	InvitationStatusCompleted  InvitationStatus = "completed"
	InvitationStatusExpired    InvitationStatus = "expired"
	InvitationStatusCancelled  InvitationStatus = "cancelled"
	InvitationStatusFailed     InvitationStatus = "failed"
)

// The account the client linked, populated once the invitation is `completed`.
// Null before that. Use it to show the partner what was connected without fetching
// the sender.
type InvitationConnectedAccount struct {
	// Provider-side identifier: the WhatsApp phone number ID, or the Facebook Page ID.
	ID string `json:"id" api:"required"`
	// Any of "whatsapp", "messenger".
	Channel string `json:"channel" api:"required"`
	// Display name of the connected account: the WhatsApp verified name, or the
	// Facebook Page name.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channel     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationConnectedAccount) RawJSON() string { return r.JSON.raw }
func (r *InvitationConnectedAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which Meta channel the client connects: `whatsapp_waba` (official WhatsApp Cloud
// API via embedded signup) or `messenger` (a Facebook Page's Messenger inbox,
// including Marketplace chats).
type InvitationConnectionType string

const (
	InvitationConnectionTypeWhatsappWaba InvitationConnectionType = "whatsapp_waba"
	InvitationConnectionTypeMessenger    InvitationConnectionType = "messenger"
)

type InvitationNewResponse struct {
	Invitation Invitation `json:"invitation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invitation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InvitationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvitationGetResponse struct {
	Invitation Invitation `json:"invitation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invitation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InvitationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvitationCancelResponse struct {
	Invitation Invitation `json:"invitation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invitation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvitationCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *InvitationCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvitationNewParams struct {
	// Email of the client being invited.
	ClientEmail param.Opt[string] `json:"clientEmail,omitzero" format:"email"`
	// Name of the client being invited.
	ClientName param.Opt[string] `json:"clientName,omitzero"`
	// Phone number of the client in E.164 format.
	ClientPhone param.Opt[string] `json:"clientPhone,omitzero"`
	// Number of days until the invitation expires.
	ExpiresInDays param.Opt[int64] `json:"expiresInDays,omitzero"`
	// ID of a Zavu phone number to pre-assign for WhatsApp registration. If provided,
	// the client will use this number instead of their own. Only valid when
	// `connectionType` is `whatsapp_waba` — sending it with `messenger` returns 400,
	// since a Facebook Page has no phone number.
	PhoneNumberID param.Opt[string] `json:"phoneNumberId,omitzero"`
	// ISO country codes for allowed phone numbers. Only valid when `connectionType` is
	// `whatsapp_waba` — sending it with `messenger` returns 400.
	AllowedPhoneCountries []string `json:"allowedPhoneCountries,omitzero"`
	// Which Meta channel the client connects, and how.
	//
	//   - `whatsapp_waba` (default): Meta's embedded signup links an official WhatsApp
	//     Business Account. Accepts `phoneNumberId` and `allowedPhoneCountries`.
	//   - `messenger`: the client authorizes with Facebook and picks a Facebook Page
	//     they administer. The Page's Messenger inbox — including Marketplace chats — is
	//     routed to Zavu. They must be an admin of at least one Page. A Page can only be
	//     connected to one Zavu project at a time: if the client picks a Page that
	//     another project already connected, the newer connection wins and the older one
	//     is disconnected.
	//
	// One invitation connects one channel. To onboard a client on several channels,
	// create one invitation per channel; each completes into its own sender.
	//
	// Any of "whatsapp_waba", "messenger".
	ConnectionType InvitationNewParamsConnectionType `json:"connectionType,omitzero"`
	paramObj
}

func (r InvitationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InvitationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvitationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which Meta channel the client connects, and how.
//
//   - `whatsapp_waba` (default): Meta's embedded signup links an official WhatsApp
//     Business Account. Accepts `phoneNumberId` and `allowedPhoneCountries`.
//   - `messenger`: the client authorizes with Facebook and picks a Facebook Page
//     they administer. The Page's Messenger inbox — including Marketplace chats — is
//     routed to Zavu. They must be an admin of at least one Page. A Page can only be
//     connected to one Zavu project at a time: if the client picks a Page that
//     another project already connected, the newer connection wins and the older one
//     is disconnected.
//
// One invitation connects one channel. To onboard a client on several channels,
// create one invitation per channel; each completes into its own sender.
type InvitationNewParamsConnectionType string

const (
	InvitationNewParamsConnectionTypeWhatsappWaba InvitationNewParamsConnectionType = "whatsapp_waba"
	InvitationNewParamsConnectionTypeMessenger    InvitationNewParamsConnectionType = "messenger"
)

type InvitationListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Current status of the partner invitation.
	//
	// `failed` means the client started the connection and it did not finish (they
	// cancelled Meta's dialog, denied a permission, or abandoned the tab). A failed
	// invitation is still usable: the same link can be retried, and it moves back to
	// `in_progress` when the client tries again.
	//
	// Any of "pending", "in_progress", "completed", "expired", "cancelled", "failed".
	Status InvitationListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvitationListParams]'s query parameters as `url.Values`.
func (r InvitationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Current status of the partner invitation.
//
// `failed` means the client started the connection and it did not finish (they
// cancelled Meta's dialog, denied a permission, or abandoned the tab). A failed
// invitation is still usable: the same link can be retried, and it moves back to
// `in_progress` when the client tries again.
type InvitationListParamsStatus string

const (
	InvitationListParamsStatusPending    InvitationListParamsStatus = "pending"
	InvitationListParamsStatusInProgress InvitationListParamsStatus = "in_progress"
	InvitationListParamsStatusCompleted  InvitationListParamsStatus = "completed"
	InvitationListParamsStatusExpired    InvitationListParamsStatus = "expired"
	InvitationListParamsStatusCancelled  InvitationListParamsStatus = "cancelled"
	InvitationListParamsStatusFailed     InvitationListParamsStatus = "failed"
)
