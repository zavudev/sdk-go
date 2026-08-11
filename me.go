// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"net/http"
	"slices"

	"github.com/zavudev/sdk-go/internal/apijson"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// MeService contains methods and other services that help with interacting with
// the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options []option.RequestOption
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.Options = opts
	return
}

// Returns the project, team, and API key metadata bound to the current Bearer
// token. Used by CLIs and SDKs to confirm which project they will operate on.
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MeGetResponse struct {
	APIKey     MeGetResponseAPIKey  `json:"apiKey" api:"required"`
	IsTestMode bool                 `json:"isTestMode" api:"required"`
	Project    MeGetResponseProject `json:"project" api:"required"`
	Team       MeGetResponseTeam    `json:"team" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		IsTestMode  respjson.Field
		Project     respjson.Field
		Team        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseAPIKey struct {
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseAPIKey) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseProject struct {
	ID           string `json:"id" api:"required"`
	IsSubAccount bool   `json:"isSubAccount" api:"required"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		IsSubAccount respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseProject) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeGetResponseTeam struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseTeam) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
