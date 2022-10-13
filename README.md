# go-json-rpc

Very small lib for calling a JSON-RPC server. I kept copy pasting this code so I just put it in its own repo :)

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	rpc "github.com/noot/go-json-rpc"
)

type MyRequestType struct {
	SomeParam string
}

type MyResponseType struct {
	SomeReturnValue string
}

func main() {
	resp, err := callSomething("http://localhost:8545", "this is an example")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func callSomething(endpoint, someParam string) (*MyResponseType, error) {
	const method = "server_someMethod"

	req := &MyRequestType{
		SomeParam: someParam,
	}

	params, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.PostRPC(endpoint, method, string(params))
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}

	var res *MyResponseType
	if err = json.Unmarshal(resp.Result, &res); err != nil {
		return nil, err
	}

	return res, nil
}
```