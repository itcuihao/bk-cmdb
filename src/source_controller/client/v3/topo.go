package v3

import (
	"configcenter/src/source_controller/client/types"
)

// ObjectGetter the Object getter interface
type ObjectGetter interface {
	Object() ObjectInterface
}

// ObjectInterface the Object interface
type ObjectInterface interface {
	CreateObject() error
}

func newObject(cli *Client) *Object {
	return &Object{
		cli: cli,
	}
}

// Object the Object interface implement
type Object struct {
	cli *types.Client
}
