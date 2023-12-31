package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	LocalnetRPCEndpoint = "http://localhost:8899"
	DevnetRPCEndpoint   = "https://api.devnet.solana.com"
	TestnetRPCEndpoint  = "https://api.testnet.solana.com"
	MainnetRPCEndpoint  = "https://api.mainnet-beta.solana.com"
)

// Commitment describes how finalized a block is at that point in time
type Commitment string

const (
	CommitmentFinalized Commitment = "finalized"
	CommitmentConfirmed Commitment = "confirmed"
	CommitmentProcessed Commitment = "processed"
)

// ErrorResponse is a error rpc response
type ErrorResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type Context struct {
	Slot uint64 `json:"slot"`
}

// GeneralResponse is a general rpc response
type GeneralResponse struct {
	JsonRPC string         `json:"jsonrpc"`
	ID      uint64         `json:"id"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type RpcClient struct {
	endpoint   string
	httpClient *http.Client
	debug      bool
}

func NewRpcClient(endpoint string) RpcClient { return New(WithEndpoint(endpoint)) }

// New applies the given options to the rpc client being created. if no options
// is passed, it defaults to a bare bone http client and solana mainnet
func New(opts ...Option) RpcClient {

	client := &RpcClient{}

	setDefaultOptions(client)

	for _, opt := range opts {
		opt(client)
	}

	return *client
}

func (c *RpcClient) SetDebug(debug bool) {
	c.debug = debug
}

// Call will return body of response. if http code beyond 200~300, the error also returns.
func (c *RpcClient) Call(ctx context.Context, params ...interface{}) ([]byte, error) {
	// prepare payload
	j, err := preparePayload(params)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare payload, err: %v", err)
	}

	// prepare request
	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(j))
	if err != nil {
		return nil, fmt.Errorf("failed to do http.NewRequestWithContext, err: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// do request
	res, err := c.httpClient.Do(req)
	if err != nil {
		if c.debug {
			{
				content, _ := httputil.DumpRequest(req, true)
				log.Println("request", string(content))
			}
			{
				content, _ := httputil.DumpResponse(res, true)
				log.Println("response", string(content))
			}
		}
		return nil, fmt.Errorf("failed to do request, err: %v", err)
	}
	defer res.Body.Close()

	// parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if c.debug {
			{
				content, _ := httputil.DumpRequest(req, true)
				log.Println("request", string(content))
			}
			{
				content, _ := httputil.DumpResponse(res, true)
				log.Println("response", string(content))
			}
		}
		return nil, fmt.Errorf("failed to read body, err: %v", err)
	}

	// check response code
	if res.StatusCode < 200 || res.StatusCode > 300 {
		if c.debug {
			{
				content, _ := httputil.DumpRequest(req, true)
				log.Println("request", string(content))
			}
			{
				content, _ := httputil.DumpResponse(res, true)
				log.Println("response", string(content))
			}
		}
		return body, fmt.Errorf("get status code: %v", res.StatusCode)
	}

	return body, nil
}

type jsonRpcRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Id      uint64        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
}

func preparePayload(params []interface{}) ([]byte, error) {
	// prepare payload
	j, err := json.Marshal(jsonRpcRequest{
		JsonRpc: "2.0",
		Id:      1,
		Method:  params[0].(string),
		Params:  params[1:],
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *RpcClient) processRpcCall(body []byte, rpcErr error, res interface{}) error {
	if rpcErr != nil {
		return fmt.Errorf("rpc: call error, err: %v, body: %v", rpcErr, string(body))
	}
	err := json.Unmarshal(body, &res)
	if err != nil {
		return fmt.Errorf("rpc: failed to json decode body, err: %v", err)
	}
	return nil
}
