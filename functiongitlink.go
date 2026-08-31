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
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// FunctionGitLinkService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionGitLinkService] method instead.
type FunctionGitLinkService struct {
	Options []option.RequestOption
}

// NewFunctionGitLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionGitLinkService(opts ...option.RequestOption) (r FunctionGitLinkService) {
	r = FunctionGitLinkService{}
	r.Options = opts
	return
}

// The link and its last deploy. Never returns the webhook secret.
func (r *FunctionGitLinkService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionGitLinkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/git-link", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Change the branch, the root directory, or whether pushes deploy. Pass at least
// one field. `rootDir: null` clears the subdirectory.
func (r *FunctionGitLinkService) Update(ctx context.Context, functionID string, body FunctionGitLinkUpdateParams, opts ...option.RequestOption) (res *FunctionGitLinkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/git-link", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Fetch the linked branch and deploy it without waiting for a push. Returns
// immediately; follow the outcome with `GET /v1/functions/{functionId}/git-link`,
// whose `lastStatus` and `lastError` describe the run.
func (r *FunctionGitLinkService) DeployNow(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionGitLinkDeployNowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/git-link/deploy", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Bind a repository to this function so every push to `branch` deploys it. A
// function holds at most one link; linking again returns 400.
//
// **The server decides how the link authenticates.** If the project has the Zavu
// GitHub App installed, the link uses that installation: private repositories work
// and there is nothing to configure in the repository. Otherwise it falls back to
// a manual link and the response carries a `webhookSecret` you add to the
// repository yourself. `connection` says which one you got.
//
// The repository is not checked against GitHub here, because it cannot be: an
// owner/repo that does not exist, or that the installation cannot see, is accepted
// and fails on the first deploy with a fetch error.
func (r *FunctionGitLinkService) Link(ctx context.Context, functionID string, body FunctionGitLinkLinkParams, opts ...option.RequestOption) (res *FunctionGitLinkLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/git-link", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove the link. The function and its deployments stay. A manual webhook left in
// the repository stops being accepted, so remove it there too.
func (r *FunctionGitLinkService) Unlink(ctx context.Context, functionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return err
	}
	path := fmt.Sprintf("v1/functions/%s/git-link", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type FunctionGitLinkGetResponse struct {
	// A GitHub repository bound to a function. A push to `branch` deploys the
	// function. A function holds at most one link.
	Link FunctionGitLinkGetResponseLink `json:"link" api:"required"`
	// Endpoint that receives GitHub's push deliveries. Only needed on a `manual` link,
	// where you add it to the repository yourself.
	WebhookURL string `json:"webhookUrl" api:"required" format:"uri"`
	// Shared secret for the repository's webhook. **Returned only when creating a
	// `manual` link, and only there** — every later read strips it, and re-linking
	// mints a new one. Absent entirely on an `app` link, which needs no secret of its
	// own.
	WebhookSecret string `json:"webhookSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link          respjson.Field
		WebhookURL    respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A GitHub repository bound to a function. A push to `branch` deploys the
// function. A function holds at most one link.
type FunctionGitLinkGetResponseLink struct {
	ID string `json:"id" api:"required"`
	// When false the link is kept and pushes are ignored.
	AutoDeploy bool `json:"autoDeploy" api:"required"`
	// Only pushes to this branch deploy.
	Branch string `json:"branch" api:"required"`
	// How this link authenticates, decided by the server rather than by the caller.
	//
	//   - `app`: the Zavu GitHub App is installed on the account. Pushes arrive on the
	//     app's webhook and private repositories work. Nothing to configure in the
	//     repository.
	//   - `manual`: no installation. The link carries its own secret and you add the
	//     webhook to the repository yourself.
	//
	// Any of "app", "manual".
	Connection string    `json:"connection" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	FunctionID string    `json:"functionId" api:"required"`
	Owner      string    `json:"owner" api:"required"`
	// Any of "github".
	Provider          string    `json:"provider" api:"required"`
	Repo              string    `json:"repo" api:"required"`
	UpdatedAt         time.Time `json:"updatedAt" api:"required" format:"date-time"`
	LastCommitMessage string    `json:"lastCommitMessage" api:"nullable"`
	LastCommitSha     string    `json:"lastCommitSha" api:"nullable"`
	LastDeployAt      time.Time `json:"lastDeployAt" api:"nullable" format:"date-time"`
	// Why the last deploy failed. Null otherwise.
	LastError string `json:"lastError" api:"nullable"`
	// Any of "deploying", "deployed", "failed".
	LastStatus string `json:"lastStatus" api:"nullable"`
	// Subdirectory holding the project, for monorepos. Null when the project is at the
	// repository root.
	RootDir string `json:"rootDir" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AutoDeploy        respjson.Field
		Branch            respjson.Field
		Connection        respjson.Field
		CreatedAt         respjson.Field
		FunctionID        respjson.Field
		Owner             respjson.Field
		Provider          respjson.Field
		Repo              respjson.Field
		UpdatedAt         respjson.Field
		LastCommitMessage respjson.Field
		LastCommitSha     respjson.Field
		LastDeployAt      respjson.Field
		LastError         respjson.Field
		LastStatus        respjson.Field
		RootDir           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkGetResponseLink) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkGetResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGitLinkUpdateResponse struct {
	// A GitHub repository bound to a function. A push to `branch` deploys the
	// function. A function holds at most one link.
	Link FunctionGitLinkUpdateResponseLink `json:"link" api:"required"`
	// Endpoint that receives GitHub's push deliveries. Only needed on a `manual` link,
	// where you add it to the repository yourself.
	WebhookURL string `json:"webhookUrl" api:"required" format:"uri"`
	// Shared secret for the repository's webhook. **Returned only when creating a
	// `manual` link, and only there** — every later read strips it, and re-linking
	// mints a new one. Absent entirely on an `app` link, which needs no secret of its
	// own.
	WebhookSecret string `json:"webhookSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link          respjson.Field
		WebhookURL    respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A GitHub repository bound to a function. A push to `branch` deploys the
// function. A function holds at most one link.
type FunctionGitLinkUpdateResponseLink struct {
	ID string `json:"id" api:"required"`
	// When false the link is kept and pushes are ignored.
	AutoDeploy bool `json:"autoDeploy" api:"required"`
	// Only pushes to this branch deploy.
	Branch string `json:"branch" api:"required"`
	// How this link authenticates, decided by the server rather than by the caller.
	//
	//   - `app`: the Zavu GitHub App is installed on the account. Pushes arrive on the
	//     app's webhook and private repositories work. Nothing to configure in the
	//     repository.
	//   - `manual`: no installation. The link carries its own secret and you add the
	//     webhook to the repository yourself.
	//
	// Any of "app", "manual".
	Connection string    `json:"connection" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	FunctionID string    `json:"functionId" api:"required"`
	Owner      string    `json:"owner" api:"required"`
	// Any of "github".
	Provider          string    `json:"provider" api:"required"`
	Repo              string    `json:"repo" api:"required"`
	UpdatedAt         time.Time `json:"updatedAt" api:"required" format:"date-time"`
	LastCommitMessage string    `json:"lastCommitMessage" api:"nullable"`
	LastCommitSha     string    `json:"lastCommitSha" api:"nullable"`
	LastDeployAt      time.Time `json:"lastDeployAt" api:"nullable" format:"date-time"`
	// Why the last deploy failed. Null otherwise.
	LastError string `json:"lastError" api:"nullable"`
	// Any of "deploying", "deployed", "failed".
	LastStatus string `json:"lastStatus" api:"nullable"`
	// Subdirectory holding the project, for monorepos. Null when the project is at the
	// repository root.
	RootDir string `json:"rootDir" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AutoDeploy        respjson.Field
		Branch            respjson.Field
		Connection        respjson.Field
		CreatedAt         respjson.Field
		FunctionID        respjson.Field
		Owner             respjson.Field
		Provider          respjson.Field
		Repo              respjson.Field
		UpdatedAt         respjson.Field
		LastCommitMessage respjson.Field
		LastCommitSha     respjson.Field
		LastDeployAt      respjson.Field
		LastError         respjson.Field
		LastStatus        respjson.Field
		RootDir           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkUpdateResponseLink) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkUpdateResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGitLinkDeployNowResponse struct {
	Scheduled bool `json:"scheduled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scheduled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkDeployNowResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkDeployNowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGitLinkLinkResponse struct {
	// A GitHub repository bound to a function. A push to `branch` deploys the
	// function. A function holds at most one link.
	Link FunctionGitLinkLinkResponseLink `json:"link" api:"required"`
	// Endpoint that receives GitHub's push deliveries. Only needed on a `manual` link,
	// where you add it to the repository yourself.
	WebhookURL string `json:"webhookUrl" api:"required" format:"uri"`
	// Shared secret for the repository's webhook. **Returned only when creating a
	// `manual` link, and only there** — every later read strips it, and re-linking
	// mints a new one. Absent entirely on an `app` link, which needs no secret of its
	// own.
	WebhookSecret string `json:"webhookSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link          respjson.Field
		WebhookURL    respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkLinkResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkLinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A GitHub repository bound to a function. A push to `branch` deploys the
// function. A function holds at most one link.
type FunctionGitLinkLinkResponseLink struct {
	ID string `json:"id" api:"required"`
	// When false the link is kept and pushes are ignored.
	AutoDeploy bool `json:"autoDeploy" api:"required"`
	// Only pushes to this branch deploy.
	Branch string `json:"branch" api:"required"`
	// How this link authenticates, decided by the server rather than by the caller.
	//
	//   - `app`: the Zavu GitHub App is installed on the account. Pushes arrive on the
	//     app's webhook and private repositories work. Nothing to configure in the
	//     repository.
	//   - `manual`: no installation. The link carries its own secret and you add the
	//     webhook to the repository yourself.
	//
	// Any of "app", "manual".
	Connection string    `json:"connection" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	FunctionID string    `json:"functionId" api:"required"`
	Owner      string    `json:"owner" api:"required"`
	// Any of "github".
	Provider          string    `json:"provider" api:"required"`
	Repo              string    `json:"repo" api:"required"`
	UpdatedAt         time.Time `json:"updatedAt" api:"required" format:"date-time"`
	LastCommitMessage string    `json:"lastCommitMessage" api:"nullable"`
	LastCommitSha     string    `json:"lastCommitSha" api:"nullable"`
	LastDeployAt      time.Time `json:"lastDeployAt" api:"nullable" format:"date-time"`
	// Why the last deploy failed. Null otherwise.
	LastError string `json:"lastError" api:"nullable"`
	// Any of "deploying", "deployed", "failed".
	LastStatus string `json:"lastStatus" api:"nullable"`
	// Subdirectory holding the project, for monorepos. Null when the project is at the
	// repository root.
	RootDir string `json:"rootDir" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AutoDeploy        respjson.Field
		Branch            respjson.Field
		Connection        respjson.Field
		CreatedAt         respjson.Field
		FunctionID        respjson.Field
		Owner             respjson.Field
		Provider          respjson.Field
		Repo              respjson.Field
		UpdatedAt         respjson.Field
		LastCommitMessage respjson.Field
		LastCommitSha     respjson.Field
		LastDeployAt      respjson.Field
		LastError         respjson.Field
		LastStatus        respjson.Field
		RootDir           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGitLinkLinkResponseLink) RawJSON() string { return r.JSON.raw }
func (r *FunctionGitLinkLinkResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGitLinkUpdateParams struct {
	RootDir    param.Opt[string] `json:"rootDir,omitzero"`
	AutoDeploy param.Opt[bool]   `json:"autoDeploy,omitzero"`
	Branch     param.Opt[string] `json:"branch,omitzero"`
	paramObj
}

func (r FunctionGitLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionGitLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionGitLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGitLinkLinkParams struct {
	Owner      string            `json:"owner" api:"required"`
	Repo       string            `json:"repo" api:"required"`
	AutoDeploy param.Opt[bool]   `json:"autoDeploy,omitzero"`
	Branch     param.Opt[string] `json:"branch,omitzero"`
	// Subdirectory holding the project, for monorepos.
	RootDir param.Opt[string] `json:"rootDir,omitzero"`
	paramObj
}

func (r FunctionGitLinkLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionGitLinkLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionGitLinkLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
