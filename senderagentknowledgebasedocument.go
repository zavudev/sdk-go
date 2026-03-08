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
	"github.com/zavudev/sdk-go/internal/apiquery"
	"github.com/zavudev/sdk-go/internal/requestconfig"
	"github.com/zavudev/sdk-go/option"
	"github.com/zavudev/sdk-go/packages/pagination"
	"github.com/zavudev/sdk-go/packages/param"
	"github.com/zavudev/sdk-go/packages/respjson"
)

// SenderAgentKnowledgeBaseDocumentService contains methods and other services that
// help with interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderAgentKnowledgeBaseDocumentService] method instead.
type SenderAgentKnowledgeBaseDocumentService struct {
	Options []option.RequestOption
}

// NewSenderAgentKnowledgeBaseDocumentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSenderAgentKnowledgeBaseDocumentService(opts ...option.RequestOption) (r SenderAgentKnowledgeBaseDocumentService) {
	r = SenderAgentKnowledgeBaseDocumentService{}
	r.Options = opts
	return
}

// Add a document to a knowledge base. The document will be automatically processed
// for RAG.
func (r *SenderAgentKnowledgeBaseDocumentService) New(ctx context.Context, kbID string, params SenderAgentKnowledgeBaseDocumentNewParams, opts ...option.RequestOption) (res *SenderAgentKnowledgeBaseDocumentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s/documents", url.PathEscape(params.SenderID), url.PathEscape(kbID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List documents in a knowledge base.
func (r *SenderAgentKnowledgeBaseDocumentService) List(ctx context.Context, kbID string, params SenderAgentKnowledgeBaseDocumentListParams, opts ...option.RequestOption) (res *pagination.Cursor[AgentDocument], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if kbID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s/documents", url.PathEscape(params.SenderID), url.PathEscape(kbID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List documents in a knowledge base.
func (r *SenderAgentKnowledgeBaseDocumentService) ListAutoPaging(ctx context.Context, kbID string, params SenderAgentKnowledgeBaseDocumentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AgentDocument] {
	return pagination.NewCursorAutoPager(r.List(ctx, kbID, params, opts...))
}

// Delete a document from a knowledge base.
func (r *SenderAgentKnowledgeBaseDocumentService) Delete(ctx context.Context, docID string, body SenderAgentKnowledgeBaseDocumentDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SenderID == "" {
		err = errors.New("missing required senderId parameter")
		return
	}
	if body.KBID == "" {
		err = errors.New("missing required kbId parameter")
		return
	}
	if docID == "" {
		err = errors.New("missing required docId parameter")
		return
	}
	path := fmt.Sprintf("v1/senders/%s/agent/knowledge-bases/%s/documents/%s", url.PathEscape(body.SenderID), url.PathEscape(body.KBID), url.PathEscape(docID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SenderAgentKnowledgeBaseDocumentNewResponse struct {
	Document AgentDocument `json:"document" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderAgentKnowledgeBaseDocumentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderAgentKnowledgeBaseDocumentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseDocumentNewParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	Content  string `json:"content" api:"required"`
	Title    string `json:"title" api:"required"`
	paramObj
}

func (r SenderAgentKnowledgeBaseDocumentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderAgentKnowledgeBaseDocumentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderAgentKnowledgeBaseDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderAgentKnowledgeBaseDocumentListParams struct {
	SenderID string            `path:"senderId" api:"required" json:"-"`
	Cursor   param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SenderAgentKnowledgeBaseDocumentListParams]'s query
// parameters as `url.Values`.
func (r SenderAgentKnowledgeBaseDocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SenderAgentKnowledgeBaseDocumentDeleteParams struct {
	SenderID string `path:"senderId" api:"required" json:"-"`
	KBID     string `path:"kbId" api:"required" json:"-"`
	paramObj
}
