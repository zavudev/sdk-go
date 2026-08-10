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

// BalanceService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r BalanceService) {
	r = BalanceService{}
	r.Options = opts
	return
}

// Get balance for the API key's team. If the API key belongs to a sub-account,
// also includes the sub-account's total spending and credit limit.
func (r *BalanceService) Get(ctx context.Context, opts ...option.RequestOption) (res *BalanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BalanceGetResponse struct {
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
func (r BalanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
