package main

import "time"

type Task struct {
	ID    uint      `gorm:"primaryKey;not null;autoIncrement;unique" json:"id"`
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
	Date  time.Time `json:"date"`
}
