# Zavu SDK

## Overview

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


### Available Operations

* [SendMessage](#sendmessage) - Send a message
* [ListMessages](#listmessages) - List messages
* [GetMessage](#getmessage) - Get message by ID
* [SendReaction](#sendreaction) - Send reaction to message
* [ListTemplates](#listtemplates) - List templates
* [CreateTemplate](#createtemplate) - Create template
* [GetTemplate](#gettemplate) - Get template
* [DeleteTemplate](#deletetemplate) - Delete template
* [ListSenders](#listsenders) - List senders
* [CreateSender](#createsender) - Create sender
* [GetSender](#getsender) - Get sender
* [UpdateSender](#updatesender) - Update sender
* [DeleteSender](#deletesender) - Delete sender
* [ListContacts](#listcontacts) - List contacts
* [GetContact](#getcontact) - Get contact
* [UpdateContact](#updatecontact) - Update contact
* [GetContactByPhone](#getcontactbyphone) - Get contact by phone number
* [IntrospectPhone](#introspectphone) - Introspect phone number

## SendMessage

Send a message to a recipient via SMS or WhatsApp.

**Channel selection:**
- If `channel` is omitted and `messageType` is `text`, defaults to SMS
- If `messageType` is anything other than `text`, WhatsApp is used automatically

**WhatsApp 24-hour window:**
- Free-form messages (non-template) require an open 24h window
- Window opens when the user messages you first
- Use template messages to initiate conversations outside the window

### Example Usage

<!-- UsageSnippet language="go" operationID="sendMessage" method="post" path="/v1/messages" -->
```go
package main

import(
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
        To: "+56912345678",
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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `body`                                                                             | [components.MessageRequest](../../models/components/messagerequest.md)             | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `zavuSender`                                                                       | **string*                                                                          | :heavy_minus_sign:                                                                 | Optional sender profile ID. If omitted, the project's default sender will be used. | sender_12345                                                                       |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.SendMessageResponse](../../models/operations/sendmessageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.Error         | 400, 401, 404, 409, 429 | application/json        |
| apierrors.Error         | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListMessages

List messages previously sent by this project.

### Example Usage

<!-- UsageSnippet language="go" operationID="listMessages" method="get" path="/v1/messages" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"github.com/zavu-dev/zavu-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ListMessages(ctx, operations.ListMessagesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListMessagesRequest](../../models/operations/listmessagesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListMessagesResponse](../../models/operations/listmessagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetMessage

Get message by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getMessage" method="get" path="/v1/messages/{messageId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GetMessage(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `messageID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMessageResponse](../../models/operations/getmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SendReaction

Send an emoji reaction to an existing WhatsApp message. Reactions are only supported for WhatsApp messages.

### Example Usage

<!-- UsageSnippet language="go" operationID="sendReaction" method="post" path="/v1/messages/{messageId}/reactions" -->
```go
package main

import(
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

    res, err := s.SendReaction(ctx, "<id>", components.ReactionRequest{
        Emoji: "👍",
    }, zavu.Pointer("sender_12345"))
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `messageID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `body`                                                                             | [components.ReactionRequest](../../models/components/reactionrequest.md)           | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `zavuSender`                                                                       | **string*                                                                          | :heavy_minus_sign:                                                                 | Optional sender profile ID. If omitted, the project's default sender will be used. | sender_12345                                                                       |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.SendReactionResponse](../../models/operations/sendreactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404      | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTemplates

List WhatsApp message templates for this project.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTemplates" method="get" path="/v1/templates" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ListTemplates(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListTemplatesResponse](../../models/operations/listtemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateTemplate

Create a WhatsApp message template. Note: Templates must be approved by Meta before use.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTemplate" method="post" path="/v1/templates" -->
```go
package main

import(
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

    res, err := s.CreateTemplate(ctx, components.TemplateCreateRequest{
        Name: "order_confirmation",
        Body: "Hi {{1}}, your order {{2}} has been confirmed and will ship within 24 hours.",
        WhatsappCategory: components.WhatsAppCategoryUtility.ToPointer(),
        Variables: []string{
            "customer_name",
            "order_id",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Template != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.TemplateCreateRequest](../../models/components/templatecreaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateTemplateResponse](../../models/operations/createtemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetTemplate

Get template

### Example Usage

<!-- UsageSnippet language="go" operationID="getTemplate" method="get" path="/v1/templates/{templateId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GetTemplate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Template != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `templateID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTemplateResponse](../../models/operations/gettemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteTemplate

Delete template

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTemplate" method="delete" path="/v1/templates/{templateId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.DeleteTemplate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `templateID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTemplateResponse](../../models/operations/deletetemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListSenders

List senders

### Example Usage

<!-- UsageSnippet language="go" operationID="listSenders" method="get" path="/v1/senders" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ListSenders(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListSendersResponse](../../models/operations/listsendersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSender

Create sender

### Example Usage

<!-- UsageSnippet language="go" operationID="createSender" method="post" path="/v1/senders" -->
```go
package main

import(
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

    res, err := s.CreateSender(ctx, components.SenderCreateRequest{
        Name: "<value>",
        PhoneNumber: "1-697-351-3400 x33934",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sender != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.SenderCreateRequest](../../models/components/sendercreaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.CreateSenderResponse](../../models/operations/createsenderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSender

Get sender

### Example Usage

<!-- UsageSnippet language="go" operationID="getSender" method="get" path="/v1/senders/{senderId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GetSender(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Sender != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `senderID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSenderResponse](../../models/operations/getsenderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateSender

Update sender

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSender" method="patch" path="/v1/senders/{senderId}" -->
```go
package main

import(
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

    res, err := s.UpdateSender(ctx, "<id>", components.SenderUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Sender != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `senderID`                                                                       | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `body`                                                                           | [components.SenderUpdateRequest](../../models/components/senderupdaterequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.UpdateSenderResponse](../../models/operations/updatesenderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404      | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteSender

Delete sender

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSender" method="delete" path="/v1/senders/{senderId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.DeleteSender(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `senderID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteSenderResponse](../../models/operations/deletesenderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404      | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListContacts

List contacts

### Example Usage

<!-- UsageSnippet language="go" operationID="listContacts" method="get" path="/v1/contacts" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.ListContacts(ctx, nil, zavu.Pointer[int64](50), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `phoneNumber`                                            | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListContactsResponse](../../models/operations/listcontactsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetContact

Get contact

### Example Usage

<!-- UsageSnippet language="go" operationID="getContact" method="get" path="/v1/contacts/{contactId}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GetContact(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Contact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `contactID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContactResponse](../../models/operations/getcontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateContact

Update contact

### Example Usage

<!-- UsageSnippet language="go" operationID="updateContact" method="patch" path="/v1/contacts/{contactId}" -->
```go
package main

import(
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

    res, err := s.UpdateContact(ctx, "<id>", components.ContactUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Contact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `contactID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | N/A                                                                                |
| `body`                                                                             | [components.ContactUpdateRequest](../../models/components/contactupdaterequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateContactResponse](../../models/operations/updatecontactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401, 404      | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetContactByPhone

Get contact by phone number

### Example Usage

<!-- UsageSnippet language="go" operationID="getContactByPhone" method="get" path="/v1/contacts/phone/{phoneNumber}" -->
```go
package main

import(
	"context"
	zavu "github.com/zavu-dev/zavu-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := zavu.New(
        zavu.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    res, err := s.GetContactByPhone(ctx, "397-335-4175 x077")
    if err != nil {
        log.Fatal(err)
    }
    if res.Contact != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `phoneNumber`                                            | *string*                                                 | :heavy_check_mark:                                       | E.164 phone number.                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetContactByPhoneResponse](../../models/operations/getcontactbyphoneresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401, 404           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## IntrospectPhone

Validate a phone number and check if a WhatsApp conversation window is open.

### Example Usage

<!-- UsageSnippet language="go" operationID="introspectPhone" method="post" path="/v1/introspect/phone" -->
```go
package main

import(
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

    res, err := s.IntrospectPhone(ctx, components.PhoneIntrospectionRequest{
        PhoneNumber: "+56912345678",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PhoneIntrospectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.PhoneIntrospectionRequest](../../models/components/phoneintrospectionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.IntrospectPhoneResponse](../../models/operations/introspectphoneresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 400, 401           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |