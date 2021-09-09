package models

import (
	"time"
)

// album represents data about a record album.
type Service struct {
    ID int `json:"id"`
    Name string `json:"name"`
    URL string `json:"url"`
    CreatedAt time.Time `json:"created_at"`
	Status string `json:"status"`
}
type ServiceList struct {
    Services []Service `json:"services"`
}