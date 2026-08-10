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

// Number10dlcCampaignService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcCampaignService] method instead.
type Number10dlcCampaignService struct {
	Options      []option.RequestOption
	PhoneNumbers Number10dlcCampaignPhoneNumberService
}

// NewNumber10dlcCampaignService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumber10dlcCampaignService(opts ...option.RequestOption) (r Number10dlcCampaignService) {
	r = Number10dlcCampaignService{}
	r.Options = opts
	r.PhoneNumbers = NewNumber10dlcCampaignPhoneNumberService(opts...)
	return
}

// Create a 10DLC campaign under an existing brand. The campaign starts in draft
// status. Submit it for carrier review using the submit endpoint.
func (r *Number10dlcCampaignService) New(ctx context.Context, body Number10dlcCampaignNewParams, opts ...option.RequestOption) (res *Number10dlcCampaignNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/10dlc/campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get 10DLC campaign
func (r *Number10dlcCampaignService) Get(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Number10dlcCampaignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a 10DLC campaign in draft status. Cannot update after submission.
func (r *Number10dlcCampaignService) Update(ctx context.Context, campaignID string, body Number10dlcCampaignUpdateParams, opts ...option.RequestOption) (res *Number10dlcCampaignUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List 10DLC campaign registrations for this project.
func (r *Number10dlcCampaignService) List(ctx context.Context, query Number10dlcCampaignListParams, opts ...option.RequestOption) (res *pagination.Cursor[TenDlcCampaign], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/10dlc/campaigns"
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

// List 10DLC campaign registrations for this project.
func (r *Number10dlcCampaignService) ListAutoPaging(ctx context.Context, query Number10dlcCampaignListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TenDlcCampaign] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete 10DLC campaign
func (r *Number10dlcCampaignService) Delete(ctx context.Context, campaignID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Submit a draft campaign for carrier review. The campaign must be in draft status
// and its brand must be verified. TCR's one-time registration fee is charged from
// your balance at submission ($15 for standard use cases, $2 for LOW_VOLUME),
// passed through at cost and refunded if the carrier rejects it. Once approved,
// the campaign's monthly TCR fee ($10 standard, $2 LOW_VOLUME) is charged from
// your balance while the campaign is active — see registrationCostCents and
// monthlyFeeCents on the campaign object.
func (r *Number10dlcCampaignService) Submit(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Number10dlcCampaignSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s/submit", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Sync the campaign status with the registration provider. Use this to check for
// approval updates after submission.
func (r *Number10dlcCampaignService) SyncStatus(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Number10dlcCampaignSyncStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/10dlc/campaigns/%s/sync", url.PathEscape(campaignID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TenDlcCampaign struct {
	ID                 string `json:"id" api:"required"`
	AffiliateMarketing bool   `json:"affiliateMarketing" api:"required"`
	AgeGated           bool   `json:"ageGated" api:"required"`
	// ID of the brand this campaign belongs to.
	BrandID   string    `json:"brandId" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Description of the messaging campaign.
	Description   string `json:"description" api:"required"`
	DirectLending bool   `json:"directLending" api:"required"`
	EmbeddedLink  bool   `json:"embeddedLink" api:"required"`
	EmbeddedPhone bool   `json:"embeddedPhone" api:"required"`
	Name          string `json:"name" api:"required"`
	NumberPooling bool   `json:"numberPooling" api:"required"`
	// Sample messages representative of campaign content.
	SampleMessages []string `json:"sampleMessages" api:"required"`
	// Status of a 10DLC campaign registration.
	//
	// Any of "draft", "pending", "approved", "rejected".
	Status           TenDlcCampaignStatus `json:"status" api:"required"`
	SubscriberHelp   bool                 `json:"subscriberHelp" api:"required"`
	SubscriberOptIn  bool                 `json:"subscriberOptIn" api:"required"`
	SubscriberOptOut bool                 `json:"subscriberOptOut" api:"required"`
	UpdatedAt        time.Time            `json:"updatedAt" api:"required" format:"date-time"`
	// Campaign use case type.
	UseCase    string    `json:"useCase" api:"required"`
	ApprovedAt time.Time `json:"approvedAt" api:"nullable" format:"date-time"`
	// Daily message limit based on brand trust score.
	DailyLimit    int64  `json:"dailyLimit" api:"nullable"`
	FailureReason string `json:"failureReason" api:"nullable"`
	HelpMessage   string `json:"helpMessage" api:"nullable"`
	MessageFlow   string `json:"messageFlow" api:"nullable"`
	// Recurring monthly fee in cents.
	MonthlyFeeCents int64    `json:"monthlyFeeCents" api:"nullable"`
	OptInKeywords   []string `json:"optInKeywords" api:"nullable"`
	OptOutKeywords  []string `json:"optOutKeywords" api:"nullable"`
	// One-time registration cost in cents.
	RegistrationCostCents int64     `json:"registrationCostCents" api:"nullable"`
	SubmittedAt           time.Time `json:"submittedAt" api:"nullable" format:"date-time"`
	SubUseCases           []string  `json:"subUseCases" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AffiliateMarketing    respjson.Field
		AgeGated              respjson.Field
		BrandID               respjson.Field
		CreatedAt             respjson.Field
		Description           respjson.Field
		DirectLending         respjson.Field
		EmbeddedLink          respjson.Field
		EmbeddedPhone         respjson.Field
		Name                  respjson.Field
		NumberPooling         respjson.Field
		SampleMessages        respjson.Field
		Status                respjson.Field
		SubscriberHelp        respjson.Field
		SubscriberOptIn       respjson.Field
		SubscriberOptOut      respjson.Field
		UpdatedAt             respjson.Field
		UseCase               respjson.Field
		ApprovedAt            respjson.Field
		DailyLimit            respjson.Field
		FailureReason         respjson.Field
		HelpMessage           respjson.Field
		MessageFlow           respjson.Field
		MonthlyFeeCents       respjson.Field
		OptInKeywords         respjson.Field
		OptOutKeywords        respjson.Field
		RegistrationCostCents respjson.Field
		SubmittedAt           respjson.Field
		SubUseCases           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenDlcCampaign) RawJSON() string { return r.JSON.raw }
func (r *TenDlcCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a 10DLC campaign registration.
type TenDlcCampaignStatus string

const (
	TenDlcCampaignStatusDraft    TenDlcCampaignStatus = "draft"
	TenDlcCampaignStatusPending  TenDlcCampaignStatus = "pending"
	TenDlcCampaignStatusApproved TenDlcCampaignStatus = "approved"
	TenDlcCampaignStatusRejected TenDlcCampaignStatus = "rejected"
)

type Number10dlcCampaignNewResponse struct {
	Campaign TenDlcCampaign `json:"campaign" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignNewResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignGetResponse struct {
	Campaign TenDlcCampaign `json:"campaign" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignUpdateResponse struct {
	Campaign TenDlcCampaign `json:"campaign" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignSubmitResponse struct {
	Campaign TenDlcCampaign `json:"campaign" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignSyncStatusResponse struct {
	Campaign TenDlcCampaign `json:"campaign" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcCampaignSyncStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcCampaignSyncStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignNewParams struct {
	AffiliateMarketing bool `json:"affiliateMarketing" api:"required"`
	AgeGated           bool `json:"ageGated" api:"required"`
	// ID of the brand to create this campaign under.
	BrandID          string   `json:"brandId" api:"required"`
	Description      string   `json:"description" api:"required"`
	DirectLending    bool     `json:"directLending" api:"required"`
	EmbeddedLink     bool     `json:"embeddedLink" api:"required"`
	EmbeddedPhone    bool     `json:"embeddedPhone" api:"required"`
	Name             string   `json:"name" api:"required"`
	NumberPooling    bool     `json:"numberPooling" api:"required"`
	SampleMessages   []string `json:"sampleMessages,omitzero" api:"required"`
	SubscriberHelp   bool     `json:"subscriberHelp" api:"required"`
	SubscriberOptIn  bool     `json:"subscriberOptIn" api:"required"`
	SubscriberOptOut bool     `json:"subscriberOptOut" api:"required"`
	// Campaign use case (e.g., ACCOUNT_NOTIFICATION, MARKETING, 2FA).
	UseCase        string            `json:"useCase" api:"required"`
	HelpMessage    param.Opt[string] `json:"helpMessage,omitzero"`
	MessageFlow    param.Opt[string] `json:"messageFlow,omitzero"`
	OptInKeywords  []string          `json:"optInKeywords,omitzero"`
	OptOutKeywords []string          `json:"optOutKeywords,omitzero"`
	SubUseCases    []string          `json:"subUseCases,omitzero"`
	paramObj
}

func (r Number10dlcCampaignNewParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcCampaignNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcCampaignNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignUpdateParams struct {
	Description    param.Opt[string] `json:"description,omitzero"`
	HelpMessage    param.Opt[string] `json:"helpMessage,omitzero"`
	MessageFlow    param.Opt[string] `json:"messageFlow,omitzero"`
	Name           param.Opt[string] `json:"name,omitzero"`
	OptInKeywords  []string          `json:"optInKeywords,omitzero"`
	OptOutKeywords []string          `json:"optOutKeywords,omitzero"`
	SampleMessages []string          `json:"sampleMessages,omitzero"`
	paramObj
}

func (r Number10dlcCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcCampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcCampaignListParams struct {
	// Filter campaigns by brand ID.
	BrandID param.Opt[string] `query:"brandId,omitzero" json:"-"`
	Cursor  param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcCampaignListParams]'s query parameters as
// `url.Values`.
func (r Number10dlcCampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
