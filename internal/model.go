package internal

import "time"

type Trace struct {
	Spans []Span `json:"spans"`
}

type Span struct {
	TraceID       string        `json:"traceId"`
	SpanID        string        `json:"spanId"`
	OperationName string        `json:"operationName"`
	References    []interface{} `json:"references"`
	StartTime     time.Time     `json:"startTime"`
	Duration      string        `json:"duration"`
	Tags          []Tag         `json:"tags"`
	Process       Process       `json:"process"`
	Logs          []Log         `json:"logs"`
}

type Log struct {
	Timestamp time.Time     `json:"timestamp"`
	Fields    []interface{} `json:"fields"`
}

type Process struct {
	ServiceName string        `json:"serviceName"`
	Tags        []interface{} `json:"tags"`
}

type Tag struct {
	Key   string `json:"key"`
	VType string `json:"vType"`
	VStr  string `json:"vStr"`
}
