package identifiers

import (
	"fmt"
	"github.com/rs/xid"
)

type ResourceType interface {
	prefix() string
}
type ID[T ResourceType] xid.ID

type App struct{}

func (a App) prefix() string { return "app" }

type Trace struct{}

func (t Trace) prefix() string { return "trace" }

func New[T ResourceType]() ID[T] {
	return ID[T](xid.New())
}

func (id ID[T]) String() string {
	var resourceType T

	return fmt.Sprintf(
		"%s_%s",
		resourceType.prefix(),
		xid.ID(id).String(),
	)
}
