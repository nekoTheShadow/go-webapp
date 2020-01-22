package trace

import "io"

type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return nil
}
