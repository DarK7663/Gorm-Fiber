package main

import "time"

type Task struct {
	ID    uint      `gorm:"primaryKey;not null;autoIncrement:true;unique" json:"id"`
	Title string    `json:"title"`
	Time  int64     `json:"time"`
	Date  time.Time `gorm:"not null" json:"date"`
}
