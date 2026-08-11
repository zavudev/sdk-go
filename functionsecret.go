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

// FunctionSecretService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionSecretService] method instead.
type FunctionSecretService struct {
	Options []option.RequestOption
}

// NewFunctionSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionSecretService(opts ...option.RequestOption) (r FunctionSecretService) {
	r = FunctionSecretService{}
	r.Options = opts
	return
}

// Lists every secret key set on the function. Plaintext is NEVER returned — only
// the last 4 characters of each value, for visual confirmation.
func (r *FunctionSecretService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionSecretListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/secrets", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create or update a secret on a function. Marks the function out-of-sync; the
// next `POST /deploy` re-publishes the Lambda with the new env. Keys must match
// `[A-Z_][A-Z0-9_]*` (uppercase env-var style) and cannot start with reserved
// prefixes (AWS*, LAMBDA*, etc).
func (r *FunctionSecretService) Set(ctx context.Context, key string, params FunctionSecretSetParams, opts ...option.RequestOption) (res *FunctionSecretSetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.FunctionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/secrets/%s", url.PathEscape(params.FunctionID), url.PathEscape(key))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Remove a secret from a function. Doesn't take effect on the running Lambda until
// the next deploy.
func (r *FunctionSecretService) Unset(ctx context.Context, key string, body FunctionSecretUnsetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.FunctionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return err
	}
	path := fmt.Sprintf("v1/functions/%s/secrets/%s", url.PathEscape(body.FunctionID), url.PathEscape(key))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type FunctionSecretListResponse struct {
	Secrets []FunctionSecretListResponseSecret `json:"secrets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secrets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionSecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionSecretListResponseSecret struct {
	ID          string  `json:"id" api:"required"`
	Key         string  `json:"key" api:"required"`
	ValueLast4  string  `json:"valueLast4" api:"required"`
	CreatedAt   float64 `json:"createdAt"`
	SyncedToAws bool    `json:"syncedToAws"`
	UpdatedAt   float64 `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Key         respjson.Field
		ValueLast4  respjson.Field
		CreatedAt   respjson.Field
		SyncedToAws respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSecretListResponseSecret) RawJSON() string { return r.JSON.raw }
func (r *FunctionSecretListResponseSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionSecretSetResponse = any

type FunctionSecretSetParams struct {
	FunctionID string `path:"functionId" api:"required" json:"-"`
	Value      string `json:"value" api:"required"`
	paramObj
}

func (r FunctionSecretSetParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionSecretSetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionSecretSetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionSecretUnsetParams struct {
	FunctionID string `path:"functionId" api:"required" json:"-"`
	paramObj
}
