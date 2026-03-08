// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zavudev

import (
	"context"
	"encoding/json"
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

// BroadcastService contains methods and other services that help with interacting
// with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBroadcastService] method instead.
type BroadcastService struct {
	Options  []option.RequestOption
	Contacts BroadcastContactService
}

// NewBroadcastService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBroadcastService(opts ...option.RequestOption) (r BroadcastService) {
	r = BroadcastService{}
	r.Options = opts
	r.Contacts = NewBroadcastContactService(opts...)
	return
}

// Create a new broadcast campaign. Add contacts after creation, then send.
func (r *BroadcastService) New(ctx context.Context, body BroadcastNewParams, opts ...option.RequestOption) (res *BroadcastNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/broadcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get broadcast
func (r *BroadcastService) Get(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *BroadcastGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a broadcast in draft status.
func (r *BroadcastService) Update(ctx context.Context, broadcastID string, body BroadcastUpdateParams, opts ...option.RequestOption) (res *BroadcastUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List broadcasts for this project.
func (r *BroadcastService) List(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) (res *pagination.Cursor[Broadcast], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/broadcasts"
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

// List broadcasts for this project.
func (r *BroadcastService) ListAutoPaging(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Broadcast] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a broadcast in draft status.
func (r *BroadcastService) Delete(ctx context.Context, broadcastID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Cancel a broadcast. Pending contacts will be skipped, but already queued
// messages may still be delivered.
func (r *BroadcastService) Cancel(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *BroadcastCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/cancel", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get real-time progress of a broadcast including delivery counts and estimated
// completion time.
func (r *BroadcastService) Progress(ctx context.Context, broadcastID string, opts ...option.RequestOption) (res *BroadcastProgress, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/progress", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the scheduled time for a broadcast. The broadcast must be in scheduled
// status.
func (r *BroadcastService) Reschedule(ctx context.Context, broadcastID string, body BroadcastRescheduleParams, opts ...option.RequestOption) (res *BroadcastRescheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/schedule", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Start sending the broadcast immediately or schedule for later. Broadcasts go
// through automated AI content review before sending. If the review passes, the
// broadcast proceeds. If rejected, use PATCH to edit content, then call POST
// /retry-review. Reserves the estimated cost from your balance.
func (r *BroadcastService) Send(ctx context.Context, broadcastID string, body BroadcastSendParams, opts ...option.RequestOption) (res *BroadcastSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastID == "" {
		err = errors.New("missing required broadcastId parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/send", url.PathEscape(broadcastID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Broadcast struct {
	ID string `json:"id" api:"required"`
	// Broadcast delivery channel. Use 'smart' for per-contact intelligent routing.
	//
	// Any of "smart", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "voice".
	Channel   BroadcastChannel `json:"channel" api:"required"`
	CreatedAt time.Time        `json:"createdAt" api:"required" format:"date-time"`
	// Type of message for broadcast.
	//
	// Any of "text", "image", "video", "audio", "document", "template".
	MessageType BroadcastMessageType `json:"messageType" api:"required"`
	Name        string               `json:"name" api:"required"`
	// Current status of the broadcast.
	//
	// Any of "draft", "pending_review", "approved", "rejected", "escalated",
	// "rejected_final", "scheduled", "sending", "paused", "completed", "cancelled",
	// "failed".
	Status BroadcastStatus `json:"status" api:"required"`
	// Total number of contacts in the broadcast.
	TotalContacts int64 `json:"totalContacts" api:"required"`
	// Actual cost so far in USD.
	ActualCost  float64   `json:"actualCost" api:"nullable"`
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Content for non-text broadcast message types.
	Content        BroadcastContent `json:"content"`
	DeliveredCount int64            `json:"deliveredCount"`
	EmailSubject   string           `json:"emailSubject"`
	// Estimated total cost in USD.
	EstimatedCost float64           `json:"estimatedCost" api:"nullable"`
	FailedCount   int64             `json:"failedCount"`
	Metadata      map[string]string `json:"metadata"`
	PendingCount  int64             `json:"pendingCount"`
	// Amount reserved from balance in USD.
	ReservedAmount float64 `json:"reservedAmount" api:"nullable"`
	// Number of review attempts (max 3).
	ReviewAttempts int64 `json:"reviewAttempts" api:"nullable"`
	// AI content review result.
	ReviewResult BroadcastReviewResult `json:"reviewResult" api:"nullable"`
	ScheduledAt  time.Time             `json:"scheduledAt" format:"date-time"`
	SenderID     string                `json:"senderId"`
	SendingCount int64                 `json:"sendingCount"`
	StartedAt    time.Time             `json:"startedAt" format:"date-time"`
	Text         string                `json:"text"`
	UpdatedAt    time.Time             `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Channel        respjson.Field
		CreatedAt      respjson.Field
		MessageType    respjson.Field
		Name           respjson.Field
		Status         respjson.Field
		TotalContacts  respjson.Field
		ActualCost     respjson.Field
		CompletedAt    respjson.Field
		Content        respjson.Field
		DeliveredCount respjson.Field
		EmailSubject   respjson.Field
		EstimatedCost  respjson.Field
		FailedCount    respjson.Field
		Metadata       respjson.Field
		PendingCount   respjson.Field
		ReservedAmount respjson.Field
		ReviewAttempts respjson.Field
		ReviewResult   respjson.Field
		ScheduledAt    respjson.Field
		SenderID       respjson.Field
		SendingCount   respjson.Field
		StartedAt      respjson.Field
		Text           respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Broadcast) RawJSON() string { return r.JSON.raw }
func (r *Broadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AI content review result.
type BroadcastReviewResult struct {
	// Policy categories violated, if any.
	Categories []string `json:"categories"`
	// Problematic text fragments, if any.
	FlaggedContent []string `json:"flaggedContent" api:"nullable"`
	// Explanation of the review decision.
	Reasoning  string    `json:"reasoning"`
	ReviewedAt time.Time `json:"reviewedAt" format:"date-time"`
	// Content safety score from 0.0 to 1.0, where 1.0 is completely safe.
	Score float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories     respjson.Field
		FlaggedContent respjson.Field
		Reasoning      respjson.Field
		ReviewedAt     respjson.Field
		Score          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastReviewResult) RawJSON() string { return r.JSON.raw }
func (r *BroadcastReviewResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Broadcast delivery channel. Use 'smart' for per-contact intelligent routing.
type BroadcastChannel string

const (
	BroadcastChannelSmart     BroadcastChannel = "smart"
	BroadcastChannelSMS       BroadcastChannel = "sms"
	BroadcastChannelSMSOneway BroadcastChannel = "sms_oneway"
	BroadcastChannelWhatsapp  BroadcastChannel = "whatsapp"
	BroadcastChannelTelegram  BroadcastChannel = "telegram"
	BroadcastChannelEmail     BroadcastChannel = "email"
	BroadcastChannelInstagram BroadcastChannel = "instagram"
	BroadcastChannelVoice     BroadcastChannel = "voice"
)

type BroadcastContact struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Recipient string    `json:"recipient" api:"required"`
	// Any of "phone", "email".
	RecipientType BroadcastContactRecipientType `json:"recipientType" api:"required"`
	// Status of a contact within a broadcast.
	//
	// Any of "pending", "queued", "sending", "delivered", "failed", "skipped".
	Status       BroadcastContactStatus `json:"status" api:"required"`
	Cost         float64                `json:"cost" api:"nullable"`
	ErrorCode    string                 `json:"errorCode"`
	ErrorMessage string                 `json:"errorMessage"`
	// Associated message ID after processing.
	MessageID         string            `json:"messageId"`
	ProcessedAt       time.Time         `json:"processedAt" format:"date-time"`
	TemplateVariables map[string]string `json:"templateVariables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Recipient         respjson.Field
		RecipientType     respjson.Field
		Status            respjson.Field
		Cost              respjson.Field
		ErrorCode         respjson.Field
		ErrorMessage      respjson.Field
		MessageID         respjson.Field
		ProcessedAt       respjson.Field
		TemplateVariables respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastContact) RawJSON() string { return r.JSON.raw }
func (r *BroadcastContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastContactRecipientType string

const (
	BroadcastContactRecipientTypePhone BroadcastContactRecipientType = "phone"
	BroadcastContactRecipientTypeEmail BroadcastContactRecipientType = "email"
)

// Status of a contact within a broadcast.
type BroadcastContactStatus string

const (
	BroadcastContactStatusPending   BroadcastContactStatus = "pending"
	BroadcastContactStatusQueued    BroadcastContactStatus = "queued"
	BroadcastContactStatusSending   BroadcastContactStatus = "sending"
	BroadcastContactStatusDelivered BroadcastContactStatus = "delivered"
	BroadcastContactStatusFailed    BroadcastContactStatus = "failed"
	BroadcastContactStatusSkipped   BroadcastContactStatus = "skipped"
)

// Content for non-text broadcast message types.
type BroadcastContent struct {
	// Filename for documents.
	Filename string `json:"filename"`
	// Media ID if already uploaded.
	MediaID string `json:"mediaId"`
	// URL of the media file.
	MediaURL string `json:"mediaUrl"`
	// MIME type of the media.
	MimeType string `json:"mimeType"`
	// Template ID for template messages.
	TemplateID string `json:"templateId"`
	// Default template variables (can be overridden per contact).
	TemplateVariables map[string]string `json:"templateVariables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename          respjson.Field
		MediaID           respjson.Field
		MediaURL          respjson.Field
		MimeType          respjson.Field
		TemplateID        respjson.Field
		TemplateVariables respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastContent) RawJSON() string { return r.JSON.raw }
func (r *BroadcastContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BroadcastContent to a BroadcastContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BroadcastContentParam.Overrides()
func (r BroadcastContent) ToParam() BroadcastContentParam {
	return param.Override[BroadcastContentParam](json.RawMessage(r.RawJSON()))
}

// Content for non-text broadcast message types.
type BroadcastContentParam struct {
	// Filename for documents.
	Filename param.Opt[string] `json:"filename,omitzero"`
	// Media ID if already uploaded.
	MediaID param.Opt[string] `json:"mediaId,omitzero"`
	// URL of the media file.
	MediaURL param.Opt[string] `json:"mediaUrl,omitzero"`
	// MIME type of the media.
	MimeType param.Opt[string] `json:"mimeType,omitzero"`
	// Template ID for template messages.
	TemplateID param.Opt[string] `json:"templateId,omitzero"`
	// Default template variables (can be overridden per contact).
	TemplateVariables map[string]string `json:"templateVariables,omitzero"`
	paramObj
}

func (r BroadcastContentParam) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of message for broadcast.
type BroadcastMessageType string

const (
	BroadcastMessageTypeText     BroadcastMessageType = "text"
	BroadcastMessageTypeImage    BroadcastMessageType = "image"
	BroadcastMessageTypeVideo    BroadcastMessageType = "video"
	BroadcastMessageTypeAudio    BroadcastMessageType = "audio"
	BroadcastMessageTypeDocument BroadcastMessageType = "document"
	BroadcastMessageTypeTemplate BroadcastMessageType = "template"
)

type BroadcastProgress struct {
	BroadcastID string `json:"broadcastId" api:"required"`
	// Successfully delivered.
	Delivered int64 `json:"delivered" api:"required"`
	// Failed to deliver.
	Failed int64 `json:"failed" api:"required"`
	// Not yet queued for sending.
	Pending int64 `json:"pending" api:"required"`
	// Percentage complete (0-100).
	PercentComplete float64 `json:"percentComplete" api:"required"`
	// Currently being sent.
	Sending int64 `json:"sending" api:"required"`
	// Skipped (broadcast cancelled).
	Skipped int64 `json:"skipped" api:"required"`
	// Current status of the broadcast.
	//
	// Any of "draft", "pending_review", "approved", "rejected", "escalated",
	// "rejected_final", "scheduled", "sending", "paused", "completed", "cancelled",
	// "failed".
	Status BroadcastStatus `json:"status" api:"required"`
	// Total contacts in broadcast.
	Total int64 `json:"total" api:"required"`
	// Actual cost so far in USD.
	ActualCost            float64   `json:"actualCost" api:"nullable"`
	EstimatedCompletionAt time.Time `json:"estimatedCompletionAt" format:"date-time"`
	// Estimated total cost in USD.
	EstimatedCost float64 `json:"estimatedCost" api:"nullable"`
	// Amount reserved from balance in USD.
	ReservedAmount float64   `json:"reservedAmount" api:"nullable"`
	StartedAt      time.Time `json:"startedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BroadcastID           respjson.Field
		Delivered             respjson.Field
		Failed                respjson.Field
		Pending               respjson.Field
		PercentComplete       respjson.Field
		Sending               respjson.Field
		Skipped               respjson.Field
		Status                respjson.Field
		Total                 respjson.Field
		ActualCost            respjson.Field
		EstimatedCompletionAt respjson.Field
		EstimatedCost         respjson.Field
		ReservedAmount        respjson.Field
		StartedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastProgress) RawJSON() string { return r.JSON.raw }
func (r *BroadcastProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the broadcast.
type BroadcastStatus string

const (
	BroadcastStatusDraft         BroadcastStatus = "draft"
	BroadcastStatusPendingReview BroadcastStatus = "pending_review"
	BroadcastStatusApproved      BroadcastStatus = "approved"
	BroadcastStatusRejected      BroadcastStatus = "rejected"
	BroadcastStatusEscalated     BroadcastStatus = "escalated"
	BroadcastStatusRejectedFinal BroadcastStatus = "rejected_final"
	BroadcastStatusScheduled     BroadcastStatus = "scheduled"
	BroadcastStatusSending       BroadcastStatus = "sending"
	BroadcastStatusPaused        BroadcastStatus = "paused"
	BroadcastStatusCompleted     BroadcastStatus = "completed"
	BroadcastStatusCancelled     BroadcastStatus = "cancelled"
	BroadcastStatusFailed        BroadcastStatus = "failed"
)

type BroadcastNewResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastGetResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastUpdateResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastCancelResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastRescheduleResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastRescheduleResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastRescheduleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastSendResponse struct {
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastSendResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastNewParams struct {
	// Broadcast delivery channel. Use 'smart' for per-contact intelligent routing.
	//
	// Any of "smart", "sms", "sms_oneway", "whatsapp", "telegram", "email",
	// "instagram", "voice".
	Channel BroadcastChannel `json:"channel,omitzero" api:"required"`
	// Name of the broadcast campaign.
	Name string `json:"name" api:"required"`
	// HTML body for email broadcasts.
	EmailHTMLBody param.Opt[string] `json:"emailHtmlBody,omitzero"`
	// Email subject line. Required for email broadcasts.
	EmailSubject param.Opt[string] `json:"emailSubject,omitzero"`
	// Idempotency key to prevent duplicate broadcasts.
	IdempotencyKey param.Opt[string] `json:"idempotencyKey,omitzero"`
	// Schedule the broadcast for future delivery.
	ScheduledAt param.Opt[time.Time] `json:"scheduledAt,omitzero" format:"date-time"`
	// Sender profile ID. Uses default sender if omitted.
	SenderID param.Opt[string] `json:"senderId,omitzero"`
	// Text content or caption. Supports template variables: {{name}}, {{1}}, etc.
	Text param.Opt[string] `json:"text,omitzero"`
	// Content for non-text broadcast message types.
	Content BroadcastContentParam `json:"content,omitzero"`
	// Type of message for broadcast.
	//
	// Any of "text", "image", "video", "audio", "document", "template".
	MessageType BroadcastMessageType `json:"messageType,omitzero"`
	Metadata    map[string]string    `json:"metadata,omitzero"`
	paramObj
}

func (r BroadcastNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastUpdateParams struct {
	EmailHTMLBody param.Opt[string] `json:"emailHtmlBody,omitzero"`
	EmailSubject  param.Opt[string] `json:"emailSubject,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	Text          param.Opt[string] `json:"text,omitzero"`
	// Content for non-text broadcast message types.
	Content  BroadcastContentParam `json:"content,omitzero"`
	Metadata map[string]string     `json:"metadata,omitzero"`
	paramObj
}

func (r BroadcastUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Current status of the broadcast.
	//
	// Any of "draft", "pending_review", "approved", "rejected", "escalated",
	// "rejected_final", "scheduled", "sending", "paused", "completed", "cancelled",
	// "failed".
	Status BroadcastStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastListParams]'s query parameters as `url.Values`.
func (r BroadcastListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastRescheduleParams struct {
	// New scheduled time for the broadcast.
	ScheduledAt time.Time `json:"scheduledAt" api:"required" format:"date-time"`
	paramObj
}

func (r BroadcastRescheduleParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastRescheduleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastRescheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastSendParams struct {
	// Schedule for future delivery. Omit to send immediately.
	ScheduledAt param.Opt[time.Time] `json:"scheduledAt,omitzero" format:"date-time"`
	paramObj
}

func (r BroadcastSendParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
