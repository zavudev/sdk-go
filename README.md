# undefined

Developer-friendly & type-safe Go SDK specifically catered to leverage *undefined* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=undefined&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/crubing/zavu). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Zavu Messaging API: Unified multi-channel messaging API for Zavu.

Supported channels:
- **SMS**: Simple text messages
- **WhatsApp**: Rich messaging with media, buttons, lists, and templates

Design goals:
- Simple `send()` entrypoint for developers
- Project-level authentication via Bearer token
- Support for all WhatsApp message types (text, image, video, audio, document, sticker, location, contact, buttons, list, reaction, template)
- If a non-text message type is sent, WhatsApp channel is used automatically
- 24-hour WhatsApp conversation window enforcement
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [undefined](#undefined)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/zavu-dev/zavu-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"))
	if err != nil {
		log.Fatal(err)
	}
	if res.MessageResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      |
| ------------ | ---- | ----------- |
| `BearerAuth` | http | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"))
	if err != nil {
		log.Fatal(err)
	}
	if res.MessageResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Zavu SDK](docs/sdks/zavu/README.md)

* [SendMessage](docs/sdks/zavu/README.md#sendmessage) - Send a message
* [ListMessages](docs/sdks/zavu/README.md#listmessages) - List messages
* [GetMessage](docs/sdks/zavu/README.md#getmessage) - Get message by ID
* [SendReaction](docs/sdks/zavu/README.md#sendreaction) - Send reaction to message
* [ListTemplates](docs/sdks/zavu/README.md#listtemplates) - List templates
* [CreateTemplate](docs/sdks/zavu/README.md#createtemplate) - Create template
* [GetTemplate](docs/sdks/zavu/README.md#gettemplate) - Get template
* [DeleteTemplate](docs/sdks/zavu/README.md#deletetemplate) - Delete template
* [ListSenders](docs/sdks/zavu/README.md#listsenders) - List senders
* [CreateSender](docs/sdks/zavu/README.md#createsender) - Create sender
* [GetSender](docs/sdks/zavu/README.md#getsender) - Get sender
* [UpdateSender](docs/sdks/zavu/README.md#updatesender) - Update sender
* [DeleteSender](docs/sdks/zavu/README.md#deletesender) - Delete sender
* [ListContacts](docs/sdks/zavu/README.md#listcontacts) - List contacts
* [GetContact](docs/sdks/zavu/README.md#getcontact) - Get contact
* [UpdateContact](docs/sdks/zavu/README.md#updatecontact) - Update contact
* [GetContactByPhone](docs/sdks/zavu/README.md#getcontactbyphone) - Get contact by phone number
* [IntrospectPhone](docs/sdks/zavu/README.md#introspectphone) - Introspect phone number

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/components"
	"github.com/zavu-dev/zavu-go/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"), operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.MessageResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/components"
	"github.com/zavu-dev/zavu-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"))
	if err != nil {
		log.Fatal(err)
	}
	if res.MessageResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `SendMessage` function may return the following errors:

| Error Type         | Status Code             | Content Type     |
| ------------------ | ----------------------- | ---------------- |
| apierrors.Error    | 400, 401, 404, 409, 429 | application/json |
| apierrors.Error    | 500                     | application/json |
| apierrors.APIError | 4XX, 5XX                | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/apierrors"
	"github.com/zavu-dev/zavu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"))
	if err != nil {

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := zavu.New(
		zavu.WithServerURL("https://api.zavu.dev"),
		zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.SendMessage(ctx, components.MessageRequest{
		To:   "+56912345678",
		Text: zavu.Pointer("Your verification code is 123456"),
	}, zavu.Pointer("sender_12345"))
	if err != nil {
		log.Fatal(err)
	}
	if res.MessageResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/zavu-dev/zavu-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = zavu.New(zavu.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=undefined&utm_campaign=go)
