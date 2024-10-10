package generator

import (
	"github.com/happsie/trace2dia/internal"
)

type OutputGenerator interface {
	Generate(trace internal.Trace) error
}