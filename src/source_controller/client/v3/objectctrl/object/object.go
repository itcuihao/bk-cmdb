package object

import (
	"configcenter/src/source_controller/client/types"
)

// Getter the Object getter interface
type Getter interface {
	Object() Interface
}

// Interface the Object api
type Interface interface {
	CreateObject(result interface{}) error
}

// Object the Object interface implement
type Object struct {
	ctx *types.Context
}

// NewObject returns the object client
func NewObject(ctx *types.Context) *Object {
	return &Object{
		ctx: ctx,
	}
}

// CreateObject create the object
func (cli *Object) CreateObject(result interface{}) error {
	return cli.ctx.Call(&result)
}
