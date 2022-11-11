# OpenTracing middleware for the cHTTP client



Adds an OpenTracing logs and headers to the request.

### Usage example

```go
package clients

import (
	"context"
	"net/http"

	"github.com/spyzhov/chttp"
	middleware "github.com/spyzhov/chttp-middleware-opentracing"
)

type Client struct {
	Client *chttp.JSONClient
}

func New() *Client {
	client := chttp.NewJSON(nil)
	client.With(middleware.Opentracing(func(request *http.Request) string {
		return request.URL.Path
	}))

	return &Client{
		Client: client,
	}
}

func (c *Client) GetCount(ctx context.Context) (result int, err error) {
    return result, c.Client.GET(ctx, "https://example.com/get/count", nil, &result)
}

```

# License

MIT licensed. See the [LICENSE](LICENSE) file for details.
