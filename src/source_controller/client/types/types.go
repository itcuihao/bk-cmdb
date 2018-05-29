package types

import (
	"configcenter/src/common"
	"configcenter/src/common/errors"
	"configcenter/src/common/http/httpclient"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
)

// Context define
type Context struct {
	// Forward the forward params
	Forward *ForwardParam
	// Module the module name, it will be used to find the server address
	Module string
	// Name named the api, it's useful in tracking calling chain
	Name string
	// URL the relative url for the api. e.g. /object/v1/object
	URL string
	// Method the http request method
	Method string
	// URLParams the url will format by the params
	URLParams []interface{}
	// Body will marshal as json and use as request body
	Body interface{}

	IsTransaction   bool
	TransactionRoot *Context
	Called          []*Context

	defErr errors.DefaultCCErrorIf
	addr   ServerGetter
}

// ServerGetter define the get server interface
type ServerGetter interface {
	GetServer(servType string) (string, error)
}

// Clone returns the clone ClientContext
func (ctx *Context) Clone() *Context {
	clone := *ctx
	return &clone
}

// ForwardParam define params should parse via request
type ForwardParam struct {
	Header http.Header
}

// Call execute the call, if mock have the same key, will just return the mockdata
func (ctx *Context) Call(result interface{}) error {
	if mockingMode {
		if err := json.Unmarshal(mockData[ctx.Name], result); err != nil {
			return ctx.defErr.Error(common.CCErrCommJSONUnmarshalFailed)
		}
	}

	url := ctx.URL
	if strings.Contains(ctx.URL, "%s") {
		url = fmt.Sprintf(ctx.URL, ctx.URLParams...)
	}

	serveraddr, err := ctx.addr.GetServer(ctx.Module)
	if err != nil {
		return ctx.defErr.Error(common.CCErrCommRelyOnServerAddressFailed)
	}

	url = serveraddr + url

	body, err := json.Marshal(ctx.Body)
	if err != nil {
		return ctx.defErr.Error(common.CCErrCommJSONMarshalFailed)
	}
	out, err := httpclient.NewHttpClient().Request(url, ctx.Method, ctx.Forward.Header, body)
	if err != nil {
		return ctx.defErr.Error(common.CCErrCommHTTPDoRequestFailed)
	}

	val := gjson.ParseBytes(out)
	if val.Get(common.HTTPBKAPIErrorCode).Int() != common.CCSuccess {
		return ctx.defErr.Error(common.CCErrCommHTTPDoRequestFailed)
	}

	err = json.Unmarshal([]byte(val.Get("data").String()), result)
	if err != nil {
		return ctx.defErr.Error(common.CCErrCommJSONUnmarshalFailed)
	}
	return nil
}

var mockingMode bool

var mockData = map[string][]byte{}

// Mock set the mocking data, only support by ctx.Name now
func (ctx *Context) Mock(data interface{}) error {
	mockingMode = true
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	mockData[ctx.Name] = out
	return nil
}
