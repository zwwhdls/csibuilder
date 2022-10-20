package machinery

import (
	"github.com/spf13/afero"
)

type Filesystem struct {
	FS afero.Fs
}
