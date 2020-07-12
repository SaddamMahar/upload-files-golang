package models

import "time"

type FilesList struct {
	Name    string
	Size    uint64
	Created time.Time
}
