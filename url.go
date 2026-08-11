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

// URLService contains methods and other services that help with interacting with
// the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLService] method instead.
type URLService struct {
	Options []option.RequestOption
}

// NewURLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewURLService(opts ...option.RequestOption) (r URLService) {
	r = URLService{}
	r.Options = opts
	return
}

// List URLs that have been verified for this project.
func (r *URLService) ListVerified(ctx context.Context, query URLListVerifiedParams, opts ...option.RequestOption) (res *pagination.Cursor[VerifiedURL], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/urls"
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

// List URLs that have been verified for this project.
func (r *URLService) ListVerifiedAutoPaging(ctx context.Context, query URLListVerifiedParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VerifiedURL] {
	return pagination.NewCursorAutoPager(r.ListVerified(ctx, query, opts...))
}

// Get details of a specific verified URL.
func (r *URLService) GetDetails(ctx context.Context, urlID string, opts ...option.RequestOption) (res *URLGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if urlID == "" {
		err = errors.New("missing required urlId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/urls/%s", url.PathEscape(urlID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit a URL for verification. URLs are automatically checked against Google Web
// Risk API. Safe URLs are auto-approved, malicious URLs are blocked. URL
// shorteners (bit.ly, t.co, etc.) are always blocked.
//
// **Important:** All SMS and Email messages containing URLs require those URLs to
// be verified before the message can be sent. This endpoint allows
// pre-verification of URLs.
func (r *URLService) SubmitForVerification(ctx context.Context, body URLSubmitForVerificationParams, opts ...option.RequestOption) (res *URLSubmitForVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/urls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VerifiedURL struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Domain extracted from the URL.
	Domain string `json:"domain" api:"required"`
	// Status of a verified URL.
	//
	// Any of "pending", "approved", "rejected", "escalated", "malicious".
	Status VerifiedURLStatus `json:"status" api:"required"`
	// The verified URL.
	URL string `json:"url" api:"required"`
	// How the URL was approved or rejected.
	//
	// Any of "manual", "auto_web_risk".
	ApprovalType VerifiedURLApprovalType `json:"approvalType"`
	UpdatedAt    time.Time               `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Domain       respjson.Field
		Status       respjson.Field
		URL          respjson.Field
		ApprovalType respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedURL) RawJSON() string { return r.JSON.raw }
func (r *VerifiedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a verified URL.
type VerifiedURLStatus string

const (
	VerifiedURLStatusPending   VerifiedURLStatus = "pending"
	VerifiedURLStatusApproved  VerifiedURLStatus = "approved"
	VerifiedURLStatusRejected  VerifiedURLStatus = "rejected"
	VerifiedURLStatusEscalated VerifiedURLStatus = "escalated"
	VerifiedURLStatusMalicious VerifiedURLStatus = "malicious"
)

// How the URL was approved or rejected.
type VerifiedURLApprovalType string

const (
	VerifiedURLApprovalTypeManual      VerifiedURLApprovalType = "manual"
	VerifiedURLApprovalTypeAutoWebRisk VerifiedURLApprovalType = "auto_web_risk"
)

type URLGetDetailsResponse struct {
	URL VerifiedURL `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLGetDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *URLGetDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLSubmitForVerificationResponse struct {
	URL VerifiedURL `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLSubmitForVerificationResponse) RawJSON() string { return r.JSON.raw }
func (r *URLSubmitForVerificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLListVerifiedParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Filter by verification status.
	//
	// Any of "pending", "approved", "rejected", "escalated", "malicious".
	Status URLListVerifiedParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [URLListVerifiedParams]'s query parameters as `url.Values`.
func (r URLListVerifiedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by verification status.
type URLListVerifiedParamsStatus string

const (
	URLListVerifiedParamsStatusPending   URLListVerifiedParamsStatus = "pending"
	URLListVerifiedParamsStatusApproved  URLListVerifiedParamsStatus = "approved"
	URLListVerifiedParamsStatusRejected  URLListVerifiedParamsStatus = "rejected"
	URLListVerifiedParamsStatusEscalated URLListVerifiedParamsStatus = "escalated"
	URLListVerifiedParamsStatusMalicious URLListVerifiedParamsStatus = "malicious"
)

type URLSubmitForVerificationParams struct {
	// The URL to submit for verification.
	URL string `json:"url" api:"required"`
	paramObj
}

func (r URLSubmitForVerificationParams) MarshalJSON() (data []byte, err error) {
	type shadow URLSubmitForVerificationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLSubmitForVerificationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
