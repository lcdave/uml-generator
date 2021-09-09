package models

import (
	"time"
)

// album represents data about a record album.
type LogEntry struct {
    ID int `json:"id"`
    ServiceId int `json:"service_id"`
    Description string `json:"description"`
    CreatedAt time.Time `json:"created_at"`
}
type Logs struct {
    Logs []LogEntry `json:"logs"`
}

type LogEntryOutput struct {
    URL string `json:"url"`
    Description string `json:"description"`
    Time time.Time `json:"time"`
}
type LogsOutput struct {
    LogsOutput []LogEntryOutput `json:"logsOutput"`
}