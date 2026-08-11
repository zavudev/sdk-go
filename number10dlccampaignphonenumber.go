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
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// Number10dlcCampaignPhoneNumberService contains methods and other services that
// help with interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcCampaignPhoneNumberService] method instead.
type Number10dlcCampaignPhoneNumberService struct {
	Options []option.RequestOption
}

// NewNumber10dlcCampaignPhoneNumberService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNumber10dlcCampaignPhoneNumberService(opts ...option.RequestOption) (r Number10dlcCampaignPhoneNumberService) {
	r = Number10dlcCampaignPhoneNumberService{}
	r.Options = opts
	return
}

// List phone numbers assigned to a 10DLC campaign.
func (r *Number10dlcCampaignPhoneNumberService) List(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Number10dlcCampaignPhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s/phone-numbers", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Assign a US phone number to an approved 10DLC campaign. The campaign must be in
// approved status.
func (r *Number10dlcCampaignPhoneNumberService) Assign(ctx context.Context, campaignID string, body Number10dlcCampaignPhoneNumberAssignParams, opts ...option.RequestOption) (res *Number10dlcCampaignPhoneNumberAssignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s/phone-numbers", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a phone number assignment from a 10DLC campaign.
func (r *Number10dlcCampaignPhoneNumberService) Unassign(ctx context.Context, assignmentID string, body Number10dlcCampaignPhoneNumberUnassignParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.CampaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return err
	}
	if assignmentID == "" {
		err = errors.New("missing required assignmentId parameter")
		return err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s/phone-numbers/%s", url.PathEscape(body.CampaignID), url.PathEscape(assignmentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TenDlcPhoneNumberAssignment struct {
	ID            string    `json:"id" api:"required"`
	CampaignID    string    `json:"campaignId" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	PhoneNumberID string    `json:"phoneNumberId" api:"required"`
	// Assignment status.
	//
	// Any of "pending", "active", "failed".
	Status        TenDlcPhoneNumberAssignmentStatus `json:"status" api:"required"`
	UpdatedAt     time.Time                         `json:"updatedAt" api:"required" format:"date-time"`
	AssignedAt    time.Time                         `json:"assignedAt" api:"nullable" format:"date-time"`
	FailureReason string                            `json:"failureReason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CampaignID    respjson.Field
		CreatedAt     respjson.Field
		PhoneNumberID respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		AssignedAt    respjson.Field
		FailureReason respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenDlcPhoneNumberAssignment) RawJSON() string { return r.JSON.raw }
func (r *TenDlcPhoneNumberAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Assignment status.
type TenDlcPhoneNumberAssignmentStatus string

const (
	TenDlcPhoneNumberAssignmentStatusPending TenDlcPhoneNumberAssignmentStatus = "pending"
	TenDlcPhoneNumberAssignmentStatusActive  TenDlcPhoneNumberAssignmentStatus = "active"
	TenDlcPhoneNumberAssignmentStatusFailed  TenDlcPhoneNumberAssignmentStatus = "failed"
)

type Number10dlcCampaignPhoneNumberListResponse struct {
	Items      []TenDlcPhoneNumberAssignment `json:"items" api:"required"`
	NextCursor string                        `json:"nextCursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignPhoneNumberAssignResponse struct {
	Assignment TenDlcPhoneNumberAssignment `json:"assignment" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignPhoneNumberAssignResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignPhoneNumberAssignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignPhoneNumberAssignParams struct {
	// ID of the phone number to assign.
	PhoneNumberID string `json:"phoneNumberId" api:"required"`
	paramObj
}

func (r Number10dlcCampaignPhoneNumberAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcCampaignPhoneNumberAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcCampaignPhoneNumberAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignPhoneNumberUnassignParams struct {
	CampaignID string `path:"campaignId" api:"required" json:"-"`
	paramObj
}
