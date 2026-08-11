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
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// SenderWhatsappSyncService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderWhatsappSyncService] method instead.
type SenderWhatsappSyncService struct {
	Options []option.RequestOption
}

// NewSenderWhatsappSyncService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSenderWhatsappSyncService(opts ...option.RequestOption) (r SenderWhatsappSyncService) {
	r = SenderWhatsappSyncService{}
	r.Options = opts
	return
}

// Get the current sync status for a sender's WhatsApp coexistence account. Only
// available for senders connected in coexistence mode (WhatsApp Business App +
// Cloud API).
func (r *SenderWhatsappSyncService) Get(ctx context.Context, senderID string, opts ...option.RequestOption) (res *SenderWhatsappSyncGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/whatsapp-sync", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Initiate contact names sync from the WhatsApp Business App. This imports contact
// names stored in the app to Zavu. Only available for coexistence accounts with
// active status.
func (r *SenderWhatsappSyncService) StartContactsSync(ctx context.Context, senderID string, opts ...option.RequestOption) (res *SenderWhatsappSyncStartContactsSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/whatsapp-sync/contacts", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Initiate message history sync from the WhatsApp Business App. This sends a
// request to the account owner to approve sharing their conversation history. Only
// available for coexistence accounts with active status.
func (r *SenderWhatsappSyncService) StartHistorySync(ctx context.Context, senderID string, opts ...option.RequestOption) (res *SenderWhatsappSyncStartHistorySyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/whatsapp-sync/history", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Contacts sync status details.
type WhatsAppSyncContacts struct {
	// Whether contacts sync can be initiated.
	CanSync bool `json:"canSync" api:"required"`
	// Status of WhatsApp contacts sync.
	//
	// Any of "not_requested", "pending", "syncing", "completed".
	Status WhatsAppSyncContactsStatus `json:"status" api:"required"`
	// When the sync was last requested.
	RequestedAt time.Time `json:"requestedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanSync     respjson.Field
		Status      respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsAppSyncContacts) RawJSON() string { return r.JSON.raw }
func (r *WhatsAppSyncContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of WhatsApp contacts sync.
type WhatsAppSyncContactsStatus string

const (
	WhatsAppSyncContactsStatusNotRequested WhatsAppSyncContactsStatus = "not_requested"
	WhatsAppSyncContactsStatusPending      WhatsAppSyncContactsStatus = "pending"
	WhatsAppSyncContactsStatusSyncing      WhatsAppSyncContactsStatus = "syncing"
	WhatsAppSyncContactsStatusCompleted    WhatsAppSyncContactsStatus = "completed"
)

// History sync status details.
type WhatsAppSyncHistory struct {
	// Whether history sync can be initiated.
	CanSync bool `json:"canSync" api:"required"`
	// Status of WhatsApp message history sync.
	//
	// Any of "not_requested", "pending", "syncing", "completed", "rejected".
	Status WhatsAppSyncHistoryStatus `json:"status" api:"required"`
	// When the sync was completed.
	CompletedAt time.Time `json:"completedAt" api:"nullable" format:"date-time"`
	// When the sync was last requested.
	RequestedAt time.Time `json:"requestedAt" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanSync     respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsAppSyncHistory) RawJSON() string { return r.JSON.raw }
func (r *WhatsAppSyncHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of WhatsApp message history sync.
type WhatsAppSyncHistoryStatus string

const (
	WhatsAppSyncHistoryStatusNotRequested WhatsAppSyncHistoryStatus = "not_requested"
	WhatsAppSyncHistoryStatusPending      WhatsAppSyncHistoryStatus = "pending"
	WhatsAppSyncHistoryStatusSyncing      WhatsAppSyncHistoryStatus = "syncing"
	WhatsAppSyncHistoryStatusCompleted    WhatsAppSyncHistoryStatus = "completed"
	WhatsAppSyncHistoryStatusRejected     WhatsAppSyncHistoryStatus = "rejected"
)

// WhatsApp coexistence sync status.
type WhatsAppSyncStatus struct {
	// Contacts sync status details.
	Contacts WhatsAppSyncContacts `json:"contacts" api:"required"`
	// History sync status details.
	History WhatsAppSyncHistory `json:"history" api:"required"`
	// Whether the account is in coexistence mode.
	IsCoexistence bool `json:"isCoexistence" api:"required"`
	// WhatsApp account status.
	//
	// Any of "pending_verification", "pending_registration", "active", "disconnected",
	// "error".
	Status WhatsAppSyncStatusStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contacts      respjson.Field
		History       respjson.Field
		IsCoexistence respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsAppSyncStatus) RawJSON() string { return r.JSON.raw }
func (r *WhatsAppSyncStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WhatsApp account status.
type WhatsAppSyncStatusStatus string

const (
	WhatsAppSyncStatusStatusPendingVerification WhatsAppSyncStatusStatus = "pending_verification"
	WhatsAppSyncStatusStatusPendingRegistration WhatsAppSyncStatusStatus = "pending_registration"
	WhatsAppSyncStatusStatusActive              WhatsAppSyncStatusStatus = "active"
	WhatsAppSyncStatusStatusDisconnected        WhatsAppSyncStatusStatus = "disconnected"
	WhatsAppSyncStatusStatusError               WhatsAppSyncStatusStatus = "error"
)

type SenderWhatsappSyncGetResponse struct {
	// WhatsApp coexistence sync status.
	Sync WhatsAppSyncStatus `json:"sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sync        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWhatsappSyncGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderWhatsappSyncGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderWhatsappSyncStartContactsSyncResponse struct {
	// Success message.
	Message string `json:"message" api:"required"`
	// WhatsApp coexistence sync status.
	Sync WhatsAppSyncStatus `json:"sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Sync        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWhatsappSyncStartContactsSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderWhatsappSyncStartContactsSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderWhatsappSyncStartHistorySyncResponse struct {
	// Success message.
	Message string `json:"message" api:"required"`
	// WhatsApp coexistence sync status.
	Sync WhatsAppSyncStatus `json:"sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Sync        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderWhatsappSyncStartHistorySyncResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderWhatsappSyncStartHistorySyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
