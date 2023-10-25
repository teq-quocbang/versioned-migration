package models

import "time"

type Semester struct {
	ID                string
	MinCredits        int
	StartTime         time.Time
	EndTime           time.Time
	RegisterStartAt   time.Time
	RegisterExpiresAt time.Time
	CreatedAt         time.Time
	CreatedBy         *uint
	UpdatedAt         time.Time
	UpdatedBy         *uint
}

func (Semester) TableName() string {
	return "semester"
}
