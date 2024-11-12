package models

import (
	"time"
)

type Job struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	Type           string    `gorm:"type:varchar(20);not null" json:"type"`
	Status         string    `gorm:"type:varchar(20);not null" json:"status"`
	ScheduledAt    time.Time `json:"scheduled_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	RecurringCron  string    `gorm:"type:varchar(50)" json:"recurring_cron"`
	ExecutionCount int       `json:"execution_count"`
	LastExecutedAt time.Time `json:"last_executed_at,omitempty"`
}

type JobDTO struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

func (j *Job) ToDTO() JobDTO {
	return JobDTO{
		ID:          j.ID,
		UserID:      j.UserID,
		Type:        j.Type,
		Status:      j.Status,
		ScheduledAt: j.ScheduledAt,
	}
}
