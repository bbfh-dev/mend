package std

import (
	"embed"
	"fmt"
	"io/fs"
)

const PREFIX = "std:"
const PREFIX_LEN = len(PREFIX)

//go:embed *.html
var FS embed.FS

func Open(name string) (fs.File, error) {
	file, err := FS.Open(fmt.Sprintf("%s.html", name))
	if err != nil {
		err = fmt.Errorf("(standard library): %w", err)
	}
	return file, err
}
