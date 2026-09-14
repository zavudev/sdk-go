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

// EmailDomainService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailDomainService] method instead.
type EmailDomainService struct {
	Options []option.RequestOption
}

// NewEmailDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailDomainService(opts ...option.RequestOption) (r EmailDomainService) {
	r = EmailDomainService{}
	r.Options = opts
	return
}

// Add a domain to send email from. Returns the DNS records to publish (DKIM CNAMEs
// are required; SPF, DMARC, and MAIL FROM are recommended). Publish them at your
// DNS provider, then verify.
func (r *EmailDomainService) New(ctx context.Context, body EmailDomainNewParams, opts ...option.RequestOption) (res *EmailDomainNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a domain with its DNS records and current status.
func (r *EmailDomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *EmailDomainGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/email-domains/%s", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List email domains
func (r *EmailDomainService) List(ctx context.Context, opts ...option.RequestOption) (res *EmailDomainListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove an email domain
func (r *EmailDomainService) Delete(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return err
	}
	path := fmt.Sprintf("v1/email-domains/%s", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Re-check the domain's published DNS records and refresh its status.
func (r *EmailDomainService) Verify(ctx context.Context, domainID string, opts ...option.RequestOption) (res *EmailDomainVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainID == "" {
		err = errors.New("missing required domainId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/email-domains/%s/verify", url.PathEscape(domainID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type EmailDomainNewResponse struct {
	Domain EmailDomainNewResponseDomain `json:"domain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainNewResponseDomain struct {
	ID         string `json:"id" api:"required"`
	DkimStatus string `json:"dkimStatus" api:"required"`
	Domain     string `json:"domain" api:"required"`
	// Overall verification status.
	Status string `json:"status" api:"required"`
	// DNS records to publish. Present when fetching a single domain or after adding
	// one.
	DNSRecords []EmailDomainNewResponseDomainDNSRecord `json:"dnsRecords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DkimStatus  respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		DNSRecords  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainNewResponseDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainNewResponseDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainNewResponseDomainDNSRecord struct {
	// Record host/name to create.
	Name string `json:"name" api:"required"`
	// What the record is for.
	//
	// Any of "dkim", "spf", "dmarc", "mail_from".
	Purpose string `json:"purpose" api:"required"`
	// Whether the record is required to verify + send (DKIM) or recommended for
	// deliverability.
	Required bool `json:"required" api:"required"`
	// DNS record type.
	Type string `json:"type" api:"required"`
	// Record value.
	Value string `json:"value" api:"required"`
	// Priority (MX records only).
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Purpose     respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainNewResponseDomainDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainNewResponseDomainDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainGetResponse struct {
	Domain EmailDomainGetResponseDomain `json:"domain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainGetResponseDomain struct {
	ID         string `json:"id" api:"required"`
	DkimStatus string `json:"dkimStatus" api:"required"`
	Domain     string `json:"domain" api:"required"`
	// Overall verification status.
	Status string `json:"status" api:"required"`
	// DNS records to publish. Present when fetching a single domain or after adding
	// one.
	DNSRecords []EmailDomainGetResponseDomainDNSRecord `json:"dnsRecords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DkimStatus  respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		DNSRecords  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetResponseDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetResponseDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainGetResponseDomainDNSRecord struct {
	// Record host/name to create.
	Name string `json:"name" api:"required"`
	// What the record is for.
	//
	// Any of "dkim", "spf", "dmarc", "mail_from".
	Purpose string `json:"purpose" api:"required"`
	// Whether the record is required to verify + send (DKIM) or recommended for
	// deliverability.
	Required bool `json:"required" api:"required"`
	// DNS record type.
	Type string `json:"type" api:"required"`
	// Record value.
	Value string `json:"value" api:"required"`
	// Priority (MX records only).
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Purpose     respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainGetResponseDomainDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainGetResponseDomainDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainListResponse struct {
	Items []EmailDomainListResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainListResponseItem struct {
	ID         string `json:"id" api:"required"`
	DkimStatus string `json:"dkimStatus" api:"required"`
	Domain     string `json:"domain" api:"required"`
	// Overall verification status.
	Status string `json:"status" api:"required"`
	// DNS records to publish. Present when fetching a single domain or after adding
	// one.
	DNSRecords []EmailDomainListResponseItemDNSRecord `json:"dnsRecords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DkimStatus  respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		DNSRecords  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainListResponseItemDNSRecord struct {
	// Record host/name to create.
	Name string `json:"name" api:"required"`
	// What the record is for.
	//
	// Any of "dkim", "spf", "dmarc", "mail_from".
	Purpose string `json:"purpose" api:"required"`
	// Whether the record is required to verify + send (DKIM) or recommended for
	// deliverability.
	Required bool `json:"required" api:"required"`
	// DNS record type.
	Type string `json:"type" api:"required"`
	// Record value.
	Value string `json:"value" api:"required"`
	// Priority (MX records only).
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Purpose     respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainListResponseItemDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainListResponseItemDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainVerifyResponse struct {
	Domain EmailDomainVerifyResponseDomain `json:"domain" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainVerifyResponseDomain struct {
	ID         string `json:"id" api:"required"`
	DkimStatus string `json:"dkimStatus" api:"required"`
	Domain     string `json:"domain" api:"required"`
	// Overall verification status.
	Status string `json:"status" api:"required"`
	// DNS records to publish. Present when fetching a single domain or after adding
	// one.
	DNSRecords []EmailDomainVerifyResponseDomainDNSRecord `json:"dnsRecords"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DkimStatus  respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		DNSRecords  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainVerifyResponseDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainVerifyResponseDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainVerifyResponseDomainDNSRecord struct {
	// Record host/name to create.
	Name string `json:"name" api:"required"`
	// What the record is for.
	//
	// Any of "dkim", "spf", "dmarc", "mail_from".
	Purpose string `json:"purpose" api:"required"`
	// Whether the record is required to verify + send (DKIM) or recommended for
	// deliverability.
	Required bool `json:"required" api:"required"`
	// DNS record type.
	Type string `json:"type" api:"required"`
	// Record value.
	Value string `json:"value" api:"required"`
	// Priority (MX records only).
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Purpose     respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		Priority    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomainVerifyResponseDomainDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *EmailDomainVerifyResponseDomainDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailDomainNewParams struct {
	// Bare domain, e.g. example.com.
	Domain string `json:"domain" api:"required"`
	paramObj
}

func (r EmailDomainNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailDomainNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
