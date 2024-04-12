package ts

import (
	"go.k6.io/k6/js/modules"
)

type extension struct{}

func init() {
	modules.Register("k6/x/fake-ts-module-just-for-k6-version-command", new(extension))
}
