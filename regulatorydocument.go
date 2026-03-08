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

// RegulatoryDocumentService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegulatoryDocumentService] method instead.
type RegulatoryDocumentService struct {
	Options []option.RequestOption
}

// NewRegulatoryDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegulatoryDocumentService(opts ...option.RequestOption) (r RegulatoryDocumentService) {
	r = RegulatoryDocumentService{}
	r.Options = opts
	return
}

// Create a regulatory document record after uploading the file. Use the upload-url
// endpoint first to get an upload URL.
func (r *RegulatoryDocumentService) New(ctx context.Context, body RegulatoryDocumentNewParams, opts ...option.RequestOption) (res *RegulatoryDocumentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific regulatory document.
func (r *RegulatoryDocumentService) Get(ctx context.Context, documentID string, opts ...option.RequestOption) (res *RegulatoryDocumentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return
	}
	path := fmt.Sprintf("v1/documents/%s", url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List regulatory documents for this project.
func (r *RegulatoryDocumentService) List(ctx context.Context, query RegulatoryDocumentListParams, opts ...option.RequestOption) (res *pagination.Cursor[RegulatoryDocument], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/documents"
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

// List regulatory documents for this project.
func (r *RegulatoryDocumentService) ListAutoPaging(ctx context.Context, query RegulatoryDocumentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[RegulatoryDocument] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a regulatory document. Cannot delete verified documents.
func (r *RegulatoryDocumentService) Delete(ctx context.Context, documentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return
	}
	path := fmt.Sprintf("v1/documents/%s", url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a presigned URL to upload a document file. After uploading, use the
// storageId to create the document record.
func (r *RegulatoryDocumentService) UploadURL(ctx context.Context, opts ...option.RequestOption) (res *RegulatoryDocumentUploadURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/documents/upload-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A regulatory document for phone number requirements.
type RegulatoryDocument struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "passport", "national_id", "drivers_license", "utility_bill", "tax_id",
	// "business_registration", "proof_of_address", "other".
	DocumentType RegulatoryDocumentDocumentType `json:"documentType" api:"required"`
	Name         string                         `json:"name" api:"required"`
	// Any of "pending", "uploaded", "verified", "rejected".
	Status          RegulatoryDocumentStatus `json:"status" api:"required"`
	FileSize        int64                    `json:"fileSize"`
	MimeType        string                   `json:"mimeType"`
	RejectionReason string                   `json:"rejectionReason" api:"nullable"`
	UpdatedAt       time.Time                `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DocumentType    respjson.Field
		Name            respjson.Field
		Status          respjson.Field
		FileSize        respjson.Field
		MimeType        respjson.Field
		RejectionReason respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryDocument) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryDocumentDocumentType string

const (
	RegulatoryDocumentDocumentTypePassport             RegulatoryDocumentDocumentType = "passport"
	RegulatoryDocumentDocumentTypeNationalID           RegulatoryDocumentDocumentType = "national_id"
	RegulatoryDocumentDocumentTypeDriversLicense       RegulatoryDocumentDocumentType = "drivers_license"
	RegulatoryDocumentDocumentTypeUtilityBill          RegulatoryDocumentDocumentType = "utility_bill"
	RegulatoryDocumentDocumentTypeTaxID                RegulatoryDocumentDocumentType = "tax_id"
	RegulatoryDocumentDocumentTypeBusinessRegistration RegulatoryDocumentDocumentType = "business_registration"
	RegulatoryDocumentDocumentTypeProofOfAddress       RegulatoryDocumentDocumentType = "proof_of_address"
	RegulatoryDocumentDocumentTypeOther                RegulatoryDocumentDocumentType = "other"
)

type RegulatoryDocumentStatus string

const (
	RegulatoryDocumentStatusPending  RegulatoryDocumentStatus = "pending"
	RegulatoryDocumentStatusUploaded RegulatoryDocumentStatus = "uploaded"
	RegulatoryDocumentStatusVerified RegulatoryDocumentStatus = "verified"
	RegulatoryDocumentStatusRejected RegulatoryDocumentStatus = "rejected"
)

type RegulatoryDocumentNewResponse struct {
	// A regulatory document for phone number requirements.
	Document RegulatoryDocument `json:"document" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryDocumentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryDocumentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryDocumentGetResponse struct {
	// A regulatory document for phone number requirements.
	Document RegulatoryDocument `json:"document" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryDocumentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryDocumentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryDocumentUploadURLResponse struct {
	// Pre-signed URL for uploading the file.
	UploadURL string `json:"uploadUrl" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegulatoryDocumentUploadURLResponse) RawJSON() string { return r.JSON.raw }
func (r *RegulatoryDocumentUploadURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryDocumentNewParams struct {
	// Any of "passport", "national_id", "drivers_license", "utility_bill", "tax_id",
	// "business_registration", "proof_of_address", "other".
	DocumentType RegulatoryDocumentNewParamsDocumentType `json:"documentType,omitzero" api:"required"`
	FileSize     int64                                   `json:"fileSize" api:"required"`
	MimeType     string                                  `json:"mimeType" api:"required"`
	Name         string                                  `json:"name" api:"required"`
	// Storage ID from the upload-url endpoint.
	StorageID string `json:"storageId" api:"required"`
	paramObj
}

func (r RegulatoryDocumentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RegulatoryDocumentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegulatoryDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RegulatoryDocumentNewParamsDocumentType string

const (
	RegulatoryDocumentNewParamsDocumentTypePassport             RegulatoryDocumentNewParamsDocumentType = "passport"
	RegulatoryDocumentNewParamsDocumentTypeNationalID           RegulatoryDocumentNewParamsDocumentType = "national_id"
	RegulatoryDocumentNewParamsDocumentTypeDriversLicense       RegulatoryDocumentNewParamsDocumentType = "drivers_license"
	RegulatoryDocumentNewParamsDocumentTypeUtilityBill          RegulatoryDocumentNewParamsDocumentType = "utility_bill"
	RegulatoryDocumentNewParamsDocumentTypeTaxID                RegulatoryDocumentNewParamsDocumentType = "tax_id"
	RegulatoryDocumentNewParamsDocumentTypeBusinessRegistration RegulatoryDocumentNewParamsDocumentType = "business_registration"
	RegulatoryDocumentNewParamsDocumentTypeProofOfAddress       RegulatoryDocumentNewParamsDocumentType = "proof_of_address"
	RegulatoryDocumentNewParamsDocumentTypeOther                RegulatoryDocumentNewParamsDocumentType = "other"
)

type RegulatoryDocumentListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RegulatoryDocumentListParams]'s query parameters as
// `url.Values`.
func (r RegulatoryDocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
