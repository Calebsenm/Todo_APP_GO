package db

import (
	"time"
)

type Task struct {
	Tittle string    `json:"tittle"`
	Text   string    `json:"text"`
	Date   time.Time `json:"date"`
}
