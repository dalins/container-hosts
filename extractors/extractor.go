package extractors

import (
	"github.com/docker/docker/api/types"
)

type Extractor interface {
	HasFilterLabel() bool
	FilterLabel() string
	HostnameFromContainer(container types.Container) (string, error)
}
