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

// SubAccountService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubAccountService] method instead.
type SubAccountService struct {
	Options []option.RequestOption
	APIKeys SubAccountAPIKeyService
}

// NewSubAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubAccountService(opts ...option.RequestOption) (r SubAccountService) {
	r = SubAccountService{}
	r.Options = opts
	r.APIKeys = NewSubAccountAPIKeyService(opts...)
	return
}

// Create a new sub-account (project) with its own API key. All charges are billed
// to the parent team's balance. Use creditLimit to set a spending cap. The
// sub-account's API key is returned only in the creation response. Requires a
// parent project API key; sub-account API keys receive HTTP 403.
func (r *SubAccountService) New(ctx context.Context, body SubAccountNewParams, opts ...option.RequestOption) (res *SubAccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sub-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get sub-account. Requires a parent project API key; sub-account API keys receive
// HTTP 403.
func (r *SubAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SubAccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update sub-account. Requires a parent project API key; sub-account API keys
// receive HTTP 403.
func (r *SubAccountService) Update(ctx context.Context, id string, body SubAccountUpdateParams, opts ...option.RequestOption) (res *SubAccountUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List sub-accounts for this team. Requires a parent project API key; sub-account
// API keys receive HTTP 403.
func (r *SubAccountService) List(ctx context.Context, query SubAccountListParams, opts ...option.RequestOption) (res *pagination.Cursor[SubAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/sub-accounts"
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

// List sub-accounts for this team. Requires a parent project API key; sub-account
// API keys receive HTTP 403.
func (r *SubAccountService) ListAutoPaging(ctx context.Context, query SubAccountListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[SubAccount] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Deactivate a sub-account. Remaining balance is returned to the parent team and
// all API keys are revoked. Requires a parent project API key; sub-account API
// keys receive HTTP 403.
func (r *SubAccountService) Deactivate(ctx context.Context, id string, opts ...option.RequestOption) (res *SubAccountDeactivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get spending information for a sub-account. Returns the parent team's balance,
// the sub-account's total spending, and its credit limit (spending cap). Requires
// a parent project API key; sub-account API keys receive HTTP 403.
func (r *SubAccountService) GetBalance(ctx context.Context, id string, opts ...option.RequestOption) (res *SubAccountGetBalanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s/balance", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SubAccount struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Any of "active", "inactive".
	Status SubAccountStatus `json:"status" api:"required"`
	// Total amount spent by this sub-account in cents.
	TotalSpent int64 `json:"totalSpent" api:"required"`
	// API key for the sub-account. Only returned on creation.
	APIKey string `json:"apiKey"`
	// Spending cap in cents. When reached, messages from this sub-account will be
	// blocked.
	CreditLimit int64 `json:"creditLimit" api:"nullable"`
	// External reference ID set by the parent account.
	ExternalID string         `json:"externalId" api:"nullable"`
	Metadata   map[string]any `json:"metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		TotalSpent  respjson.Field
		APIKey      respjson.Field
		CreditLimit respjson.Field
		ExternalID  respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccount) RawJSON() string { return r.JSON.raw }
func (r *SubAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountStatus string

const (
	SubAccountStatusActive   SubAccountStatus = "active"
	SubAccountStatusInactive SubAccountStatus = "inactive"
)

type SubAccountNewResponse struct {
	SubAccount SubAccount `json:"subAccount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubAccount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountGetResponse struct {
	SubAccount SubAccount `json:"subAccount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubAccount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountUpdateResponse struct {
	SubAccount SubAccount `json:"subAccount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubAccount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountDeactivateResponse struct {
	// Number of API keys revoked.
	KeysRevoked int64  `json:"keysRevoked" api:"required"`
	Message     string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeysRevoked respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountGetBalanceResponse struct {
	// Team balance in cents. All charges are billed to the parent team.
	Balance  int64  `json:"balance" api:"required"`
	Currency string `json:"currency" api:"required"`
	// Spending cap in cents (only for sub-accounts).
	CreditLimit int64 `json:"creditLimit" api:"nullable"`
	// Whether this API key belongs to a sub-account.
	IsSubAccount bool `json:"isSubAccount"`
	// Total amount spent by this sub-account in cents (only for sub-accounts).
	TotalSpent int64 `json:"totalSpent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance      respjson.Field
		Currency     respjson.Field
		CreditLimit  respjson.Field
		IsSubAccount respjson.Field
		TotalSpent   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountGetBalanceResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountGetBalanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountNewParams struct {
	// Name of the sub-account.
	Name string `json:"name" api:"required"`
	// Spending cap in cents. When reached, messages from this sub-account will be
	// blocked. Omit or set to 0 for no limit.
	CreditLimit param.Opt[int64] `json:"creditLimit,omitzero"`
	// External reference ID for your own tracking.
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	Metadata   map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r SubAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountUpdateParams struct {
	CreditLimit param.Opt[int64]  `json:"creditLimit,omitzero"`
	ExternalID  param.Opt[string] `json:"externalId,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Metadata    map[string]any    `json:"metadata,omitzero"`
	// Any of "active", "inactive".
	Status SubAccountUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r SubAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountUpdateParamsStatus string

const (
	SubAccountUpdateParamsStatusActive   SubAccountUpdateParamsStatus = "active"
	SubAccountUpdateParamsStatusInactive SubAccountUpdateParamsStatus = "inactive"
)

type SubAccountListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubAccountListParams]'s query parameters as `url.Values`.
func (r SubAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
