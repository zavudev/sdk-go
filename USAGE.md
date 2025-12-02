<!-- Start SDK Example Usage [usage] -->
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