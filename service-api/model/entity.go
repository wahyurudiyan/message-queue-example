package model

type Notification struct {
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	Data      map[string]string `json:"data,omitempty"`
	Receivers []string          `json:"receivers"`
}
