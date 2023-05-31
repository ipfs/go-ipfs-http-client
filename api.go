// Deprecated: This repo has moved inside Kubo in order to make it easier to
// keep kubo and the http-client in sync, please use github.com/ipfs/kubo/client/rpc instead.
package httpapi

import (
	"context"
	"net/http"

	"github.com/ipfs/kubo/client/rpc"
	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use [rpc.DefaultPathName] instead.
const DefaultPathName = rpc.DefaultPathName

// Deprecated: use [rpc.DefaultPathRoot] instead.
const DefaultPathRoot = rpc.DefaultPathRoot

// Deprecated: use [rpc.DefaultApiFile] instead.
const DefaultApiFile = rpc.DefaultApiFile

// Deprecated: use [rpc.EnvDir] instead.
const EnvDir = rpc.EnvDir

// Deprecated: use [rpc.ErrApiNotFound] instead.
var ErrApiNotFound = rpc.ErrApiNotFound

// Deprecated: use [rpc.HttpApi] instead.
type HttpApi = rpc.HttpApi

// Deprecated: use [rpc.BlockAPI] instead.
type BlockAPI = rpc.BlockAPI

// Deprecated: use [rpc.HttpDagServ] instead.
type HttpDagServ = rpc.HttpDagServ

// Deprecated: use [rpc.DhtAPI] instead.
type DhtAPI = rpc.DhtAPI

// Deprecated: use [rpc.KeyAPI] instead.
type KeyAPI = rpc.KeyAPI

// Deprecated: use [rpc.NameAPI] instead.
type NameAPI = rpc.NameAPI

// Deprecated: use [rpc.ObjectAPI] instead.
type ObjectAPI = rpc.ObjectAPI

// Deprecated: use [rpc.PinAPI] instead.
type PinAPI = rpc.PinAPI

// Deprecated: use [rpc.PubsubAPI] instead.
type PubsubAPI = rpc.PubsubAPI

// Deprecated: use [rpc.RoutingAPI] instead.
type RoutingAPI = rpc.RoutingAPI

// Deprecated: use [rpc.SwarmAPI] instead.
type SwarmAPI = rpc.SwarmAPI

// Deprecated: use [rpc.UnixfsAPI] instead.
type UnixfsAPI = rpc.UnixfsAPI

// Deprecated: use [rpc.NewLocalApi] instead.
func NewLocalApi() (*HttpApi, error) {
	return rpc.NewLocalApi()
}

// Deprecated: use [rpc.NewPathApi] instead.
func NewPathApi(ipfspath string) (*HttpApi, error) {
	return rpc.NewPathApi(ipfspath)
}

// Deprecated: use [rpc.ApiAddr] instead.
func ApiAddr(ipfspath string) (ma.Multiaddr, error) {
	return rpc.ApiAddr(ipfspath)
}

// Deprecated: use [rpc.NewApi] instead.
func NewApi(a ma.Multiaddr) (*HttpApi, error) {
	return rpc.NewApi(a)
}

// Deprecated: use [rpc.NewApiWithClient] instead.
func NewApiWithClient(a ma.Multiaddr, c *http.Client) (*HttpApi, error) {
	return rpc.NewApiWithClient(a, c)
}

// Deprecated: use [rpc.NewURLApiWithClient] instead.
func NewURLApiWithClient(url string, c *http.Client) (*HttpApi, error) {
	return rpc.NewURLApiWithClient(url, c)
}

// Deprecated: use [rpc.Request] instead.
type Request = rpc.Request

// Deprecated: use [rpc.NewRequest] instead.
func NewRequest(ctx context.Context, url, command string, args ...string) *Request {
	return rpc.NewRequest(ctx, url, command, args...)
}

// Deprecated: use [rpc.RequestBuilder] instead.
type RequestBuilder = rpc.RequestBuilder

// Deprecated: use [rpc.Error] instead.
type Error = rpc.Error

// Deprecated: use [rpc.Response] instead.
type Response = rpc.Response
