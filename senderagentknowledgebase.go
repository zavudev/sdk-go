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

	"github.com/stainless-sdks/zavudev-go/internal/apijson"
	"github.com/stainless-sdks/zavudev-go/internal/apiquery"
	"github.com/stainless-sdks/zavudev-go/internal/requestconfig"
	"github.com/stainless-sdks/zavudev-go/option"
	"github.com/stainless-sdks/zavudev-go/packages/pagination"
	"github.com/stainless-sdks/zavudev-go/packages/param"
	"github.com/stainless-sdks/zavudev-go/packages/respjson"
)

// SenderAgentKnowledgeBaseService contains methods and other services that help
// with interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentKnowledgeBaseService] method instead.
type SenderAgentKnowledgeBaseService struct {
	Options   []option.RequestOption
	Documents SenderAgentKnowledgeBaseDocumentService
}

// NewSenderAgentKnowledgeBaseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSenderAgentKnowledgeBaseService(opts ...option.RequestOption) (r SenderAgentKnowledgeBaseService) {
	r = SenderAgentKnowledgeBaseService{}
	r.Options = opts
	r.Documents = NewSenderAgentKnowledgeBaseDocumentService(opts...)
	return
}

// Create a new knowledge base for an agent.
func (r *SenderAgentKnowledgeBaseService) New(ctx context.Context, senderID string, body SenderAgentKnowledgeBaseNewParams, opts ...option.RequestOption) (res *SenderAgentKnowledgeBaseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific knowledge base.
func (r *SenderAgentKnowledgeBaseService) Get(ctx context.Context, kbID string, query SenderAgentKnowledgeBaseGetParams, opts ...option.RequestOption) (res *SenderAgentKnowledgeBaseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s", url.PathEscape(query.SenderID), url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a knowledge base.
func (r *SenderAgentKnowledgeBaseService) Update(ctx context.Context, kbID string, params SenderAgentKnowledgeBaseUpdateParams, opts ...option.RequestOption) (res *SenderAgentKnowledgeBaseUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s", url.PathEscape(params.SenderID), url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List knowledge bases for an agent.
func (r *SenderAgentKnowledgeBaseService) List(ctx context.Context, senderID string, query SenderAgentKnowledgeBaseListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentKnowledgeBase], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases", url.PathEscape(senderID))
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

// List knowledge bases for an agent.
func (r *SenderAgentKnowledgeBaseService) ListAutoPaging(ctx context.Context, senderID string, query SenderAgentKnowledgeBaseListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentKnowledgeBase] {
	return pagination.NewCursorAutoPager(r.List(ctx, senderID, query, opts...))
}

// Delete a knowledge base and all its documents.
func (r *SenderAgentKnowledgeBaseService) Delete(ctx context.Context, kbID string, body SenderAgentKnowledgeBaseDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s", url.PathEscape(body.SenderID), url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AgentDocument struct {
	ID string `json:"id" api:"required"`
	// Number of chunks created from this document.
	ChunkCount int64 `json:"chunkCount" api:"required"`
	// Length of the document content in characters.
	ContentLength int64     `json:"contentLength" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the document has been processed for RAG.
	IsProcessed     bool      `json:"isProcessed" api:"required"`
	KnowledgeBaseID string    `json:"knowledgeBaseId" api:"required"`
	Title           string    `json:"title" api:"required"`
	UpdatedAt       time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChunkCount      respjson.Field
		ContentLength   respjson.Field
		CreatedAt       respjson.Field
		IsProcessed     respjson.Field
		KnowledgeBaseID respjson.Field
		Title           respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDocument) RawJSON() string { return r.JSON.raw }
func (r *AgentDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentKnowledgeBase struct {
	ID            string    `json:"id" api:"required"`
	AgentID       string    `json:"agentId" api:"required"`
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	DocumentCount int64     `json:"documentCount" api:"required"`
	Name          string    `json:"name" api:"required"`
	TotalChunks   int64     `json:"totalChunks" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	Description   string    `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AgentID       respjson.Field
		CreatedAt     respjson.Field
		DocumentCount respjson.Field
		Name          respjson.Field
		TotalChunks   respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentKnowledgeBase) RawJSON() string { return r.JSON.raw }
func (r *AgentKnowledgeBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseNewResponse struct {
	KnowledgeBase AgentKnowledgeBase `json:"knowledgeBase" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBase respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentKnowledgeBaseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentKnowledgeBaseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseGetResponse struct {
	KnowledgeBase AgentKnowledgeBase `json:"knowledgeBase" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBase respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentKnowledgeBaseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentKnowledgeBaseGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseUpdateResponse struct {
	KnowledgeBase AgentKnowledgeBase `json:"knowledgeBase" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KnowledgeBase respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentKnowledgeBaseUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentKnowledgeBaseUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseNewParams struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r SenderAgentKnowledgeBaseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentKnowledgeBaseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentKnowledgeBaseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseGetParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}

type SenderAgentKnowledgeBaseUpdateParams struct {
	SenderID    string            `path:"senderId" api:"required" json:"-"`
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r SenderAgentKnowledgeBaseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentKnowledgeBaseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentKnowledgeBaseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentKnowledgeBaseListParams]'s query parameters as
// `url.Values`.
func (r SenderAgentKnowledgeBaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderAgentKnowledgeBaseDeleteParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	paramObj
}
