package objectctrl

import (
	"configcenter/src/source_controller/client/types"
	"configcenter/src/source_controller/client/v3/objectctrl/object"
)

// Getter the Object getter interface
type Getter interface {
	Object() Interface
}

// Interface the objectcontroller api
type Interface interface {
	object.Getter
}

// ObjectCtrl the Object interface implement
type ObjectCtrl struct {
	ctx *types.Context
}

// NewObjectCtrl returns objectcontroller client
func NewObjectCtrl(ctx *types.Context) *ObjectCtrl {
	ctx.Module = "objectctrl"
	return &ObjectCtrl{
		ctx: ctx,
	}
}

// Object returns object client
func (cli *ObjectCtrl) Object() object.Interface {
	return object.NewObject(cli.ctx)
}
