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

// SenderTelegramService contains methods and other services that help with
// interacting with the zavudev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSenderTelegramService] method instead.
type SenderTelegramService struct {
	Options []option.RequestOption
}

// NewSenderTelegramService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSenderTelegramService(opts ...option.RequestOption) (r SenderTelegramService) {
	r = SenderTelegramService{}
	r.Options = opts
	return
}

// Connect a Telegram bot to a sender. Provide the bot token from @BotFather; Zavu
// validates it, registers the webhook, and routes the sender's Telegram messages
// through it.
func (r *SenderTelegramService) Connect(ctx context.Context, senderID string, body SenderTelegramConnectParams, opts ...option.RequestOption) (res *SenderTelegramConnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/senders/%s/telegram", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Disconnect Telegram from a sender and remove the webhook.
func (r *SenderTelegramService) Disconnect(ctx context.Context, senderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if senderID == "" {
		err = errors.New("missing required senderId parameter")
		return err
	}
	path := fmt.Sprintf("v1/senders/%s/telegram", url.PathEscape(senderID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type SenderTelegramConnectResponse struct {
	Telegram SenderTelegramConnectResponseTelegram `json:"telegram" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Telegram    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderTelegramConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *SenderTelegramConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderTelegramConnectResponseTelegram struct {
	Connected   bool   `json:"connected" api:"required"`
	BotID       string `json:"botId"`
	BotUsername string `json:"botUsername"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connected   respjson.Field
		BotID       respjson.Field
		BotUsername respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SenderTelegramConnectResponseTelegram) RawJSON() string { return r.JSON.raw }
func (r *SenderTelegramConnectResponseTelegram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SenderTelegramConnectParams struct {
	// Bot token from @BotFather.
	BotToken string `json:"botToken" api:"required"`
	paramObj
}

func (r SenderTelegramConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow SenderTelegramConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SenderTelegramConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
