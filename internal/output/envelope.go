package output

import (
	"encoding/json"
)

type Envelope struct {
	OK    bool        `json:"ok"`
	Data  interface{} `json:"data,omitempty"`
	Error *ErrorInfo  `json:"error,omitempty"`
	Meta  *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code       interface{} `json:"code,omitempty"`
	Message    string      `json:"message"`
	Suggestion string      `json:"suggestion,omitempty"`
}

type Meta struct {
	Page       int    `json:"page,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
	Identity   string `json:"identity,omitempty"`
}

func SuccessEnvelope(data interface{}, meta *Meta) *Envelope {
	return &Envelope{OK: true, Data: data, Meta: meta}
}

func ErrorEnvelope(code interface{}, message, suggestion string) *Envelope {
	return &Envelope{
		OK: false,
		Error: &ErrorInfo{
			Code:       code,
			Message:    message,
			Suggestion: suggestion,
		},
	}
}

func (e *Envelope) JSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}
