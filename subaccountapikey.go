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

// SubAccountAPIKeyService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubAccountAPIKeyService] method instead.
type SubAccountAPIKeyService struct {
	Options []option.RequestOption
}

// NewSubAccountAPIKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubAccountAPIKeyService(opts ...option.RequestOption) (r SubAccountAPIKeyService) {
	r = SubAccountAPIKeyService{}
	r.Options = opts
	return
}

// Create sub-account API key. Requires a parent project API key; sub-account API
// keys receive HTTP 403.
func (r *SubAccountAPIKeyService) New(ctx context.Context, id string, body SubAccountAPIKeyNewParams, opts ...option.RequestOption) (res *SubAccountAPIKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s/api-keys", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List sub-account API keys. Requires a parent project API key; sub-account API
// keys receive HTTP 403.
func (r *SubAccountAPIKeyService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *SubAccountAPIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s/api-keys", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke sub-account API key. Requires a parent project API key; sub-account API
// keys receive HTTP 403.
func (r *SubAccountAPIKeyService) Revoke(ctx context.Context, keyID string, body SubAccountAPIKeyRevokeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if keyID == "" {
		err = errors.New("missing required keyId parameter")
		return err
	}
	path := fmt.Sprintf("v1/sub-accounts/%s/api-keys/%s", url.PathEscape(body.ID), url.PathEscape(keyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type SubAccountAPIKeyNewResponse struct {
	APIKey SubAccountAPIKeyNewResponseAPIKey `json:"apiKey" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountAPIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountAPIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountAPIKeyNewResponseAPIKey struct {
	ID string `json:"id" api:"required"`
	// Any of "live", "test".
	Environment string `json:"environment" api:"required"`
	Key         string `json:"key" api:"required"`
	Name        string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountAPIKeyNewResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *SubAccountAPIKeyNewResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountAPIKeyListResponse struct {
	Items []SubAccountAPIKeyListResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountAPIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubAccountAPIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountAPIKeyListResponseItem struct {
	ID        string  `json:"id" api:"required"`
	CreatedAt float64 `json:"createdAt" api:"required"`
	// Any of "live", "test".
	Environment string `json:"environment" api:"required"`
	// First characters of the key for identification.
	KeyPrefix string `json:"keyPrefix" api:"required"`
	Name      string `json:"name" api:"required"`
	// Full API key. Only returned on creation.
	Key         string   `json:"key"`
	LastUsedAt  float64  `json:"lastUsedAt" api:"nullable"`
	Permissions []string `json:"permissions"`
	RevokedAt   float64  `json:"revokedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Environment respjson.Field
		KeyPrefix   respjson.Field
		Name        respjson.Field
		Key         respjson.Field
		LastUsedAt  respjson.Field
		Permissions respjson.Field
		RevokedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAccountAPIKeyListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *SubAccountAPIKeyListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountAPIKeyNewParams struct {
	Name string `json:"name" api:"required"`
	// Any of "live", "test".
	Environment SubAccountAPIKeyNewParamsEnvironment `json:"environment,omitzero"`
	Permissions []string                             `json:"permissions,omitzero"`
	paramObj
}

func (r SubAccountAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubAccountAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubAccountAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAccountAPIKeyNewParamsEnvironment string

const (
	SubAccountAPIKeyNewParamsEnvironmentLive SubAccountAPIKeyNewParamsEnvironment = "live"
	SubAccountAPIKeyNewParamsEnvironmentTest SubAccountAPIKeyNewParamsEnvironment = "test"
)

type SubAccountAPIKeyRevokeParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
