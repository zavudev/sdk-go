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
	"github.com/zavudev/sdk-go/internal/apiquery"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/pagination"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// BroadcastContactService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBroadcastContactService] method instead.
type BroadcastContactService struct {
	Options []option.RequestOption
}

// NewBroadcastContactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBroadcastContactService(opts ...option.RequestOption) (r BroadcastContactService) {
	r = BroadcastContactService{}
	r.Options = opts
	return
}

// List contacts in a broadcast with optional status filter.
func (r *BroadcastContactService) List(ctx context.Context, broadcastID string, query BroadcastContactListParams, opts ...option.RequestOption) (res *pagination.Cursor[BroadcastContact], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/contacts", url.PathEscape(broadcastID))
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

// List contacts in a broadcast with optional status filter.
func (r *BroadcastContactService) ListAutoPaging(ctx context.Context, broadcastID string, query BroadcastContactListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[BroadcastContact] {
	return pagination.NewCursorAutoPager(r.List(ctx, broadcastID, query, opts...))
}

// Add contacts to a broadcast in batch. Maximum 1000 contacts per request.
func (r *BroadcastContactService) Add(ctx context.Context, broadcastID string, body BroadcastContactAddParams, opts ...option.RequestOption) (res *BroadcastContactAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/contacts", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove a contact from a broadcast in draft status.
func (r *BroadcastContactService) Remove(ctx context.Context, contactID string, body BroadcastContactRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.BroadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/contacts/%s", url.PathEscape(body.BroadcastID), url.PathEscape(contactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BroadcastContactAddResponse struct {
	// Number of contacts successfully added.
	Added int64 `json:"added" api:"required"`
	// Number of duplicate contacts skipped.
	Duplicates int64 `json:"duplicates" api:"required"`
	// Number of invalid contacts rejected.
	Invalid int64 `json:"invalid" api:"required"`
	// Details about invalid contacts.
	Errors []BroadcastContactAddResponseError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Duplicates  respjson.Field
		Invalid     respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastContactAddResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastContactAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastContactAddResponseError struct {
	Reason    string `json:"reason"`
	Recipient string `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastContactAddResponseError) RawJSON() string { return r.JSON.raw }
func (r *BroadcastContactAddResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastContactListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Status of a contact within a broadcast.
	//
	// Any of "pending", "queued", "sending", "delivered", "failed", "skipped".
	Status BroadcastContactStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastContactListParams]'s query parameters as
// `url.Values`.
func (r BroadcastContactListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastContactAddParams struct {
	// List of contacts to add (max 1000 per request).
	Contacts []BroadcastContactAddParamsContact `json:"contacts,omitzero" api:"required"`
	paramObj
}

func (r BroadcastContactAddParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastContactAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastContactAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Recipient is required.
type BroadcastContactAddParamsContact struct {
	// Phone number (E.164) or email address.
	Recipient string `json:"recipient" api:"required"`
	// Per-contact template variables to personalize the message.
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r BroadcastContactAddParamsContact) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastContactAddParamsContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastContactAddParamsContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastContactRemoveParams struct {
	BroadcastID string `path:"broadcastId" api:"required" json:"-"`
	paramObj
}
