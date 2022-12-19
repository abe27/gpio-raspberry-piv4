package models

import "time"

type Response struct {
	StatusCode   int         `json:"status_code,omitempty"`
	IsSuccessful bool        `json:"is_successful,omitempty"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	OnTime       time.Time   `json:"on_time" default:"now"`
}
