// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"github.com/zavudev/sdk-go/option"
)

// Number10dlcService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcService] method instead.
type Number10dlcService struct {
	Options   []option.RequestOption
	Brands    Number10dlcBrandService
	Campaigns Number10dlcCampaignService
}

// NewNumber10dlcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumber10dlcService(opts ...option.RequestOption) (r Number10dlcService) {
	r = Number10dlcService{}
	r.Options = opts
	r.Brands = NewNumber10dlcBrandService(opts...)
	r.Campaigns = NewNumber10dlcCampaignService(opts...)
	return
}
