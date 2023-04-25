# Go API client for MvcApi

API definition for MvcApi provided apis


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/mvc-labs/mvcapi-sdk-go
```

## Documentation For Authorization


### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), mvcapi.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

## Usage Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/mvc-labs/mvcapi-sdk-go"
	"os"
)

func main() {
	authCtx := context.WithValue(context.Background(), mvcapi.ContextAccessToken, "YOUR_MVCAPI_KEY")
	// testnet config
	authCtx = context.WithValue(authCtx, mvcapi.ContextServerIndex, 0)
	// mainnet config
	//authCtx = context.WithValue(authCtx, mvcapi.ContextServerIndex, 1)
	configuration := mvcapi.NewConfiguration()
	client := mvcapi.NewAPIClient(configuration)

	_, response, err := client.BlockApi.BlockInfoGet(authCtx).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockApi.BlockInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", response)
	}
	// response from `BlockInfoGet`: BlockchainInfo
	fmt.Fprintf(os.Stdout, "Response from `BlockApi.BlockInfoGet`: %v\n", response)
}

```

## Author

heqiming@mvcapi.com

