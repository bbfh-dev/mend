package attrs

import "io"

type writer interface {
	io.Writer
	io.StringWriter
}
