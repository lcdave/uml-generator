package models

import (
	"time"
)

// album represents data about a record album.
type MailLogEntry struct {
    ID int `json:"id"`
    ServiceId int `json:"service_id"`
    Recipient string `json:"recipient"`
    CreatedAt time.Time `json:"created_at"`
}
type MailLogEntryList struct {
    MailLogEntryList []MailLogEntry `json:"mailLogEntryList"`
}