package generator

import (
	"fmt"
	"slices"
	"strings"

	"github.com/happsie/trace2dia/internal"
)

type PlantUML struct{}

func NewPlantUML() PlantUML {
	return PlantUML{}
}

func (p PlantUML) Generate(trace internal.Trace) error {
	spans := trace.Spans
	sb := strings.Builder{}

	sb.WriteString("@startuml\n")
	sb.WriteString(fmt.Sprintf("title %s\n", spans[0].TraceID))
	slices.SortFunc(spans, func(a internal.Span, b internal.Span) int {
		return a.StartTime.Compare(b.StartTime)
	})
	for i, s := range spans {
		if i+1 >= len(spans) {
			continue
		}
		nextSpan := spans[i+1]
		sb.WriteString(fmt.Sprintf("%s->%s: %s\n", strings.ReplaceAll(s.Process.ServiceName, "-", ""), strings.ReplaceAll(nextSpan.Process.ServiceName, "-", ""), nextSpan.OperationName))
	}
	sb.WriteString("@enduml")
	fmt.Print(sb.String())
	return nil
}
