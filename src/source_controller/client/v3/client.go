package v3

import (
	"configcenter/src/source_controller/client/types"
	"configcenter/src/source_controller/client/v3/objectctrl"
)

// Getter define V3 client getter
type Getter interface {
	V3() Interface
}

// Interface define v3 api
type Interface interface {
	objectctrl.Getter
}

// Client implements the V3 Interface
type Client struct {
	ctx *types.Context
}

// NewClient returns the new V3 Client
func NewClient(ctx *types.Context) *Client {
	return &Client{ctx: ctx}
}

// Object returns the objectcontrller api
func (cli *Client) Object() objectctrl.Interface {
	return objectctrl.NewObjectCtrl(cli.ctx)
}
