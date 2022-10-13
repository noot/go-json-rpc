package rpc

import (
	"encoding/json"
	"fmt"
)

// DefaultJSONRPCVersion ...
const DefaultJSONRPCVersion = "2.0"

// Request represents a JSON-RPC request
type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      uint64          `json:"id"`
}

// Response represents a JSON-RPC response
type Response struct {
	Version string           `json:"jsonrpc"`
	Result  json.RawMessage  `json:"result"`
	Error   *Error           `json:"error"`
	ID      *json.RawMessage `json:"id"`
}

// ErrCode used for RPC error codes
type ErrCode int

// Error contains an error message and the error code for a error, as well as
// any additional data
type Error struct {
	Message   string                 `json:"message"`
	ErrorCode ErrCode                `json:"code"`
	Data      map[string]interface{} `json:"data"`
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("message=%s; code=%d; data=%v", e.Message, e.ErrorCode, e.Data)
}
