package models

import (
	"time"
)

// album represents data about a record album.
type Request struct {
    ID int `json:"id"`
    ServiceId int `json:"service_id"`
    ResponseTime float32 `json:"response_time"`
    CreatedAt time.Time `json:"created_at"`
}
type RequestList struct {
    Requests []Request `json:"requests"`
}