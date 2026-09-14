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

// FunctionTriggerService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionTriggerService] method instead.
type FunctionTriggerService struct {
	Options []option.RequestOption
}

// NewFunctionTriggerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionTriggerService(opts ...option.RequestOption) (r FunctionTriggerService) {
	r = FunctionTriggerService{}
	r.Options = opts
	return
}

// Subscribe a function to one or more event types, optionally scoped to specific
// senders. Provide eventTypes and senderIds (use null in senderIds for all
// senders); a trigger is created for each event type and sender combination.
//
// The special event type `cron` runs the function on a schedule instead of a
// messaging event: include a `cron` field with a 5-field UTC cron expression
// (minimum granularity one minute). A cron trigger ignores the sender axis, and a
// function may hold several cron triggers with different expressions. The function
// receives an event with `type: "cron"` and `data.cron`.
func (r *FunctionTriggerService) New(ctx context.Context, functionID string, body FunctionTriggerNewParams, opts ...option.RequestOption) (res *FunctionTriggerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/triggers", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Enable or disable a trigger
func (r *FunctionTriggerService) Update(ctx context.Context, triggerID string, body FunctionTriggerUpdateParams, opts ...option.RequestOption) (res *FunctionTriggerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/triggers/%s", url.PathEscape(triggerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List function triggers
func (r *FunctionTriggerService) List(ctx context.Context, functionID string, opts ...option.RequestOption) (res *FunctionTriggerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionID == "" {
		err = errors.New("missing required functionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/functions/%s/triggers", url.PathEscape(functionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a trigger
func (r *FunctionTriggerService) Delete(ctx context.Context, triggerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if triggerID == "" {
		err = errors.New("missing required triggerId parameter")
		return err
	}
	path := fmt.Sprintf("v1/functions/triggers/%s", url.PathEscape(triggerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type FunctionTriggerNewResponse struct {
	Added int64 `json:"added" api:"required"`
	// Number of triggers that already existed.
	Skipped  int64                               `json:"skipped" api:"required"`
	Triggers []FunctionTriggerNewResponseTrigger `json:"triggers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Skipped     respjson.Field
		Triggers    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTriggerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionTriggerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A subscription that runs a Zavu Function when a messaging event fires.
type FunctionTriggerNewResponseTrigger struct {
	ID        string    `json:"id" api:"required"`
	Active    bool      `json:"active" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Event type that fires the function. See GET /v1/functions/event-types for the
	// supported list. The special type `cron` fires on a schedule instead of a
	// messaging event and carries a `cron` expression.
	EventType  string    `json:"eventType" api:"required"`
	FunctionID string    `json:"functionId" api:"required"`
	UpdatedAt  time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// 5-field cron expression (minute hour day-of-month month day-of-week), evaluated
	// in UTC. Present only on `cron` triggers.
	Cron string `json:"cron" api:"nullable"`
	// Last time the schedule fired. Null until the first fire.
	LastRunAt time.Time `json:"lastRunAt" api:"nullable" format:"date-time"`
	// Next scheduled fire time. Present only on `cron` triggers.
	NextRunAt time.Time `json:"nextRunAt" api:"nullable" format:"date-time"`
	// Restrict the trigger to a single sender. Null means all senders in the project.
	SenderID string `json:"senderId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		EventType   respjson.Field
		FunctionID  respjson.Field
		UpdatedAt   respjson.Field
		Cron        respjson.Field
		LastRunAt   respjson.Field
		NextRunAt   respjson.Field
		SenderID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTriggerNewResponseTrigger) RawJSON() string { return r.JSON.raw }
func (r *FunctionTriggerNewResponseTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTriggerUpdateResponse struct {
	Active bool `json:"active" api:"required"`
	Ok     bool `json:"ok" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Ok          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTriggerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionTriggerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTriggerListResponse struct {
	Triggers []FunctionTriggerListResponseTrigger `json:"triggers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Triggers    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTriggerListResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionTriggerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A subscription that runs a Zavu Function when a messaging event fires.
type FunctionTriggerListResponseTrigger struct {
	ID        string    `json:"id" api:"required"`
	Active    bool      `json:"active" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Event type that fires the function. See GET /v1/functions/event-types for the
	// supported list. The special type `cron` fires on a schedule instead of a
	// messaging event and carries a `cron` expression.
	EventType  string    `json:"eventType" api:"required"`
	FunctionID string    `json:"functionId" api:"required"`
	UpdatedAt  time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// 5-field cron expression (minute hour day-of-month month day-of-week), evaluated
	// in UTC. Present only on `cron` triggers.
	Cron string `json:"cron" api:"nullable"`
	// Last time the schedule fired. Null until the first fire.
	LastRunAt time.Time `json:"lastRunAt" api:"nullable" format:"date-time"`
	// Next scheduled fire time. Present only on `cron` triggers.
	NextRunAt time.Time `json:"nextRunAt" api:"nullable" format:"date-time"`
	// Restrict the trigger to a single sender. Null means all senders in the project.
	SenderID string `json:"senderId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		EventType   respjson.Field
		FunctionID  respjson.Field
		UpdatedAt   respjson.Field
		Cron        respjson.Field
		LastRunAt   respjson.Field
		NextRunAt   respjson.Field
		SenderID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTriggerListResponseTrigger) RawJSON() string { return r.JSON.raw }
func (r *FunctionTriggerListResponseTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTriggerNewParams struct {
	// Event types to subscribe to.
	EventTypes []string `json:"eventTypes,omitzero" api:"required"`
	// Senders to scope the triggers to. Use null for all senders.
	SenderIDs []string `json:"senderIds,omitzero" api:"required"`
	// Required when eventTypes includes `cron`: a 5-field cron expression (minute hour
	// day-of-month month day-of-week), evaluated in UTC.
	Cron param.Opt[string] `json:"cron,omitzero"`
	paramObj
}

func (r FunctionTriggerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionTriggerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionTriggerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTriggerUpdateParams struct {
	Active bool `json:"active" api:"required"`
	paramObj
}

func (r FunctionTriggerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionTriggerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionTriggerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
