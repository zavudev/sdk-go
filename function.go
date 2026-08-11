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
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// FunctionService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	Options []option.RequestOption
	Secrets FunctionSecretService
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r FunctionService) {
	r = FunctionService{}
	r.Options = opts
	r.Secrets = NewFunctionSecretService(opts...)
	return
}

// Create a new Zavu Function. The function starts in `draft` status. A dedicated
// API key is auto-provisioned and injected as the `ZAVU_API_KEY` secret so the
// function can call back into the Zavu API without manual setup.
//
// Provide `sourceCode` to seed the draft. Call
// `POST /v1/functions/{functionId}/deploy` afterwards to publish.
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *FunctionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get function
func (r *FunctionService) Get(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing function. `sourceCode` / `dependencies` edit the draft
// without triggering a build — they go live on the next
// `POST /v1/functions/{functionId}/deploy`. `httpEnabled` is applied to the
// deployed function immediately, so turning the public endpoint on or off does not
// require a redeploy.
func (r *FunctionService) Update(ctx context.Context, functionID string, body FunctionUpdateParams, opts ...option.RequestOption) (res *FunctionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Permanently delete a function and cascade: triggers, secrets, deployment
// history, managed agents+tools, and revoke the auto-provisioned API key. The AWS
// Lambda + log group are torn down asynchronously.
func (r *FunctionService) Delete(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Publish the function. If `sourceCode` or `dependencies` are provided in the
// body, they replace the current draft before deployment. Returns immediately with
// a deployment ID — poll `GET /v1/functions/deployments/{deploymentId}` until
// status is `active` or `failed`.
func (r *FunctionService) Deploy(ctx context.Context, functionID string, body FunctionDeployParams, opts ...option.RequestOption) (res *FunctionDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/deploy", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a deployment to poll its status during a deploy.
func (r *FunctionService) GetDeployment(ctx context.Context, deploymentID string, opts ...option.RequestOption) (res *FunctionGetDeploymentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if deploymentID == "" {
		err = errors.New("missing required deploymentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/deployments/%s", url.PathEscape(deploymentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch invocation logs for a function. Logs are paginated via `nextToken`. Pass
// `startTime` / `endTime` (Unix epoch milliseconds) to bound the window, or
// `filterPattern` to filter messages.
func (r *FunctionService) TailLogs(ctx context.Context, functionID string, query FunctionTailLogsParams, opts ...option.RequestOption) (res *FunctionTailLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/logs", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FunctionNewResponse struct {
	// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
	// messaging events or HTTP requests.
	Function FunctionNewResponseFunction `json:"function" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
// messaging events or HTTP requests.
type FunctionNewResponseFunction struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// npm dependencies installed in the function bundle. Keys are package names,
	// values are semver ranges.
	Dependencies map[string]string `json:"dependencies" api:"required"`
	// Whether the function can be invoked over HTTPS via its public URL.
	HTTPEnabled bool `json:"httpEnabled" api:"required"`
	// Memory allocation in MB.
	MemoryMB int64  `json:"memoryMb" api:"required"`
	Name     string `json:"name" api:"required"`
	// Runtime the function is deployed on.
	//
	// Any of "nodejs24".
	Runtime string `json:"runtime" api:"required"`
	// URL-safe identifier, unique per project.
	Slug string `json:"slug" api:"required"`
	// Lifecycle status of a Zavu Function.
	//
	// Any of "draft", "bundling", "deploying", "active", "failed", "disabled".
	Status string `json:"status" api:"required"`
	// Per-invocation timeout in seconds.
	TimeoutSec int64     `json:"timeoutSec" api:"required"`
	UpdatedAt  time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the deployment currently serving traffic.
	ActiveDeploymentID string `json:"activeDeploymentId" api:"nullable"`
	Description        string `json:"description" api:"nullable"`
	// HTTPS endpoint, present only while httpEnabled is true. Null otherwise,
	// including for a function that was previously exposed — the stored URL stops
	// serving the moment HTTP is turned off, so it is never returned.
	PublicURL string `json:"publicUrl" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Dependencies       respjson.Field
		HTTPEnabled        respjson.Field
		MemoryMB           respjson.Field
		Name               respjson.Field
		Runtime            respjson.Field
		Slug               respjson.Field
		Status             respjson.Field
		TimeoutSec         respjson.Field
		UpdatedAt          respjson.Field
		ActiveDeploymentID respjson.Field
		Description        respjson.Field
		PublicURL          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionNewResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *FunctionNewResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetResponse struct {
	// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
	// messaging events or HTTP requests.
	Function FunctionGetResponseFunction `json:"function" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
// messaging events or HTTP requests.
type FunctionGetResponseFunction struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// npm dependencies installed in the function bundle. Keys are package names,
	// values are semver ranges.
	Dependencies map[string]string `json:"dependencies" api:"required"`
	// Whether the function can be invoked over HTTPS via its public URL.
	HTTPEnabled bool `json:"httpEnabled" api:"required"`
	// Memory allocation in MB.
	MemoryMB int64  `json:"memoryMb" api:"required"`
	Name     string `json:"name" api:"required"`
	// Runtime the function is deployed on.
	//
	// Any of "nodejs24".
	Runtime string `json:"runtime" api:"required"`
	// URL-safe identifier, unique per project.
	Slug string `json:"slug" api:"required"`
	// Lifecycle status of a Zavu Function.
	//
	// Any of "draft", "bundling", "deploying", "active", "failed", "disabled".
	Status string `json:"status" api:"required"`
	// Per-invocation timeout in seconds.
	TimeoutSec int64     `json:"timeoutSec" api:"required"`
	UpdatedAt  time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the deployment currently serving traffic.
	ActiveDeploymentID string `json:"activeDeploymentId" api:"nullable"`
	Description        string `json:"description" api:"nullable"`
	// HTTPS endpoint, present only while httpEnabled is true. Null otherwise,
	// including for a function that was previously exposed — the stored URL stops
	// serving the moment HTTP is turned off, so it is never returned.
	PublicURL string `json:"publicUrl" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Dependencies       respjson.Field
		HTTPEnabled        respjson.Field
		MemoryMB           respjson.Field
		Name               respjson.Field
		Runtime            respjson.Field
		Slug               respjson.Field
		Status             respjson.Field
		TimeoutSec         respjson.Field
		UpdatedAt          respjson.Field
		ActiveDeploymentID respjson.Field
		Description        respjson.Field
		PublicURL          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateResponse struct {
	// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
	// messaging events or HTTP requests.
	Function FunctionUpdateResponseFunction `json:"function" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Zavu Function — user-supplied TypeScript that runs in Zavu Cloud and reacts to
// messaging events or HTTP requests.
type FunctionUpdateResponseFunction struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// npm dependencies installed in the function bundle. Keys are package names,
	// values are semver ranges.
	Dependencies map[string]string `json:"dependencies" api:"required"`
	// Whether the function can be invoked over HTTPS via its public URL.
	HTTPEnabled bool `json:"httpEnabled" api:"required"`
	// Memory allocation in MB.
	MemoryMB int64  `json:"memoryMb" api:"required"`
	Name     string `json:"name" api:"required"`
	// Runtime the function is deployed on.
	//
	// Any of "nodejs24".
	Runtime string `json:"runtime" api:"required"`
	// URL-safe identifier, unique per project.
	Slug string `json:"slug" api:"required"`
	// Lifecycle status of a Zavu Function.
	//
	// Any of "draft", "bundling", "deploying", "active", "failed", "disabled".
	Status string `json:"status" api:"required"`
	// Per-invocation timeout in seconds.
	TimeoutSec int64     `json:"timeoutSec" api:"required"`
	UpdatedAt  time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// ID of the deployment currently serving traffic.
	ActiveDeploymentID string `json:"activeDeploymentId" api:"nullable"`
	Description        string `json:"description" api:"nullable"`
	// HTTPS endpoint, present only while httpEnabled is true. Null otherwise,
	// including for a function that was previously exposed — the stored URL stops
	// serving the moment HTTP is turned off, so it is never returned.
	PublicURL string `json:"publicUrl" api:"nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Dependencies       respjson.Field
		HTTPEnabled        respjson.Field
		MemoryMB           respjson.Field
		Name               respjson.Field
		Runtime            respjson.Field
		Slug               respjson.Field
		Status             respjson.Field
		TimeoutSec         respjson.Field
		UpdatedAt          respjson.Field
		ActiveDeploymentID respjson.Field
		Description        respjson.Field
		PublicURL          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionUpdateResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *FunctionUpdateResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionDeleteResponse struct {
	Deleted bool   `json:"deleted" api:"required"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionDeployResponse struct {
	Deployment FunctionDeployResponseDeployment `json:"deployment" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionDeployResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionDeployResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionDeployResponseDeployment struct {
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	FunctionID string    `json:"functionId" api:"required"`
	// Stage of a function deployment.
	//
	// Any of "pending", "bundling", "uploading", "publishing", "active", "failed",
	// "superseded".
	Status string `json:"status" api:"required"`
	// Monotonically increasing deployment version, starting at 1.
	Version int64 `json:"version" api:"required"`
	// What the build printed: dependency installation, the bundler's output, and the
	// compiler's message when it failed. Returned when fetching a single deployment,
	// omitted from the list. Read this first when a deploy fails — `errorMessage` is
	// often the outer wrapper's summary, and the line that names the broken import or
	// the syntax error is here.
	BuildLogs string `json:"buildLogs" api:"nullable"`
	// Size of the built bundle in bytes. Null until the build finishes.
	BundleBytes int64     `json:"bundleBytes" api:"nullable"`
	DeployedAt  time.Time `json:"deployedAt" api:"nullable" format:"date-time"`
	// Failure reason when status is 'failed'.
	ErrorMessage string `json:"errorMessage" api:"nullable"`
	// Total size of the deployed source tree in bytes.
	SourceCodeBytes int64 `json:"sourceCodeBytes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FunctionID      respjson.Field
		Status          respjson.Field
		Version         respjson.Field
		BuildLogs       respjson.Field
		BundleBytes     respjson.Field
		DeployedAt      respjson.Field
		ErrorMessage    respjson.Field
		SourceCodeBytes respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionDeployResponseDeployment) RawJSON() string { return r.JSON.raw }
func (r *FunctionDeployResponseDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetDeploymentResponse struct {
	Deployment FunctionGetDeploymentResponseDeployment `json:"deployment" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deployment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetDeploymentResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetDeploymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetDeploymentResponseDeployment struct {
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"createdAt" api:"required" format:"date-time"`
	FunctionID string    `json:"functionId" api:"required"`
	// Stage of a function deployment.
	//
	// Any of "pending", "bundling", "uploading", "publishing", "active", "failed",
	// "superseded".
	Status string `json:"status" api:"required"`
	// Monotonically increasing deployment version, starting at 1.
	Version int64 `json:"version" api:"required"`
	// What the build printed: dependency installation, the bundler's output, and the
	// compiler's message when it failed. Returned when fetching a single deployment,
	// omitted from the list. Read this first when a deploy fails — `errorMessage` is
	// often the outer wrapper's summary, and the line that names the broken import or
	// the syntax error is here.
	BuildLogs string `json:"buildLogs" api:"nullable"`
	// Size of the built bundle in bytes. Null until the build finishes.
	BundleBytes int64     `json:"bundleBytes" api:"nullable"`
	DeployedAt  time.Time `json:"deployedAt" api:"nullable" format:"date-time"`
	// Failure reason when status is 'failed'.
	ErrorMessage string `json:"errorMessage" api:"nullable"`
	// Total size of the deployed source tree in bytes.
	SourceCodeBytes int64 `json:"sourceCodeBytes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FunctionID      respjson.Field
		Status          respjson.Field
		Version         respjson.Field
		BuildLogs       respjson.Field
		BundleBytes     respjson.Field
		DeployedAt      respjson.Field
		ErrorMessage    respjson.Field
		SourceCodeBytes respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetDeploymentResponseDeployment) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetDeploymentResponseDeployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTailLogsResponse struct {
	Events []FunctionTailLogsResponseEvent `json:"events" api:"required"`
	// Pass to the next request to fetch the following page of logs.
	NextToken string `json:"nextToken" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		NextToken   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTailLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionTailLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTailLogsResponseEvent struct {
	Message   string    `json:"message" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTailLogsResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *FunctionTailLogsResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParams struct {
	Name string `json:"name" api:"required"`
	// URL-safe identifier (lowercase, digits, hyphens). Must be unique per project.
	Slug        string            `json:"slug" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Which file in `files` is the entry point. Defaults to `index.ts`.
	Entrypoint param.Opt[string] `json:"entrypoint,omitzero"`
	// Whether to expose a public HTTPS URL for this function.
	HTTPEnabled param.Opt[bool] `json:"httpEnabled,omitzero"`
	// Shortcut for a single-file function: exactly equivalent to sending `files` with
	// one entry named after `entrypoint` (`index.ts` by default). Fully supported —
	// use whichever fits. If both are sent, `files` wins.
	SourceCode param.Opt[string] `json:"sourceCode,omitzero"`
	// Per-invocation timeout in seconds. Event and cron invocations are asynchronous,
	// so a long timeout only bounds cost; a tool called during a live conversation
	// holds up the reply, and a function exposed over HTTP is additionally bounded by
	// the platform's HTTP response limit.
	TimeoutSec param.Opt[int64] `json:"timeoutSec,omitzero"`
	// npm dependencies. Keys are package names, values are semver ranges.
	Dependencies map[string]string `json:"dependencies,omitzero"`
	// The project's source files, keyed by path relative to the project root (e.g.
	// `index.ts`, `lib/orders.ts`). Imports between them are resolved when the
	// function is built, so a function can be split across as many files as it needs.
	//
	// Paths must be relative and use forward slashes; `..`, `node_modules/` and
	// `package.json` are rejected. npm packages are not uploaded here — declare them
	// under `dependencies` and Zavu installs them. Limits: 200 files and 900,000 bytes
	// for the whole tree.
	Files map[string]string `json:"files,omitzero"`
	// Any of 128, 256, 512, 1024.
	MemoryMB int64 `json:"memoryMb,omitzero"`
	// Runtime the function is deployed on.
	//
	// Any of "nodejs24".
	Runtime FunctionNewParamsRuntime `json:"runtime,omitzero"`
	paramObj
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime the function is deployed on.
type FunctionNewParamsRuntime string

const (
	FunctionNewParamsRuntimeNodejs24 FunctionNewParamsRuntime = "nodejs24"
)

type FunctionUpdateParams struct {
	// Which file in `files` is the entry point. Defaults to `index.ts`.
	Entrypoint param.Opt[string] `json:"entrypoint,omitzero"`
	// Expose the function on its public HTTPS URL, or take it down. Applies to the
	// already-deployed function without redeploying; the URL is returned as
	// `publicUrl`.
	HTTPEnabled param.Opt[bool] `json:"httpEnabled,omitzero"`
	// Shortcut for a single-file function: exactly equivalent to sending `files` with
	// one entry named after `entrypoint` (`index.ts` by default). Fully supported —
	// use whichever fits. If both are sent, `files` wins.
	SourceCode param.Opt[string] `json:"sourceCode,omitzero"`
	// New dependency map (replaces existing dependencies).
	Dependencies map[string]string `json:"dependencies,omitzero"`
	// The project's source files, keyed by path relative to the project root (e.g.
	// `index.ts`, `lib/orders.ts`). Imports between them are resolved when the
	// function is built, so a function can be split across as many files as it needs.
	//
	// Paths must be relative and use forward slashes; `..`, `node_modules/` and
	// `package.json` are rejected. npm packages are not uploaded here — declare them
	// under `dependencies` and Zavu installs them. Limits: 200 files and 900,000 bytes
	// for the whole tree.
	Files map[string]string `json:"files,omitzero"`
	paramObj
}

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionDeployParams struct {
	// Which file in `files` is the entry point. Defaults to `index.ts`.
	Entrypoint param.Opt[string] `json:"entrypoint,omitzero"`
	// Shortcut for a single-file function: exactly equivalent to sending `files` with
	// one entry named after `entrypoint` (`index.ts` by default). Fully supported —
	// use whichever fits. If both are sent, `files` wins.
	SourceCode param.Opt[string] `json:"sourceCode,omitzero"`
	// New dependency map (replaces existing dependencies).
	Dependencies map[string]string `json:"dependencies,omitzero"`
	// The project's source files, keyed by path relative to the project root (e.g.
	// `index.ts`, `lib/orders.ts`). Imports between them are resolved when the
	// function is built, so a function can be split across as many files as it needs.
	//
	// Paths must be relative and use forward slashes; `..`, `node_modules/` and
	// `package.json` are rejected. npm packages are not uploaded here — declare them
	// under `dependencies` and Zavu installs them. Limits: 200 files and 900,000 bytes
	// for the whole tree.
	Files map[string]string `json:"files,omitzero"`
	paramObj
}

func (r FunctionDeployParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionDeployParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionDeployParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTailLogsParams struct {
	// End of the log window in Unix epoch milliseconds.
	EndTime       param.Opt[int64]  `query:"endTime,omitzero" json:"-"`
	FilterPattern param.Opt[string] `query:"filterPattern,omitzero" json:"-"`
	Limit         param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	NextToken     param.Opt[string] `query:"nextToken,omitzero" json:"-"`
	// Start of the log window in Unix epoch milliseconds.
	StartTime param.Opt[int64] `query:"startTime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunctionTailLogsParams]'s query parameters as `url.Values`.
func (r FunctionTailLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
